# Simple Bank

##### source: https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/

Due to unknown error, i start version 2 from (after dockerfile created at course 25) : https://github.com/techschool/simplebank/tree/fc041eb8f380482600f872cd6bd52fc19b7da4f2

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

1. ?????? sqlc.init ?????? teminal
   ????????? sqlc.yaml
2. ??? db ?????????????????? query ??? sqlc ???????????????
3. ??? query ????????????????????? sql ??????????????????????????? account, entry ??? transfer ?????? sql ???
4. ?????? makefile ????????? make sqlc ???????????? go ????????????????????????????????????????????????

## A clean way to implement database transaction in Golang

### What is a db transaction?

- A single unit of work
- Often made up of multiple db operations

Example:
Transfer 10 USD from bank account 1 to bank account 2

1. Create a transfer record with amount 10
2. Create an account entry for account 1 with amount = -10
3. Create an account entry for account 2 with amount = +10
4. Subtract 10 from the balance of account 1
5. Add 10 to the balance of account2

### Why do we need db transaction?

1. To provide a reliable and consistent unit of work, even in case of system failure
2. To provide isolation between programs that access the database concurrently.

### ACID PROPERTY

1. Atomicity(A): Either all operations complete successfully or the transaction fails and the db is unchanged.
2. Consistency(C): The db state must be valid after the transaction. All constranints must be satisfied.
3. Isolation(I): Concurrency transactions must not affect each other.
4. Durability(D): Data written by a successful transaction must be recorded in persistent storage.

### How to run sql TX?

```
BEGIN;
...
COMMIT
```

```
BEGIN
...
ROLLBACK;
```

## Deeply understand transaction isolation levels & read phenomena

### Read Phenomena

- DIRTY READ
  A transaction **reads** data written by other concurrent **uncommitted** transaction
- NON-REPEATABLE READ
  A transaction **reads** the **same row twice** and sees different value because it has been **modified** by other **committed** transaction
- PHANTOM READ
  A transaction **re-executes** a query to **find rows** that satisfy a condition and sees a **different set** of rows, due to changes by other **committed** transaction
- SERIALIZATION ANOMALY
  The result of a **group** of concurrent **committed transactions** is **impossible to achieve** if we try to run them **sequentially** in any order without overlapping.

### 4 Standard Isolation Levels(American National Standards Institure-ANSI)

- Low Level

1. READ UNCOMMITTED: Can see data written by uncommitted transaction
2. READ COMMITTED: Only see data written by committed transaction
3. REPEATABLE READ: Same read query always reutrns same result.
4. SERIALIZABLE: Can achieve same result if execute transactions serially in some order instead of concurrently.

### See 4 isolation Levels in MySQL & postgreSQL

- MySQL

1. READ UNCOMMITTED:
   ![](https://i.imgur.com/KuvBIPs.jpg)
2. READ COMMITTED
   ![](https://i.imgur.com/9lRWthS.jpg)
3. REPEATABLE COMMITTED
   ???????????????????????? serializable anomaly ?????????
   ![](https://i.imgur.com/FQWPyS0.jpg)
4. SERIALIZABLE
   MySQL ????????????????????? serializable anomaly ????????????????????? commit;?????????????????? update
   ![](https://i.imgur.com/5dZQG7n.jpg)

- postgreSQL
  ?????? postgre ?????? read uncommitted ??? read committed ????????????????????????????????? isolation level.

1. READ UNCOMMITTED & READ COMMITTED:
   ???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? PHANTOM READ
   ![](https://i.imgur.com/HWQTbbs.jpg)
2. REPEATABLE READ
   ????????????????????? REAPEATABLE READ ????????????????????????(could not serialize access due to concurrent update)
   ![](https://i.imgur.com/X8ZxFfc.jpg)
   ????????? REAPEATABLE READ ???????????? SERIALIZATION ANOMALY(????????????)
   ![](https://i.imgur.com/zEvZHQ1.jpg)

3. SERIALIZABLE
   ??????????????????(CHECK)(HINT: The transaction might succeed if retried.)????????? SERIALIZATION ANOMALY
   ![](https://i.imgur.com/6AcgIXz.jpg)

**Summary**
level 4 ???????????? ??????
level 3 ???????????? SERIALIZATION ANOMALY ???????????????
...

mysql default mode is REPEATABLE COMMITTED
postgreSql mode is READ COMMITTED

mysql ?????? Loking mechanism
postgresql ?????? dependencies detection

# Building RESTful HTTP JSON API1

###### tags: `golang_backend`

## Why mock database?

- INDEPENDENT TESTS
  Isolate tests data to avoid conflicts

- FASTER TESTS
  Reduce a lot of time talking to the database

- 100% COVERAGE
  Easily setup edge cases: unexpected errors

### Is it good enough to test our API with a mock DB?

Yes, our real db store is already tested
-> MOCK DB & REAL DB SHOULD IMPLEMENT THE SAME INTERFACE

### How to mock ?

- USE FACK DB: MEMORY
  Implement a fake version of DB: store data in memory

- USE DB STUBS: GOMOCK
  Generate and build stubs that returns hard-coded values.

### steps

1. After modify file in store.go in sqlc

```go=
// Store provides all functions to execute db queries and transaction
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transaction
type SQLStore struct {
	db *sql.DB
	*Queries
}
```

modify the other term using above in the code.

2. at sqlc.ymal
   set

```
emit_interface: true
```

and in terminal

```
make sqlc
```

you will see a code called querier.go in sqlc
it can be called in the Store interface

3. generate a folder in db, called mock(db/mock):
   generate a code using mockdb typing:

```
mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store
```

in terminal
you can see store.go in mock folder

## Why PASETO is better than JWT for token-based authentication?

### JWT SIGNING ALGORITHMS

**Symmetric digital signature algorithm**

- The same secreat key is used to sign & verify token
- For local use: internal services, where the secret key can be shared
- HS256, HS384, HS512
  - HS256 = HMAC + SHA256
  - HMAC: Hash-based Message Authentication Code
  - SHA: Secure Hash Algorithm
  - 256/384/512:number of output bits

**Asymmetric digital signature algorithm**

- The private key is used to sign token
- The public key is used to verify token
- For public use: internal service signs token, but external services needs to verify it
- RS256, RS384, RS512 | PS256, PS384, PS512 | PS256, ES384, ES512
  - RS256 = RSA PKCSv1.5 + SHA256 [PKCS:Public-Key Cryptography Standards]
  - PS256 = RSA PSS + SHA256 [PSS: Probabilistic Signature Scheme]
  - ES256 = ECDSA + SHA256 [ECDSA:Elliptic Curve Digital Signature Algorithm]

**What's the problem of JWT?**
![](https://i.imgur.com/qEjfAwi.png)

### Platform-Agnostic SEcurity TOkens[PASETO]

**Stronger algorithms**

- Developers don't have to choose the algorithm
- Only need to select the version of PASETO
- Each version has 1 strong cipher suite
- Only 2 most recent PASTO versions are accepted

**Non-trivial Forgery**

- No more "alg" header or "none" algorithm
- Everything is authenticated
- Encrypted payload for local use <symmetric key>
  ![](https://i.imgur.com/rH7c2Ec.png)

## Implement authentication middleware and authorization rules in Golang using Gin

### What is a middleware?

![](https://i.imgur.com/DKAjabZ.png)
![](https://i.imgur.com/JaBnHqp.png)
![](https://i.imgur.com/4M6JK6G.png)

### AUTHORIZATION RULES

![](https://i.imgur.com/RXLCh1p.png)

# Deploying the application to production

###### tags: `golang_backend`

## How to deploy an application?

????????????????????????????????????
https://www.youtube.com/watch?v=p1dwLKAxUxA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=23
![](https://i.imgur.com/lvRt8O6.png)

```dockerfile=
# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]
```

1. ????????? dockerfile:

```
docker build -t simplebank:latest .
```

2. ?????? docker images
   ![](https://i.imgur.com/oZzr2zD.png)
   ????????????????????????????????? simplebank ????????? images

3. ????????? docker

```
docker run --name simplebank -p 8080:8080 simplebank:latest
```

![](https://i.imgur.com/cDaTOok.png)

?????? postgres12 ??? ip ??? simplebank ??? ip ????????????????????????????????????????????????????????????????????? port:8080 ??????????????????????????????????????????

**????????????**
??? create ?????? network:

```
docker network create bank-network
```

???????????? postgres12 ?????? network

```
docker network connect bank-network postgres12
```

???????????????????????????????????????

```
docker network inspect bank-network
```

![](https://i.imgur.com/ZqOYfw4.png)
???????????? postgres12 ??????????????????????????????????????????????????????
![](https://i.imgur.com/EZBS9oN.png)

????????????????????? simplebank container ???????????????(docker rm simplebank) ????????????????????????????????????

```
docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" simplebank:latest
```

??????????????????????????????

```
--network bank-network
```

??? postgres12 ?????? @ ????????? ip

## ???????????? 30 ??????????????????????????????????????????????????? AWS ?????????)

## Kubernetes architecture & How to create an EKS cluster on AWS

### What is Kubernetes?

- An open-source container orchestration engine
- For automating deployment, scaling, and management of containerized applications.
  ![](https://i.imgur.com/FGIuFA2.png)
  ![](https://i.imgur.com/OPUeKEP.png)
  ![](https://i.imgur.com/z3Ii5Y7.png)
