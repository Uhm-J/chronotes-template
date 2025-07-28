import React from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './ui/card';
import { Button } from './ui/button';
import { User, Mail, Calendar, Shield, LogOut } from 'lucide-react';
import type { User as UserType } from '../types';

interface ProfileProps {
  user: UserType;
  onLogout: () => void;
  onBack: () => void;
}

export const Profile = ({ user, onLogout, onBack }: ProfileProps) => {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-center">
        <Card className="w-full max-w-2xl">
          <CardHeader className="text-center">
            <div className="mx-auto mb-4 flex h-20 w-20 items-center justify-center rounded-full bg-primary/10">
              <User className="h-10 w-10 text-primary" />
            </div>
            <CardTitle className="text-3xl font-bold">{user.name}</CardTitle>
            <CardDescription>Your Profile Information</CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            {/* User Information */}
            <div className="space-y-4">
              <div className="flex items-center space-x-3 p-4 bg-muted rounded-lg">
                <Mail className="h-5 w-5 text-primary" />
                <div className="flex-1">
                  <p className="text-sm font-medium text-muted-foreground">Email Address</p>
                  <p className="text-base font-semibold">{user.email}</p>
                </div>
              </div>

              <div className="flex items-center space-x-3 p-4 bg-muted rounded-lg">
                <Shield className="h-5 w-5 text-primary" />
                <div className="flex-1">
                  <p className="text-sm font-medium text-muted-foreground">User ID</p>
                  <p className="text-base font-semibold">{user.id || 'N/A'}</p>
                </div>
              </div>

              {user.avatar && (
                <div className="flex items-center space-x-3 p-4 bg-muted rounded-lg">
                  <User className="h-5 w-5 text-primary" />
                  <div className="flex-1">
                    <p className="text-sm font-medium text-muted-foreground">Avatar</p>
                    <img 
                      src={user.avatar} 
                      alt="User avatar" 
                      className="h-10 w-10 rounded-full"
                    />
                  </div>
                </div>
              )}
            </div>

            {/* Account Status */}
            <Card>
              <CardHeader>
                <CardTitle className="text-lg flex items-center space-x-2">
                  <Shield className="h-5 w-5" />
                  <span>Account Status</span>
                </CardTitle>
              </CardHeader>
              <CardContent>
                <div className="flex items-center space-x-2">
                  <div className="h-2 w-2 rounded-full bg-green-500"></div>
                  <span className="text-sm font-medium">Active</span>
                </div>
                <p className="text-sm text-muted-foreground mt-1">
                  Your account is active and authenticated via Google OAuth
                </p>
              </CardContent>
            </Card>

            {/* Actions */}
            <div className="flex flex-col space-y-3">
              <Button 
                variant="outline" 
                onClick={onBack}
                className="w-full"
              >
                Back to Home
              </Button>
              
              <Button 
                variant="outline" 
                onClick={() => window.location.reload()}
                className="w-full"
              >
                Refresh Profile
              </Button>
              
              <Button 
                variant="destructive" 
                onClick={onLogout}
                className="w-full flex items-center space-x-2"
              >
                <LogOut className="h-4 w-4" />
                <span>Sign Out</span>
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}; 