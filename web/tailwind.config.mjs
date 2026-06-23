/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#e6f7ff',
          100: '#ccefff',
          200: '#99dfff',
          300: '#66cfff',
          400: '#33bfff',
          500: '#00BFFF', // Deep Sky Blue
          600: '#0099cc',
          700: '#007399',
          800: '#004d66',
          900: '#002633',
          950: '#001319',
        },
        secondary: {
          50: '#eff6ff',
          100: '#dbeafe',
          200: '#bfdbfe',
          300: '#93c5fd',
          400: '#60a5fa',
          500: '#3b82f6',
          600: '#2563eb',
          700: '#1d4ed8',
          800: '#1e40af',
          900: '#1e3a8a',
          950: '#172554',
        },
        accent: {
          50: '#fdf4ff',
          100: '#fae8ff',
          200: '#f5d0fe',
          300: '#f0abfc',
          400: '#e879f9',
          500: '#d946ef',
          600: '#c026d3',
          700: '#a21caf',
          800: '#86198f',
          900: '#701a75',
          950: '#4a044e',
        },
        terminal: {
          dark: '#121212',
          light: '#f8f8f8',
          accent: '#00BFFF',
          border: '#2a2a2a',
          header: '#1e1e1e',
        },
        code: {
          bg: '#1e1e1e',
          text: '#d4d4d4',
          comment: '#6a9955',
          keyword: '#569cd6',
          string: '#ce9178',
          function: '#dcdcaa',
          variable: '#9cdcfe',
          number: '#b5cea8',
        },
      },
      fontFamily: {
        sans: ['Inter', 'sans-serif'],
        mono: ['JetBrains Mono', 'monospace'],
      },
      animation: {
        'terminal-cursor': 'blink 1s step-end infinite',
        'fade-in': 'fadeIn 0.5s ease-in-out',
        'slide-up': 'slideUp 0.5s ease-out',
        'slide-down': 'slideDown 0.3s ease-out',
        'pulse-slow': 'pulse 3s cubic-bezier(0.4, 0, 0.6, 1) infinite',
      },
      keyframes: {
        blink: {
          '0%, 100%': { opacity: 1 },
          '50%': { opacity: 0 },
        },
        fadeIn: {
          '0%': { opacity: 0 },
          '100%': { opacity: 1 },
        },
        slideUp: {
          '0%': { transform: 'translateY(10px)', opacity: 0 },
          '100%': { transform: 'translateY(0)', opacity: 1 },
        },
        slideDown: {
          '0%': { transform: 'translateY(-10px)', opacity: 0 },
          '100%': { transform: 'translateY(0)', opacity: 1 },
        },
      },
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-conic': 'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
        'hero-pattern': 'url("/hero-pattern.svg")',
      },
      boxShadow: {
        'glow': '0 0 15px 5px rgba(0, 191, 255, 0.3)',
        'glow-lg': '0 0 25px 10px rgba(0, 191, 255, 0.25)',
      },
      typography: {
        DEFAULT: {
          css: {
            maxWidth: '100ch',
            color: 'inherit',
            a: {
              color: '#00BFFF',
              '&:hover': {
                color: '#0099cc',
              },
            },
            'h1,h2,h3,h4': {
              color: 'inherit',
              fontWeight: '700',
            },
            code: {
              color: '#00BFFF',
              backgroundColor: '#e6f7ff',
              borderRadius: '0.25rem',
              padding: '0.2em 0.4em',
            },
          },
        },
      },
    },
  },
  plugins: [],
}
