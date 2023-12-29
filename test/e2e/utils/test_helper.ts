import { BrowserContext, Page, expect } from "@playwright/test"
import * as config from '../config.ts';
import * as utils from './utils.js';

type prelRequestInput = {
    projectId: string,
    email: string,
    roles: string[],
    reason: string,
    status: string,
    ctx: BrowserContext,
}

export type replRequestOutput = {
    page: Page,
}

// add request
export const addRequest = async (req: prelRequestInput): Promise<replRequestOutput> => {

    const page = await req.ctx.newPage();
    const iamResponsePromise = page.waitForResponse(response =>
        response.url().includes('/api/iam-roles') && response.status() === 200
    );
    Promise.all([
        await page.goto(`${config.url}/request-form`),
        await page.getByRole('textbox', { name: 'Select Project' }).click(),
        await page.getByRole('searchbox').nth(1).fill('prel'),

        await page.getByRole('option', { name: 'prel-test' }).click(),
        expect(await iamResponsePromise).toBeTruthy(),
        expect(await page.textContent('#select2-project_id-container')).toContain(req.projectId),
        // open role dropdown
        await page.evaluate((roles) => {
            $('#role').val(roles);
            $('#role').trigger('change');
        }, req.roles),
        await page.getByRole('textbox', { name: 'minutes' }).click(),
        await page.getByRole('option', { name: '10 minutes' }).click(),
        await page.getByLabel('Reason').click(),
        await page.getByLabel('Reason').fill('test reason'),
        await page.getByRole('button', { name: 'Request' }).click(),
    ]);

    // move to pending request page
    Promise.all([
        await page.waitForURL(`${config.url}/request/*`),
        expect(await page.textContent('h2')).toBe('Pending Requests'),
    ]);

    return { page: page };
}

type prelJudgeRequestInput = {
    judgeAction: utils.JudgeAction,
    requestUrl: string,
    projectId: string,
    email: string,
    roles: string[],
    reason: string,
    ctx: BrowserContext,
}

type prelJudgeRequestOutput = {
    page: Page,
}

export const judgeRequestInSpecificPage = async (req: prelJudgeRequestInput): Promise<prelJudgeRequestOutput> => {
    const page = await req.ctx.newPage();


    Promise.all([
        await page.goto(req.requestUrl, { waitUntil: 'networkidle' }),
        await page.waitForSelector('h2'),
        expect(await page.textContent('h2')).toBe('Pending Requests'),
        expect(await page.textContent('.email')).toBe(req.email),
        expect(await page.textContent('.project-id')).toBe(req.projectId),
        req.roles.forEach(async (role, _) => {
            expect(await page.textContent('.iam-roles')).toContain(role);
        }),
        expect(await page.textContent('.reason')).toBe(req.reason),
        expect(await page.textContent('.status')).toBe('pending'),
    ]);

    let btn: string;
    switch (req.judgeAction) {
        case 'approve':
            btn = '#btn-approve';
            break;
        case 'reject':
            btn = '#btn-reject';
            break;
        case 'delete':
            btn = '#btn-delete';
            break;
        default:
            throw new Error('invalid judge status');
    }

    // wait api response
    const requestResponsePromise = page.waitForResponse(response =>
        response.url().includes('/api/requests/') && response.status() === 204
    )
    Promise.all([
        await page.click(btn),
        expect(await requestResponsePromise).toBeTruthy(),
    ]);

    Promise.all([
        page.on('dialog', async dialog => {
            await dialog.accept();
        }),
        await page.waitForLoadState('domcontentloaded'),
    ]);
    return { page: page };
}
