package show

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/mock"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateService(t *testing.T) {

	comm := communication.New()

	useCases := map[string]struct {
		urlParamId       string
		expectedResponse communication.Response
		prepare          func(repostitoryMock *mock.MockIPlanetRepository, loggerMock *mock.MockILoggerProvider)
	}{
		"success": {
			urlParamId: "1",
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success"].Code,
				Message: comm.Mapping["success"].Message,
			},
			prepare: func(repostitoryMock *mock.MockIPlanetRepository, loggerMock *mock.MockILoggerProvider) {
				expectedData := entities.Planet{
					ID:             "1",
					Name:           "testName",
					Climate:        "testName",
					Ground:         "testName",
					FilmApparences: 5,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}

				repostitoryMock.EXPECT().GetById(gomock.Any()).Return(expectedData, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		}, "error: should return status 404 if error to find document not found in database": {
			urlParamId: "1",
			expectedResponse: communication.Response{
				Status:  404,
				Code:    comm.Mapping["not_found"].Code,
				Message: comm.Mapping["not_found"].Message,
			},
			prepare: func(repostitoryMock *mock.MockIPlanetRepository, loggerMock *mock.MockILoggerProvider) {
				expectedData := entities.Planet{}

				repostitoryMock.EXPECT().GetById(gomock.Any()).Return(expectedData, mongo.ErrNoDocuments)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		}, "error: should return status 400 if error to find document in database": {
			urlParamId: "1",
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_list"].Code,
				Message: comm.Mapping["error_list"].Message,
			},
			prepare: func(repostitoryMock *mock.MockIPlanetRepository, loggerMock *mock.MockILoggerProvider) {
				expectedData := entities.Planet{}

				repostitoryMock.EXPECT().GetById(gomock.Any()).Return(expectedData, errors.New("error"))

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mock.NewMockIPlanetRepository(ctrl)
			logger := mock.NewMockILoggerProvider(ctrl)

			useCase.prepare(repository, logger)

			service := Service{
				Logger:     logger,
				Repository: repository,
			}

			response := service.Execute(useCase.urlParamId)

			if response.Status != useCase.expectedResponse.Status {
				t.Errorf("Expected %d, but got %d", useCase.expectedResponse.Status, response.Status)
			}

			if response.Message != useCase.expectedResponse.Message {
				t.Errorf("Expected %s, but got %s", useCase.expectedResponse.Message, response.Message)
			}
		})
	}
}
