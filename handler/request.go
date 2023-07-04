package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func isEmpty(value string) bool {
	return value == ""
}

func isNotEmpty(value string) bool {
	return value != ""
}

func exists(value any) bool {
	return value == nil
}

type CreateOpenningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpenningRequest) validate() error {
	if isEmpty(r.Role) {
		return errParamIsRequired("role", "string")
	}

	if isEmpty(r.Company) {
		return errParamIsRequired("company", "string")
	}

	if isEmpty(r.Location) {
		return errParamIsRequired("location", "string")
	}

	if isEmpty(r.Link) {
		return errParamIsRequired("link", "string")
	}

	if exists(r.Remote) {
		return errParamIsRequired("remote", "bool")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

type UpdateOpenningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpenningRequest) validate() error {
	if isNotEmpty(r.Role) || isNotEmpty(r.Company) || isNotEmpty(r.Location) || isNotEmpty(r.Link) || exists(r.Remote) || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}
