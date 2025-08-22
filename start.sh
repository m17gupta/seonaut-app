#!/bin/bash

# Render Deployment Script for Seonaut
echo "Starting Seonaut deployment on Render..."

# The application will automatically run database migrations on startup
echo "Starting Seonaut server..."
exec ./seonaut -c config.prod
