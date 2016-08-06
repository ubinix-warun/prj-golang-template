package main

import (
	"echo"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {
	fmt.Println(":8080/api/.status")
	fmt.Println(":8080/api/hello 'hello world'")

	api := rest.NewApi()
	statusMw := &rest.StatusMiddleware{}
	api.Use(statusMw)
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/.status", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(statusMw.GetStatus())
		}),
		rest.Get("/hello", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": echo.Echo("hello world")})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}


	api.SetApp(router)
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
