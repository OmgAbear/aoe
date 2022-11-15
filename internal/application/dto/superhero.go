package dto

type SuperheroOutputDto struct {
	Name        string   `json:"name"`
	Identity    string   `json:"identity"`
	Birthday    string   `json:"birthday"`
	Superpowers []string `json:"superpowers"`
}

type SuperheroInputDto struct {
	Name        string           `json:"name"`
	Identity    IdentityInputDto `json:"identity"`
	Birthday    string           `json:"birthday"`
	Superpowers []string         `json:"superpowers"`
}
