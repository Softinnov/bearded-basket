#!/bin/bash

# NET=$(netstat -nr | grep '^0\.0\.0\.0' | awk '{print $2}')
#
# echo "$NET back" >> /etc/hosts

exec nginx
