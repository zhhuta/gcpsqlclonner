package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	//"google.golang.org/api/sqladmin/v1beta4"
)

type Resourse struct {
	Project     string `json:"project"`
	SQLInstance string `json:"sqlInstance"`
}

func init() {

	if _, ok := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS"); ok {
		fmt.Println("Credentials is set")
	} else {
		log.Fatal("GOOGLE_APPLICATION_CREDENTIALS is Not set")
	}

	if _, ok := os.LookupEnv("GOOGLE_PARENT"); ok {
		fmt.Println("Google Parent is set")
	} else {
		log.Fatal("Google Parent system variable GOOGLE_PARENT is NOt set")
	}
	if _, err := os.Stat("assets.json"); errors.Is(err, os.ErrNotExist) {
		writeData2File("assets.json", listSQLAssets(os.Getenv("GOOGLE_PARENT")))
	}

}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/csqlall", httpGetSQLAll).Methods("GET")
	r.HandleFunc("/api/v1/csql/{project}", httpGetProjectsSQLs).Methods("GET")
	r.HandleFunc("/api/v1/csql/{project}/{instance}/clone/{arbitrary_name}", httpCloneSQLInstance).Methods("POST")
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("End of execution")
}

func httpGetSQLAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	lapstime := 5 * time.Minute
	now := time.Now()
	diff := now.Sub(checkFileTimeStemp("assets.json"))
	var list []Resourse
	if diff > lapstime {
		list = listSQLAssets(os.Getenv("GOOGLE_PARENT"))
		writeData2File("assets.json", list)
	} else {
		list = readData4File("assets.json")

	}
	json.NewEncoder(w).Encode(list)
	//list := listSQLAssets(os.Getenv("GOOGLE_PARENT"))
}
func httpGetProjectsSQLs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sqlInstance, err := ListSQLInstances(params["project"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(sqlInstance)

}
func httpCloneSQLInstance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	clone, err := CloneSQLInstance(params["project"], params["instance"], params["arbitrary_name"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(clone)

}
