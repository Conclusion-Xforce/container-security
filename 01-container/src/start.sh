#!/bin/bash

# Start Redis
echo "Starting Redis..."
redis-server --daemonize yes

export REDIS_HOST=localhost

# Start the application server
echo "Starting app..."
exec /app/server