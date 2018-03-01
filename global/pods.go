package global

import (
	"context"
	"fmt"
	"log"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// Pod - kubectl get pod
type Pod struct {
	Status    string
	Name      string
	Namespace string
	Worker    string
	Ip        string
	Image     string
}

// PodList - list of pod
type PodList []Pod

// ListPods - return a list of pod
func ListPods() PodList {
	pl := make(PodList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	var pods corev1.PodList
	if err := client.List(context.Background(), "", &pods); err != nil {
		log.Fatal(err)
	}
	for _, pods := range pods.Items {
		//Status
		s := fmt.Sprintf("%q", *pods.Status.Phase)
		sc := TrimQuotes(s)
		si := ChooseStatusFaIcon(sc)
		//Name
		n := fmt.Sprintf("%q", *pods.Metadata.Name)
		nc := TrimQuotes(n)
		// Namespace
		ns := fmt.Sprintf("%q", *pods.Metadata.Namespace)
		nsc := TrimQuotes(ns)
		//Worker
		w := fmt.Sprintf("%q", *pods.Spec.NodeName)
		wc := TrimQuotes(w)
		// Put in slice
		p := Pod{Status: si, Name: nc, Namespace: nsc, Worker: wc}
		pl = append(pl, p)
	}
	return pl
}

// GetPod - describe a pod
func GetPod(ns string, name string) Pod {
	var p Pod
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	var pod corev1.Pod
	if err := client.Get(context.Background(), ns, name, &pod); err != nil {
		log.Fatal(err)
	}
	//Status
	s := fmt.Sprintf("%q", *pod.Status.Phase)
	sc := TrimQuotes(s)
	//Name
	n := fmt.Sprintf("%q", *pod.Metadata.Name)
	nc := TrimQuotes(n)
	// Namespace
	nss := fmt.Sprintf("%q", *pod.Metadata.Namespace)
	nsc := TrimQuotes(nss)
	//Worker
	w := fmt.Sprintf("%q", *pod.Spec.NodeName)
	wc := TrimQuotes(w)
	//IP
	ip := pod.Status.GetPodIP()
	//Image
	ci := pod.Spec.Containers[0]
	image := ci.GetImage()
	// Put in slice
	p = Pod{Status: sc, Name: nc, Namespace: nsc, Worker: wc, Ip: ip, Image: image}
	return p
}
