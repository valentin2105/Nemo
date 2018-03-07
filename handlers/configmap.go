package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetConfigmaps - Generate the Configmaps list view
func GetConfigmaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	configmaps := global.ListConfigmaps()
	type ClusterVars struct {
		Configmap global.ConfigmapList
	}
	ClusterDatas := ClusterVars{Configmap: configmaps}
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/configmaps.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET /configmaps from " + ip)
}
