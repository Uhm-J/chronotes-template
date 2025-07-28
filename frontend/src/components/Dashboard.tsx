import React from 'react';
import { Button } from './ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './ui/card';
import { User, Mail, Settings } from 'lucide-react';
import type { User as UserType } from '../types';

interface DashboardProps {
  user: UserType;
  onViewProfile: () => void;
}

export const Dashboard = ({ user, onViewProfile }: DashboardProps) => {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-center">
        <Card className="w-full max-w-2xl">
          <CardHeader className="text-center">
            <CardTitle className="text-4xl font-bold">Welcome!</CardTitle>
            <CardDescription>
              You're successfully authenticated
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            <div className="flex items-center space-x-4 p-4 bg-muted rounded-lg">
              <div className="flex-shrink-0">
                <User className="h-8 w-8 text-primary" />
              </div>
              <div className="flex-1 space-y-1">
                <p className="text-sm font-medium leading-none">
                  {user?.name}
                </p>
                <div className="flex items-center space-x-2">
                  <Mail className="h-4 w-4 text-muted-foreground" />
                  <p className="text-sm text-muted-foreground">
                    {user?.email}
                  </p>
                </div>
              </div>
            </div>
            
            <div className="grid gap-4 md:grid-cols-2">
              <Card>
                <CardHeader>
                  <CardTitle className="text-lg">Frontend</CardTitle>
                  <CardDescription>React with shadcn/ui</CardDescription>
                </CardHeader>
                <CardContent>
                  <p className="text-sm text-muted-foreground">
                    Modern React application with TypeScript, hooks, providers, and component library.
                  </p>
                </CardContent>
              </Card>
              
              <Card>
                <CardHeader>
                  <CardTitle className="text-lg">Backend</CardTitle>
                  <CardDescription>Go with Google OAuth2</CardDescription>
                </CardHeader>
                <CardContent>
                  <p className="text-sm text-muted-foreground">
                    Minimal Go backend with Google OAuth2 authentication and PSQL database.
                  </p>
                </CardContent>
              </Card>
            </div>
            
            <div className="flex justify-center space-x-4">
              <Button 
                variant="outline" 
                onClick={onViewProfile}
                className="flex items-center space-x-2"
              >
                <Settings className="h-4 w-4" />
                <span>View Profile</span>
              </Button>
              <Button variant="outline" onClick={() => window.location.reload()}>
                Refresh Page
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}; 