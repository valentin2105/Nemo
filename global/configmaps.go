package global

import (
	"context"
	"fmt"
	"log"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type Configmap struct {
	Name      string
	Namespace string
	Age       string
}

type ConfigmapList []Configmap

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
		n := fmt.Sprintf("%q", *configmaps.Metadata.Name)
		nc := TrimQuotes(n)
		// Namespace
		ns := fmt.Sprintf("%q", *configmaps.Metadata.Namespace)
		nsc := TrimQuotes(ns)
		//Age
		a := fmt.Sprintf("%q", *configmaps.Metadata.CreationTimestamp)
		ac := TrimQuotes(a)
		// Put in slice
		p := Configmap{Name: nc, Namespace: nsc, Age: ac}
		pl = append(pl, p)
	}
	return pl
}
