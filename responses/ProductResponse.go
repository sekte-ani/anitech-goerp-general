package responses

type ProductResponse struct {
	ID 				uint 		`json:"id"`
  	Name     		string 		`json:"name"`
	Description     string 		`json:"description"`
	Images     		string 		`json:"images"`
}