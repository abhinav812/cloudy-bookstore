#!/usr/bin/env bash
echo 'Deleting postgresql-client...'
apk del postgresql-client

echo 'Start application...'
/app/bookstore