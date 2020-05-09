package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"	
)

func main() {
	api := rest.NewApi()
	statusMw := &rest.StatusMiddleware{}
	api.Use(statusMw)
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/countries", GetAllCountries),
		rest.Get("/.status", func(w rest.ResponseWriter, r *rest.Request){
	w.WriteJson(statusMw.GetStatus())
}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":3001",api.MakeHandler()))
}

type Country struct {
	Code string 
	Name string
}

func GetAllCountries(w rest.ResponseWriter, r *rest.Request){
	w.WriteJson(
		[]Country{
			Country{
				Code: "FR",
				Name: "France",
			},
			Country{
				Code: "US",
				Name: "United States",
			},
		},
	)
}
