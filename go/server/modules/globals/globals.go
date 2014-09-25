package globals

import(
  "html/template"
  "regexp"
  "gopkg.in/mgo.v2"
)

type Config struct {
  Adapter struct {
    Server string
    Username string
    Password string
    Database string
    Collection string
  }
}

var Cfg Config

var Templates = template.Must(template.ParseFiles(
  "templates/edit.html",
  "templates/view.html",
  "templates/index.html"))

var ValidPath = regexp.MustCompile("^/(edit|save|view|delete)/([a-zA-Z0-9]+)$")

var Collection *mgo.Collection
