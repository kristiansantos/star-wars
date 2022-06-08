package index

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/mocks"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
)

func TestIndexService(t *testing.T) {

	comm := communication.New()

	useCases := map[string]struct {
		urlParamId       string
		expectedResponse communication.Response
		prepare          func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"success: should return status 200": {
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success"].Code,
				Message: comm.Mapping["success"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planets{
					{
						ID:             "1",
						Name:           "testNameOne",
						Climate:        "testClimateOne",
						Ground:         "testGroundOne",
						FilmApparences: 5,
						CreatedAt:      time.Now(),
						UpdatedAt:      time.Now(),
					},
					{
						ID:             "2",
						Name:           "testNameTwo",
						Climate:        "testClimateTwo",
						Ground:         "testGroundTwo",
						FilmApparences: 10,
						CreatedAt:      time.Now(),
						UpdatedAt:      time.Now(),
					},
				}

				repostitoryMock.EXPECT().GetPlanets(gomock.Any()).Return(expectedData, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		}, "error: should return status 400 when error to find documents in database": {
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_list"].Code,
				Message: comm.Mapping["error_list"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planets{}

				repostitoryMock.EXPECT().GetPlanets(gomock.Any()).Return(expectedData, errors.New("error"))

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

			response := service.Execute()

			if response.Status != useCase.expectedResponse.Status {
				t.Errorf("Expected %d, but got %d", useCase.expectedResponse.Status, response.Status)
			}

			if response.Message != useCase.expectedResponse.Message {
				t.Errorf("Expected %s, but got %s", useCase.expectedResponse.Message, response.Message)
			}
		})
	}
}
