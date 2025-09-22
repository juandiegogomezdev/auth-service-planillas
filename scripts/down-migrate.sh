#! /bin/bash

cd ../migrations

archivos=$(ls *.down.sql | sort -r)

read -s -p "Enter password : " password
echo " "
echo ""
echo "MIGRATING DOWN FILES"


for archivo in $archivos; do
    echo ""
    echo "<<------------------------->>"
    echo ""
    echo "Migrating: $archivo"

    PGPASSWORD=$password psql -h localhost -p 5433 -U juan -d juan -f "$archivo"

    echo "Done migrating: $archivo"
done