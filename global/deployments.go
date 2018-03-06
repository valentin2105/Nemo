package global

import (
	"context"

	"github.com/Sirupsen/logrus"
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
		logrus.Warn("Error " + err.Error())
	}
	var deployments appsv1.DeploymentList
	if err := client.List(context.Background(), k8s.AllNamespaces, &deployments); err != nil {
		logrus.Warn("Error " + err.Error())

		var deployments appsv1beta1.DeploymentList
		if err := client.List(context.Background(), k8s.AllNamespaces, &deployments); err != nil {
			logrus.Warn("Error " + err.Error())
		}

	}
	for _, deployments := range deployments.Items {
		//Name
		n := *deployments.Metadata.Name
		nc := TrimQuotes(n)
		// Namespace
		ns := *deployments.Metadata.Namespace
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

// GetDeployment - describe a deployment
func GetDeployment(ns string, name string) Deployment {
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		logrus.Warn("Error " + err.Error())
	}
	var deployment appsv1.Deployment
	if err := client.Get(context.Background(), ns, name, &deployment); err != nil {

		var deployment appsv1beta1.Deployment
		if err := client.Get(context.Background(), ns, name, &deployment); err != nil {
			logrus.Warn("Error " + err.Error())
		}

	}
	//Name
	n := *deployment.Metadata.Name
	nc := TrimQuotes(n)
	// Namespace
	ns = *deployment.Metadata.Namespace
	nsc := TrimQuotes(ns)
	// PodWanted
	pw := *deployment.Status.Replicas
	// PodRunning
	pr := *deployment.Status.AvailableReplicas
	st := "Ready"
	if pw != pr {
		st = "NotReady"
	}
	// Put in slice
	d := Deployment{Status: st, Name: nc, Namespace: nsc, PodWanted: pw, PodRunning: pr}
	return d
}
