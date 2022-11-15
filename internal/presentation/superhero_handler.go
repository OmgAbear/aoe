package presentation

import (
	"encoding/json"
	"github.com/OmgAbear/aoe/internal/application"
	"github.com/OmgAbear/aoe/internal/application/dto"
	"net/http"
)

type SuperheroHandler struct {
	appService application.Service
}

// NewSuperheroHandler - creates a new SuperheroHandler instance
func NewSuperheroHandler() SuperheroHandler {
	return SuperheroHandler{
		appService: application.NewService(),
	}
}

// List -http handler that lists all superheroes
func (handler SuperheroHandler) List(writer http.ResponseWriter, request *http.Request) {
	superheroes := handler.appService.List(request.URL.Query())
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(superheroes)
}

// Create - http handler that creates superhero entries
func (handler SuperheroHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var input dto.SuperheroInputDto

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		//normally I'd translate the error into custom errors of the app
		//generally I create some sort of app specific "http response errors" that "wrap" internal domain errors
		//with a mix of custom error codes as well as messages depending on case, but that also translates down to domain
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := handler.appService.Create(input)
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		//Another case for custom errors, both internal and external, that should be handled differently
		_ = json.NewEncoder(writer).Encode(struct {
			Err string `json:"err"`
		}{Err: err.Error()})
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(res)
}
