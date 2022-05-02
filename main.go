package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"google.golang.org/api/sqladmin/v1beta4"
)

type Resourse struct {
	Project     string `json:"project"`
	SQLInstance string `json:"sqlInstance"`
}

func main() {
	fmt.Println("Starting .....")
	listSQLAssets("projects/32517729315")
	fmt.Println("End of execution")

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/csqlall", httpGetSQLAll).Methods("GET")
	r.HandleFunc("/api/v1/csql/{project}", httpGetProjectsSQLs).Methods("GET")
	r.HandleFunc("/api/v1/csql/{project}/{instance}/clone", httpCloneSQLInstance).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("End of execution")
}

func httpGetSQLAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := listSQLAssets("projects/32517729315")
	json.NewEncoder(w).Encode(list)
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
	clone, err := CloneSQLInstance(params["project"], params["instance"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(clone)

}

/* example of handled request with parameters

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/device/{clusterID}/job", getJob)
    http.ListenAndServe(":8000", r)
}

func getJob(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fmt.Println(params["clusterID"])
    fmt.Println(params)
} */

/* sqlInstance, err := ListSQLInstances("or2-msq-epm-gcp-t1iylu")
if err != nil {
	log.Fatal(err)
}

for _, sqlInstance := range sqlInstance {
	fmt.Println(sqlInstance.SelfLink)
} */

/* clone, err := cloneSQLInstance("or2-msq-epm-gcp-t1iylu", "zhhuta")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("%#v\n", clone.TargetLink)
*/
