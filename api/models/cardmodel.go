package models

// swagger:model Card
type Card struct {
	// Id of the card
	// in: int64
	Id int64 `json:"id"`
	// Name of the card
	// in: string
	Name string `json:"name"`
	// Description of the card
	// in: string
	Description string `json:"description"`
	// Status of the card
	// in: int64
	Status int64 `json:"status"`

	// Creation Date of the card
	// in: string
	CreatedDate int64 `json:"createdDate"`
}

type Cards []Card

// swagger:parameters deleteCard
type ReqDeleteCard struct {
	// name: id
	// in: path
	// type: integer
	// required: true
	Id string `json:"id"`
}

type ReqAddCard struct {
	// Name of the card
	// in: string
	Name string `json:"name" validate:"required,min=2,max=100,alpha_space"`
	// Description of the card
	// in: string
	Description string `json:"description" validate:"required,min=2,max=100,alpha_space"`
}

// swagger:parameters editCard
type ReqCard struct {
	// name: id
	// in: path
	// type: integer
	// required: true
	Id int64 `json:"id"`
	// name: Name
	// in: formData
	// type: string
	// required: true
	Name        string `json:"name" validate:"required,min=2,max=100,alpha_space"`
	Description string `json:"description" validate:"required,min=2,max=100,alpha_space"`
}

// swagger:parameters addCard
type ReqCardBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqAddCard"
	//  required: true
	Body ReqAddCard `json:"body"`
}
