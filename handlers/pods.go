package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetPods - Generate the Pods list view
func GetPods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	pods := global.ListPods()
	type ClusterVars struct {
		Pod global.PodList
	}
	ClusterDatas := ClusterVars{Pod: pods}
	tmpl, err := template.ParseFiles("templates/_navbar.html.tmpl", "templates/pods.html.tmpl")
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
	pod := global.GetPod(ns, name)
	tmpl, err := template.ParseFiles("templates/_navbar.html.tmpl", "templates/get/pod.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, pod)
	ip := r.RemoteAddr
	logrus.Infoln("GET /get/" + ns + "/" + name + " from " + ip)

}
