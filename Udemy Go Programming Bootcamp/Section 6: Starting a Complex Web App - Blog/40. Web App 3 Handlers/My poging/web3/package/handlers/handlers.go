package handlers

import (
	"net/http"
	"web3/package/render"
)

func HomeHandler(w http.ResponseWriter,
	request *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func AboutHandler(w http.ResponseWriter,
	request *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
