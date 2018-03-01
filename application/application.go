package application

import (
	"net/http"

	"github.com/carbocation/interpose"
	gorilla_mux "github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"

	"github.com/valentin2105/Nemo/handlers"
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

	// HTTP Routes
	router.Handle("/", http.HandlerFunc(handlers.GetHome)).Methods("GET")
	router.Handle("/pods", http.HandlerFunc(handlers.GetPods)).Methods("GET")
	router.Handle("/deployments", http.HandlerFunc(handlers.GetDeployments)).Methods("GET")
	router.Handle("/volumes", http.HandlerFunc(handlers.GetVolumes)).Methods("GET")
	router.Handle("/configmaps", http.HandlerFunc(handlers.GetConfigmaps)).Methods("GET")
	router.Handle("/secrets", http.HandlerFunc(handlers.GetSecrets)).Methods("GET")
	router.HandleFunc("/get/{namespace}/pod/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := gorilla_mux.Vars(r)
		ns := vars["namespace"]
		name := vars["name"]
		handlers.GetAnyPod(w, r, ns, name)
	}).Methods("GET")

	// Path of static files must be last!
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	return router
}
