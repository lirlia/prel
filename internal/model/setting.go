package model

import "github.com/google/uuid"

type Setting struct {
	id                            string
	notificationMessageForRequest string
	notificationMessageForJudge   string
}

func ReconstructSetting(
	id string,
	notificationMessageForRequest string,
	notificationMessageForJudge string) *Setting {

	return newSetting(
		id,
		notificationMessageForRequest,
		notificationMessageForJudge,
	)
}

func newSetting(
	id string,
	notificationMessageForRequest string,
	notificationMessageForJudge string) *Setting {
	return &Setting{
		id:                            id,
		notificationMessageForRequest: notificationMessageForRequest,
		notificationMessageForJudge:   notificationMessageForJudge,
	}
}

func NewSetting(
	notificationMessageForRequest string,
	notificationMessageForJudge string) *Setting {
	return newSetting(
		uuid.New().String(),
		notificationMessageForRequest,
		notificationMessageForJudge,
	)
}

func (s *Setting) ID() string {
	return s.id
}

func (s *Setting) NotificationMessageForRequest() string {
	return s.notificationMessageForRequest
}

func (s *Setting) NotificationMessageForJudge() string {
	return s.notificationMessageForJudge
}

func (s *Setting) UpdateNotificationMessageForRequest(message string) {
	s.notificationMessageForRequest = message
}

func (s *Setting) UpdateNotificationMessageForJudge(message string) {
	s.notificationMessageForJudge = message
}
