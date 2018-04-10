package global

import (
	"context"
	"time"

	"github.com/Sirupsen/logrus"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/valentin2105/Nemo/global"
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
		logrus.Warn("Error " + err.Error())
	}

	var configmaps corev1.ConfigMapList
	if err := client.List(context.Background(), "", &configmaps); err != nil {
		logrus.Warn("Error " + err.Error())
	}
	for _, configmaps := range configmaps.Items {
		//Name
		n := *configmaps.Metadata.Name
		nc := global.TrimQuotes(n)
		// Namespace
		ns := *configmaps.Metadata.Namespace
		nsc := global.TrimQuotes(ns)
		//Created
		c := configmaps.Metadata.GetCreationTimestamp()
		cs := c.GetSeconds()
		csc := time.Unix(cs, 0)
		cscf := csc.Format(global.DefaultDateFormat)

		// Put in slice
		p := Configmap{Name: nc, Namespace: nsc, Created: cscf}
		pl = append(pl, p)
	}
	return pl
}
