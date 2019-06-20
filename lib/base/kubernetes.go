package base

import (
	"flag"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

type K8sOp struct {
	clientset *kubernetes.Clientset
}

func (e *K8sOp) Init() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	e.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func (e *K8sOp) GetPodInfo() *v1.PodList {
	pods, err := e.clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	return pods
}

func (e *K8sOp) GetNodeInfo() *v1.NodeList {
	nodes, err := e.clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	return nodes
}
