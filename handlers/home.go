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
	nodes := global.ListNodes()
	components := global.ListComponentStatus()
	type ClusterVars struct {
		Node            global.NodeList
		ComponentStatus global.ComponentStatusList
	}
	ClusterDatas := ClusterVars{ComponentStatus: components, Node: nodes}
	tmpl, err := template.ParseFiles("templates/_navbar.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET / from " + ip)
}
