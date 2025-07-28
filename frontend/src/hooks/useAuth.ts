import { useState, useEffect } from 'react';
import { authApi } from '../lib/auth/api';
import type { User, AuthState, UserResponse } from '../types';

export const useAuth = () => {
  const [state, setState] = useState<AuthState>({
    user: null,
    isLoading: true,
    isAuthenticated: false,
  });

  const checkAuth = async () => {
    try {
      setState(prev => ({ ...prev, isLoading: true }));
      const { isAuthenticated, user } = await authApi.checkAuthStatus();
      setState({
        user: user || null,
        isLoading: false,
        isAuthenticated,
      });
    } catch (error) {
      setState({
        user: null,
        isLoading: false,
        isAuthenticated: false,
      });
    }
  };

  const login = () => {
    authApi.initiateGoogleLogin();
  };

  const logout = async () => {
    try {
      await authApi.logout();
      setState({
        user: null,
        isLoading: false,
        isAuthenticated: false,
      });
    } catch (error) {
      console.error('Logout error:', error);
    }
  };

  const refreshUser = async () => {
    try {
      const user = await authApi.getCurrentUser();
      setState(prev => ({
        ...prev,
        user: user?.data || null,
        isAuthenticated: true,
      }));
    } catch (error) {
      setState(prev => ({
        ...prev,
        user: null,
        isAuthenticated: false,
      }));
    }
  };

  useEffect(() => {
    checkAuth();
  }, []);

  return {
    ...state,
    login,
    logout,
    refreshUser,
    checkAuth,
  };
}; 