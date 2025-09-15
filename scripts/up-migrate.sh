#! /bin/bash

cd ../migrations

archivos=$(ls *.up.sql)

read -s -p "Enter password : " password
echo ""
echo "MIGRATING UP FILES"


for archivo in $archivos; do
    echo ""
    echo "<<------------------------->>"
    echo "Migrating: $archivo"
    PGPASSWORD=$password psql -h localhost -p 5433 -U juan -d juan -f "$archivo"
    echo "Done migrating: $archivo" 
done