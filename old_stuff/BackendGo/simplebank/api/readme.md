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
