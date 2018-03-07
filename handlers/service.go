package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetServices - Generate the Services list view
func GetServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	services, err := global.ListServices()
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	type ClusterVars struct {
		Service global.ServiceList
	}
	ClusterDatas := ClusterVars{Service: services}
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/services.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET /services from " + ip)
}

// GetAnyService - Generate the Service describe view
func GetAnyService(w http.ResponseWriter, r *http.Request, ns string, name string) {
	w.Header().Set("Content-Type", "text/html")
	service, err := global.GetService(ns, name)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/get/service.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, service)
	ip := r.RemoteAddr
	logrus.Infoln("GET /get/" + ns + "/service/" + name + " from " + ip)

}
