package main

import (
	"echo"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"encoding/json"
)

func echoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}
}

type Link struct {
	Url string
}

type Config struct {
	Links []Link
}

var in = `{
    "links": [
        {
            "Url": "127.0.0.1:17001"
        },
        {
            "Url": "127.0.0.1:17002"
        }
    ]
}`


func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("view")
	var conf Config
	err := json.Unmarshal([]byte(in), &conf)

	if err != nil {
		fmt.Print("Error:", err)
	}

	fmt.Printf("%#v\n", conf)
	fmt.Printf("%s\n",conf.Links[0].Url)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func main() {
	fmt.Println(":8080/api/.status")
	fmt.Println(":8080/api/hello 'hello world'")
	fmt.Println(":8080/echo/hello/john ...")
	fmt.Println(":8080/ file")

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
	http.Handle("/echo/", http.StripPrefix("/echo", echoHandler()))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/view", makeHandler(viewHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
