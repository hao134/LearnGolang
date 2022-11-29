# Working with database (Postgres + SQLC)

###### tags: `golang_backend`

## 1. install & use docker + porstgres +tableplus to create DB schemas

### what will we do

1. install docker desktop
2. run postgres container

```dockerfile
docker pull postgres:12-alpine
docker pull <image>:<tag>
```

**start a container**

```dockerfile=
docker run
--name <container_name>
-e <environment_variable>
        -d
    <image>:<tag>

docker run
--name some-postgres
-e POSTGRES_PASSWORD=mysecretpassword
        -d
    postgres


// our case
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

// stop and start container
docker stop [container-name]
docker start [container-name]
```

Basically, a container is 1 instance of the application contained in the image.
we can start multiple containers from 1 single image.
**Run command in container**

```dockerfile=
docker exec -it
<container_name_or_id>
    <commnad> [args]
// in our case
docker exec -it postgres12 psql -U root
// log in postgres console using root as user


// go into containers bash run any linux command
docker exec -it postgres12 /bin/sh
```

**View container logs**

```dockerfile=
docker logs
<container_name_or_id>
```

3. install table plus

![](https://i.imgur.com/OC9zU4D.png)

4. create database schema

## 2. How to write & run database migration in Golang

1. install https://github.com/golang-migrate/migrate

```
brew install golang-migrate
```

useful command
![](https://i.imgur.com/zc1BMSf.png)

1. create folder /db/migration
   use code

```
migrate create -ext sql -dir db/migration -seq init_schema
```

2. go into docker postgres12's terminal

```
docker exec -it postgres12 createdb --username=root --owner=root simple_bank

//see if it work or not
docker exec -it postgres12 psql -U root simple_bank
```

in the Makefile

```
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

.PHONY: postgres createdb dropdb
```

In terminal run:

```
make postgres // to set up docker postgres
make createdb // to create db
...
```

3. migrate setup

To set up the database we created using the sql code in migrate up

```
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

## 3. Generate CRUD Golang code from SQL

:::info

### Database / SQL

- Very fast & straightforward
- Manual mapping SQL fields to variables
- Easy to make mistakes, not caught until runtimes

### Gorm

- CRUD functions already implemented, very short production code
- Must learn to write queries using gorm's function
- Run slowly on high load

### SQLX

- Quite fast & easy to use.
- Fields mapping via query text & struct tags
- Failure won't occur until runtime.

### SQLC

- Very fast & easy to use
- Automatic code generation
- Catch SQL query errors before generating codes
- Full support Postgres. MySQL is experimental.
  :::

Makefile

```
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
# -v verbose -cover -> measure code coverage ./... to run unit tests in all of them

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
```

1. 首先 sqlc.init 打在 teminal
   會產生 sqlc.yaml
2. 在 db 資料夾中創立 query 和 sqlc 兩個資料夾
3. 在 query 中按照資料庫寫 sql 語言，在此例中寫了 account, entry 和 transfer 三個 sql 檔
4. 回到 makefile 中打下 make sqlc 就會生成 go 檔，以此可以寫下測試檔來測試它們
