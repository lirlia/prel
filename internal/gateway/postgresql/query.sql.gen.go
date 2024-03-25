// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countInvitation = `-- name: CountInvitation :one
SELECT COUNT(*) FROM invitations
`

func (q *Queries) CountInvitation(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countInvitation)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countRequest = `-- name: CountRequest :one
SELECT COUNT(*) FROM requests
`

func (q *Queries) CountRequest(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countRequest)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countUser = `-- name: CountUser :one
SELECT COUNT(*) FROM users
`

func (q *Queries) CountUser(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countUser)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteIamRoleFilteringRule = `-- name: DeleteIamRoleFilteringRule :exec
DELETE FROM iam_role_filtering_rules
WHERE id = $1
`

func (q *Queries) DeleteIamRoleFilteringRule(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteIamRoleFilteringRule, id)
	return err
}

const deleteInvitation = `-- name: DeleteInvitation :exec
DELETE FROM invitations
WHERE id = $1
`

func (q *Queries) DeleteInvitation(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteInvitation, id)
	return err
}

const deleteRequest = `-- name: DeleteRequest :exec
DELETE FROM requests
WHERE id = $1
`

func (q *Queries) DeleteRequest(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteRequest, id)
	return err
}

const findIamRoleFilteringRule = `-- name: FindIamRoleFilteringRule :many
SELECT id, pattern, user_id, created_at, updated_at FROM iam_role_filtering_rules
`

func (q *Queries) FindIamRoleFilteringRule(ctx context.Context) ([]IamRoleFilteringRule, error) {
	rows, err := q.db.Query(ctx, findIamRoleFilteringRule)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IamRoleFilteringRule
	for rows.Next() {
		var i IamRoleFilteringRule
		if err := rows.Scan(
			&i.ID,
			&i.Pattern,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findIamRoleFilteringRuleByID = `-- name: FindIamRoleFilteringRuleByID :one
SELECT id, pattern, user_id, created_at, updated_at FROM iam_role_filtering_rules
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindIamRoleFilteringRuleByID(ctx context.Context, id string) (IamRoleFilteringRule, error) {
	row := q.db.QueryRow(ctx, findIamRoleFilteringRuleByID, id)
	var i IamRoleFilteringRule
	err := row.Scan(
		&i.ID,
		&i.Pattern,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findInvitationByInviteeMail = `-- name: FindInvitationByInviteeMail :one
SELECT id, inviter_user_id, invitee_mail, invitee_role, invited_at, expired_at, created_at, updated_at FROM invitations
WHERE invitee_mail = $1 LIMIT 1
`

func (q *Queries) FindInvitationByInviteeMail(ctx context.Context, inviteeMail string) (Invitation, error) {
	row := q.db.QueryRow(ctx, findInvitationByInviteeMail, inviteeMail)
	var i Invitation
	err := row.Scan(
		&i.ID,
		&i.InviterUserID,
		&i.InviteeMail,
		&i.InviteeRole,
		&i.InvitedAt,
		&i.ExpiredAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findInvitationByInviteeMailsAndExpiredAt = `-- name: FindInvitationByInviteeMailsAndExpiredAt :many
SELECT id, inviter_user_id, invitee_mail, invitee_role, invited_at, expired_at, created_at, updated_at FROM invitations
WHERE invitee_mail = ANY($1::text[]) AND expired_at >= $2
`

type FindInvitationByInviteeMailsAndExpiredAtParams struct {
	Column1   []string
	ExpiredAt pgtype.Timestamptz
}

func (q *Queries) FindInvitationByInviteeMailsAndExpiredAt(ctx context.Context, arg FindInvitationByInviteeMailsAndExpiredAtParams) ([]Invitation, error) {
	rows, err := q.db.Query(ctx, findInvitationByInviteeMailsAndExpiredAt, arg.Column1, arg.ExpiredAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Invitation
	for rows.Next() {
		var i Invitation
		if err := rows.Scan(
			&i.ID,
			&i.InviterUserID,
			&i.InviteeMail,
			&i.InviteeRole,
			&i.InvitedAt,
			&i.ExpiredAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findRequestByID = `-- name: FindRequestByID :one
SELECT
    requests.id, requests.requester_user_id, requests.judger_user_id, requests.status, requests.project_id, requests.iam_roles, requests.period, requests.reason, requests.requested_at, requests.expired_at, requests.judged_at, requests.created_at, requests.updated_at,
    users.email AS requester_email,
    judger.email AS judger_email
FROM
    requests
LEFT OUTER JOIN
    users ON requests.requester_user_id = users.id
LEFT OUTER JOIN
    users AS judger ON requests.judger_user_id = judger.id
WHERE
    requests.id = $1 LIMIT 1
`

type FindRequestByIDRow struct {
	ID              string
	RequesterUserID string
	JudgerUserID    pgtype.Text
	Status          string
	ProjectID       string
	IamRoles        string
	Period          int32
	Reason          string
	RequestedAt     pgtype.Timestamptz
	ExpiredAt       pgtype.Timestamptz
	JudgedAt        pgtype.Timestamptz
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	RequesterEmail  pgtype.Text
	JudgerEmail     pgtype.Text
}

func (q *Queries) FindRequestByID(ctx context.Context, id string) (FindRequestByIDRow, error) {
	row := q.db.QueryRow(ctx, findRequestByID, id)
	var i FindRequestByIDRow
	err := row.Scan(
		&i.ID,
		&i.RequesterUserID,
		&i.JudgerUserID,
		&i.Status,
		&i.ProjectID,
		&i.IamRoles,
		&i.Period,
		&i.Reason,
		&i.RequestedAt,
		&i.ExpiredAt,
		&i.JudgedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RequesterEmail,
		&i.JudgerEmail,
	)
	return i, err
}

const findRequestByRequestUserID = `-- name: FindRequestByRequestUserID :many
SELECT
    requests.id, requests.requester_user_id, requests.judger_user_id, requests.status, requests.project_id, requests.iam_roles, requests.period, requests.reason, requests.requested_at, requests.expired_at, requests.judged_at, requests.created_at, requests.updated_at,
    users.email AS requester_email,
    judger.email AS judger_email
FROM
    requests
LEFT OUTER JOIN
    users ON requests.requester_user_id = users.id
LEFT OUTER JOIN
    users AS judger ON requests.judger_user_id = judger.id
WHERE
    requester_user_id = $1
`

type FindRequestByRequestUserIDRow struct {
	ID              string
	RequesterUserID string
	JudgerUserID    pgtype.Text
	Status          string
	ProjectID       string
	IamRoles        string
	Period          int32
	Reason          string
	RequestedAt     pgtype.Timestamptz
	ExpiredAt       pgtype.Timestamptz
	JudgedAt        pgtype.Timestamptz
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	RequesterEmail  pgtype.Text
	JudgerEmail     pgtype.Text
}

func (q *Queries) FindRequestByRequestUserID(ctx context.Context, requesterUserID string) ([]FindRequestByRequestUserIDRow, error) {
	rows, err := q.db.Query(ctx, findRequestByRequestUserID, requesterUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindRequestByRequestUserIDRow
	for rows.Next() {
		var i FindRequestByRequestUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.RequesterUserID,
			&i.JudgerUserID,
			&i.Status,
			&i.ProjectID,
			&i.IamRoles,
			&i.Period,
			&i.Reason,
			&i.RequestedAt,
			&i.ExpiredAt,
			&i.JudgedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RequesterEmail,
			&i.JudgerEmail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findRequestByStatusAndExpiredAt = `-- name: FindRequestByStatusAndExpiredAt :many
SELECT
    requests.id, requests.requester_user_id, requests.judger_user_id, requests.status, requests.project_id, requests.iam_roles, requests.period, requests.reason, requests.requested_at, requests.expired_at, requests.judged_at, requests.created_at, requests.updated_at,
    users.email AS requester_email,
    judger.email AS judger_email
FROM
    requests
LEFT OUTER JOIN
    users ON requests.requester_user_id = users.id
LEFT OUTER JOIN
    users AS judger ON requests.judger_user_id = judger.id
WHERE
    requests.status = $1 AND requests.expired_at >= $2
`

type FindRequestByStatusAndExpiredAtParams struct {
	Status    string
	ExpiredAt pgtype.Timestamptz
}

type FindRequestByStatusAndExpiredAtRow struct {
	ID              string
	RequesterUserID string
	JudgerUserID    pgtype.Text
	Status          string
	ProjectID       string
	IamRoles        string
	Period          int32
	Reason          string
	RequestedAt     pgtype.Timestamptz
	ExpiredAt       pgtype.Timestamptz
	JudgedAt        pgtype.Timestamptz
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	RequesterEmail  pgtype.Text
	JudgerEmail     pgtype.Text
}

func (q *Queries) FindRequestByStatusAndExpiredAt(ctx context.Context, arg FindRequestByStatusAndExpiredAtParams) ([]FindRequestByStatusAndExpiredAtRow, error) {
	rows, err := q.db.Query(ctx, findRequestByStatusAndExpiredAt, arg.Status, arg.ExpiredAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindRequestByStatusAndExpiredAtRow
	for rows.Next() {
		var i FindRequestByStatusAndExpiredAtRow
		if err := rows.Scan(
			&i.ID,
			&i.RequesterUserID,
			&i.JudgerUserID,
			&i.Status,
			&i.ProjectID,
			&i.IamRoles,
			&i.Period,
			&i.Reason,
			&i.RequestedAt,
			&i.ExpiredAt,
			&i.JudgedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RequesterEmail,
			&i.JudgerEmail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findRequestPaged = `-- name: FindRequestPaged :many
SELECT
    requests.id, requests.requester_user_id, requests.judger_user_id, requests.status, requests.project_id, requests.iam_roles, requests.period, requests.reason, requests.requested_at, requests.expired_at, requests.judged_at, requests.created_at, requests.updated_at,
    users.email AS requester_email,
    judger.email AS judger_email
FROM
    requests
LEFT OUTER JOIN
    users ON requests.requester_user_id = users.id
LEFT OUTER JOIN
    users AS judger ON requests.judger_user_id = judger.id
ORDER BY
    requests.requested_at DESC
LIMIT $1 OFFSET $2
`

type FindRequestPagedParams struct {
	Limit  int32
	Offset int32
}

type FindRequestPagedRow struct {
	ID              string
	RequesterUserID string
	JudgerUserID    pgtype.Text
	Status          string
	ProjectID       string
	IamRoles        string
	Period          int32
	Reason          string
	RequestedAt     pgtype.Timestamptz
	ExpiredAt       pgtype.Timestamptz
	JudgedAt        pgtype.Timestamptz
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	RequesterEmail  pgtype.Text
	JudgerEmail     pgtype.Text
}

func (q *Queries) FindRequestPaged(ctx context.Context, arg FindRequestPagedParams) ([]FindRequestPagedRow, error) {
	rows, err := q.db.Query(ctx, findRequestPaged, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindRequestPagedRow
	for rows.Next() {
		var i FindRequestPagedRow
		if err := rows.Scan(
			&i.ID,
			&i.RequesterUserID,
			&i.JudgerUserID,
			&i.Status,
			&i.ProjectID,
			&i.IamRoles,
			&i.Period,
			&i.Reason,
			&i.RequestedAt,
			&i.ExpiredAt,
			&i.JudgedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RequesterEmail,
			&i.JudgerEmail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findSetting = `-- name: FindSetting :one
SELECT id, notification_message_for_request, notification_message_for_judge, created_at, updated_at FROM setting LIMIT 1
`

func (q *Queries) FindSetting(ctx context.Context) (Setting, error) {
	row := q.db.QueryRow(ctx, findSetting)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.NotificationMessageForRequest,
		&i.NotificationMessageForJudge,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserAndInvitationPagedByExpiredAt = `-- name: FindUserAndInvitationPagedByExpiredAt :many
SELECT id, google_id, email, role, is_available, session_id, session_expired_at, last_signin_at, expired_at FROM (
    SELECT
        id,
        google_id,
        email as email,
        role,
        is_available,
        session_id,
        session_expired_at,
        last_signin_at,
        NULL::timestamp as expired_at
    FROM
        users
    UNION ALL
    SELECT
        id,
        '' as google_id,
        invitee_mail as email,
        invitee_role as role,
        TRUE as is_available,
        '' as session_id,
        NULL::timestamp as session_expired_at,
        NULL::timestamp as last_signin_at,
        expired_at
    FROM
        invitations
    WHERE
        expired_at >= $1
) as combined
ORDER BY combined.email DESC
LIMIT $2 OFFSET $3
`

type FindUserAndInvitationPagedByExpiredAtParams struct {
	ExpiredAt pgtype.Timestamptz
	Limit     int32
	Offset    int32
}

type FindUserAndInvitationPagedByExpiredAtRow struct {
	ID               string
	GoogleID         string
	Email            string
	Role             string
	IsAvailable      bool
	SessionID        string
	SessionExpiredAt pgtype.Timestamptz
	LastSigninAt     pgtype.Timestamptz
	ExpiredAt        pgtype.Timestamp
}

func (q *Queries) FindUserAndInvitationPagedByExpiredAt(ctx context.Context, arg FindUserAndInvitationPagedByExpiredAtParams) ([]FindUserAndInvitationPagedByExpiredAtRow, error) {
	rows, err := q.db.Query(ctx, findUserAndInvitationPagedByExpiredAt, arg.ExpiredAt, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindUserAndInvitationPagedByExpiredAtRow
	for rows.Next() {
		var i FindUserAndInvitationPagedByExpiredAtRow
		if err := rows.Scan(
			&i.ID,
			&i.GoogleID,
			&i.Email,
			&i.Role,
			&i.IsAvailable,
			&i.SessionID,
			&i.SessionExpiredAt,
			&i.LastSigninAt,
			&i.ExpiredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, google_id, email, is_available, role, session_id, session_expired_at, last_signin_at, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.GoogleID,
		&i.Email,
		&i.IsAvailable,
		&i.Role,
		&i.SessionID,
		&i.SessionExpiredAt,
		&i.LastSigninAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByGoogleID = `-- name: FindUserByGoogleID :one
SELECT id, google_id, email, is_available, role, session_id, session_expired_at, last_signin_at, created_at, updated_at FROM users
WHERE google_id = $1 LIMIT 1
`

func (q *Queries) FindUserByGoogleID(ctx context.Context, googleID string) (User, error) {
	row := q.db.QueryRow(ctx, findUserByGoogleID, googleID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.GoogleID,
		&i.Email,
		&i.IsAvailable,
		&i.Role,
		&i.SessionID,
		&i.SessionExpiredAt,
		&i.LastSigninAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT id, google_id, email, is_available, role, session_id, session_expired_at, last_signin_at, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindUserByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.GoogleID,
		&i.Email,
		&i.IsAvailable,
		&i.Role,
		&i.SessionID,
		&i.SessionExpiredAt,
		&i.LastSigninAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByIDs = `-- name: FindUserByIDs :many
SELECT id, google_id, email, is_available, role, session_id, session_expired_at, last_signin_at, created_at, updated_at FROM users
WHERE id = ANY($1::text[])
`

func (q *Queries) FindUserByIDs(ctx context.Context, dollar_1 []string) ([]User, error) {
	rows, err := q.db.Query(ctx, findUserByIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.GoogleID,
			&i.Email,
			&i.IsAvailable,
			&i.Role,
			&i.SessionID,
			&i.SessionExpiredAt,
			&i.LastSigninAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUserBySessionID = `-- name: FindUserBySessionID :one
SELECT id, google_id, email, is_available, role, session_id, session_expired_at, last_signin_at, created_at, updated_at FROM users
WHERE session_id = $1 LIMIT 1
`

func (q *Queries) FindUserBySessionID(ctx context.Context, sessionID string) (User, error) {
	row := q.db.QueryRow(ctx, findUserBySessionID, sessionID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.GoogleID,
		&i.Email,
		&i.IsAvailable,
		&i.Role,
		&i.SessionID,
		&i.SessionExpiredAt,
		&i.LastSigninAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertIamRoleFilteringRule = `-- name: UpsertIamRoleFilteringRule :exec
INSERT INTO iam_role_filtering_rules (
    id,
    pattern,
    user_id
) VALUES (
    $1,
    $2,
    $3
) ON CONFLICT (id) DO UPDATE SET
    pattern = excluded.pattern,
    user_id = excluded.user_id
`

type UpsertIamRoleFilteringRuleParams struct {
	ID      string
	Pattern string
	UserID  string
}

func (q *Queries) UpsertIamRoleFilteringRule(ctx context.Context, arg UpsertIamRoleFilteringRuleParams) error {
	_, err := q.db.Exec(ctx, upsertIamRoleFilteringRule, arg.ID, arg.Pattern, arg.UserID)
	return err
}

const upsertInvitation = `-- name: UpsertInvitation :exec
INSERT INTO invitations (
    id,
    inviter_user_id,
    invitee_mail,
    invitee_role,
    invited_at,
    expired_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) ON CONFLICT (id) DO UPDATE SET
    inviter_user_id = excluded.inviter_user_id,
    invitee_mail = excluded.invitee_mail,
    invitee_role = excluded.invitee_role,
    invited_at = excluded.invited_at,
    expired_at = excluded.expired_at
`

type UpsertInvitationParams struct {
	ID            string
	InviterUserID string
	InviteeMail   string
	InviteeRole   string
	InvitedAt     pgtype.Timestamptz
	ExpiredAt     pgtype.Timestamptz
}

func (q *Queries) UpsertInvitation(ctx context.Context, arg UpsertInvitationParams) error {
	_, err := q.db.Exec(ctx, upsertInvitation,
		arg.ID,
		arg.InviterUserID,
		arg.InviteeMail,
		arg.InviteeRole,
		arg.InvitedAt,
		arg.ExpiredAt,
	)
	return err
}

const upsertRequest = `-- name: UpsertRequest :exec
INSERT INTO requests (
    id,
    requester_user_id,
    judger_user_id,
    status,
    project_id,
    iam_roles,
    period,
    reason,
    requested_at,
    expired_at,
    judged_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11
) ON CONFLICT (id) DO UPDATE SET
    requester_user_id = excluded.requester_user_id,
    judger_user_id = excluded.judger_user_id,
    status = excluded.status,
    project_id = excluded.project_id,
    iam_roles = excluded.iam_roles,
    period = excluded.period,
    reason = excluded.reason,
    requested_at = excluded.requested_at,
    expired_at = excluded.expired_at,
    judged_at = excluded.judged_at
`

type UpsertRequestParams struct {
	ID              string
	RequesterUserID string
	JudgerUserID    pgtype.Text
	Status          string
	ProjectID       string
	IamRoles        string
	Period          int32
	Reason          string
	RequestedAt     pgtype.Timestamptz
	ExpiredAt       pgtype.Timestamptz
	JudgedAt        pgtype.Timestamptz
}

func (q *Queries) UpsertRequest(ctx context.Context, arg UpsertRequestParams) error {
	_, err := q.db.Exec(ctx, upsertRequest,
		arg.ID,
		arg.RequesterUserID,
		arg.JudgerUserID,
		arg.Status,
		arg.ProjectID,
		arg.IamRoles,
		arg.Period,
		arg.Reason,
		arg.RequestedAt,
		arg.ExpiredAt,
		arg.JudgedAt,
	)
	return err
}

const upsertSetting = `-- name: UpsertSetting :exec
INSERT INTO setting (
    id,
    notification_message_for_request,
    notification_message_for_judge
) VALUES (
    $1,
    $2,
    $3
) ON CONFLICT (id) DO UPDATE SET
    notification_message_for_request = excluded.notification_message_for_request,
    notification_message_for_judge = excluded.notification_message_for_judge
`

type UpsertSettingParams struct {
	ID                            string
	NotificationMessageForRequest pgtype.Text
	NotificationMessageForJudge   pgtype.Text
}

func (q *Queries) UpsertSetting(ctx context.Context, arg UpsertSettingParams) error {
	_, err := q.db.Exec(ctx, upsertSetting, arg.ID, arg.NotificationMessageForRequest, arg.NotificationMessageForJudge)
	return err
}

const upsertUser = `-- name: UpsertUser :exec
INSERT INTO users (
    id,
    google_id,
    email,
    role,
    is_available,
    session_id,
    session_expired_at,
    last_signin_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) ON CONFLICT (id) DO UPDATE SET
    google_id = excluded.google_id,
    email = excluded.email,
    role = excluded.role,
    is_available = excluded.is_available,
    session_id = excluded.session_id,
    session_expired_at = excluded.session_expired_at,
    last_signin_at = excluded.last_signin_at
`

type UpsertUserParams struct {
	ID               string
	GoogleID         string
	Email            string
	Role             string
	IsAvailable      bool
	SessionID        string
	SessionExpiredAt pgtype.Timestamptz
	LastSigninAt     pgtype.Timestamptz
}

func (q *Queries) UpsertUser(ctx context.Context, arg UpsertUserParams) error {
	_, err := q.db.Exec(ctx, upsertUser,
		arg.ID,
		arg.GoogleID,
		arg.Email,
		arg.Role,
		arg.IsAvailable,
		arg.SessionID,
		arg.SessionExpiredAt,
		arg.LastSigninAt,
	)
	return err
}
