package handlers

import (
	"MiddleApp/internal/repositories/interfaces/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

var (
	doesNotExist = errors.New("does not exist")
)

func TestHandlers_Delete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRep := mock_interfaces.NewMockRepository(controller)

	tests := []struct {
		name            string
		id              uint32
		expectedIdValue uint32
		expectedError   error
	}{
		{
			name:            "Ok",
			id:              0,
			expectedIdValue: 0,
			expectedError:   nil,
		},
		{
			name:            "Not ok",
			id:              6,
			expectedIdValue: 6,
			expectedError:   doesNotExist,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			mockRep.EXPECT().Delete(testCase.id).Return(nil, testCase.id)
			h := New(mockRep)
			var w http.ResponseWriter
			var r *http.Request
			h.Delete(w, r)

		})
	}
}
