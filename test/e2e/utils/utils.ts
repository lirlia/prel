import postgres from 'postgres'
import { v4 as uuidv4 } from 'uuid';
import * as config from '../config.ts';

const sql = postgres('postgres://postgres:password@localhost:5432/prel_e2e')
import * as query from '../db/query_ts_sql.js';
import { BrowserContext, expect } from '@playwright/test';

export type Role = 'requester' | 'judger' | 'admin';
export type JudgeAction = 'approve' | 'reject' | 'delete';
export type User = {
    id: string,
    googleId: string,
    email: string,
    role: Role,
    isAvailable: boolean,
    sessionId: string,
    sessionExpiredAt: Date,
    lastSigninAt: Date
};

export const createUser = async ({ role, sessionId, email, expiredAt, isAvailable }: { role: Role, sessionId?: string, email?: string, expiredAt?: Date, isAvailable?: boolean }): Promise<User> => {
    const now = new Date();
    email = email || randomStr(10) + '@example.com';
    sessionId = sessionId || uuidv4();
    expiredAt = expiredAt || new Date(now.setHours(now.getHours() + 1));
    isAvailable = isAvailable || true;
    const userData: query.UpsertUserArgs = {
        id: uuidv4(),
        googleId: uuidv4(),
        email: email,
        role: role,
        isAvailable: isAvailable,
        sessionId: sessionId,
        sessionExpiredAt: expiredAt,
        lastSigninAt: now
    };

    await query.upsertUser(sql, userData);
    return {
        id: userData.id,
        googleId: userData.googleId,
        email: userData.email,
        role: role,
        isAvailable: userData.isAvailable,
        sessionId: userData.sessionId,
        sessionExpiredAt: userData.sessionExpiredAt,
        lastSigninAt: userData.lastSigninAt
    };
};

export const saveUser = async (user: User) => {
    const userData: query.UpsertUserArgs = {
        id: user.id,
        googleId: user.googleId,
        email: user.email,
        role: user.role,
        isAvailable: user.isAvailable,
        sessionId: user.sessionId,
        sessionExpiredAt: user.sessionExpiredAt,
        lastSigninAt: new Date()
    };

    await query.upsertUser(sql, userData);
}

export const deleteUserByEmail = async (email: string) => {
    await query.deleteUserByEmail(sql, { email });
}

export const deleteAllIamRoleFilteringRules = async () => {
    await query.deleteAllIamRoleFilteringRules(sql);
}

export const setCookie = async (key: string, val: string, ctx: BrowserContext) => {
    await ctx.addCookies([
        {
            name: key,
            value: val,
            domain: config.domain,
            path: "/",
        }
    ])
};

const randomStr = (length) => {
    const s = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    return Array.from(Array(length)).map(() => s[Math.floor(Math.random() * s.length)]).join('');
}
