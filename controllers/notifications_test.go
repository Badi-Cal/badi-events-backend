package controllers

import (
	"badi-cal/badi-events-backend/mock_models"
	"badi-cal/badi-events-backend/orm"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNotificationsIndex(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := CreateMockApp(ctrl)
	m := mock_models.NewMockINotifications(ctrl)
	m.
		EXPECT().
		GetAll().
		Return([]*orm.Notification{}, nil).
		AnyTimes()
	app.Models.Notifications = m

	r := httptest.NewRequest(
		http.MethodGet,
		"/notifications",
		nil,
	)
	w := httptest.NewRecorder()
	app.Route(w, r)

	assert.Equal(w.Code, http.StatusOK)
	assert.Equal(w.Body.String(), "[]")
}
