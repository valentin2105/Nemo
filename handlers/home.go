package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global"
	"github.com/valentin2105/Nemo/global/k8s"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetHome - Generate the home view
func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	type ClusterVars struct {
		Node            k8s.NodeList
		ComponentStatus k8s.ComponentStatusList
	}
	components := k8s.ListComponentStatus()
	nodes, err := k8s.ListNodes()
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
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

// NotFound - Generate the home view
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/404.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	var data string
	tmpl.Execute(w, data)
	ip := r.RemoteAddr
	path := r.RequestURI
	logrus.Infoln("GET " + path + " from " + ip)
}
