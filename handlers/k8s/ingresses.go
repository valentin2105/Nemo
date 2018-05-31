package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global/k8s"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetIngresses - Generate the Ingress list view
func GetIngresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ingresses, err := global.ListIngresses()
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	type ClusterVars struct {
		Ingress global.IngressList
	}
	ClusterDatas := ClusterVars{Ingress: ingresses}
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/list/ingresses.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET /ingresses from " + ip)
}
