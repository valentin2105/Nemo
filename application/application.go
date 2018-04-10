package application

import (
	"net/http"

	"github.com/carbocation/interpose"
	gorilla_mux "github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"

	handlers "github.com/valentin2105/Nemo/handlers"
	k8s "github.com/valentin2105/Nemo/handlers/k8s"
	"github.com/valentin2105/Nemo/middlewares"
)

// New is the constructor for Application struct.
func New(config *viper.Viper) (*Application, error) {
	cookieStoreSecret := config.Get("cookie_secret").(string)

	app := &Application{}
	app.config = config
	app.sessionStore = sessions.NewCookieStore([]byte(cookieStoreSecret))
	return app, nil
}

// Application is the application object that runs HTTP server.
type Application struct {
	config       *viper.Viper
	sessionStore sessions.Store
}

// MiddlewareStruct - Call middlewares
func (app *Application) MiddlewareStruct() (*interpose.Middleware, error) {
	middle := interpose.New()
	middle.Use(middlewares.SetSessionStore(app.sessionStore))

	middle.UseHandler(app.mux())

	return middle, nil
}

func (app *Application) mux() *gorilla_mux.Router {
	router := gorilla_mux.NewRouter()

	// All other Routes -> 404
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	// HTTP Routes
	router.Handle("/", http.HandlerFunc(handlers.GetHome)).Methods("GET")
	// Create
	router.Handle("/create", http.HandlerFunc(handlers.Create)).Methods("GET")
	// List
	router.Handle("/pods", http.HandlerFunc(k8s.GetPods)).Methods("GET")
	router.Handle("/deployments", http.HandlerFunc(k8s.GetDeployments)).Methods("GET")
	router.Handle("/services", http.HandlerFunc(k8s.GetServices)).Methods("GET")
	router.Handle("/volumes", http.HandlerFunc(k8s.GetVolumes)).Methods("GET")
	router.Handle("/configmaps", http.HandlerFunc(k8s.GetConfigmaps)).Methods("GET")
	router.Handle("/secrets", http.HandlerFunc(k8s.GetSecrets)).Methods("GET")
	// Get
	router.HandleFunc("/get/{namespace}/pod/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := gorilla_mux.Vars(r)
		ns := vars["namespace"]
		name := vars["name"]
		k8s.GetAnyPod(w, r, ns, name)
	}).Methods("GET")
	router.HandleFunc("/get/{namespace}/deployment/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := gorilla_mux.Vars(r)
		ns := vars["namespace"]
		name := vars["name"]
		k8s.GetAnyDeployment(w, r, ns, name)
	}).Methods("GET")
	router.HandleFunc("/get/{namespace}/service/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := gorilla_mux.Vars(r)
		ns := vars["namespace"]
		name := vars["name"]
		k8s.GetAnyService(w, r, ns, name)
	}).Methods("GET")
	router.HandleFunc("/get/node/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := gorilla_mux.Vars(r)
		name := vars["name"]
		k8s.GetAnyNode(w, r, name)
	}).Methods("GET")
	// Delete
	router.HandleFunc("/delete/{namespace}/pod/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := gorilla_mux.Vars(r)
		name := vars["name"]
		ns := vars["namespace"]
		k8s.DeleteAnyPod(w, r, ns, name)
	}).Methods("DELETE")
	router.HandleFunc("/delete/{namespace}/deployment/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := gorilla_mux.Vars(r)
		name := vars["name"]
		ns := vars["namespace"]
		k8s.DeleteAnyDeployment(w, r, ns, name)
	}).Methods("DELETE")

	// Path of static files must be last!
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	return router
}
