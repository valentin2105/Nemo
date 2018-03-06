package global

import (
	"context"

	"github.com/Sirupsen/logrus"
	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

// Pod - kubectl get pod
type Pod struct {
	Status    string
	Name      string
	Namespace string
	Worker    string
	IP        string
	Image     string
}

// PodList - list of pod
type PodList []Pod

// ListPods - return a list of pod
func ListPods() (PodList, error) {
	pl := make(PodList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		logrus.Warn("Error " + err.Error())
		return pl, err
	}

	var pods corev1.PodList
	if err := client.List(context.Background(), "", &pods); err != nil {
		return pl, err
		logrus.Warn("Error " + err.Error())
	}
	for _, pods := range pods.Items {
		//Status
		s := *pods.Status.Phase
		sc := TrimQuotes(s)
		//Name
		n := *pods.Metadata.Name
		nc := TrimQuotes(n)
		// Namespace
		ns := *pods.Metadata.Namespace
		nsc := TrimQuotes(ns)
		//Worker
		w := *pods.Spec.NodeName
		wc := TrimQuotes(w)
		// Put in slice
		p := Pod{Status: sc, Name: nc, Namespace: nsc, Worker: wc}
		pl = append(pl, p)
	}
	return pl, err
}

// GetPod - describe a pod
func GetPod(ns string, name string) (Pod, error) {
	var p Pod
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		logrus.Warn("Error " + err.Error())
		return p, err
	}
	var pod corev1.Pod
	if err := client.Get(context.Background(), ns, name, &pod); err != nil {
		logrus.Warn("Error " + err.Error())
		return p, err
	}
	//Status
	s := *pod.Status.Phase
	sc := TrimQuotes(s)
	//Name
	n := *pod.Metadata.Name
	nc := TrimQuotes(n)
	// Namespace
	nsc := TrimQuotes(ns)
	//Worker
	w := *pod.Spec.NodeName
	wc := TrimQuotes(w)
	//IP
	ip := pod.Status.GetPodIP()
	//Image
	ci := pod.Spec.Containers[0]
	image := ci.GetImage()
	// Put in slice
	p = Pod{Status: sc, Name: nc, Namespace: nsc, Worker: wc, IP: ip, Image: image}
	return p, err
}

// DeletePod - describe a pod
func DeletePod(ns string, name string) error {
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		logrus.Warn("Error " + err.Error())
		return err
	}
	Pod := &corev1.Pod{
		Metadata: &metav1.ObjectMeta{
			Name:      k8s.String(name),
			Namespace: k8s.String(ns),
		},
	}
	err = client.Delete(context.Background(), Pod)
	return err
}
