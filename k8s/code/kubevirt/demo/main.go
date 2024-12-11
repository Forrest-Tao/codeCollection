package main

import (
	vm2 "codeCollection/k8s/code/kubevirt/demo/pkg/vm"
	"codeCollection/k8s/code/kubevirt/demo/utils"
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
	virtcorev1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"
	"log"
)

var virtClient kubecli.KubevirtClient
var ctx = context.Background()
var err error

func init() {
	config, err := clientcmd.BuildConfigFromFlags("", utils.GetKubePath())
	if err != nil {
		log.Fatalf(err.Error())
	}
	virtClient, err = kubecli.GetKubevirtClientFromRESTConfig(config)
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt client: %v\n", err)
	}
}

const (
	namespaceDefault = "kubevirt"
	virtName         = `cirros`
)

func main() {
	router := gin.Default()
	//vm
	router.POST("/vm", func(c *gin.Context) {
		_, err = virtClient.VirtualMachine(namespaceDefault).Create(ctx, vm2.CreateDefaultVM(namespaceDefault, virtName), metav1.CreateOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"message": "VM created successfully"})
		}
	})
	router.DELETE("/vm", func(c *gin.Context) {
		err = virtClient.VirtualMachine(namespaceDefault).Delete(ctx, virtName, metav1.DeleteOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"message": "VM deleted successfully"})
		}
	})
	router.GET("/vm", func(c *gin.Context) {
		vm, err := virtClient.VirtualMachine(namespaceDefault).Get(ctx, virtName, metav1.GetOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"vm": vm})
		}
	})
	router.GET("/vm/status", func(c *gin.Context) {
		//获取vm的详细状态
		vm, err := virtClient.VirtualMachine(namespaceDefault).Get(ctx, virtName, metav1.GetOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"status": vm.Status.Conditions[0]})
	})

	//restart
	router.POST("/vm/restart", func(c *gin.Context) {
		//重启vm
		err = virtClient.VirtualMachine(namespaceDefault).Restart(ctx, virtName, &virtcorev1.RestartOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"message": "VM restarted successfully"})
		}
	})
	//stop
	router.POST("/vm/stop", func(c *gin.Context) {
		err = virtClient.VirtualMachine(namespaceDefault).Stop(ctx, virtName, &virtcorev1.StopOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"message": "VM stopped successfully"})
		}
	})
	//start
	router.POST("/vm/start", func(c *gin.Context) {
		err = virtClient.VirtualMachine(namespaceDefault).Start(ctx, virtName, &virtcorev1.StartOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"message": "VM started successfully"})
		}
	})

	//snapshot
	router.POST("/snapshot/:name", func(c *gin.Context) {
		name := c.Param("name")
		snap, err2 := virtClient.VirtualMachineSnapshot(namespaceDefault).Create(ctx, vm2.CreateSnapshot(namespaceDefault, name), metav1.CreateOptions{})
		if err2 != nil {
			c.JSON(500, gin.H{"error": err2.Error()})
		} else {
			c.JSON(200, gin.H{"message": "Snapshot created successfully", "snapshot": snap})
		}
	})

	router.GET("/snapshot/:name", func(c *gin.Context) {
		name := c.Param("name")
		snapshot, err := virtClient.VirtualMachineSnapshot(namespaceDefault).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"phase": snapshot.Status.Phase})
	})
	router.DELETE("/snapshot/:name", func(c *gin.Context) {
		name := c.Param("name")
		err := virtClient.VirtualMachineSnapshot(namespaceDefault).Delete(ctx, name, metav1.DeleteOptions{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"message": "Snapshot deleted successfully"})
		}
	})

	//TODO： addvolume

	runtime.Must(router.Run(":8989"))
}
