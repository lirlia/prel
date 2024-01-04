import { defineConfig, devices } from "@playwright/test";

export default defineConfig({
    // Look for test files in the "tests" directory, relative to this configuration file.
    testDir: ".",

    // Run all tests in parallel.
    fullyParallel: true,

    // Fail the build on CI if you accidentally left test.only in the source code.
    forbidOnly: !!process.env.CI,

    // Retry on CI only.
    retries: process.env.CI ? 2 : 0,

    // Opt out of parallel tests on CI.
    workers: 5,

    // Reporter to use
    reporter: "html",

    use: {
        // Base URL to use in actions like `await page.goto('/')`.
        baseURL: "http://127.0.0.1:8182",

        // Collect trace when retrying the failed test.
        trace: "on-first-retry",
    },
    // Configure projects for major browsers.
    projects: [
        {
            name: "chromium",
            use: { ...devices["Desktop Chrome"] },
        },
    ],
    // Run your local dev server before starting the tests.
    webServer: {
        command: "(cd $(git rev-parse --show-toplevel); make run-e2e 1>&2)",
        url: "http://127.0.0.1:8182",
    },

    timeout: 15000,
});
