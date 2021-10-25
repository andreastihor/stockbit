Please create database and run the migration query manually on your sql editor

### Environment Variable Example
```sh

DB_HOST=localhost
DB_PORT=5432
DB_NAME=stockbit
DB_USERNAME=tihor
DB_PASSWORD=12345
SSL_MODE=false
```
### Explantaion For Env
The description for each variables is presentes in the following table,
|Variables|Description|Example value|
|-------------------------|------------------------------------------------|-----|
| DB_HOST | the target host for postgres server | localhost |
| DB_PORT | the target port for postgres server | 5432 |
| DB_NAME | the database name for postgres server | stockbit |
| DB_USERNAME | the database username credential for postgres server | tihor |
| DB_PASSWORD | the database username credential for postgres server | 12345 |
| SSL_MODE | whether or not the postgres server using ssl technology. default = false| true |
