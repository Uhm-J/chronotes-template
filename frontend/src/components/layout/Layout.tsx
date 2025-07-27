import { ReactNode } from 'react';
import type { BaseComponentProps } from '../../types';
import { cn } from '../../lib/utils';

interface LayoutProps extends BaseComponentProps {
  children: ReactNode;
  header?: ReactNode;
  footer?: ReactNode;
  sidebar?: ReactNode;
  className?: string;
}

export const Layout = ({ 
  children, 
  header, 
  footer, 
  sidebar, 
  className,
  ...props 
}: LayoutProps) => {
  return (
    <div className={cn('min-h-screen flex flex-col', className)} {...props}>
      {header && (
        <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
          {header}
        </header>
      )}
      
      <div className="flex flex-1">
        {sidebar && (
          <aside className="w-64 border-r bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
            {sidebar}
          </aside>
        )}
        
        <main className="flex-1 overflow-auto">
          {children}
        </main>
      </div>
      
      {footer && (
        <footer className="border-t bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
          {footer}
        </footer>
      )}
    </div>
  );
}; 