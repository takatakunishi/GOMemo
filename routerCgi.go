package main

import (
	"net/http"
	"net/http/cgi"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
)

type Country struct {
	Code string
	Name string
}

var store = map[string]*Country{}

var lock = sync.RWMutex{}

func GetCountry(w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")

	lock.RLock()

	var country *Country
	if store[code] != nil {
		country = &Country{}
		*country = *store[code]
	}
	lock.RUnlock()

	if country == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(country)
}

func GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	lock.RLock()
	countries := make([]Country, len(store))
	i := 0
	for _, country := range store {
		countries[i] = *country
		i++
	}
	lock.RUnlock()
	w.WriteJson(&countries)
}

func PostCountry(w rest.ResponseWriter, r *rest.Request) {
	country := Country{}

	err := r.DecodeJsonPayload(&country)
	//送られてきたデータをCountry型に落とし込む。
	//落とし込め得ない場合はたぶんエラー
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if country.Code == "" {
		rest.Error(w, "country code required", 400)
		return
	}
	if country.Name == "" {
		rest.Error(w, "country name required", 400)
		return
	}
	lock.Lock()
	store[country.Code] = &country
	lock.Unlock()
	w.WriteJson(&country)
}

func DeleteCountry(w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")
	lock.Lock()
	delete(store, code)
	lock.Unlock()
	w.WriteHeader(http.StatusOK)
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/cgi-bin/routerCgi.cgi/countries", GetAllCountries),
		rest.Post("/cgi-bin/routerCgi.cgi/countries", PostCountry),
		rest.Get("/cgi-bin/routerCgi.cgi/countries/:code", GetCountry),
		rest.Delete("/cgi-bin/routerCgi.cgi/countries/:code", DeleteCountry),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	api.SetApp(router)
	cgi.Serve(api.MakeHandler())
}
