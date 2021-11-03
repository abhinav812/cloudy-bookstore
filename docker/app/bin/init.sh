#!/usr/bin/env bash
echo 'Running migrations...'
/app/migrate up > /dev/null 2>&1 &
echo 'Deleting mysql-client...'
apk del postgresql-client
echo 'Start application...'
/app/bookstore