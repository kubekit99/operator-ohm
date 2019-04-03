#!/bin/bash

{{/*
Copyright 2017 The Openstack-Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/}}

set -ex

# Creates a backup for a db for an OpenStack Service:
# Set DB_USER, DB_PASSWORD, DB_HOST environment variables to contain strings
# for connection to the database.
# Set DB_NAME environment variable to contain the name of the database to back up.
# Alternateively, leave DB_NAME blank to back up all databases

function fail_if_not_exists() {
  if [ -z "${!1}" ];
  then
    echo "$1 not set"
    exit 1
  fi
}

fail_if_not_exists DB_USER
fail_if_not_exists DB_HOST
fail_if_not_exists DB_PASSWORD

if [ -z ${DB_NAME} ]
then
  echo "Backing up all databases"
  SQL_FILE=all_databases_backup.sql
  mysqldump --single-transaction -u ${DB_USER} --password=${DB_PASSWORD} -h ${DB_HOST} --all-databases > ${SQL_FILE}
else
  echo "Backing up ${DB_NAME} database"
  SQL_FILE=${DB_NAME}_backup.sql
  mysqldump --single-transaction -u ${DB_USER} --password=${DB_PASSWORD} -h ${DB_HOST} ${DB_NAME} > ${SQL_FILE}
fi

BACKUP_DIR=/etc/keystone/backups
BACKUP_FILE=${BACKUP_DIR}/$(date -u +%Y%m%dT%H%M%SZbackup.tar.gz)
echo "Dumped database(s) to ${SQL_FILE}"
tar -czf ${BACKUP_FILE} ${SQL_FILE}
echo "Backed up database(s) in ${BACKUP_FILE}"

echo "Deleting old backups"
rm -f $(ls -t1 ${BACKUP_DIR} | tail -n+11)
