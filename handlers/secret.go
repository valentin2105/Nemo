package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetSecrets - Generate the Secrets list view
func GetSecrets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	secrets := global.ListSecrets()
	type ClusterVars struct {
		Secret global.SecretList
	}
	ClusterDatas := ClusterVars{Secret: secrets}
	tmpl, err := template.ParseFiles("templates/_navbar.html.tmpl", "templates/secrets.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET /secrets from " + ip)
}
