package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/libhttp"
)

// Create - Generate the Creation view
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles("templates/_head.tmpl.html", "templates/create.tmpl.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, "")
	ip := r.RemoteAddr
	logrus.Infoln("GET /create from " + ip)
}
