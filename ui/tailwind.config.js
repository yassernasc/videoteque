/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.js'],
  plugins: [require('@headlessui/tailwindcss')],
  theme: {
    fontFamily: { futura: ['futura'], georgia: ['georgia'], sans: ['dm-sans'] },
  },
}
