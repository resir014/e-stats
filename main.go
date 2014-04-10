package main

import (
        "net/http"
        "github.com/go-martini/martini"
        "html/template"
)

const dsLoc = "./prod.db"

func main() {
        startServer()
}

func startServer() {
        m := martini.Classic()
	indexHandler := createIndexHandler()

        m.Get("/", indexHandler)

        m.Run()
}
func createIndexHandler() func(http.ResponseWriter){
	renderer := createRenderer()

        dataStore, err := createDataStore(dsLoc)
        if err != nil {
                panic(err)
        }

        return func(w http.ResponseWriter) {
                data, err := dataStore.GetLatestRecords()
                if err != nil {
                        http.Error(w, "Could not fetch records", 500)
                }
		renderer(w, data)
        }
}

func createRenderer() func(http.ResponseWriter, []Record){
        templ, err := template.ParseGlob("templates/*.tpl")
        if err != nil {
                panic(err)
        }

        return func(w http.ResponseWriter, rec []Record) {
                templ.ExecuteTemplate(w,"indexPage",  rec)
        }
}
