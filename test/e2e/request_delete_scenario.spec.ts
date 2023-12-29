// @ts-check
import { test, expect, chromium } from '@playwright/test';
import * as utils from './utils/utils.js';
import * as helper from './utils/test_helper.ts';
import * as config from './config.ts';

test('send request and delete it by admin', async () => {
    const projectId = 'prel-test';
    const roles = ['roles/spanner.admin', 'roles/bigquery.admin'];
    const now = new Date();
    const reason = 'test reason';

    const expiredAt = new Date(now.setHours(now.getHours() + 1));
    const user = await utils.createUser({ role: "requester", expiredAt: expiredAt });

    // await createUser(userData)
    const browser = await chromium.launch();
    const context = await browser.newContext();
    utils.setCookie("token", user.sessionId, context);

    const res = helper.addRequest({
        projectId: projectId,
        email: user.email,
        roles: roles,
        reason: reason,
        status: 'pending',
        ctx: context,
    })

    const page = (await res).page;

    // move to pending request page
    Promise.all([
        await page.waitForURL(`${config.url}/request/*`),
        expect(await page.textContent('h2')).toBe('Pending Requests'),
    ]);

    // reject request by judger
    const requestUrl = page.url();
    const admin = await utils.createUser({ role: "admin", expiredAt: expiredAt });
    const adminContext = await browser.newContext();
    utils.setCookie("token", admin.sessionId, adminContext);

    const adminRes = await helper.judgeRequestInSpecificPage({
        judgeAction: 'delete',
        requestUrl: requestUrl,
        projectId: projectId,
        email: user.email,
        roles: roles,
        reason: reason,
        ctx: adminContext,
    })

    const _ = adminRes.page;
    const response = await page.goto(requestUrl);
    expect(response?.status()).toBe(404);
});
