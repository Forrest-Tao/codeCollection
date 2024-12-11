package vm

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	v1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/core/v1"
	virtCli "kubevirt.io/client-go/kubecli"
	"os"
	"os/signal"
)

func ConsoleHandler(conn *websocket.Conn, client virtCli.KubevirtClient, namespace string, name string) (err error) {
	pipeInReader, pipeInWriter := io.Pipe()
	pipeOutReader, pipeOutWriter := io.Pipe()
	resChan := make(chan error)
	runningChan := make(chan error)
	waitInterrupt := make(chan os.Signal, 1)
	signal.Notify(waitInterrupt, os.Interrupt)
	go func() {
		consoleStream, err := client.VirtualMachineInstance(namespace).SerialConsole(name, nil)
		runningChan <- err
		resChan <- consoleStream.Stream(v1.StreamOptions{
			In:  pipeInReader,
			Out: pipeOutWriter,
		})
	}()

	if err = <-runningChan; err != nil {
		return err
	}

	return attachConsole(conn, pipeInReader, pipeOutReader, pipeInWriter, pipeOutWriter, "", resChan)
}

func attachConsole(
	conn *websocket.Conn, // 使用 gorilla/websocket 的连接对象
	stdinReader, stdoutReader *io.PipeReader,
	stdinWriter, stdoutWriter *io.PipeWriter,
	message string, resChan <-chan error,
) (err error) {

	stopChan := make(chan struct{}, 1)
	writeStop := make(chan error)
	readStop := make(chan error)

	fmt.Fprint(os.Stderr, message)

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		<-interrupt
		close(stopChan)
	}()

	// 从 stdoutReader 读取数据，写入 WebSocket 发送给浏览器
	go func() {
		defer close(readStop)
		buf := make([]byte, 1024)
		for {
			n, err := stdoutReader.Read(buf)
			if err != nil {
				readStop <- err
				return
			}
			err = conn.WriteMessage(websocket.BinaryMessage, buf[:n])
			if err != nil {
				readStop <- err
				return
			}
		}
	}()

	// 从 WebSocket 读取数据，写入 stdinWriter 发送给虚拟机控制台
	go func() {
		defer close(writeStop)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				writeStop <- err
				return
			}

			// 检查是否是转义序列
			if len(message) > 0 && message[0] == 29 {
				return
			}

			_, err = stdinWriter.Write(message)
			if err != nil {
				writeStop <- err
				return
			}
		}
	}()

	select {
	case <-stopChan:
	case err = <-readStop:
	case err = <-writeStop:
	case err = <-resChan:
	}

	return err
}
