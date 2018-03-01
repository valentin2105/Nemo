package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/valentin2105/Nemo/global"
	"github.com/valentin2105/Nemo/libhttp"
)

// GetVolumes - Generate the Volumes list view
func GetVolumes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	pvcs := global.ListPVC()
	pvs := global.ListPV()
	type ClusterVars struct {
		PVC global.PVClist
		PV  global.PVlist
	}
	ClusterDatas := ClusterVars{PVC: pvcs, PV: pvs}
	tmpl, err := template.ParseFiles("templates/_navbar.html.tmpl", "templates/volumes.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	tmpl.Execute(w, ClusterDatas)
	ip := r.RemoteAddr
	logrus.Infoln("GET /volumes from " + ip)
}
