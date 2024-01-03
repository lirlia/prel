import { Sql } from "postgres";

export const upsertUserQuery = `-- name: UpsertUser :exec
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
    last_signin_at = excluded.last_signin_at`;

export interface UpsertUserArgs {
    id: string;
    googleId: string;
    email: string;
    role: string;
    isAvailable: boolean;
    sessionId: string;
    sessionExpiredAt: Date;
    lastSigninAt: Date;
}

export async function upsertUser(sql: Sql, args: UpsertUserArgs): Promise<void> {
    await sql.unsafe(upsertUserQuery, [args.id, args.googleId, args.email, args.role, args.isAvailable, args.sessionId, args.sessionExpiredAt, args.lastSigninAt]);
}

export const deleteUserByEmailQuery = `-- name: DeleteUserByEmail :exec
DELETE FROM users
WHERE email = $1`;

export interface DeleteUserByEmailArgs {
    email: string;
}

export async function deleteUserByEmail(sql: Sql, args: DeleteUserByEmailArgs): Promise<void> {
    await sql.unsafe(deleteUserByEmailQuery, [args.email]);
}

export const deleteAllIamRoleFilteringRulesQuery = `-- name: DeleteAllIamRoleFilteringRules :exec
DELETE FROM iam_role_filtering_rules`;

export async function deleteAllIamRoleFilteringRules(sql: Sql): Promise<void> {
    await sql.unsafe(deleteAllIamRoleFilteringRulesQuery, []);
}
