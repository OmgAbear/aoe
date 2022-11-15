package dto

type IdentityOutputDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type IdentityInputDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
