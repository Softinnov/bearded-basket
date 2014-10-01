#!/bin/bash

if [[ $# -lt 4 ]]; then
	echo "Usage: $0 <username> <password> <database> <sql_file1> <...>"
	echo "sql_file: file names without .sql or .txt.gz"
	exit 1
fi

echo "=> Starting MySQL Server"
/usr/bin/mysqld_safe > /dev/null 2>&1 &
PID=$!

RET=1
while [[ RET -ne 0 ]]; do
    echo "=> returned ${RET}"
    echo "=> Waiting for confirmation of MySQL service startup"
    sleep 5
    mysql -u"$1" -p"$2" -e "status" > /dev/null 2>&1
RET=$?
done

echo "   Started with PID ${PID}"

DATA="/data/"
echo "=> importing files inside $DATA"
for ARG in ${*:4}
do
    echo "  -> importing $ARG"
    gunzip -f "$DATA""$ARG".txt.gz
    cat "$DATA""$ARG".sql | mysql -u"$1" -p"$2" "$3" || exit $?
    mysqlimport -u"$1" -p"$2" "$3" "$DATA""$ARG".txt || exit $?
done

echo "=> Stopping MySQL Server"
mysqladmin -u"$1" -p"$2" shutdown

echo "=> Done!"
