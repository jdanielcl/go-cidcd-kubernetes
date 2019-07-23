package main

import (
  "net/http"
  "log"
  "html/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" { return }

  index, err := template.ParseFiles("./index.gohtml")
  if err != nil {
    panic(err)
  }

  if err := index.Execute(w, nil); err != nil {
    panic(err)
  }
  //fmt.Fprintln(w, "new home")
}

func main() {
  http.HandleFunc("/", homeHandler)
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
  http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
  http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("vendor"))))
  http.Handle("/scss/", http.StripPrefix("/scss/", http.FileServer(http.Dir("scss"))))
  log.Fatal(http.ListenAndServe(":8181", nil))
}
