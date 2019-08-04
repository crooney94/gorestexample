package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"gotutorials/app/model"

	"gotutorials/app/config"

	"gotutorials/app/controller"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) addRoutes() {
	a.Router.HandleFunc("/user", a.handleRequest(controller.GetAllUsers))

}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)
	fmt.Print(dbURI)
	db, err := gorm.Open(config.DB.Dialect, dbURI)

	if err != nil {
		log.Fatal(err)
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.addRoutes()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
