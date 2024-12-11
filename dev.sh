#!/bin/bash

# Trap Ctrl+C and kill all child processes
trap 'kill $(jobs -p); exit' SIGINT

# Create tmp directory if it doesn't exist
mkdir -p tmp

# Run these in parallel using & and wait
($HOME/go/bin/templ generate --watch &)
(tailwindcss -i ./web/global.css -o ./public/output.css --watch &)
($HOME/go/bin/air &)

# Wait for all background processes
wait