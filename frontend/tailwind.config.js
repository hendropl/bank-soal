/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}", // Pastikan path ini benar
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Poppins', 'sans-serif'],
      },
      colors: {
        'primary': '#6C63FF',
        'background': '#F0F2FF',
        'dark-text': '#1E293B',
        'medium-text': '#475569',
        'exam-green': '#28A745',
      },
    },
  },
  plugins: [],
}