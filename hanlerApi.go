package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main (){
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/message", func(w rest.ResponseWriter, r *rest.Request){
			w.WriteJson(map[string]string{"Body":"Hello World!"})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))
	// http.Dir型 は FileSystemインターフェースを満たすので
	// これを http.FileServer() に渡して 静的ファイルを返すハンドラを得る
	// fileHandler := http.FileServer(http.Dir(".")) で変数化できる
	//ファイルの読み出し
	// -> ブラウザで http://localhost:3001/file/index.html にアクセスすると
	//    静的ファイル static/index.html が得られる
	// StripPrefixについては => https://tech-up.hatenablog.com/entry/2018/12/28/120517

	log.Fatal(http.ListenAndServe(":3001", nil))
	// サーバーを起動する
    // nilを指定すると DefaultServeMux が使われる
}