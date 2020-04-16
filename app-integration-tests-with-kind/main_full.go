package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func TestMain(m *testing.M) {
	os.Exit(testStartup(m))
}

var clientset *kubernetes.Clientset

func testStartup(m *testing.M) int {

	// Kubeconfig
	var kubeconfig string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Provide path to kubeconfig")
	flag.Parse()

	if kubeconfig == "" {
		panic("-kubeconfig flag can't be empty")
	}

	// Cluster
	deleteCmd := exec.Command("kind", "delete", "cluster")
	deleteCmd.Stdout = os.Stdout
	deleteCmd.Stderr = os.Stderr
	defer deleteCmd.Run()
	defer log.Println("DELETE cluster...")

	createCmd := exec.Command("kind", "create", "cluster")
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr
	log.Println("Starting cluster...")

	err := createCmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Client
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return m.Run()
}

func TestPod(t *testing.T) {
	t.Log("running test TestPod()")

	podList, err := clientset.CoreV1().Pods("kube-system").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, item := range podList.Items {
		t.Log(item.Name)
	}
}

func TestNetpol(t *testing.T) {
	t.Log("running test TestNetpol()")
}

func TestDeploy(t *testing.T) {
	t.Log("running test TestDeploy()")
}

func TestSvc(t *testing.T) {
	t.Log("running test TestSvc()")
}
