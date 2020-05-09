package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"	
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.JsonpMiddleware{
		CallbackNameKey: "cb",
	})
	api.SetApp(rest.AppSimple(jsonHandler))
	log.Fatal(http.ListenAndServe(":3000", api.MakeHandler()))
}

func jsonHandler(w rest.ResponseWriter, r *rest.Request){
	w.WriteJson(map[string]string{"Body": "Hello World!"})
}
