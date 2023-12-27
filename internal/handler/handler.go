package handler

import (
	"net/http"
	"prel/config"
	"prel/internal"
	"prel/internal/usecase"

	"golang.org/x/oauth2"
)

type Handler struct {
	config             *config.Config
	oauthConfig        *oauth2.Config
	client             internal.GoogleCloudClient
	httpClient         *http.Client
	notificationClient internal.NotificationClient
	usecase            *usecase.Usecase
}

func NewHandler(
	c *config.Config,
	o *oauth2.Config,
	g internal.GoogleCloudClient,
	h *http.Client,
	n internal.NotificationClient) *Handler {
	return &Handler{
		config:             c,
		oauthConfig:        o,
		client:             g,
		httpClient:         h,
		notificationClient: n,
		usecase:            usecase.NewUsecase(c, g, h, n),
	}
}
