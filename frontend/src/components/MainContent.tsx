import React, { useState } from 'react';
import { Dashboard } from './Dashboard';
import { Profile } from './Profile';
import type { User as UserType } from '../types';

interface MainContentProps {
  user: UserType;
  onLogout: () => void;
}

export const MainContent = ({ user, onLogout }: MainContentProps) => {
  const [showProfile, setShowProfile] = useState(false);

  const handleViewProfile = () => setShowProfile(true);
  const handleBackToHome = () => setShowProfile(false);

  if (showProfile) {
    return (
      <Profile 
        user={user} 
        onLogout={onLogout} 
        onBack={handleBackToHome} 
      />
    );
  }

  return (
    <Dashboard 
      user={user} 
      onViewProfile={handleViewProfile} 
    />
  );
}; 