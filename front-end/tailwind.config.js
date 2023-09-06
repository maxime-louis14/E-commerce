/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    screens: {
    },
    fontFamily: {
    },
    extend: {
    },
    plugins: [
      // ...
      require("@tailwindcss/aspect-ratio")
    ]
  }
};
