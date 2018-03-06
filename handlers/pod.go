package handlers

import (
	"html/template"
	"io"
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
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/pods.tmpl.html")
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
	global.DeletePod(ns, name)
	ip := r.RemoteAddr
	io.WriteString(w, "Done")
	logrus.Infoln("DELETE /delete/" + ns + "/pod/" + name + " from " + ip)
}
