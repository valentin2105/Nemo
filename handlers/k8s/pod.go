package handlers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global/k8s"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetPods - Generate the Pods list view
func GetPods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	pods, err := global.ListPods()
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	type ClusterVars struct {
		Pod global.PodList
	}
	ClusterDatas := ClusterVars{Pod: pods}
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/list/pods.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET /pods from " + ip)
}

// GetAnyPod - Generate the Pod describe view
func GetAnyPod(w http.ResponseWriter, r *http.Request, ns string, name string) {
	w.Header().Set("Content-Type", "text/html")
	pod, err := global.GetPod(ns, name)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/get/pod.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, pod)
	ip := r.RemoteAddr
	logrus.Infoln("GET /get/" + ns + "/pod/" + name + " from " + ip)
}

// DeleteAnyPod - Generate the Pod describe view
func DeleteAnyPod(w http.ResponseWriter, r *http.Request, ns string, name string) {
	w.Header().Set("Content-Type", "text/html")
	err := global.DeletePod(ns, name)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	ip := r.RemoteAddr
	io.WriteString(w, "Done")
	logrus.Infoln("DELETE /delete/" + ns + "/pod/" + name + " from " + ip)
}
