package config

import "os"

var STATIC_CONFIRM_EMAIL_URL = os.Getenv("REGISTER_CONFIRM_URL")
