#!/bin/bash

psql -U postgres < ./drop_tables.sql
psql -U postgres < ./shema.sql
psql -U postgres < ./triggers.sql
psql -U postgres < ./data.sql
