// @ts-check
import { test, expect, chromium } from '@playwright/test';
import * as utils from './utils/utils.js';
import * as config from './config.ts';


test('cant login after invite', async () => {

    // add first user
    // next user must be invited by this user
    const inviter = await utils.createUser({ role: "admin" });
    console.log(inviter);

    const inviteeEmail = "invitee@example.com";
    await utils.deleteUserByEmail(inviteeEmail);

    const browser = await chromium.launch();
    const ctx = await browser.newContext();
    const state = "state-123";
    utils.setCookie("state", state, ctx)
    const page = await ctx.newPage();
    const callbackUrl = `${config.url}/auth/google/callback?code=code-123&state=${state}`

    const response = await page.goto(callbackUrl);

    // not invited yet
    expect(response?.status()).toBe(403);

    const inviterCtx = await browser.newContext();
    utils.setCookie("token", inviter.sessionId, inviterCtx);
    const inviterPage = await inviterCtx.newPage();

    // inviter invite user
    Promise.all([
        await inviterPage.goto(`${config.url}/admin/user`),
        await inviterPage.waitForSelector('h2'),
        expect(await inviterPage.textContent('h2')).toBe('User Management'),
        await inviterPage.fill('#inviteeEmail', inviteeEmail),
        await inviterPage.click('#invite'),
    ]);

    await inviterPage.waitForResponse(response =>
        response.url().includes('/api/invitations') && response.status() === 204
    );

    // login as invitee
    const callbackRes = await page.goto(callbackUrl);
    Promise.all([
        expect(callbackRes?.status()).toBe(200),
        expect(await page.textContent('h2')).toBe('IAM Role Request Form'),
    ]);
});
