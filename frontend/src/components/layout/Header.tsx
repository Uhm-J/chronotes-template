import { Button } from '../ui/button';
import { User, LogOut } from 'lucide-react';
import type { User as UserType } from '../../types';
import { APP_CONFIG } from '../../constants/app';

interface HeaderProps {
  user?: UserType | null;
  onLogout?: () => void;
}

export const Header = ({ user, onLogout }: HeaderProps) => {
  return (
    <div className="container mx-auto px-4 py-3">
      <div className="flex items-center justify-between">
        {/* Logo and brand */}
        <div className="flex items-center space-x-2">
          <h1 className="text-xl font-bold">{APP_CONFIG.NAME}</h1>
        </div>

        {/* User menu */}
        {user ? (
          <div className="flex items-center space-x-4">
            <div className="flex items-center space-x-3 text-sm">
              <User className="h-4 w-4" />
              <div className="flex flex-col">
                <span className="font-medium">{user.name}</span>
                <span className="text-xs text-muted-foreground">{user.email}</span>
              </div>
            </div>
            
            {onLogout && (
              <Button 
                variant="ghost" 
                size="sm" 
                onClick={onLogout}
                className="flex items-center space-x-2"
              >
                <LogOut className="h-4 w-4" />
                <span>Logout</span>
              </Button>
            )}
          </div>
        ) : (
          <div className="text-sm text-muted-foreground">
            Not signed in
          </div>
        )}
      </div>
    </div>
  );
}; 