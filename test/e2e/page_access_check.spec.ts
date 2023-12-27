// @ts-check
import { test, expect, chromium } from '@playwright/test';
import * as utils from './utils/utils.js';
import * as config from './config.ts';

type Role = 'unauthenticated' | 'requester' | 'judger' | 'admin';

type TestCase = {
    role: Role;
    url: string;
    expectedStatus: number;
    expectedText?: string;
};

const testCases: TestCase[] = [
    { role: 'unauthenticated', url: `${config.url}/request-form`, expectedStatus: 401 },
    { role: 'requester', url: `${config.url}/request-form`, expectedStatus: 200 },
    { role: 'judger', url: `${config.url}/request-form`, expectedStatus: 200 },
    { role: 'admin', url: `${config.url}/request-form`, expectedStatus: 200 },

    { role: 'unauthenticated', url: `${config.url}/request`, expectedStatus: 401 },
    { role: 'requester', url: `${config.url}/request`, expectedStatus: 200 },
    { role: 'judger', url: `${config.url}/request`, expectedStatus: 200 },
    { role: 'admin', url: `${config.url}/request`, expectedStatus: 200 },

    // 401: unauthenticated
    // 403: authenticated but not allowed
    { role: 'unauthenticated', url: `${config.url}/admin/request`, expectedStatus: 401 },
    { role: 'requester', url: `${config.url}/admin/request`, expectedStatus: 403 },
    { role: 'judger', url: `${config.url}/admin/request`, expectedStatus: 403 },
    { role: 'admin', url: `${config.url}/admin/request`, expectedStatus: 200 },

    { role: 'unauthenticated', url: `${config.url}/admin/user`, expectedStatus: 401 },
    { role: 'requester', url: `${config.url}/admin/user`, expectedStatus: 403 },
    { role: 'judger', url: `${config.url}/admin/user`, expectedStatus: 403 },
    { role: 'admin', url: `${config.url}/admin/user`, expectedStatus: 200 },
];

for (const testCase of testCases) {
    test(`[${testCase.role}] access test: ${testCase.url}`, async () => {
        test.setTimeout(5000);
        const browser = await chromium.launch();
        const context = await browser.newContext();

        if (testCase.role !== 'unauthenticated') {
            const user = await utils.createUser({ role: testCase.role });
            utils.setCookie("token", user.sessionId, context);
        }

        const page = await context.newPage();
        const response = await page.goto(testCase.url);
        expect(response?.status()).toBe(testCase.expectedStatus);

        if (testCase.expectedText) {
            expect(await page.textContent('h1')).toBe(testCase.expectedText);
        }

        await browser.close();
    });
}
