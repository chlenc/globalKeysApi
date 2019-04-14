#!/usr/bin/env bash

cat drop_tables.sql | heroku pg:psql --app globalkeys
cat shema.sql | heroku pg:psql --app globalkeys
cat triggers.sql | heroku pg:psql --app globalkeys
cat data.sql | heroku pg:psql --app globalkeys

echo "select * from table;" | heroku pg:psql --app globalkeys

heroku pg:psql --app globalkeys < drop_tables.sql
heroku pg:psql --app globalkeys < shema.sql
heroku pg:psql --app globalkeys < triggers.sql
heroku pg:psql --app globalkeys < data.sql
