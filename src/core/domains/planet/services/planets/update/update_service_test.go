package update

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/mocks"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
)

func TestUpdateService(t *testing.T) {

	comm := communication.New()

	useCases := map[string]struct {
		urlParamId       string
		inputData        Dto
		expectedResponse communication.Response
		prepare          func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"success: should return status 200": {
			urlParamId: "1",
			inputData: Dto{
				Name:    "testName",
				Climate: "testClimate",
				Ground:  "testGround",
			},
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success"].Code,
				Message: comm.Mapping["success"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{
					ID:             "1",
					Name:           "testName",
					Climate:        "testClimate",
					Ground:         "testGround",
					FilmApparences: 5,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}

				repostitoryMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(expectedData, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		}, "Error: should return status 400 when database update return error": {
			urlParamId: "1",
			inputData: Dto{
				Name:    "testName",
				Climate: "testClimate",
				Ground:  "testGround",
			},
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_update"].Code,
				Message: comm.Mapping["error_update"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{}
				repostitoryMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(expectedData, errors.New("error"))

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

			response := service.Execute(useCase.urlParamId, useCase.inputData)

			if response.Status != useCase.expectedResponse.Status {
				t.Errorf("Expected %d, but got %d", useCase.expectedResponse.Status, response.Status)
			}

			if response.Message != useCase.expectedResponse.Message {
				t.Errorf("Expected %s, but got %s", useCase.expectedResponse.Message, response.Message)
			}
		})
	}
}
