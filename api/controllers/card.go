package controllers

import (
	"api/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CardApi struct{}

func NewCardApi() *CardApi {
	return &CardApi{}
}

// swagger:model HTTPError
type HTTPError struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:model GetCards
type CardsResponse struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string        `json:"message"`
	Data    *models.Cards `json:"data"`
}

// swagger:model GetCard
type CardResponse struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string `json:"message"`
	// Cards for this user
	Data *models.Card `json:"data"`
}

// swagger:model GetCard
type CardRequest struct {
	// Cards for this user
	Data *models.ReqAddCard `json:"data"`
}

// GetCards godoc
//
//	@Summary		Get  a list of cards
//	@Description	Get  a list of cards
//	@Tags			/cards
//	@Accept			json
//	@Produce		json
//
// @Success 200 {object} CardsResponse
//
//	@Failure		400		{object}	HTTPError
//	@Router			/cards [get]
func GetCards(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Welcome to the Article REST API!")
	fmt.Println("Article REST API")
	fmt.Println(vars)
}

// GetCard godoc
// @Summary Get details for a given cardId
// @Description Get details of card corresponding to the input cardId
//
//	@Tags			/cards
//	@Accept			json
//	@Produce		json
//
// @Param cardId path int true "ID of the card"
//
// @Success 200 {object} CardResponse
//
//	@Failure		400		{object}	HTTPError
//		@Router			/cards/{id} [get]
func GetCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Welcome to the Article REST API!")
	fmt.Println("Article REST API")
	fmt.Println(vars)
}

// CreateCard godoc
//
//	@Summary		Create a card
//	@Description	Create a new card item
//	@Tags			/cards
//	@Accept			json
//	@Produce		json
//	@Param			CardRequest	body	CardRequest	true	"Add a New Card Request"
//
// @Success 200 {object} CardResponse
//
//	@Failure		400		{object}	HTTPError
//		@Router			/cards [post]
func CreateCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Welcome to the Article REST API!")
	fmt.Println("Article REST API")
	fmt.Println(vars)
}

// UpdateCard godoc
//
// @Summary Update card identified by the given cardId
// @Description Update the card corresponding to the input cardId
//
//	@Tags			/cards
//	@Accept			json
//	@Produce		json
//
// @Param cardId path int true "ID of the card to be updated"
//
// @Success 200 {object} CardResponse
//
//	@Failure		400		{object}	HTTPError
//		@Router			/cards/{id} [put]
func UpdateCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Welcome to the Article REST API!")
	fmt.Println("Article REST API")
	fmt.Println(vars)
}

// DeleteCard godoc
//
//	@Summary		Delete a card
//	@Description	Delete a new card item
//	@Tags			/cards
//	@Accept			json
//	@Produce		json
//
// @Param cardId path int true "ID of the card to be deleted"
//
// @Success 204 "No Content"
//
//	@Failure		400		{object}	HTTPError
//	@Router			/cards/{id} [delete]
func DeleteCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Welcome to the Article REST API!")
	fmt.Println("Article REST API")
	fmt.Println(vars)
}
