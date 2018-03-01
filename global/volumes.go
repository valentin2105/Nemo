package global

import (
	"context"
	"fmt"
	"log"
	"strconv"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// PVC - kubectl get pvc
type PVC struct {
	Name      string
	Namespace string
	Size      string
	Status    string
	Scope     string
}

// PV - kubectl get pv
type PV struct {
	Name   string
	Size   string
	Status string
	Scope  string
}

// PVClist - list of pvc
type PVClist []PVC

// PVlist - list of pv
type PVlist []PV

// ListPVC - return pvc list
func ListPVC() PVClist {
	pvcl := make(PVClist, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	var PVCs corev1.PersistentVolumeClaimList
	if err := client.List(context.Background(), "", &PVCs); err != nil {
		log.Fatal(err)
	}
	for _, PVCs := range PVCs.Items {
		//Name
		n := fmt.Sprintf("%q", *PVCs.Metadata.Name)
		nc := TrimQuotes(n)
		// Namespace
		ns := fmt.Sprintf("%q", *PVCs.Metadata.Namespace)
		nsc := TrimQuotes(ns)
		// Status
		si := PVCs.Spec.Size()
		sic := strconv.Itoa(si)
		// Status
		s := fmt.Sprintf("%q", PVCs.Status.GetPhase())
		sc := TrimQuotes(s)
		// Put in slice
		p := PVC{Name: nc, Namespace: nsc, Status: sc, Size: sic}
		pvcl = append(pvcl, p)
	}
	return pvcl
}

// ListPV - return pv list
func ListPV() PVlist {
	pvl := make(PVlist, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	var PVs corev1.PersistentVolumeClaimList
	if err := client.List(context.Background(), "", &PVs); err != nil {
		log.Fatal(err)
	}
	for _, PVs := range PVs.Items {
		//Name
		n := fmt.Sprintf("%q", *PVs.Metadata.Name)
		nc := TrimQuotes(n)
		// Status
		si := PVs.Spec.Size()
		sic := strconv.Itoa(si)
		// Status
		s := fmt.Sprintf("%q", PVs.Status.GetPhase())
		sc := TrimQuotes(s)
		// Put in slice
		p := PV{Name: nc, Status: sc, Size: sic}
		pvl = append(pvl, p)
	}
	return pvl
}
