package main

import (
    "net/http"
)


func main() {

    // serve the index page
    http.Handle("/", http.FileServer(http.Dir("./static")))

    // serve the contact page
    http.HandleFunc("/contact/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/contact.html")
    })

    // server the static css
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // start listening to incomming connections and serve the appropriate files
    panic(http.ListenAndServe(":80", nil))
}
