// @ts-check
const { test, expect } = require('@playwright/test')

test.skip('late subtitle and another text appears', async ({
  context,
  page,
}) => {
  const settings = await context.newPage()

  await page.goto('/')
  await settings.goto('/settings')

  const player = page.getByRole('main')
  const subtitle = page.getByTitle('subtitle')
  const tooLateBtn = settings.getByRole('button', { name: 'too late' })

  await player.click() // play
  await expect(subtitle).toHaveText('Hildy!')
  await tooLateBtn.click()
  await expect(subtitle).toContainText('lord of the universe')
})
