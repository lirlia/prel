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
    Promise.all([
        await page.goto(config.url + "/request-form", { waitUntil: 'networkidle' }),
        await page.waitForSelector('h2'),
        expect(await page.textContent('h2')).toBe('IAM Role Request Form'),
        expect(await page.textContent('#select2-email-container')).toBe(req.email),
        // open project id dropdown
        await page.evaluate((projectId) => {
            $('#project_id').val(projectId);
            $('#project_id').trigger('change');
        }, req.projectId),
        await page.waitForResponse(response =>
            response.url().includes('/api/iam-roles') && response.status() === 200
        ),
        expect(await page.textContent('#select2-project_id-container')).toContain(req.projectId),
        // open role dropdown
        await page.evaluate((roles) => {
            $('#role').val(roles);
            $('#role').trigger('change');
        }, req.roles),
        //
        // fill reason
        //
        await page.fill('#reason', req.reason),
    ]);
    Promise.all([
        expect(await page.textContent('#role')).toContain(req.roles[0]),
        expect(await page.textContent('#role')).toContain(req.roles[1]),
        // submit
        await page.click('#submit-request')
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
    Promise.all([
        await page.click(btn),
        await page.waitForResponse(response =>
            response.url().includes('/api/requests/') && response.status() === 204
        ),
    ]);
    const navigationPromise = page.waitForNavigation();
    Promise.all([
        page.on('dialog', async dialog => {
            await dialog.accept();
        }),
        await page.waitForLoadState('domcontentloaded'),
        await navigationPromise,
    ]);
    return { page: page };
}
