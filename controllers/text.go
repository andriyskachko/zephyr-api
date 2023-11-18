package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/andriyskachko/zephyr-api/repositories"
	"github.com/andriyskachko/zephyr-api/services"
)

type TextController struct {
	TextService services.TextService
}

type textsGetResponse struct {
	Count int                 `json:"count"`
	Texts []repositories.Text `json:"texts"`
}

type textCreateResponse struct {
	Text       *repositories.Text `json:"text"`
	Successful bool               `json:"successful"`
}

func NewTextController(textService services.TextService) *TextController {
	return &TextController{
		TextService: textService,
	}
}

func (c *TextController) TextGETHandler(w http.ResponseWriter, r *http.Request) {
	texts, err := c.TextService.GetTexts()
	if err != nil {
		// Handle error
	}

	response, err := json.Marshal(&textsGetResponse{
		Count: len(texts),
		Texts: texts,
	})

	if err != nil {
		// Handle error
	}

	WriteJsonResponse(w, http.StatusOK, response)
}

func (c *TextController) TextPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	text, err := c.TextService.CreateText(title, content)

	if err != nil {
		// handle error
	}

	response, err := json.Marshal(&textCreateResponse{Text: text, Successful: true})

	if err != nil {
		// handle error
	}

	WriteJsonResponse(w, http.StatusOK, response)
}
