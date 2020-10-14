package kube

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func LoadKubeConfigMap() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("failed to load kubernetes config %v\n", err)
		return
	}
	// creates the clientset
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	configMap, err := client.CoreV1().ConfigMaps("default").Get(context.Background(), "cerberus", metav1.GetOptions{})
	if err != nil {
		return
	}
	for key, value := range configMap.Data {
		viper.Set(key, value)
	}
}
