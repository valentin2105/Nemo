package global

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/ghodss/yaml"
)

var (
	//Kubeconfig - Kubernetes configuration file path
	Kubeconfig = flag.String("kubeconfig", "/root/.kube/config", "KubeConfig Path")
)

// ComponentStatus - kubectl get cs
type ComponentStatus struct {
	Name   string
	Status string
}

// ComponentStatusList - List of cs
type ComponentStatusList []ComponentStatus

// LoadClient - Create kubernetes connexion client
func LoadClient(kubeconfigPath *string) (*k8s.Client, error) {
	flag.Parse()
	path := string(*kubeconfigPath)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read kubeconfig: %v", err)
	}
	// Unmarshal YAML into a Kubernetes config object.
	var config k8s.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
	}
	return k8s.NewClient(&config)
}

// ListComponentStatus - Return a cs list
func ListComponentStatus() ComponentStatusList {
	nl := make(ComponentStatusList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	var components corev1.ComponentStatusList
	if err := client.List(context.Background(), "", &components); err != nil {
		log.Fatal(err)
	}
	for _, component := range components.Items {
		//Status
		status := component.GetConditions()
		var st string
		for _, stat := range status {
			if *stat.Type == "Healthy" || *stat.Type == "Unhealthy" {
				st = TrimQuotes(*stat.Type)
			}
		}
		//Name
		n := fmt.Sprintf("%q", *component.Metadata.Name)
		nc := TrimQuotes(n)
		// Put in slice
		no := ComponentStatus{Status: st, Name: nc}
		nl = append(nl, no)
	}
	return nl
}
