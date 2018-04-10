package handlers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global/k8s"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetDeployments - Generate the deployment list view
func GetDeployments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	deployment := global.ListDeployments()
	type ClusterVars struct {
		Deployment global.DeploymentList
	}
	ClusterDatas := ClusterVars{Deployment: deployment}
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/list/deployments.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET /deployments from " + ip)
}

// GetAnyDeployment - Generate the Deployment describe view
func GetAnyDeployment(w http.ResponseWriter, r *http.Request, ns string, name string) {
	w.Header().Set("Content-Type", "text/html")
	deployment := global.GetDeployment(ns, name)
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/get/deployment.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, deployment)
	ip := r.RemoteAddr
	logrus.Infoln("GET /get/" + ns + "/deployment/" + name + " from " + ip)

}

// DeleteAnyDeployment - Generate the Deployment describe view
func DeleteAnyDeployment(w http.ResponseWriter, r *http.Request, ns string, name string) {
	w.Header().Set("Content-Type", "text/html")
	err := global.DeleteDeployment(ns, name)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	ip := r.RemoteAddr
	io.WriteString(w, "Done")
	logrus.Infoln("DELETE /delete/" + ns + "/deployment/" + name + " from " + ip)
}
