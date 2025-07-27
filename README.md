# Chronotes Template

This repository provides a minimal yet **extensible** template inspired by the
[Chronotes](https://github.com/Uhm-J/Chronotes) project. It demonstrates how to set up a Go backend with Google OAuth2 authentication, a SQLite user table and a **modern, scalable React frontend built with shadcn/ui**. The goal is to serve as a robust starting point for building production-ready full‑stack applications.

## Features

- **Go backend** using `net/http` and `golang.org/x/oauth2`. A
  minimal `/v1/auth/google/login` endpoint initiates the OAuth2
  flow, `/v1/auth/google/callback` processes the Google callback and
  stores the user in a SQLite database, and `/v1/auth/me` returns the
  authenticated user.
- **SQLite user store** powered by the pure Go driver
  `modernc.org/sqlite`. A `User` table with `id`, `email` and `name`
  columns is created automatically on startup.
- **Extensible React frontend** built with modern patterns and best practices:
  - **Component architecture** with shadcn/ui components
  - **Custom hooks** for state management and reusable logic
  - **Context providers** for app-wide state (auth, theme, etc.)
  - **TypeScript** throughout with comprehensive type definitions
  - **Utility libraries** for validation, formatting, and API calls
  - **Layout system** with reusable header, footer, and sidebar components

## Tech Stack

### Frontend Architecture
- **React 18** with TypeScript and modern hooks patterns
- **Vite** for fast development and optimized production builds
- **Tailwind CSS** for utility-first styling with design system
- **shadcn/ui** for beautiful, accessible, and customizable UI components
- **Lucide React** for a comprehensive icon library
- **Custom hooks** for auth (`useAuth`), localStorage (`useLocalStorage`), and more
- **Context providers** for global state management
- **Utility functions** for validation, formatting, and API interactions
- **ESLint & TypeScript** for code quality and type safety

### Backend
- **Go** with standard library HTTP server
- **Google OAuth2** for secure authentication
- **SQLite** with pure Go driver for data persistence
- **Session-based** authentication with secure cookies

## Project Structure

```
frontend/src/
├── components/           # React components
│   ├── ui/              # shadcn/ui base components (Button, Card, etc.)
│   ├── layout/          # Layout components (Header, Layout, etc.)
│   └── auth/            # Auth-specific components
├── hooks/               # Custom React hooks
│   ├── useAuth.ts       # Authentication state management
│   ├── useLocalStorage.ts # localStorage management
│   └── index.ts         # Hook exports
├── lib/                 # Utility libraries
│   ├── api/             # API client and HTTP utilities
│   ├── auth/            # Authentication API functions
│   ├── utils/           # General utility functions
│   └── utils.ts         # shadcn/ui className utilities (cn function)
├── types/               # TypeScript type definitions
│   ├── auth.ts          # Authentication types
│   ├── common.ts        # Common types (API responses, pagination, etc.)
│   └── index.ts         # Type exports
├── providers/           # React context providers
│   └── AuthProvider.tsx # Authentication context provider
├── config/              # Configuration files
│   └── env.ts           # Environment variable management
├── constants/           # Application constants
│   └── app.ts           # API endpoints, UI config, etc.
├── App.tsx              # Main application component
├── main.tsx             # React entry point
└── index.css            # Global styles with Tailwind and shadcn/ui variables
```

## Running Locally

1. **Clone the repository** (or create it from this template).

2. **Backend:**
   - Install Go ≥1.20.
   - Navigate to the `backend` directory and run `go mod tidy` to
     download dependencies.
   - Set up Google OAuth credentials (see below) and export the
     environment variables:

     ```bash
     export GOOGLE_CLIENT_ID=your-google-client-id
     export GOOGLE_CLIENT_SECRET=your-google-client-secret
     export GOOGLE_REDIRECT_URL=http://localhost:8080/v1/auth/google/callback
     export PORT=8080
     ```
   - Start the server: `go run ./main.go`.

3. **Frontend:**
   - Install Node.js (18+ recommended).
   - Navigate to the `frontend` directory and run `npm install`.
   - Start the dev server: `npm run dev`. The app will be available
     at http://localhost:5173. For production builds run
     `npm run build` and then serve the files in `dist` via the Go
     backend.

## Development

### Frontend Development Commands
```bash
cd frontend
npm run dev        # Start development server with hot reload
npm run build      # Build for production (TypeScript + Vite)
npm run lint       # Run ESLint for code quality
npm run preview    # Preview production build locally
```

### Extending the Frontend

#### Adding New shadcn/ui Components
1. Visit [shadcn/ui components](https://ui.shadcn.com/docs/components)
2. Copy the component code into `src/components/ui/`
3. Install any additional dependencies as needed
4. Import and use in your components

#### Creating Custom Hooks
```typescript
// src/hooks/useCustomHook.ts
import { useState, useEffect } from 'react';

export const useCustomHook = () => {
  // Your hook logic here
  return { /* your return values */ };
};

// Export in src/hooks/index.ts
export { useCustomHook } from './useCustomHook';
```

#### Adding New API Endpoints
```typescript
// src/lib/api/newFeature.ts
import { apiClient } from './client';

export const newFeatureApi = {
  getData: () => apiClient.get('/api/data'),
  createData: (data) => apiClient.post('/api/data', data),
};
```

#### Creating New Types
```typescript
// src/types/newFeature.ts
export interface NewFeatureType {
  id: string;
  name: string;
}

// Add to src/types/index.ts
export * from './newFeature';
```

#### Adding Layout Components
The `Layout` component supports:
- **Header**: Navigation, user menu, branding
- **Sidebar**: Navigation menus, filters
- **Footer**: Links, copyright, additional info
- **Main content**: Your page content

```typescript
<Layout 
  header={<CustomHeader />}
  sidebar={<CustomSidebar />}
  footer={<CustomFooter />}
>
  <YourPageContent />
</Layout>
```

### Design System Customization

#### Updating Theme Colors
Edit `src/index.css` to modify CSS variables:
```css
:root {
  --primary: 221.2 83.2% 53.3%;
  --secondary: 210 40% 96%;
  /* Add your custom colors */
}
```

#### Extending Tailwind Config
Update `tailwind.config.js` to add custom utilities:
```javascript
module.exports = {
  theme: {
    extend: {
      // Your custom theme extensions
    },
  },
}
```

#### Adding Utility Functions
Create new utilities in `src/lib/utils/`:
```typescript
// src/lib/utils/newUtils.ts
export const newUtilFunction = (input: string): string => {
  // Your utility logic
  return result;
};

// Export in src/lib/utils/index.ts
export * from './newUtils';
```

## Obtaining Google OAuth Credentials

Follow the steps in the [Google OAuth Setup Guide](https://developers.google.com/identity/protocols/oauth2/web-server#creatingcred) to create OAuth2 credentials:

1. Go to the [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Enable the Google+ API
4. Go to "Credentials" and create an OAuth 2.0 Client ID
5. Set the authorized redirect URI to: `http://localhost:8080/v1/auth/google/callback`
6. Use the generated Client ID and Client Secret in your environment variables

## Architecture Benefits

### Scalability
- **Modular structure**: Easy to add new features without affecting existing code
- **Type safety**: Comprehensive TypeScript types prevent runtime errors
- **Component reusability**: shadcn/ui components can be customized and extended
- **Hook-based logic**: Reusable business logic across components

### Developer Experience
- **Fast feedback**: Vite dev server with hot reload
- **Code quality**: ESLint, TypeScript, and consistent patterns
- **Easy imports**: Barrel exports and path aliases (`@/*`)
- **Documentation**: Well-documented code with examples

### Production Ready Features
- **Error handling**: API client with timeout and error management
- **State management**: Context providers for global state
- **Performance**: Optimized builds and code splitting
- **Accessibility**: shadcn/ui components follow accessibility best practices

## Notes

- This template is **production-ready** but you should add additional security measures like CSRF protection, rate limiting, and input sanitization before deploying.
- The backend serves the built frontend from `frontend/dist` if it exists. You can keep the frontend dev server separate during development.
- The database file `users.db` is created in the working directory. Delete it to reset the user list.
- The frontend uses modern React patterns with TypeScript for better developer experience and maintainability.
- All components are fully typed and documented for easy extension and modification.

## Getting Started with Development

1. **Start with the existing components** in `src/components/ui/`
2. **Create new features** by adding components, hooks, and API functions
3. **Use the provided utilities** for validation, formatting, and API calls
4. **Follow the TypeScript patterns** established in the codebase
5. **Leverage the layout system** for consistent page structure
6. **Add tests** as you build new features (testing setup can be added)

This template provides a solid foundation for building modern, scalable React applications with Go backends!
