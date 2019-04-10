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
fail_if_not_exists DB_PASSWORD
fail_if_not_exists DB_HOST

# NOTE(howell): Revisions count backward. This needs to be fixed
revision=${DB_RESTORE_REVISION:-1}

backup_dir=/etc/keystone/backups
backup_file_name=$(ls ${backup_dir} -1t | tail -n+$revision | head -n1)

if [ -z ${backup_file_name} ]
then
  echo "Revision $revision does not exist"
  exit
fi

backup_file=${backup_dir}/${backup_file_name}
echo "Using backup file ${backup_file}"

# NOTE(howell): This assumes that the only file stored in ${backup_file} is
# called "tmp/keystone_backup.sql". It would be ideal to come up with a
# smarter way of checking this
tar --overwrite --directory=/ -xzf ${backup_file}
sql_file=/tmp/keystone_backup.sql

echo "Restoring keystone database"
mysql --host ${DB_HOST} --user=${DB_USER} --password=${DB_PASSWORD} --database=keystone < ${sql_file}
echo "Restored keystone database"
