package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetHome - Generate the home view
func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	type ClusterVars struct {
		Node            global.NodeList
		ComponentStatus global.ComponentStatusList
	}
	components := global.ListComponentStatus()
	nodes := global.ListNodes()
	ClusterDatas := ClusterVars{ComponentStatus: components, Node: nodes}
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/home.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET / from " + ip)
}
