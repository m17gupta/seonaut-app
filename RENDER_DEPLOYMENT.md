# Deploying Seonaut to Render

This guide walks you through deploying the Seonaut SEO crawler to Render.

## Prerequisites

1. A GitHub account with your Seonaut code pushed to a repository
2. A Render account (free tier available)

## Deployment Steps

### 1. Prepare Your Repository

Ensure your code is pushed to GitHub with the following files:
- `Dockerfile` (already configured)
- `render.yaml` (Render configuration)
- `config.prod` (production configuration)
- `start.sh` (startup script)

### 2. Create a New Web Service on Render

1. Go to [Render Dashboard](https://dashboard.render.com/)
2. Click "New +" and select "Web Service"
3. Connect your GitHub repository
4. Configure the service:
   - **Name**: `seonaut-web`
   - **Environment**: `Docker`
   - **Region**: Choose your preferred region
   - **Branch**: `main` (or your main branch)
   - **Dockerfile Path**: `./Dockerfile`

### 3. Configure Environment Variables

Add the following environment variables in Render:

#### Required Variables:
- `PORT`: `10000` (automatically set by Render)
- `DATABASE_URL`: Your MySQL database connection string

#### Optional Variables:
- `SEONAUT_SERVER_URL`: Your app's public URL (e.g., `https://your-app.onrender.com`)
- `SEONAUT_CRAWLER_AGENT`: Custom user agent for crawling

### 4. Database Setup

#### Option A: Use Render's Managed Database
1. Create a new PostgreSQL database in Render
2. Connect it to your web service
3. The `DATABASE_URL` will be automatically provided

#### Option B: Use External MySQL Database
1. Set up MySQL database elsewhere (e.g., PlanetScale, Railway, etc.)
2. Set the `DATABASE_URL` environment variable:
   ```
   mysql://username:password@hostname:3306/database_name
   ```

### 5. Deploy

1. Click "Create Web Service"
2. Render will automatically build and deploy your application
3. The deployment process includes:
   - Building the Go application
   - Building frontend assets with esbuild
   - Creating the final Docker image
   - Starting the service

### 6. Run Database Migrations

After deployment, you may need to run database migrations. You can do this by:
1. Connecting to your service's shell in Render dashboard
2. Running migration commands manually

## Environment Variables Reference

| Variable | Description | Example |
|----------|-------------|---------|
| `PORT` | Server port (set by Render) | `10000` |
| `DATABASE_URL` | Database connection string | `mysql://user:pass@host:3306/db` |
| `SEONAUT_SERVER_URL` | Public URL of your app | `https://your-app.onrender.com` |
| `SEONAUT_CRAWLER_AGENT` | User agent for web crawling | `SEOnautBot/1.0` |

## Health Check

The application exposes a health check endpoint at `/` that Render will use to monitor the service health.

## Troubleshooting

### Common Issues:

1. **Build Failures**: Check the build logs in Render dashboard
2. **Database Connection**: Verify `DATABASE_URL` is correct
3. **Port Issues**: Ensure the app listens on the `PORT` environment variable

### Logs

View application logs in the Render dashboard under your service's "Logs" tab.

## Scaling

Render's starter plan includes:
- 512 MB RAM
- 0.1 CPU
- Automatic HTTPS
- Custom domains

You can upgrade to higher plans for better performance and resources.

## Custom Domain

To use a custom domain:
1. Go to your service settings in Render
2. Add your custom domain
3. Update your DNS records as instructed by Render

## Support

For Render-specific issues, check:
- [Render Documentation](https://render.com/docs)
- [Render Community](https://community.render.com/)
- Render support through their dashboard
