// @ts-check
import { test, expect, chromium } from '@playwright/test';
import * as utils from './utils/utils.js';
import * as config from './config.ts';

test('cant login after user is not available', async () => {
    test.setTimeout(10000);

    const admin = await utils.createUser({ role: "admin" });
    console.log(admin);
    const browser = await chromium.launch();
    const context = await browser.newContext();
    utils.setCookie("token", admin.sessionId, context);

    const page = await context.newPage();

    // move to request form page (can access)
    await page.goto(`${config.url}/request-form`);

    // set unavailable
    admin.isAvailable = false;
    await utils.saveUser(admin);

    const response = await page.goto(`${config.url}/request-form`);
    expect(response?.status()).toBe(403);
    expect(await page.textContent('h1')).toBe('Account Unavailable');
});
