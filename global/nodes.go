package global

import (
	"context"
	"log"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// Node - kubectl get node
type Node struct {
	Status      string
	Name        string
	Schedulable bool
}

// NodeList - list of node
type NodeList []Node

// ListNodes - return a list of nodes
func ListNodes() NodeList {
	nl := make(NodeList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	var nodes corev1.NodeList
	if err := client.List(context.Background(), "", &nodes); err != nil {
		log.Fatal(err)
	}
	for _, node := range nodes.Items {
		//Status
		status := node.Status.GetConditions()
		var st string
		for _, stat := range status {
			if *stat.Type == "Ready" || *stat.Type == "NotReady" {
				st = TrimQuotes(*stat.Type)
			}
		}
		//Name
		n := *node.Metadata.Name
		nc := TrimQuotes(n)

		//Spec
		//sp := node.Status.GetAllocatable()

		sch := !*node.Spec.Unschedulable
		// Put in slice
		no := Node{Status: st, Name: nc, Schedulable: sch}
		nl = append(nl, no)
	}
	return nl
}

// GetNode - describe a node
func GetNode(name string) Node {
	var no Node
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	var node corev1.Node
	if err := client.Get(context.Background(), "", name, &node); err != nil {
		log.Fatal(err)
	}

	//Status
	status := node.Status.GetConditions()
	var st string
	for _, stat := range status {
		if *stat.Type == "Ready" || *stat.Type == "NotReady" {
			st = TrimQuotes(*stat.Type)
		}
	}
	//Name
	n := *node.Metadata.Name
	nc := TrimQuotes(n)

	//Spec
	//sp := node.Status.GetAllocatable()

	sch := !*node.Spec.Unschedulable
	// Put in slice
	no = Node{Status: st, Name: nc, Schedulable: sch}

	return no
}
