#!/bin/bash

# Start Redis
echo "Starting Redis..."
redis-server --daemonize yes

# Start the application server
export REDIS_HOST=localhost
export REDIS_PORT=6379
export SERVICE_PORT=80 

exec /app/server