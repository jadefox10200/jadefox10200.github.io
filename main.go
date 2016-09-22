package main

import (
	"log"
	"net/http"
)

//THIS IS FOR USE AS A LOCAL TEMP SERVER FOR PAGE TESTING, NOT FOR PRODUCTION!

func main() {

	//Handle route for all assets of local site
	http.Handle("/assets/",
		http.StripPrefix("/assets",
			http.FileServer(http.Dir("assets/"))))

	//Handle route for index.html
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.Error(res, "Page not found", 404)
		} else {
			http.ServeFile(res, req, "index.html")
		}

	})

	//Start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ERROR Starting server: ", err)
	}
}
