package create

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/repositories"
	httphelper "gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/http_helper"
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

type Results struct {
	Films []string `json:"films"`
}

type ResponseBody struct {
	Results []Results `json:"results"`
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

	resul, _ := getFilmApparences(planet.Name)
	planet.FilmApparences = resul

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

	ctx, execTimeOut := context.WithTimeout(context.Background(), time.Second*60)
	defer execTimeOut()

	request, err := http.NewRequestWithContext(ctx, "GET", baseUrl+"planets/?search="+name+"&format=json", nil)

	if err != nil {
		return 0, err
	}

	if resp, err := http.DefaultClient.Do(request); err != nil {
		return 0, err
	} else if resp.StatusCode == 200 {
		var ParseResp ResponseBody

		httphelper.GetBody(resp.Body, &ParseResp)

		if len(ParseResp.Results) == 1 {
			return len(ParseResp.Results[0].Films), nil
		} else {
			return 0, nil
		}
	} else {
		return 0, errors.New("Request Error code:" + strconv.Itoa(resp.StatusCode))
	}
}
