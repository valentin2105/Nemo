package global

import (
	"context"
	"fmt"
	"log"

	"github.com/ericchiang/k8s"
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
	appsv1beta1 "github.com/ericchiang/k8s/apis/apps/v1beta1"
)

// Deployment - kubectl get deployments
type Deployment struct {
	Status     string
	Name       string
	Namespace  string
	PodWanted  int32
	PodRunning int32
	Image      string
}

// DeploymentList - list of deployments
type DeploymentList []Deployment

// ListDeployments - return a list of deploys
func ListDeployments() DeploymentList {
	dl := make(DeploymentList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	version := GetEnv("KUBERNETES_VERSION", "v1.9")
	if version == "v1.8" || version == "v1.7" || version == "v1.6" {
		var deployments appsv1beta1.DeploymentList

		if err := client.List(context.Background(), k8s.AllNamespaces, &deployments); err != nil {
			log.Fatal(err)
		}
		for _, deployments := range deployments.Items {
			//Name
			n := fmt.Sprintf("%q", *deployments.Metadata.Name)
			nc := TrimQuotes(n)
			// Namespace
			ns := fmt.Sprintf("%q", *deployments.Metadata.Namespace)
			nsc := TrimQuotes(ns)
			// PodWanted
			pw := *deployments.Status.Replicas
			// PodRunning
			pr := *deployments.Status.AvailableReplicas
			st := "Ready"
			if pw != pr {
				st = "NotReady"
			}
			// Put in slice
			d := Deployment{Status: st, Name: nc, Namespace: nsc, PodWanted: pw, PodRunning: pr}
			dl = append(dl, d)
		}
		return dl

	} else {
		var deployments appsv1.DeploymentList
		if err := client.List(context.Background(), k8s.AllNamespaces, &deployments); err != nil {
			log.Fatal(err)
		}
		for _, deployments := range deployments.Items {
			//Name
			n := fmt.Sprintf("%q", *deployments.Metadata.Name)
			nc := TrimQuotes(n)
			// Namespace
			ns := fmt.Sprintf("%q", *deployments.Metadata.Namespace)
			nsc := TrimQuotes(ns)
			// PodWanted
			pw := *deployments.Status.Replicas
			// PodRunning
			pr := *deployments.Status.AvailableReplicas
			st := "Ready"
			if pw != pr {
				st = "NotReady"
			}
			// Put in slice
			d := Deployment{Status: st, Name: nc, Namespace: nsc, PodWanted: pw, PodRunning: pr}
			dl = append(dl, d)
		}
		return dl
	}
}
