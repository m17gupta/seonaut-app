# Frontend-Backend Integration Guide

## ✅ CORS Configuration Complete

Your backend is now configured to allow requests from your frontend at `https://codified-seo.vercel.app/`.

### What's Been Added:

1. **CORS Middleware** - Allows cross-origin requests from your frontend
2. **API Routes** - RESTful API endpoints for your frontend to consume
3. **Environment Configuration** - Configurable allowed origins
4. **Health Check** - API endpoint for monitoring

## API Endpoints Available

### Base URL
When deployed on Render: `https://your-backend-app.onrender.com`
For development: `http://localhost:3000`

### Authentication Endpoints
```
POST /api/auth/signup      - User registration
POST /api/auth/signin      - User login
POST /api/auth/signout     - User logout
GET  /api/auth/user        - Get current user info
```

### Project Management
```
GET    /api/projects           - Get all user projects
POST   /api/projects           - Create new project
GET    /api/projects/{id}      - Get specific project
PUT    /api/projects/{id}      - Update project
DELETE /api/projects/{id}      - Delete project
```

### Crawling Operations
```
POST /api/projects/{id}/crawl/start   - Start crawling project
POST /api/projects/{id}/crawl/stop    - Stop crawling project
GET  /api/projects/{id}/crawl/status  - Get crawl status
```

### SEO Data & Analytics
```
GET /api/projects/{id}/dashboard  - Get dashboard data
GET /api/projects/{id}/pages      - Get crawled pages
GET /api/projects/{id}/issues     - Get SEO issues
GET /api/projects/{id}/issues/{issue_id} - Get specific issue
```

### Data Export
```
GET /api/projects/{id}/export/csv     - Export data as CSV
GET /api/projects/{id}/export/sitemap - Export sitemap
```

### Health Check
```
GET /health - Backend health status
```

## Frontend Integration

### 1. API Base URL Configuration

In your frontend (Vercel app), configure the API base URL:

```javascript
// config/api.js
const API_BASE_URL = process.env.NODE_ENV === 'production' 
  ? 'https://your-backend-app.onrender.com'  // Replace with your Render URL
  : 'http://localhost:3000';

export default API_BASE_URL;
```

### 2. API Client Setup

Create an API client to handle requests:

```javascript
// services/api.js
import axios from 'axios';
import API_BASE_URL from '../config/api';

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  withCredentials: true, // Important for session cookies
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add request interceptor for error handling
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Handle unauthorized access
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default apiClient;
```

### 3. Example API Usage

```javascript
// services/projects.js
import apiClient from './api';

export const projectService = {
  // Get all projects
  async getProjects() {
    const response = await apiClient.get('/api/projects');
    return response.data;
  },

  // Create new project
  async createProject(projectData) {
    const response = await apiClient.post('/api/projects', projectData);
    return response.data;
  },

  // Start crawling
  async startCrawl(projectId) {
    const response = await apiClient.post(`/api/projects/${projectId}/crawl/start`);
    return response.data;
  },

  // Get dashboard data
  async getDashboardData(projectId) {
    const response = await apiClient.get(`/api/projects/${projectId}/dashboard`);
    return response.data;
  },
};
```

### 4. Authentication Service

```javascript
// services/auth.js
import apiClient from './api';

export const authService = {
  async login(credentials) {
    const response = await apiClient.post('/api/auth/signin', credentials);
    return response.data;
  },

  async signup(userData) {
    const response = await apiClient.post('/api/auth/signup', userData);
    return response.data;
  },

  async logout() {
    const response = await apiClient.post('/api/auth/signout');
    return response.data;
  },

  async getCurrentUser() {
    const response = await apiClient.get('/api/auth/user');
    return response.data;
  },
};
```

## Environment Variables

### Backend (Render)
```
PORT=10000
DATABASE_URL=<provided_by_render>
SEONAUT_FRONTEND_URL=https://codified-seo.vercel.app
SEONAUT_ALLOWED_ORIGINS=https://codified-seo.vercel.app,http://localhost:3000
```

### Frontend (Vercel)
```
NEXT_PUBLIC_API_URL=https://your-backend-app.onrender.com
NODE_ENV=production
```

## CORS Headers Configured

Your backend now sends these CORS headers:

```
Access-Control-Allow-Origin: https://codified-seo.vercel.app
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type, Authorization, X-Requested-With
Access-Control-Allow-Credentials: true
Access-Control-Max-Age: 86400
```

## Next Steps

1. **Deploy Backend**: Push your changes and deploy to Render
2. **Get Backend URL**: Note your Render app URL (e.g., `https://seonaut-web.onrender.com`)
3. **Update Frontend**: Configure your frontend to use the backend URL
4. **Implement API Handlers**: Complete the TODO implementations in `/internal/routes/api.go`
5. **Test Integration**: Test the connection between frontend and backend

## Testing the Connection

You can test the backend from your frontend with:

```javascript
// Test health check
fetch('https://your-backend-app.onrender.com/health')
  .then(response => response.json())
  .then(data => console.log('Backend health:', data));
```

## Important Notes

- ✅ Sessions/cookies will work across domains
- ✅ All API endpoints have CORS support
- ✅ Authentication is handled via session cookies
- ✅ Environment variables allow easy configuration
- ⚠️ API handlers need implementation (currently returning "Not implemented")

The backend is ready to receive requests from your frontend at `https://codified-seo.vercel.app/`!
