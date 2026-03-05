#!/bin/bash

read -sp "Enter Postgresql password: " PGPASSWORD
export PGPASSWORD

echo -e "\ncreating db"

psql -h 127.0.0.1 -U postgres -c "CREATE DATABASE what_to_code;"

echo "applying schema"

psql -h 127.0.0.1 -U postgres -d ideas -f init.sql

echo "done"