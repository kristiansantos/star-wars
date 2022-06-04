package create

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/repositories"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/namespace"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/parse"
)

var Namespace = namespace.New("core.domains.format.services.create.create_service")

type Dto struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Ground  string `json:"ground"`
}

type Service struct {
	Context    context.Context
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

func (service *Service) Execute(dto Dto) communication.Response {
	service.Logger.Info(Namespace.Concat("Execute"), "")

	comm := communication.New()

	var planet entities.Planet

	parse.Unmarshal(dto, &planet)
	planet.Populate()
	getFilmApparences(planet.Name)

	// Create document in database
	document, err := service.Repository.Create(planet)
	if err != nil {
		service.Logger.Error(Namespace.Concat("Execute", "Create"), err.Error())
		return comm.ResponseError(400, "error_create", err)
	}

	return comm.Response(201, "success_create", document)
}

func getFilmApparences(name string) (apparences int, err error) {
	var baseUrl = "https://swapi.dev/api/"

	if resp, err := http.Get(baseUrl + "planets/?name=" + name + "&format=json"); err != nil {
		return 0, err
	} else if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(respBody))

		return 0, nil
	} else {
		fmt.Println("Error:" + strconv.Itoa(resp.StatusCode))

		return 0, errors.New("Request Error code:" + strconv.Itoa(resp.StatusCode))
	}
}
