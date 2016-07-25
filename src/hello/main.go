package main

import (
	"echo"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {
	fmt.Println(":8080 'hello world'")

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": echo.Echo("hello world")})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
