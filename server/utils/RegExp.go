package utils

import "regexp"

var EmailRegex = regexp.MustCompile(`^[^\s@]+@(gmail\.com|yahoo\.com|outlook\.com|icloud\.com|fastmail\.com)$`)

var UsernameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{2,15}$`)
