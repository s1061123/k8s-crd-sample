package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	apiv1 "github.com/s1061123/k8s-crd-sample/pkg/apis/k8s.cni.cncf.io/v1"
	clientset "github.com/s1061123/k8s-crd-sample/pkg/client/clientset/versioned"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func PrintSampleCRD (crd *apiv1.SampleCRD) {
	fmt.Printf("Name: %s\n", crd.ObjectMeta.Name)
	fmt.Printf("Protocol: %s\n", string(*crd.Spec.Protocol))
	fmt.Printf("Port: %s\n", crd.Spec.Port.String())
	fmt.Printf("LabelSelector: %s\n", crd.Spec.Selector.String())
	fmt.Printf("\tMatchLabels: %v\n", crd.Spec.Selector.MatchLabels)
}

func main() {
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %v", err)
	}

	list, err := exampleClient.K8sCniCncfIoV1().SampleCRDs("default").List(
		context.TODO(), metav1.ListOptions{})
	if err != nil {
		glog.Fatalf("Error listing all network attachment definitions: %v", err)
	}

	for _, sample := range list.Items {
		//fmt.Printf("sample: %v", sample)
		PrintSampleCRD(&sample)
	}
}
