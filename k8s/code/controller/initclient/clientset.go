package initclient

import (
	"forrest/codeCollection/k8s/code/controller/utils"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

var ClientSet = Client{}

type Client struct {
	Client        kubernetes.Interface
	DynamicClient dynamic.Interface
}

func GetClient() Client {
	return ClientSet
}

func getKubeConfigPath() *string {
	var kubeConfigPath string
	kubeConfigPath = filepath.Join(utils.GetHomeDir(), ".kube", "config")
	return utils.GetStringPtr(kubeConfigPath)
}

func getInternalConfig() (*rest.Config, error) {
	return rest.InClusterConfig()
}

func getClusterConfig() (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", *getKubeConfigPath())
}

func getConfig() *rest.Config {
	internalClient, err := getInternalConfig()
	if err != nil {
		clusterConfig, err := getClusterConfig()
		if err != nil {
			panic(err)
		}
		return clusterConfig
	}
	return internalClient
}

func init() {
	cfg := getConfig()
	ClientSet.Client = InitClient(cfg)
	ClientSet.DynamicClient = InitDynamicClient(cfg)
}

func InitClient(cfg *rest.Config) kubernetes.Interface {
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	return clientset
}

func InitDynamicClient(cfg *rest.Config) dynamic.Interface {
	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	return dynamicClient
}
