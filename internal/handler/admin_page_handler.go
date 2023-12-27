package handler

import (
	"bytes"
	"context"
	"html/template"
	api "prel/api/prel_api"
	"prel/config"
	"prel/internal/model"
	tpl "prel/web/template"

	"github.com/cockroachdb/errors"
)

func (h *Handler) AdminRequestGet(ctx context.Context) (api.AdminRequestGetRes, error) {

	tmpl, err := template.ParseFS(tpl.Files, tpl.AdminRequestPageTpl, tpl.HeaderTpl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}

	user := model.GetUser(ctx)
	b := &bytes.Buffer{}
	err = tmpl.Execute(b, &tpl.PageData{
		HeaderData: tpl.NewHeaderData(user.IsAdmin(), config.AppName),
		AdminListPage: tpl.AdminListPage{
			Options: api.AllPageSize(),
		},
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}

	return &api.AdminRequestGetOK{
		Data: b,
	}, nil
}

func (h *Handler) AdminUserGet(ctx context.Context) (api.AdminUserGetRes, error) {

	tmpl, err := template.ParseFS(tpl.Files, tpl.AdminUserPageTpl, tpl.HeaderTpl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}

	b := &bytes.Buffer{}
	user := model.GetUser(ctx)
	err = tmpl.Execute(b, &tpl.PageData{
		HeaderData: tpl.NewHeaderData(user.IsAdmin(), config.AppName),
		AdminListPage: tpl.AdminListPage{
			Options:   api.AllPageSize(),
			UserRoles: convertUserRoles(model.SortedUserRoles()),
		},
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}

	return &api.AdminUserGetOK{
		Data: b,
	}, nil
}
