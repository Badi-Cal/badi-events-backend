package controllers

import (
	"badi-cal/badi-events-backend/models"

	"github.com/golang/mock/gomock"
)

func CreateMockApp(ctrl *gomock.Controller) *App {
	return &App{
		Models: &models.Models{
			Notifications: nil,
			Users:         nil,
		},
	}
}
