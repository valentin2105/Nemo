package global

import (
	"context"
	"fmt"

	"github.com/Sirupsen/logrus"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// Service - kubectl get service
type Service struct {
	Name      string
	Namespace string
	Type      string
	IP        string
	Port      string
	Label     string
}

// ServiceList - list of service
type ServiceList []Service

// ListServices - return a list of service
func ListServices() (ServiceList, error) {
	sl := make(ServiceList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		logrus.Warn("Error " + err.Error())
		return sl, err
	}

	var services corev1.ServiceList
	if err := client.List(context.Background(), "", &services); err != nil {
		logrus.Warn("Error " + err.Error())
		return sl, err
	}
	for _, services := range services.Items {
		//Name
		n := *services.Metadata.Name
		nc := TrimQuotes(n)
		// Namespace
		ns := *services.Metadata.Namespace
		nsc := TrimQuotes(ns)
		// IP
		i := services.Spec.GetClusterIP()
		ic := TrimQuotes(i)
		// Port
		po := fmt.Sprintf("%q", services.Spec.GetPorts())
		poc := TrimQuotes(po)
		// Type
		t := services.Spec.GetType()
		tc := TrimQuotes(t)
		// Put in slice
		s := Service{Name: nc, Namespace: nsc, IP: ic, Port: poc, Type: tc}
		sl = append(sl, s)
	}
	return sl, err
}

// GetService - describe a service
func GetService(ns string, name string) (Service, error) {
	var s Service
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		logrus.Warn("Error " + err.Error())
		return s, err
	}
	var service corev1.Service
	if err := client.Get(context.Background(), ns, name, &service); err != nil {
		logrus.Warn("Error " + err.Error())
		return s, err
	}
	//Name
	n := *service.Metadata.Name
	nc := TrimQuotes(n)
	// Namespace
	ns = *service.Metadata.Namespace
	nsc := TrimQuotes(ns)
	// IP
	i := service.Spec.GetClusterIP()
	ic := TrimQuotes(i)
	// Port
	po := fmt.Sprintf("%q", service.Spec.GetPorts())
	poc := TrimQuotes(po)
	// Type
	t := service.Spec.GetType()
	tc := TrimQuotes(t)
	// Put in slice
	s = Service{Name: nc, Namespace: nsc, IP: ic, Port: poc, Type: tc}
	return s, err
}
