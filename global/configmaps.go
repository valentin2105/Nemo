package global

import (
	"context"
	"log"
	"time"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// Configmap - kubectl get configmap
type Configmap struct {
	Name      string
	Namespace string
	Created   string
}

// ConfigmapList - list of configmaps
type ConfigmapList []Configmap

// ListConfigmaps - return a cm list
func ListConfigmaps() ConfigmapList {
	pl := make(ConfigmapList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	var configmaps corev1.ConfigMapList
	if err := client.List(context.Background(), "", &configmaps); err != nil {
		log.Fatal(err)
	}
	for _, configmaps := range configmaps.Items {
		//Name
		n := *configmaps.Metadata.Name
		nc := TrimQuotes(n)
		// Namespace
		ns := *configmaps.Metadata.Namespace
		nsc := TrimQuotes(ns)
		//Created
		c := configmaps.Metadata.GetCreationTimestamp()
		cs := c.GetSeconds()
		csc := time.Unix(cs, 0)
		cscf := csc.Format(DefaultDateFormat)

		// Put in slice
		p := Configmap{Name: nc, Namespace: nsc, Created: cscf}
		pl = append(pl, p)
	}
	return pl
}
