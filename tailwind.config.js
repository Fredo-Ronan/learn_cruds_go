/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.html", // Semua file html di folder templates
    "./assets/js/**/*.js",    // File JS eksternal Anda
    "./main.go",              // Jika Anda menulis HTML di main.go
    "./**/*.go",              // Scan seluruh file .go untuk class string
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}