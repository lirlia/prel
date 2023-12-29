// @ts-check
import { test, expect, chromium } from '@playwright/test';
import * as utils from './utils/utils.js';
import * as helper from './utils/test_helper.ts';
import * as config from './config.ts';

test('send request and approve by judger', async () => {
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
    console.log(user);
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

    // approve request by judger
    const requestUrl = page.url();
    const judger = await utils.createUser({ role: "judger", expiredAt: expiredAt });
    const judgerContext = await browser.newContext();
    utils.setCookie("token", judger.sessionId, judgerContext);

    const judgeRes = await helper.judgeRequestInSpecificPage({
        judgeAction: 'approve',
        requestUrl: requestUrl,
        projectId: projectId,
        email: user.email,
        roles: roles,
        reason: reason,
        ctx: judgerContext,
    })

    await judgeRes.page.locator('.status:has-text("approved")').waitFor();
});
