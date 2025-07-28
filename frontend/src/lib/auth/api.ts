import { apiClient } from '../api/client';
import { API_ENDPOINTS } from '../../constants/app';
import type { User, UserResponse } from '../../types';

export const authApi = {
  /**
   * Get current user information
   */
  getCurrentUser: async (): Promise<User | null> => {
    const userResponse = await apiClient.get<UserResponse>(API_ENDPOINTS.AUTH.ME);
    return userResponse.data || null;
  },

  /**
   * Initiate Google OAuth login
   * Note: This will redirect the user to Google's OAuth page
   */
  initiateGoogleLogin: (): void => {
    window.location.href = API_ENDPOINTS.AUTH.LOGIN;
  },

  /**
   * Logout the current user
   */
  logout: async (): Promise<void> => {
    return apiClient.post<void>(API_ENDPOINTS.AUTH.LOGOUT);
  },

  /**
   * Check if user is authenticated by trying to fetch user data
   */
  checkAuthStatus: async (): Promise<{ isAuthenticated: boolean; user?: User }> => {
    try {
      const user = await authApi.getCurrentUser();
      return { isAuthenticated: true, user: user };
    } catch (error) {
      return { isAuthenticated: false };
    }
  },
}; 