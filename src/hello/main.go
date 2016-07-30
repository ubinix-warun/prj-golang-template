package main

import (
	"echo"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"encoding/json"
	//"golang.org/x/net/html"
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
            "Url": "http://www.google.com"
        },
        {
            "Url": "http://kapook.com"
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

	for _,element := range conf.Links {
		// element is the element from someSlice for where we are
		//fmt.Fprintf(w, "<h1>%s</h1>", element.Url)
		resp, _ := http.Get(element.Url)
		//b, _ := ioutil.ReadAll(resp.Body)

		//fmt.Println("HTML:\n\n", string(bytes))
		//

		//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
		// Parse HTML for Anchor Tags //
		//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//

		//z := html.NewTokenizer(resp.Body)
		//
		//for {
		//	tt := z.Next()
		//
		//	switch {
		//	case tt == html.ErrorToken:
		//		// End of the document, we're done
		//		return
		//	//break;
		//	case tt == html.StartTagToken:
		//		t := z.Token()
		//
		//		isAnchor := t.Data == "title"
		//		if isAnchor {
		//			fmt.Println("We found a link!")
		//		}
		//	}
		//}

		resp.Body.Close()

	}

	//fmt.Printf("%#v\n", conf)
	//fmt.Printf("%s\n",conf.Links[0].Url)
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
