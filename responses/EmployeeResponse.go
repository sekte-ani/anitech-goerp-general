package responses

type EmployeeResponse struct {
	ID       		uint 		`json:"id"`
	Name     		string 		`json:"name"`
	RoleName 		[]string 	`json:"role_name"`
	Images 			string 		`json:"images"`
}