// @ts-check
const { defineConfig, devices } = require('@playwright/test')

const url = 'http://localhost:1201'

module.exports = defineConfig({
  testDir: './tests',
  fullyParallel: true,
  use: { baseURL: url, trace: 'off' },
  projects: [
    { name: 'chrome', use: { ...devices['Desktop Chrome'] } },
    { name: 'firefox', use: { ...devices['Desktop Firefox'] } },
    { name: 'safari', use: { ...devices['Desktop Safari'] } },
  ],
  webServer: { command: 'make test-ui --directory ..', url },
})
