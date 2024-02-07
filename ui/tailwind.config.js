/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{ts,tsx}'],
  plugins: [require('@headlessui/tailwindcss')],
  theme: {
    extend: {
      fontFamily: { futura: ['futura'], georgia: ['georgia'], dm: ['dm-sans'] },
    },
    screens: {
      // https://en.wikipedia.org/wiki/Graphics_display_resolution
      vga: '640px',
      xga: '1024px',
      hd: '1280px',
      fhd: '1920px',
      uhd: '3840px',
    },
  },
}
