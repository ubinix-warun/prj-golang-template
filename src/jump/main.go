package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"golang.org/x/net/html"
	"strconv"
	"io/ioutil"
	"os"
)

type Link struct {
	Url string
}

type Config struct {
	Links []Link
}

/*
var in = `{
    "links": [
        {
            "Url": "http://sanook.com"
        },
        {
            "Url": "http://kapook.com"
        }
    ]
}`
*/

var conf Config

func pageHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Query().Get("id"))
	idx,_ := strconv.ParseInt(r.URL.Query().Get("id"),10,0)
	if int(idx) < len(conf.Links) {
		//fmt.Println(conf.Links[idx])

		// READ title

		//fmt.Fprintf(w, "<title>File's Not Found</title>")
	} else {
		fmt.Fprintf(w, "<h1>File's Not Found</h1>")
	}

}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	for idx,element := range conf.Links {
		// element is the element from someSlice for where we are
		//fmt.Fprintf(w, "<h1>%s</h1>", element.Url)
		resp, _ := http.Get(element.Url)
		//b, _ := ioutil.ReadAll(resp.Body)

		//fmt.Println("HTML:\n\n", string(bytes))
		//

		//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
		// Parse HTML for Anchor Tags //
		//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//

		z := html.NewTokenizer(resp.Body)
		flag:= 0

		for {
			tt := z.Next()
			if tt == html.ErrorToken {
				break
			}

			switch {
			//case tt == html.ErrorToken:
				// End of the document, we're done
				//break
			//break;
			case tt == html.StartTagToken:
				t := z.Token()

				flag = 0
				isAnchor := t.Data == "title"
				if isAnchor {
					flag = 1
					//fmt.Println("We found a link!")
				}
			case tt == html.TextToken:
				t:= z.Token()
				if flag == 1 {
					fmt.Fprintf(w, "<h1><a href=\"/page?id=%d\">%s</a></h1>", idx, t)
					//fmt.Println(t)
				}
			}

		}

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

	fmt.Println(":8080/ file")
	fmt.Println(":8080/view")
	fmt.Println(":8080/page?id=x")

	raw, err := ioutil.ReadFile("./pages.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(raw, &conf)

	if err != nil {
		fmt.Print("Error:", err)
	}
	fmt.Println(":8080/x config.ed")


	//api.SetApp(router)
	//http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	//http.Handle("/echo/", http.StripPrefix("/echo", echoHandler()))

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/view", makeHandler(viewHandler))
	http.HandleFunc("/page", makeHandler(pageHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
