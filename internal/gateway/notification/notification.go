package notification

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"prel/internal/model"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
)

type SlackClient struct {
	httpClient *http.Client
	url        string
}

func NewSlackClient(c *http.Client, url string) *SlackClient {
	return &SlackClient{
		httpClient: c,
		url:        url,
	}
}

var (
	defaultTimeout time.Duration = 5
)

func SetDefaultTimeout(timeout time.Duration) {
	defaultTimeout = timeout
}

func (c *SlackClient) CanSend() bool {
	return c.url != ""
}

func (c *SlackClient) SendRequestMessage(ctx context.Context, targetId string, requestUrl, projectID, period, reason string, roles []string, requestExpiredAt time.Time) (*http.Response, error) {
	msg, err := c.requestMessage(targetId, requestUrl, projectID, period, reason, roles, requestExpiredAt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request message")
	}

	return c.send(ctx, msg)
}

func (c *SlackClient) SendJudgeMessage(ctx context.Context, judge model.RequestStatus, requesterId, judgerId, requestUrl, projectID, reason string, roles []string, until time.Time) (*http.Response, error) {
	msg, err := c.judgeMessage(judge, requesterId, judgerId, requestUrl, projectID, reason, roles, until)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create judge message")
	}

	return c.send(ctx, msg)
}
func (c *SlackClient) send(ctx context.Context, msg io.Reader) (*http.Response, error) {
	_, cancel := context.WithTimeout(ctx, defaultTimeout*time.Second)
	defer cancel()

	r, err := c.httpClient.Post(c.url, "application/json", msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send message")
	}

	return r, nil
}

func (c *SlackClient) requestMessage(targetId string, requestUrl, projectID, period, reason string, roles []string, requestExpiredAt time.Time) (io.Reader, error) {
	role := strings.Join(roles, "\n")
	message := map[string]interface{}{
		"text":   "",
		"blocks": []map[string]interface{}{},
		"attachments": []map[string]interface{}{
			{
				"color": "#000000",
				"blocks": []map[string]interface{}{
					{
						"type": "header",
						"text": map[string]string{
							"type": "plain_text",
							"text": "New request",
						},
					},
					{
						"type": "section",
						"fields": []map[string]string{
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Account:*\n%s", targetId),
							},
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Project:*\n%s", projectID),
							},
						},
					},
					{
						"type": "section",
						"fields": []map[string]string{
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*IAM Role:*\n%s", role),
							},
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Period:*\n%s", period),
							},
						},
					},
					{
						"type": "section",
						"fields": []map[string]string{
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Reason:*\n%s", reason),
							},
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Request Expiration Time:*\n%s", requestExpiredAt.Format("2006/01/02 15:04:05 MST")),
							},
						},
					},
					{
						"type": "section",
						"text": map[string]string{
							"type": "mrkdwn",
							"text": fmt.Sprintf("*Link*\n<%s|%s>", requestUrl, requestUrl),
						},
					},
				},
			},
		},
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(message)

	if err != nil {
		return nil, errors.Wrap(err, "failed to encode message")
	}

	return b, nil
}

func (c *SlackClient) judgeMessage(judge model.RequestStatus, requesterId, judgerId, requestUrl, projectID, reason string, roles []string, until time.Time) (io.Reader, error) {
	var color string
	switch judge {
	case model.RequestStatusApproved:
		color = "#00DF74"
	case model.RequestStatusRejected:
		color = "#DF0500"
	default:
		return nil, errors.Newf("invalid status(%s), status must be approved or rejected", judge)
	}

	role := strings.Join(roles, "\n")
	message := map[string]interface{}{
		"text":   "",
		"blocks": []map[string]interface{}{},
		"attachments": []map[string]interface{}{
			{
				"color": color,
				"blocks": []map[string]interface{}{
					{
						"type": "header",
						"text": map[string]string{
							"type": "plain_text",
							"text": fmt.Sprintf("Judged By %s", judgerId),
						},
					},
					{
						"type": "section",
						"fields": []map[string]string{
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Account:*\n%s", requesterId),
							},
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Project:*\n%s", projectID),
							},
						},
					},
					{
						"type": "section",
						"fields": []map[string]string{
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*IAM Roles:*\n%s", role),
							},
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Until:*\n%s", until.Format("2006/01/02 15:04:05 MST")),
							},
						},
					},
					{
						"type": "section",
						"fields": []map[string]string{
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Reason:*\n%s", reason),
							},
							{
								"type": "mrkdwn",
								"text": fmt.Sprintf("*Judge*\n%s", judge),
							},
						},
					},
					{
						"type": "section",
						"text": map[string]string{
							"type": "mrkdwn",
							"text": fmt.Sprintf("*Link*\n<%s|%s>", requestUrl, requestUrl),
						},
					},
				},
			},
		},
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(message)

	if err != nil {
		return nil, errors.Wrap(err, "failed to encode message")
	}

	return b, nil
}
