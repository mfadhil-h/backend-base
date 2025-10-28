#!/bin/bash

# This script sets up the environment for the Go project.

# Exit immediately if a command exits with a non-zero status.
set -e

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go to proceed."
    exit 1
fi

# Initialize Go modules
echo "Initializing Go modules..."
go mod tidy

# Run any database migrations or setup tasks here
# echo "Running database migrations..."
# Uncomment and add your migration command if needed

echo "Setup completed successfully."