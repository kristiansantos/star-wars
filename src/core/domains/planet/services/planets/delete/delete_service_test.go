package delete

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/mocks"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
)

func TestDeleteService(t *testing.T) {

	comm := communication.New()

	useCases := map[string]struct {
		urlParamId       string
		expectedResponse communication.Response
		prepare          func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"success: should return status 204": {
			urlParamId: "1",
			expectedResponse: communication.Response{
				Status:  202,
				Code:    comm.Mapping["success"].Code,
				Message: comm.Mapping["success"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				repostitoryMock.EXPECT().Delete(gomock.Any()).Return(nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		}, "error: should return status 400 if error in delete": {
			urlParamId: "1",
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_delete"].Code,
				Message: comm.Mapping["error_delete"].Message,
			},
			prepare: func(repostitoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				repostitoryMock.EXPECT().Delete(gomock.Any()).Return(errors.New("error"))

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
