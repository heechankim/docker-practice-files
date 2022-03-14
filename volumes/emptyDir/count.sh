#!/bin/sh
trap "exit" SIGINT
mkdir /var/htdocs

SET=$(seq 99999)

for i in $SET; do
	echo "Running loop seq $i" > /var/htdocs/index.html
	sleep 10
done
