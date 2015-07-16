package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Structure de mise en place dans le datastore
type Data struct {
	Date     time.Time
	Json     string
	Hidden   bool
	Type		 string
	Cap 		 int
	Price 	 int
}

type Object struct {
	Index int64
	Title string
	Model string
	Type  string
	Desc  string
	Cap   int
	Price int
	Own   string
}

var object Object

// Function ReqObject -> Gere toutes les requetes liées aux objects (Voiture, Imprimante 3D, Cours)
func getObject(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req) // récupère les valeurs d'url
	// vars["id"] contient l'id utilisateur...
	if req.Method == "GET" {
		var obj []Object
		var tmp Object

		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		lilist, kikeys, err := getData(req, id)
		if lilist == nil {
			io.WriteString(w, err)
		} else {
			for i, v := range lilist {
				if (lilist[i].Hidden == false) {
					json.Unmarshal([]byte(v.Json), &tmp)
					tmp.Index = kikeys[i].IntID()
					obj = append(obj, tmp)
				}
			}
			datas, _ := json.Marshal(obj)
			io.WriteString(w, string(datas))
		}
	}
}

func modObject(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &object)
	d := Data{
		Date:     time.Now(),
		Json:     string(body),
		Type:			object.Type,
		Cap:			object.Cap,
		Price:		object.Price,
	}
	if req.Method == "POST" {
		//Appelle l'api pour creer un object
		addData(d, req)
	}
	if req.Method == "PUT" {
		//Appelle l'api pour modifier le /object/{id}
		object.Index, _ = strconv.ParseInt(vars["id"], 10, 64)
		putData(d, req, object.Index)
	}
}

func delObject(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	lilist, _, _ := getData(req, id)
	lilist[0].Hidden = true;
	putData(lilist[0], req, id)
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/api/objects/", getObject).Methods("GET")
	r.HandleFunc("/api/objects/{id:[0-9]*}", getObject).Methods("GET")
	r.HandleFunc("/api/objects/", modObject).Methods("POST")
	r.HandleFunc("/api/objects/{id:[0-9]*}", modObject).Methods("PUT")
	r.HandleFunc("/api/objects/{id:[0-9]*}", delObject).Methods("DELETE")
	http.Handle("/", r)
}
