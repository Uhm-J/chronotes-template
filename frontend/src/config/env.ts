interface EnvConfig {
  API_BASE_URL: string;
  APP_NAME: string;
  NODE_ENV: string;
  DEV_MODE: boolean;
}

const env: EnvConfig = {
  API_BASE_URL: import.meta.env.VITE_API_BASE_URL || '',
  APP_NAME: import.meta.env.VITE_APP_NAME || 'Template',
  NODE_ENV: import.meta.env.MODE || 'development',
  DEV_MODE: import.meta.env.DEV || false,
};

// Validate required environment variables in production
if (env.NODE_ENV === 'production') {
  const requiredVars = [];
  
  requiredVars.forEach((varName) => {
    if (!env[varName as keyof EnvConfig]) {
      throw new Error(`Missing required environment variable: ${varName}`);
    }
  });
}

export default env; 