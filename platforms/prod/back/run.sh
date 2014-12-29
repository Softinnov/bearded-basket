#!/bin/sh

NET=$(netstat -nr | grep '^0\.0\.0\.0' | awk '{print $2}')

echo "$NET chey" >> /etc/hosts

/exe/bearded-basket -db "admin:admin@(db:3306)/prod" -chey "http://chey:80" > logs/go.log 2>&1
