# gotrade
An small e-commerce app written in Go

#### Live version (frontend)
https://trader-cyan.vercel.app

### Other Parts

#### Frontend
[github.com/alf248/trade-frontend](https://github.com/alf248/trade-frontend)

#### User service
[github.com/alf248/user-service](https://github.com/alf248/user-service)


## Features
- MongoDB
- Firebase Authentication

## Todo
- GraphQL
- Payment system

## Setup
Golang and MongoDB must be installed.

```sh
go get github.com/alf248/gotrade
```

## Environment Variables
Provide these or an env file
| Var | Description                                                                                                                                                      |
|------------|----------------|
| PORT | the port for this HTTP server
| FRONTEND | URL of the frontend (needed for CORS to work)
| MONGO | URL to mongo database (example: mongodb://localhost:27017)

## env file (optional)
Instead of environment variables, you can add an "env" in the root folder:

```sh
{
    "PORT": "1323",
    "FRONTEND": "http://localhost:3000",
    "MONGO": "mongodb://localhost:27017",
    "UseAtlasSearch": false,
}
```
Set "UseAtlasSearch" true when connecting to MongoDB Atlas.
