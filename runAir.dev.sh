#! /bin/bash


# Static URLs
export ROOT_URL="http://localhost:8080"
export LOGIN_URL="$ROOT_URL/static/login/index.html"
export LOGIN_CODE_URL="$ROOT_URL/static/confirm-code/index.html"
export REGISTER_URL="$ROOT_URL/static/register/index.html"
export REGISTER_CONFIRM_URL="$ROOT_URL/static/confirm-email/index.html"

# Database
export DB_HOST="localhost"
export DB_PORT="5433"
export DB_USER="juan"
export DB_PASSWORD="tunclave"
export DB_NAME="juan"


# Mailer 
export RESEND_API_KEY="resend_api_key_here"
air