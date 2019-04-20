#!/usr/bin/env bash

godep save

git add .
git commit -m "auto save"
git push heroku master

cd sql
./heroku_reload.sh
