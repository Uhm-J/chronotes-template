import React from 'react';
import { AuthProvider, useAuthContext } from './providers/AuthProvider';
import { Layout, Header, MainContent, Button, Card, CardContent, CardDescription, CardHeader, CardTitle } from './components';
import { Chrome, Loader2 } from 'lucide-react';

function AuthenticatedApp() {
  const { user, logout } = useAuthContext();

  return (
    <Layout 
      header={<Header user={user} onLogout={logout} />}
      className="bg-gradient-to-br from-background to-muted"
    >
      <MainContent user={user!} onLogout={logout} />
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