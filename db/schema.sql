CREATE TABLE users (
    id TEXT PRIMARY KEY NOT NULL,
    google_id TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    is_available BOOLEAN NOT NULL,
    role TEXT NOT NULL,
    session_id TEXT NOT NULL,
    session_expired_at TIMESTAMPTZ NOT NULL,
    last_signin_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE requests (
    id TEXT PRIMARY KEY NOT NULL,
    requester_user_id TEXT NOT NULL,
    judger_user_id TEXT,
    status TEXT NOT NULL,
    project_id TEXT NOT NULL,
    iam_roles TEXT NOT NULL,
    reason TEXT NOT NULL,
    requested_at TIMESTAMPTZ NOT NULL,
    expired_at TIMESTAMPTZ NOT NULL,
    judged_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (requester_user_id) REFERENCES users(id),
    FOREIGN KEY (judger_user_id) REFERENCES users(id)
);

CREATE TABLE invitations (
    id TEXT PRIMARY KEY NOT NULL,
    inviter_user_id TEXT NOT NULL,
    invitee_mail TEXT NOT NULL UNIQUE,
    invitee_role TEXT NOT NULL,
    invited_at TIMESTAMPTZ NOT NULL,
    expired_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (inviter_user_id) REFERENCES users(id)
);
