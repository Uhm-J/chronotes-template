/**
 * Common validation utilities
 */

export const isEmail = (email: string): boolean => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
};

export const isUrl = (url: string): boolean => {
  try {
    new URL(url);
    return true;
  } catch {
    return false;
  }
};

export const isPhoneNumber = (phone: string): boolean => {
  const phoneRegex = /^\+?[\d\s\-\(\)]{10,}$/;
  return phoneRegex.test(phone);
};

export const isStrongPassword = (password: string): boolean => {
  // At least 8 characters, 1 uppercase, 1 lowercase, 1 number
  const strongPasswordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d@$!%*?&]{8,}$/;
  return strongPasswordRegex.test(password);
};

export const isNotEmpty = (value: any): boolean => {
  if (value == null) return false;
  if (typeof value === 'string') return value.trim().length > 0;
  if (Array.isArray(value)) return value.length > 0;
  if (typeof value === 'object') return Object.keys(value).length > 0;
  return true;
};

export const isNumeric = (value: string): boolean => {
  return !isNaN(Number(value)) && !isNaN(parseFloat(value));
};

export const isInRange = (value: number, min: number, max: number): boolean => {
  return value >= min && value <= max;
};

export const isValidDate = (date: any): boolean => {
  return date instanceof Date && !isNaN(date.getTime());
};

// Validation result type
export interface ValidationResult {
  isValid: boolean;
  message?: string;
}

// Validation function type
export type ValidatorFn = (value: any) => ValidationResult;

// Common validators
export const validators = {
  required: (message = 'This field is required'): ValidatorFn => 
    (value) => ({
      isValid: isNotEmpty(value),
      message: isNotEmpty(value) ? undefined : message,
    }),

  email: (message = 'Please enter a valid email address'): ValidatorFn =>
    (value) => ({
      isValid: isEmail(value),
      message: isEmail(value) ? undefined : message,
    }),

  minLength: (min: number, message?: string): ValidatorFn =>
    (value) => {
      const isValid = typeof value === 'string' && value.length >= min;
      return {
        isValid,
        message: isValid ? undefined : message || `Minimum ${min} characters required`,
      };
    },

  maxLength: (max: number, message?: string): ValidatorFn =>
    (value) => {
      const isValid = typeof value === 'string' && value.length <= max;
      return {
        isValid,
        message: isValid ? undefined : message || `Maximum ${max} characters allowed`,
      };
    },
}; 