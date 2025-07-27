#!/bin/bash

# Chronotes Backend Setup Script
# This script helps you set up the development environment

set -e  # Exit on any error

echo "🚀 Setting up Chronotes Backend..."

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go 1.20 or later.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Go is installed: $(go version)${NC}"

# Check if PostgreSQL is installed
if ! command -v psql &> /dev/null; then
    echo -e "${YELLOW}⚠️  PostgreSQL CLI (psql) not found. Please install PostgreSQL.${NC}"
    echo "   macOS: brew install postgresql"
    echo "   Ubuntu: sudo apt-get install postgresql postgresql-contrib"
    echo "   Windows: Download from https://www.postgresql.org/download/"
else
    echo -e "${GREEN}✅ PostgreSQL CLI is available${NC}"
fi

# Create .env file from template
if [ ! -f .env ]; then
    echo -e "${YELLOW}📄 Creating .env file from template...${NC}"
    cp env.development.template .env
    echo -e "${GREEN}✅ .env file created${NC}"
    echo -e "${YELLOW}⚠️  Please edit .env file with your actual values:${NC}"
    echo "   - Database password (DB_PASSWORD)"
    echo "   - Google OAuth credentials (GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET)"
else
    echo -e "${GREEN}✅ .env file already exists${NC}"
fi

# Download Go dependencies
echo -e "${YELLOW}📦 Downloading Go dependencies...${NC}"
go mod tidy
echo -e "${GREEN}✅ Dependencies downloaded${NC}"

# Check if database exists
DB_NAME=$(grep DB_NAME .env | cut -d '=' -f2)
DB_USER=$(grep DB_USER .env | cut -d '=' -f2)

echo -e "${YELLOW}🗄️  Checking database connection...${NC}"
if command -v psql &> /dev/null; then
    # Try to connect to PostgreSQL
    if psql -h localhost -U "$DB_USER" -d postgres -c '\q' 2>/dev/null; then
        echo -e "${GREEN}✅ PostgreSQL connection successful${NC}"
        
        # Check if database exists, create if not
        if psql -h localhost -U "$DB_USER" -d postgres -lqt | cut -d \| -f 1 | grep -qw "$DB_NAME"; then
            echo -e "${GREEN}✅ Database '$DB_NAME' already exists${NC}"
        else
            echo -e "${YELLOW}📊 Creating database '$DB_NAME'...${NC}"
            createdb -h localhost -U "$DB_USER" "$DB_NAME"
            echo -e "${GREEN}✅ Database '$DB_NAME' created${NC}"
        fi
    else
        echo -e "${YELLOW}⚠️  Could not connect to PostgreSQL. Please ensure:${NC}"
        echo "   1. PostgreSQL is running"
        echo "   2. User '$DB_USER' exists and has proper permissions"
        echo "   3. Password in .env is correct"
    fi
fi

echo ""
echo -e "${GREEN}🎉 Setup complete!${NC}"
echo ""
echo "Next steps:"
echo "1. Edit .env file with your actual values"
echo "2. Make sure PostgreSQL is running"
echo "3. Set up Google OAuth credentials"
echo "4. Run the server: go run cmd/server/main.go"
echo ""
echo "Google OAuth Setup:"
echo "1. Go to https://console.cloud.google.com/"
echo "2. Create a new project or select existing"
echo "3. Enable Google+ API"
echo "4. Go to Credentials > Create OAuth 2.0 Client ID"
echo "5. Set redirect URI to: http://localhost:8080/v1/auth/google/callback"
echo "6. Copy Client ID and Secret to .env file" 