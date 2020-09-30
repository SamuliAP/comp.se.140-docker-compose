package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func hello(w http.ResponseWriter, req *http.Request) {

	res, err := http.Get("http://server")
	
	if(err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not connect")
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)


	if(err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Invalid server response")
		return
	}

	fmt.Fprintf(w, "Hello from " + req.RemoteAddr + "\nto " +  req.Host + "\n" + string(body))
}

func main() {

	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
