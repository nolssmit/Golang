package handlers

import (
	"net/http"
	"web3/models"
	"web3/package/config"
	"web3/package/render"
)

type Repository struct{
	App *config.AppConfig	
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository {
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter,
	request *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl",
	&models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter,
	request *http.Request) {

	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. " +
					  "We love talking about ourselves."

	render.RenderTemplate(w, "about.page.tmpl",
			&models.PageData{StrMap: strMap})  // StrMap is defied in pagedata.go
}
