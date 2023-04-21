package dto

type CustomerResponse struct {
	Id       string `json:"id"`
	Name     string	`json:"full_name"`
	City     string	`json:"city"`
	Zipcode  string	`json:"zipcode"`
	Birthday string	`json:"birthday"`
	Status   string	`json:"status"`
}