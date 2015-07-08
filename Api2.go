package hello

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"io/ioutil"
//	"log"
	"net/http"
	"strconv"
)

// Structure Response -> Valeur(s) renvoyé a la requete GET

type Response struct {
	Hello string
}

// Structure Person -> Valeur(s) envoyé depuis le client (POST & PUT)

type Person struct {
	Id             int
	Name           string
	Age            int
	ServerResponse bool
}

type Vehicule struct {
	Title       string
	Model       string
	Price       int
	Description string
	Type        string
	Capacity    int
	Owner       string
}

// Function ReqObject -> Gere toutes les requetes liées aux objects (Voiture, Imprimante 3D, Cours)

var vehicule Vehicule

func ReqObject(w http.ResponseWriter, req *http.Request) {
	//	vars := mux.Vars(req) // récupère les valeurs d'url
	// vars["id"] contient l'id utilisateur...
	if req.Method == "GET" {
		//Appelle l'api pour recueper le /object/{id}
		b, err := json.Marshal(vehicule)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		io.WriteString(w, string(b))
	}
	if req.Method == "POST" {
		//Appelle l'api pour creer un object
		req.ParseForm()
		vehicule.Title = req.Form["title"][0]
		vehicule.Model = req.Form["model"][0]
		vehicule.Price, _ = strconv.Atoi(req.Form["price"][0])
		vehicule.Description = req.Form["desc"][0]
		vehicule.Type = req.Form["type"][0]
		vehicule.Capacity, _ = strconv.Atoi(req.Form["cap"][0])
		vehicule.Owner = req.Form["own"][0]
		http.Redirect(w, req, "/", http.StatusFound)
		//	body, _ := ioutil.ReadAll(req.Body)
		//	json.Unmarshal(body, &vehicule)

		//	responseXML, _ := json.Marshal(vehicule)
		//	fmt.Print(string(responseXML))
	}
	if req.Method == "PUT" {
		//Appelle l'api pour modifier le /object/{id}
		var vehicule Vehicule
		body, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &vehicule)

		responseXML, _ := json.Marshal(vehicule)
		fmt.Print(string(responseXML))
	}
	if req.Method == "DELETE" {
		//Appelle l'api pour del le /object/{id}
	}
}

// Function ReqProfile -> Gere toutes les requetes liées aux utilisateurs

func ReqProfil(w http.ResponseWriter, req *http.Request) {
	var response Response
	vars := mux.Vars(req) // récupère les valeurs d'url
	// vars["id"] contient l'id utilisateur...
	fmt.Println("On demande le profil %s", vars["id"])
	if req.Method == "GET" {
		//Appelle l'api pour recueper le /profile/{id}
		response.Hello = "Hello world"
		b, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		io.WriteString(w, string(b))
		io.WriteString(w, "\n")
		io.WriteString(w, req.Method)
	}
	if req.Method == "POST" {
		fmt.Println("Post")
		//Appelle l'api pour creer un profile
		var person Person
		body, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &person)
		person.ServerResponse = true

		responseXML, _ := json.Marshal(person)
		fmt.Print(string(responseXML))
	}
	if req.Method == "PUT" {
		fmt.Println("Put")
		//Appelle l'api pour modifier le /profile/{id}
		var person Person
		body, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &person)
		person.ServerResponse = true

		responseXML, _ := json.Marshal(person)
		fmt.Print(string(responseXML))
	}
	if req.Method == "DELETE" {
		//Appelle l'api pour del le /profile/{id}
	}
}

func deleteVehicule(w http.ResponseWriter, req *http.Request) {
	vehicule = Vehicule{}
	http.Redirect(w, req, "/", http.StatusFound)
}

func index(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

// Coeur du progamme

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/object/{id:[0-9]*}", ReqObject).Methods("GET")
	r.HandleFunc("/profil/{id:[0-9]*}", ReqProfil).Methods("GET")
	r.HandleFunc("/delete/", deleteVehicule)
	r.HandleFunc("/profil/{id:[0-9]*}", ReqProfil).Methods("DELETE")
	r.HandleFunc("/object/", ReqObject)
	r.HandleFunc("/profil/", ReqProfil)
	http.Handle("/", r)
/*	err := http.ListenAndServe(":9000", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
}
