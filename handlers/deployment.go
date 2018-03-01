package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global"
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
	tmpl, err := template.ParseFiles("templates/_navbar.html.tmpl", "templates/deployments.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET /deployments from " + ip)
}
