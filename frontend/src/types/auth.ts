export interface User {
  id?: string;
  name: string;
  email: string;
  avatar?: string;
}

export interface UserResponse {
  data?: User;
  
}

export interface AuthState {
  user: User | null;
  isLoading: boolean;
  isAuthenticated: boolean;
}

export interface LoginResponse {
  user: User;
  message?: string;
}

export interface AuthError {
  message: string;
  code?: string;
} 