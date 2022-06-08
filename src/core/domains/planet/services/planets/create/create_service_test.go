package create

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/mocks"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
)

func TestCreateService(t *testing.T) {

	comm := communication.New()

	useCases := map[string]struct {
		inputData        Dto
		expectedResponse communication.Response
		prepare          func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"success: should return status 200": {
			inputData: Dto{
				Name:    "testName",
				Climate: "testName",
				Ground:  "testName",
			},
			expectedResponse: communication.Response{
				Status:  201,
				Code:    comm.Mapping["success_create"].Code,
				Message: comm.Mapping["success_create"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{
					ID:             "1",
					Name:           "testName",
					Climate:        "testName",
					Ground:         "testName",
					FilmApparences: 5,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}

				repostitoryMock.EXPECT().Create(gomock.Any()).Return(expectedData, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		}, "Error: should return status 400 when database save return error": {
			inputData: Dto{
				Name:    "testName",
				Climate: "testName",
				Ground:  "testName",
			},
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_create"].Code,
				Message: comm.Mapping["error_create"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{}
				repostitoryMock.EXPECT().Create(gomock.Any()).Return(expectedData, errors.New("error"))

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mocks.NewMockIPlanetRepository(ctrl)
			logger := mocks.NewMockILoggerProvider(ctrl)

			useCase.prepare(repository, logger)

			service := Service{
				Logger:     logger,
				Repository: repository,
			}

			response := service.Execute(useCase.inputData)

			if response.Status != useCase.expectedResponse.Status {
				t.Errorf("Expected %d, but got %d", useCase.expectedResponse.Status, response.Status)
			}

			if response.Message != useCase.expectedResponse.Message {
				t.Errorf("Expected %s, but got %s", useCase.expectedResponse.Message, response.Message)
			}
		})
	}
}
