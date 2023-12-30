// @ts-check
import { test, expect, chromium, ElementHandle } from '@playwright/test';
import * as utils from './utils/utils.js';
import * as config from './config.ts';

test('after filter role, request form role is filtered', async () => {

    await utils.deleteAllIamRoleFilteringRules();

    // add first user
    // next user must be invited by this user
    const admin = await utils.createUser({ role: "admin" });
    console.log(admin);

    const browser = await chromium.launch();
    const ctx = await browser.newContext();

    utils.setCookie("token", admin.sessionId, ctx);
    const page = await ctx.newPage();

    page.goto(`${config.url}/admin/iam-role-filtering`);

    expect(await page.textContent('h2')).toBe('IAM Role Filtering');
    await page.getByPlaceholder('Enter a keyword between 3 and').fill('spanner');
    await page.getByRole('button', { name: 'Add Keyword' }).click();
    // Check if 'spanner' is added to the <td>
    expect(await (await page.waitForSelector('td')).innerText()).toContain('spanner');

    // Add 'bigquery' as a keyword
    await page.getByPlaceholder('Enter a keyword between 3 and').fill('bigquery');
    await page.getByRole('button', { name: 'Add Keyword' }).click();

    // Wait for the table row to increase
    await page.waitForFunction(() => document.querySelectorAll('tr').length === 3);

    // check tr[1].td[0] is 'bigquery' and tr[2].td[0] is 'spanner'
    const trs = await page.$$('tr');
    const td = await trs[1].$('td');
    expect(td).not.toBeNull();
    expect(await td?.innerText()).toContain('bigquery');

    const td2 = await trs[2].$('td');
    expect(td2).not.toBeNull();
    expect(await td2?.innerText()).toContain('spanner');

    // add second user
    const requester = await utils.createUser({ role: "requester" });
    console.log(requester);

    utils.setCookie("token", requester.sessionId, ctx);
    const requesterPage = await ctx.newPage();
    await requesterPage.goto(`${config.url}/request-form`);

    // check iam role filtering
    const iamResponsePromise = page.waitForResponse(response =>
        response.url().includes('/api/iam-roles')
        && response.status() === 200
        // request json.iamRoles has only spanner / bigquery names roles
        && response.json().then(json => json.iamRoles.map(
            (iamRole: string) => iamRole.includes('spanner') || iamRole.includes('bigquery')).every((v: boolean) => v))
    );

    const roles = ['roles/spanner.admin', 'roles/bigquery.admin'];
    const projectId = 'prel-test';
    Promise.all([
        await page.goto(`${config.url}/request-form`),
        await page.getByRole('textbox', { name: 'Select Project' }).click(),
        await page.getByRole('searchbox').nth(1).fill('prel'),

        await page.getByRole('option', { name: 'prel-test' }).click(),
        expect(await iamResponsePromise).toBeTruthy(),
        expect(await page.textContent('#select2-project_id-container')).toContain(projectId),
        // open role dropdown
        await page.evaluate((roles) => {
            $('#role').val(roles);
            $('#role').trigger('change');
        }, roles),
        await page.getByRole('textbox', { name: 'minutes' }).click(),
        await page.getByRole('option', { name: '10 minutes' }).click(),
        await page.getByLabel('Reason').click(),
        await page.getByLabel('Reason').fill('test reason'),
        await page.getByRole('button', { name: 'Request' }).click(),
    ]);

    await page.waitForURL(`${config.url}/request/*`);
    expect(await page.textContent('h2')).toBe('Pending Requests');

    await utils.deleteAllIamRoleFilteringRules();
});
