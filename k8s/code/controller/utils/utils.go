package utils

import "k8s.io/client-go/util/homedir"

func GetStringPtr(s string) *string {
	return &s
}

func GetHomeDir() string {
	return homedir.HomeDir()
}
