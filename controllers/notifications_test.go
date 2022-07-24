package controllers

import (
	"badi-cal/badi-events-backend/mock_models"
	"badi-cal/badi-events-backend/models"
	"badi-cal/badi-events-backend/orm"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNotificationsCreate(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := CreateMockApp(ctrl)
	now := time.Now()
	now_str := now.Format(time.RFC3339)
	m := mock_models.NewMockINotifications(ctrl)
	m.
		EXPECT().
		Create(models.NotificationCreatePayload{
			Event_id:       "abc",
			Body:           "my_body",
			Scheduled_time: now_str,
		}).
		Return(&orm.Notification{
			ID:             1,
			Event_id:       "abc",
			Body:           "my_body",
			Scheduled_time: now,
			Status:         orm.Unsent,
			Attempts:       0,
			CreatedAt:      now,
			UpdatedAt:      now,
		}, nil).
		AnyTimes()
	app.Models.Notifications = m

	bodyReader := strings.NewReader(fmt.Sprintf(`{"event_id": "abc", "body": "my_body", "scheduled_time": "%s"}`, now_str))
	r := httptest.NewRequest(
		http.MethodPut,
		"/notifications",
		bodyReader,
	)
	w := httptest.NewRecorder()
	app.Route(w, r)

	assert.Equal(w.Code, http.StatusOK)
	assert.Equal(w.Body.String(), fmt.Sprintf(`{"id":1,"event_id":"abc","body":"my_body","scheduled_time":"%s","status":"unsent","attempts":0}`, now_str))
}

func TestNotificationsIndex(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := CreateMockApp(ctrl)
	now := time.Now()
	now_str := now.Format(time.RFC3339)
	records := []*orm.Notification{{
		ID:             1,
		Event_id:       "abc",
		Body:           "my_body",
		Scheduled_time: now,
		Status:         orm.Unsent,
		Attempts:       0,
		CreatedAt:      now,
		UpdatedAt:      now,
	}}
	m := mock_models.NewMockINotifications(ctrl)
	m.
		EXPECT().
		GetAll().
		Return(records, nil).
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
	assert.Equal(w.Body.String(), fmt.Sprintf(`[{"id":1,"event_id":"abc","body":"my_body","scheduled_time":"%s","status":"unsent","attempts":0}]`, now_str))
}
