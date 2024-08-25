/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["**/*.templ"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["dracula"],
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
};
