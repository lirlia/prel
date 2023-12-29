// @ts-check
import { test, expect, chromium } from '@playwright/test';
import * as utils from './utils/utils.js';
import * as config from './config.ts';


test('can login after invite', async () => {

    // add first user
    // next user must be invited by this user
    const inviter = await utils.createUser({ role: "admin" });
    console.log(inviter);

    const inviteeEmail = "invitee@example.com";

    const browser = await chromium.launch();
    const ctx = await browser.newContext();
    const state = "state-123";
    utils.setCookie("state", state, ctx)
    const page = await ctx.newPage();
    const callbackUrl = `${config.url}/auth/google/callback?code=code-123&state=${state}`

    await utils.deleteUserByEmail(inviteeEmail).then(async () => {
        const response = await ctx.request.get(callbackUrl, { maxRedirects: 0 });
        console.log(response);

        // not invited yet
        expect(response.status()).toBe(403);
    })


    const inviterCtx = await browser.newContext();
    utils.setCookie("token", inviter.sessionId, inviterCtx);
    const inviterPage = await inviterCtx.newPage();

    // inviter invite user
    const inviteResponsePromise = inviterPage.waitForResponse(response =>
        response.url().includes('/api/invitations') && response.status() === 204
    );

    Promise.all([
        await inviterPage.goto(`${config.url}/admin/user`),
        await inviterPage.waitForSelector('h2'),
        expect(await inviterPage.textContent('h2')).toBe('User Management'),
        await inviterPage.fill('#inviteeEmail', inviteeEmail),
        await inviterPage.click('#invite'),
    ]);

    expect(await inviteResponsePromise).toBeTruthy();

    // login as invitee
    const callbackRes = await ctx.request.get(callbackUrl, { maxRedirects: 10 });
    Promise.all([
        expect(callbackRes?.status()).toBe(200),
        console.log(callbackRes),
        expect(await callbackRes?.text()).toContain('IAM Role Request Form'),
    ]);
});
