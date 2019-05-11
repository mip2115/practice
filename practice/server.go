package main

import (
  "net/http"
  "log"
  "html/template"
  "fmt"
)

func handlebundle(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hit")
  w.Header().Set("Content-Type", "application/json")

}

func handlepost (w http.ResponseWriter, r *http.Request){
  
  fmt.Println("Hit the post")
  q := r.FormValue("query")
  fmt.Println("VALUE IS: ", q)

}

func handler(w http.ResponseWriter, r *http.Request) {

   //w.Header().Set("Content-Type", "application/json")
 // w.Header().Set("Content-Type", "text/html")
  t, _ := template.ParseFiles("./index.html")
  t.Execute(w, "")
}

func main() {
  // Set the router as the default one shipped with Gin

  // http.HandleFunc("/public/", handlebundle)
  http.HandleFunc("/api/postform", handlepost)
  http.HandleFunc("/", handler)

  fmt.Println("hello this is server")
  
  fs := http.FileServer(http.Dir("./public/"))
  
  http.Handle("/public/", http.StripPrefix("/public", fs)) 
 

  log.Fatal(http.ListenAndServe(":8080", nil))
}

// should be able to retrieve the value netered into the form and print it out.




