package main

import (
	 "github.com/madetomeasure/api/endpoints/subscribers"
        //"github.com/madetomeasure/api/endpoints/lists"
        //"github.com/madetomeasure/api/endpoints/campaigns"
  "fmt"
  "path/filepath"
        "github.com/jinzhu/gorm"
         _ "github.com/lib/pq"
        "log"
	"github.com/gorilla/mux"
	"net/http"
        "flag"
"github.com/kylelemons/go-gypsy/yaml"
)

var Env *string
var Port *int

func main() {

      Env = flag.String("env", "development", "Defines the working environment for the API")
      Port = flag.Int("port", 8080, "What port to run the API on")
      flag.Parse()


        startLogger()
        establishDB()
	http.ListenAndServe(fmt.Sprintf(":%d",*Port), startRouter())
}

// This will try to establish a connection with the database
// It also tries a lot of things...
func establishDB() {
  cfgFile, err := filepath.Abs("db/dbconf.yml")

  giveUp(err, "db/dbconf.yml is missing")

  f, err := yaml.ReadFile(cfgFile)
  
  giveUp(err, "Problem parsing YAML, look at syntax of db/dbconf.yml")

  drv, err := f.Get(fmt.Sprintf("%s.driver", *Env))

  giveUp(err, "Could not detect driver properly, make sure driver is defined")

  open, err := f.Get(fmt.Sprintf("%s.open", *Env))

  giveUp(err, "Could not detect open string for database connection, check the dbconf.yml")

  db, err := gorm.Open(drv, open)

  db.DB()
  db.DB().Ping()
  db.DB().SetMaxIdleConns(10)
  db.DB().SetMaxOpenConns(100)


  giveUp(err, "Could not establish connection to db")
}

func startRouter() http.Handler {
	r := mux.NewRouter()
        // Subscriber Endpoints
	r.HandleFunc("/subscribers", subscriber.IndexHandler).Methods("GET")
	r.HandleFunc("/subscribers", subscriber.CreateHandler).Methods("POST")
        r.HandleFunc("/subscribers/{subscriberId}", subscriber.UpdateHandler).Methods("POST", "PATCH")
        r.HandleFunc("/subscribers/{subscriberId}", subscriber.DestroyHandler).Methods("DELETE")

        // Human readable spec
        r.Handle("/spec", http.FileServer(http.Dir("./public/")))
        return r
}

func startLogger() {
        log.Println("=> Booting...")
	log.Println(fmt.Sprintf("=> Made to Measure Alpha application starting in %s on http://0.0.0.0:%d", *Env, *Port))
	log.Println("=> Ctrl-C to shutdown Server")
}

func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func giveUp(err interface{}, message string) {
  if err != nil {
    log.Fatal(message, err)
  }
}

