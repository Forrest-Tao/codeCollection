package utils

import (
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func getHomeDir() string {
	return homedir.HomeDir()
}

func GetKubePath() string {
	return filepath.Join(getHomeDir(), ".kube", "tx.config")
}
