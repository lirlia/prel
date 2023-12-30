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

-- name: DeleteUserByEmail :exec
DELETE FROM users
WHERE email = $1;

-- name: DeleteAllIamRoleFilteringRules :exec
DELETE FROM iam_role_filtering_rules;
