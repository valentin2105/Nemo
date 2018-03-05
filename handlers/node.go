package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetAnyNode - Generate the Node describe view
func GetAnyNode(w http.ResponseWriter, r *http.Request, name string) {
	w.Header().Set("Content-Type", "text/html")
	node := global.GetNode(name)
	tmpl, err := template.ParseFiles("templates/_head.html.tmpl", "templates/get/node.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, node)
	ip := r.RemoteAddr
	logrus.Infoln("GET /get/node/" + name + " from " + ip)

}
