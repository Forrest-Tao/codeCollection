package service

import (
	"context"
	"encoding/json"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"os"
)

var ctx = context.Background()

func ReadServiceYaml(filename string) coreV1.Service {
	serviceYaml, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	serviceJson, err := yaml.ToJSON(serviceYaml)
	if err != nil {
		panic(err)
	}
	var service coreV1.Service
	if json.Unmarshal(serviceJson, &service) != nil {
		panic(err)
	}
	return service
}

func ApplyServcice(clientset *kubernetes.Clientset, new_service coreV1.Service) (err error) {
	//err！=nil，说明service不存在或查询失败
	if _, err = clientset.CoreV1().Services(new_service.Namespace).Get(ctx, new_service.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			//说明查询出错
			return
		}
		//err！=nil&&IsNotFound()
		//创建service
		_, err = clientset.CoreV1().Services(new_service.Namespace).Create(ctx, &new_service, metav1.CreateOptions{})
		return
	}
	//err==nil 说明service存在，则更新
	_, err = clientset.CoreV1().Services(new_service.Namespace).Update(ctx, &new_service, metav1.UpdateOptions{})
	return
}
