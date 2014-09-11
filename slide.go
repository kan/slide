package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/markdown", mdHandler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	listen := make(chan bool)
	go func() {
		<-listen
		openUrl("http://localhost:17901")
	}()
	listen <- true
	panic(http.ListenAndServe(":17901", nil))
}

func openUrl(url string) {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	}

	fmt.Printf("open %s\n", url)
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
