// @ts-check
import { test, expect, chromium, ElementHandle } from "@playwright/test";
import * as utils from "./utils/utils.js";
import * as config from "./config.ts";

test("can set notification message", async () => {
    const admin = await utils.createUser({ role: "admin" });
    console.log(admin);

    const browser = await chromium.launch();
    const ctx = await browser.newContext();

    utils.setCookie("token", admin.sessionId, ctx);
    const page = await ctx.newPage();

    page.goto(`${config.url}/admin/setting`);

    // Request Notification Message
    expect(await page.textContent("h2")).toBe("Setting");
    await page.fill("#notification-message-for-request", "test request notification message");
    await page.click('[data-related-field="notification-message-for-request"]');

    // Judge Notification Message
    await page.fill("#notification-message-for-judge", "test judge notification message");
    await page.click('[data-related-field="notification-message-for-judge"]');

    // check
    await page.reload();
    expect(await page.textContent("#notification-message-for-request")).
        toBe("test request notification message");

    expect(await page.textContent("#notification-message-for-judge")).
        toBe("test judge notification message");
});
