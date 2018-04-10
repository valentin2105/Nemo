package global

import (
	"context"
	"strconv"

	"github.com/Sirupsen/logrus"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/valentin2105/Nemo/global"
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
		logrus.Warn("Error " + err.Error())
	}

	var PVCs corev1.PersistentVolumeClaimList
	if err := client.List(context.Background(), "", &PVCs); err != nil {
		logrus.Warn("Error " + err.Error())
	}
	for _, PVCs := range PVCs.Items {
		//Name
		n := *PVCs.Metadata.Name
		nc := global.TrimQuotes(n)
		// Namespace
		ns := *PVCs.Metadata.Namespace
		nsc := global.TrimQuotes(ns)
		// Status
		si := PVCs.Spec.Size()
		sic := strconv.Itoa(si)
		// Status
		s := PVCs.Status.GetPhase()
		sc := global.TrimQuotes(s)
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
		logrus.Warn("Error " + err.Error())
	}

	var PVs corev1.PersistentVolumeClaimList
	if err := client.List(context.Background(), "", &PVs); err != nil {
		logrus.Warn("Error " + err.Error())
	}
	for _, PVs := range PVs.Items {
		//Name
		n := *PVs.Metadata.Name
		nc := global.TrimQuotes(n)
		// Status
		si := PVs.Status.Size()
		sic := strconv.Itoa(si)
		// Status
		s := PVs.Status.GetPhase()
		sc := global.TrimQuotes(s)
		// Put in slice
		p := PV{Name: nc, Status: sc, Size: sic}
		pvl = append(pvl, p)
	}
	return pvl
}
