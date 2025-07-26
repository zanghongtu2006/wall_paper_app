package model

type Category struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
