package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/markdown", mdHandler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	panic(http.ListenAndServe(":17901", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	src, err := Asset("static/lite.html")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, string(src))
	}
}

func mdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("open %s\n", os.Args[1])
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, string(contents))
	}
}
