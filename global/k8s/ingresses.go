package global

import (
	"context"
	"fmt"

	"github.com/Sirupsen/logrus"
	//"github.com/ericchiang/k8s"
	v1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
	"github.com/valentin2105/Nemo/global"
)

// Ingress - kubectl get ingress
type Ingress struct {
	Name      string
	Namespace string
	Host      string
	//BackendName string
	//BackendPort string
}

// IngressList - list of ingress
type IngressList []Ingress

// ListIngresses - return a list of ingress
func ListIngresses() (IngressList, error) {
	il := make(IngressList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		logrus.Warn("Error " + err.Error())
		return il, err
	}

	var ingresses v1beta1.IngressList
	if err := client.List(context.Background(), "", &ingresses); err != nil {
		logrus.Warn("Error " + err.Error())
		return il, err
	}
	for _, ingress := range ingresses.Items {
		//Name
		n := *ingress.Metadata.Name
		nc := global.TrimQuotes(n)
		// Namespace
		ns := *ingress.Metadata.Namespace
		nsc := global.TrimQuotes(ns)
		// Host
		i := fmt.Sprintf("%q", ingress.Spec.GetRules())
		ic := global.TrimQuotes(i)
		// Port
		//		po := fmt.Sprintf("%q", services.Spec.GetPorts())
		//		poc := global.TrimQuotes(po)
		// Type
		//		t := services.Spec.GetType()
		//		tc := global.TrimQuotes(t)
		// Put in slice
		ia := Ingress{Name: nc, Namespace: nsc, Host: ic}
		il = append(il, ia)
	}
	return il, err
}
