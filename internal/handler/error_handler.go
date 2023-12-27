package handler

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"prel/config"
	"prel/pkg/custom_error"
	"prel/pkg/logger"
	tpl "prel/web/template"

	"github.com/cockroachdb/errors"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
)

func ErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {

	var code int = http.StatusInternalServerError
	var customErr *custom_error.CustomError

	handleCustomError := func(err error) *custom_error.CustomError {
		details := errors.GetAllDetails(err)
		if len(details) == 0 {
			return nil
		}
		ec := custom_error.ErrorCode(details[0])
		customErr, ok := custom_error.GetError(ec)
		if !ok {
			return nil
		}
		return &customErr
	}

	switch {
	case errors.Is(err, ht.ErrNotImplemented):
		code = http.StatusNotImplemented
	default:
		switch e := err.(type) {
		case *ogenerrors.DecodeRequestError:
			code = e.Code()
		case *ogenerrors.DecodeParamsError:
			code = e.Code()
		case *ogenerrors.SecurityError:
			customErr = handleCustomError(e.Err)
			if customErr != nil {
				code = customErr.HttpStatusCode
			} else {
				code = http.StatusUnauthorized
			}
		default:
			customErr = handleCustomError(err)
			if customErr != nil {
				code = customErr.HttpStatusCode
			}
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)

	var title, description string

	switch code {
	case http.StatusBadRequest:
		title = "400 Bad Request"
		description = custom_error.DefaultMessageBadRequest
	case http.StatusUnauthorized:
		title = "401 Unauthorized"
		description = custom_error.DefaultMessageUnauthorized
	case http.StatusForbidden:
		title = "403 Forbidden"
		description = custom_error.DefaultMessageForbidden

	// not found(404) is handled by generated handler files by ogen
	// see: NotFoundHandler

	default:
		title = custom_error.DefaultTitle
		description = custom_error.DefaultInternalServerError
	}

	if customErr != nil {
		description = customErr.Message
		title = customErr.Title
	}

	if code >= 500 {
		logger.Get(ctx).Error(fmt.Sprintf("%+v", err))
	}

	if config.IsDebug() {
		// display beautiful stack trace
		slog.Error(fmt.Sprintf("%+v", err))
	}

	returnHtml(code, w, title, description)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	returnHtml(http.StatusNotFound, w, "404 Not Found", custom_error.DefaultNotFound)
}

func returnHtml(status int, w http.ResponseWriter, name, description string) {
	_ = generateErrorHTML(name, description, w)
}

func generateErrorHTML(name, description string, w io.Writer) error {
	tmpl, err := template.ParseFS(tpl.Files, tpl.ErrorPageTpl)
	if err != nil {
		return errors.Wrap(err, "failed to parse template")
	}

	data := &tpl.ErrorPageData{
		Name:        name,
		Description: template.HTML(description),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return errors.Wrap(err, "failed to execute template")
	}

	return nil
}
