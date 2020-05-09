package main 

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net"
    "net/http"
)

// func jsonHandler(w rest.ResponseWriter, r *rest.Request){
// 	w.WriteJson(map[string]string{"Body": "Hello World!"})
// }  @@@ 1

func GetRoute (w rest.ResponseWriter, req *rest.Request){
		ip, err := net.LookupIP(req.PathParam("host"))
		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteJson(&ip)
	}

func main () {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	

	router, err := rest.MakeRouter(
		rest.Get("/lookup/#host", GetRoute),
	)
	if err != nil {
		log.Fatal(err)
	}
	// api.SetApp(rest.AppSimple(jsonHandler))  @@@ 1
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":3001", api.MakeHandler()))
}