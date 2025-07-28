// Application constants
export const APP_CONFIG = {
  NAME: "Chronotes Template",
  DESCRIPTION:
    "A modern full-stack template with Go backend and React frontend",
  VERSION: "1.0.0",
} as const;

// API endpoints
export const API_ENDPOINTS = {
  AUTH: {
    ME: "/v1/profile",
    LOGIN: "/v1/auth/google/login",
    LOGOUT: "/v1/auth/logout",
  },
} as const;

// Local storage keys
export const STORAGE_KEYS = {
  THEME: "chronotes-theme",
  USER_PREFERENCES: "chronotes-user-prefs",
} as const;

// UI constants
export const UI_CONFIG = {
  ANIMATION_DURATION: 200,
  TOAST_DURATION: 3000,
  DEBOUNCE_DELAY: 300,
} as const;

// Breakpoints (matching Tailwind CSS)
export const BREAKPOINTS = {
  SM: "640px",
  MD: "768px",
  LG: "1024px",
  XL: "1280px",
  "2XL": "1536px",
} as const;
