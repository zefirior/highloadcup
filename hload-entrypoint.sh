#!/usr/bin/env bash

service mysql start
mysql < /app/db/migrate.sql

DATA_DIR=/app/data

mkdir ${DATA_DIR}
unzip -o /tmp/data/data.zip -d ${DATA_DIR}
python /app/src/parse.py ${DATA_DIR}

uwsgi --http :80 --wsgi-file /src/app.py --callable app --process 4
