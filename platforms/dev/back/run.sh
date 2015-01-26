#!/bin/sh

NET=$(netstat -nr | grep '^0\.0\.0\.0' | awk '{print $2}')

echo "$NET chey" >> /etc/hosts

/exe/bearded-basket -dbuspw "admin:admin" -dbhost "db" -dbname "prod" -chey "http://chey:8000" > logs/go.log 2>&1
