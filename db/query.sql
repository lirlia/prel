-- name: FindUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: FindUserByIDs :many
SELECT * FROM users
WHERE id = ANY($1::text[]);

-- name: FindUserByGoogleID :one
SELECT * FROM users
WHERE google_id = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: FindUserBySessionID :one
SELECT * FROM users
WHERE session_id = $1 LIMIT 1;

-- name: CountUser :one
SELECT COUNT(*) FROM users;

-- name: UpsertUser :exec
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
    last_signin_at = excluded.last_signin_at;

-- name: FindUserAndInvitationPagedByExpiredAt :many
SELECT * FROM (
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
LIMIT $2 OFFSET $3;

-- name: FindRequestByRequestUserID :many
SELECT
    requests.*,
    users.email AS requester_email,
    judger.email AS judger_email
FROM
    requests
LEFT OUTER JOIN
    users ON requests.requester_user_id = users.id
LEFT OUTER JOIN
    users AS judger ON requests.judger_user_id = judger.id
WHERE
    requester_user_id = $1;

-- name: FindRequestByID :one
SELECT
    requests.*,
    users.email AS requester_email,
    judger.email AS judger_email
FROM
    requests
LEFT OUTER JOIN
    users ON requests.requester_user_id = users.id
LEFT OUTER JOIN
    users AS judger ON requests.judger_user_id = judger.id
WHERE
    requests.id = $1 LIMIT 1;

-- name: FindRequestByStatusAndExpiredAt :many
SELECT
    requests.*,
    users.email AS requester_email,
    judger.email AS judger_email
FROM
    requests
LEFT OUTER JOIN
    users ON requests.requester_user_id = users.id
LEFT OUTER JOIN
    users AS judger ON requests.judger_user_id = judger.id
WHERE
    requests.status = $1 AND requests.expired_at >= $2;

-- name: FindRequestPaged :many
SELECT
    requests.*,
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
LIMIT $1 OFFSET $2;

-- name: CountRequest :one
SELECT COUNT(*) FROM requests;

-- name: UpsertRequest :exec
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
    judged_at = excluded.judged_at;

-- name: DeleteRequest :exec
DELETE FROM requests
WHERE id = $1;

-- name: FindInvitationByInviteeMail :one
SELECT * FROM invitations
WHERE invitee_mail = $1 LIMIT 1;

-- name: FindInvitationByInviteeMailsAndExpiredAt :many
SELECT * FROM invitations
WHERE invitee_mail = ANY($1::text[]) AND expired_at >= $2;

-- name: CountInvitation :one
SELECT COUNT(*) FROM invitations;

-- name: UpsertInvitation :exec
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
    expired_at = excluded.expired_at;

-- name: DeleteInvitation :exec
DELETE FROM invitations
WHERE id = $1;

-- name: FindIamRoleFilteringRuleByID :one
SELECT * FROM iam_role_filtering_rules
WHERE id = $1 LIMIT 1;

-- name: FindIamRoleFilteringRule :many
SELECT * FROM iam_role_filtering_rules;

-- name: UpsertIamRoleFilteringRule :exec
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
    user_id = excluded.user_id;

-- name: DeleteIamRoleFilteringRule :exec
DELETE FROM iam_role_filtering_rules
WHERE id = $1;
