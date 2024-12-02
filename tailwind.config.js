/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./components/**/*.templ"],
  theme: {
    extend: {
      colors: {
        secondary: "#61758A",
        divider: "#DBE0E5",
      },
      container: {
        screens: {
          DEFAULT: "960px",
        },
      },
    },
  },
  plugins: [],
};
