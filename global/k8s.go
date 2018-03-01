package global

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ericchiang/k8s"
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
	appsv1beta1 "github.com/ericchiang/k8s/apis/apps/v1beta1"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/ghodss/yaml"
)

const Kubeconfig = "/Users/Valentin/.kube/config"

type ComponentStatus struct {
	Name   string
	Status string
}

type Node struct {
	Status      string
	Name        string
	Schedulable bool
}

type Pod struct {
	Status    string
	Name      string
	Namespace string
	Worker    string
	Ip        string
	Image     string
}

type Deployment struct {
	Status     string
	Name       string
	Namespace  string
	PodWanted  int32
	PodRunning int32
	Image      string
}

type ComponentStatusList []ComponentStatus
type NodeList []Node
type PodList []Pod
type DeploymentList []Deployment

func LoadClient(kubeconfigPath string) (*k8s.Client, error) {
	data, err := ioutil.ReadFile(kubeconfigPath)
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

func ListNodes() NodeList {
	nl := make(NodeList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	var nodes corev1.NodeList
	if err := client.List(context.Background(), "", &nodes); err != nil {
		log.Fatal(err)
	}
	for _, node := range nodes.Items {
		//Status
		status := node.Status.GetConditions()
		var st string
		for _, stat := range status {
			if *stat.Type == "Ready" || *stat.Type == "NotReady" {
				st = TrimQuotes(*stat.Type)
			}
		}
		//Name
		n := fmt.Sprintf("%q", *node.Metadata.Name)
		nc := TrimQuotes(n)
		// Schedule
		sch := !*node.Spec.Unschedulable
		// Put in slice
		no := Node{Status: st, Name: nc, Schedulable: sch}
		nl = append(nl, no)
	}
	return nl
}

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
		si := ChooseFaIcon(sc)
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
