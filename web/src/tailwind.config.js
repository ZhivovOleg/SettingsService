/** @type {import('tailwindcss').Config} */
export default {
  content: [],
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        "primary-color": "var(--primary-color)",
        "secondary-color": "var(--secondary-color)",
        "select-color": "var(--select-color)",
        "hover-color": "var(--hover-color)",
        "gray-color": "var(--gray)",
        "error-color": "var(--error)",
      },
    },
  },
  plugins: [],
}

