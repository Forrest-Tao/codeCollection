package init

import (
	"codeCollection/k8s/code/client-go/utils"
	"k8s.io/client-go/kubernetes"
)

var ClientSet *kubernetes.Clientset

func init() {
	var err error
	ClientSet, err = utils.GetKubeClientSet(utils.GetHomeDir())
	if err != nil {
		panic(err)
	}
}
