package handlers

import (
	"net/http"
	"web3/models"
	"web3/package/render"
)

func HomeHandler(w http.ResponseWriter,
	request *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl",
	&models.PageData{})
}

func AboutHandler(w http.ResponseWriter,
	request *http.Request) {

	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. " +
					  "We love talking about ourselves."

	render.RenderTemplate(w, "about.page.tmpl",
			&models.PageData{StrMap: strMap})  // StrMap is defied in pagedata.go
}
