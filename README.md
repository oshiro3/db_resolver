# db_resolver
db_resolver is an independent subsystem for versioning and migrating multiple databases.
This system refers to the database schema managed by the application and can migrate the database synchronously with the update of the application.

Currently it only supports local filesystems and gitbucket, but we plan to allow multiple filehosts to be treated as schema hosts in the future.
Also, regarding the target DB, only mysql is currently supported, but we are considering future expansion.

## Usage

### Library

Not Suported

### CLI

```
$db_resolver set ${ database_name }
-> (C) $DB_RESOLVER_ROOT/${ database_name }/sqls

$db_resolver up ${ database_name } \
  --src ${ yourSchemaManagingSystem } \
  --path ${ path/to/your/schema | url }
-> (U) $DB_RESOLVER_ROOT/${ database_name }/out_of_date_schema.sql
-> (U) $DB_RESOLVER_ROOT/${ database_name }/schema.sql
-> (U) $DB_RESOLVER_ROOT/${ database_name }/version
```

## Demo

```
$make up
$make init
-> Not exists dev column in test_table

$make run
-> Created dev column in test_table and migrate_version

$make clean
```
