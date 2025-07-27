import React from 'react';
import { AuthProvider, useAuthContext } from './providers/AuthProvider';
import { Layout } from './components/layout/Layout';
import { Header } from './components/layout/Header';
import { Button } from './components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './components/ui/card';
import { Chrome, User, Mail, Loader2 } from 'lucide-react';

function AuthenticatedApp() {
  const { user, logout } = useAuthContext();

  return (
    <Layout 
      header={<Header user={user} onLogout={logout} />}
      className="bg-gradient-to-br from-background to-muted"
    >
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
                      Minimal Go backend with Google OAuth2 authentication and SQLite database.
                    </p>
                  </CardContent>
                </Card>
              </div>
              
              <div className="text-center">
                <Button variant="outline" onClick={() => window.location.reload()}>
                  Refresh Page
                </Button>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </Layout>
  );
}

function LoginApp() {
  const { login } = useAuthContext();

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-background to-muted p-4">
      <Card className="w-full max-w-md">
        <CardHeader className="text-center">
          <CardTitle className="text-3xl font-bold">Chronotes Template</CardTitle>
          <CardDescription>
            A modern full-stack template with extensible architecture
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="text-center space-y-2">
            <p className="text-sm text-muted-foreground">
              Sign in with your Google account to continue
            </p>
          </div>
          <Button 
            onClick={login}
            className="w-full"
            size="lg"
          >
            <Chrome className="mr-2 h-4 w-4" />
            Continue with Google
          </Button>
          <div className="text-center">
            <p className="text-xs text-muted-foreground">
              Built with React, TypeScript, shadcn/ui, and Go
            </p>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

function LoadingApp() {
  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="flex items-center space-x-2">
        <Loader2 className="h-8 w-8 animate-spin" />
        <span className="text-lg">Loading...</span>
      </div>
    </div>
  );
}

function AppContent() {
  const { isLoading, isAuthenticated } = useAuthContext();

  if (isLoading) {
    return <LoadingApp />;
  }

  if (isAuthenticated) {
    return <AuthenticatedApp />;
  }

  return <LoginApp />;
}

function App() {
  return (
    <AuthProvider>
      <AppContent />
    </AuthProvider>
  );
}

export default App; 