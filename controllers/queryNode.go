package controllers

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type QueryNodeController struct {
	MainController
}

func (c *QueryNodeController) Get() {
	nodes, err := ClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Nodes().List() failed: %v \n", err)
	}

	c.Data["json"] = nodes
	c.ServeJSON()
}

// func (c *QueryNodeController) Get() {
// 	var kubeconfig *string
// 	if home := homedir.HomeDir(); home != "" {
// 		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
// 	} else {
// 		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
// 	}
// 	flag.Parse()

// 	*kubeconfig = "/etc/rancher/k3s/k3s.yaml"

// 	// use the current context in kubeconfig
// 	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// create the clientset
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for {
// 		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

// 		for _, pod := range pods.Items {
// 			fmt.Println("ns: ", pod.GetNamespace())
// 		}

// 		nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		for _, node := range nodes.Items {
// 			fmt.Println("name: ", node.GetName(), " gName: ", node.GetGenerateName())
// 		}

// 		// Examples for error handling:
// 		// - Use helper functions like e.g. errors.IsNotFound()
// 		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
// 		namespace := "default"
// 		pod := "example-xxxxx"
// 		_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
// 		if errors.IsNotFound(err) {
// 			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
// 		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
// 			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
// 				pod, namespace, statusError.ErrStatus.Message)
// 		} else if err != nil {
// 			panic(err.Error())
// 		} else {
// 			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
// 		}

// 		time.Sleep(10 * time.Second)
// 	}
// }
