package main

import (
	"log"
	"net/http"

	"html/template"

	"webapp.com/m/backfin"
)


var tpl * template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	portf := backfin.RnPCSV()
	sectors := backfin.SectorVals(portf)

	data := struct {
		Portfolio []backfin.HoldingInfo
		Sectors map[string]float64
	} {
		Portfolio: portf,
		Sectors: sectors,
	}

	// TODO Adjust grouping of companies such that they are grouped by sector

	err := tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w , err.Error(), 500)
		log.Fatal()
	}	
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	t := r.FormValue("searchq")
	prof := backfin.InitCompanyProfile(t)
	tpl.ExecuteTemplate(w, "search.html", prof)
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)
	http.ListenAndServe(":3000", nil)
}

