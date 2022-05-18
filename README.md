# gotrade
A trading app. Prototype.

## Features
- Uses MongoDB as database
- Uses MongoDB Atlas Search, based on Apache Lucene
- The frontend is hosted separately (like on a CDN)

## Setup
Golang and MongoDB must be installed.

```sh
go get github.com/alf248/gotrade
```

## Environment Variables
This app needs the following variables to run:

| Var | Description                                                                                                                                                                                                                                                                                                                                                                                                    |
|------------|----------------|
| PORT | the port that this HTTP server will listen for requests
| FRONTEND | URL of the frontend, that sends requests to this server (needed for CORS to work)
| MONGO | URL to mongo database (example: mongodb://localhost:27017)

## dev file (optional)
Instead of environment variables, you can add a file named "dev" into the root folder, where main.go is.\
Here is the json syntax:

```sh
{
    "PORT": "1323",

    "FRONTEND": "http://localhost:3000",

    "MONGO": "mongodb://localhost:27017",

    "UseAtlasSearch": false,
}
```
It has an additional var, "UseAtlasSearch".\
Set it to true only if you connect to MongoDB Atlas, which is hosted on the cloud.

## mockdata file (optional)
There is a file called "mockdata" in the root folder.\
It is used to add mock data into the database.\
Read the file for instructions.

## API
- REQUESTS to this HTTP server are always POST with JSON body (Content-Type: application/json)
- RESPONSES are always with JSON body
- A success is always a response with HTTP OK 200
- A fail is any other HTTP status code, and usually an error message string
- URL paths can be found in server/server.go
- Clues of valid inputs can be found in the forms folder

## Login / Logout
- URL path to login and logut is "/login" and "/logout"
- Format for the json body is in forms/login.go
- When logging in, the server will respond with a json "user" object that contains a TOKEN
- That TOKEN must be used togehter with the users EMAIL in requests that need authentication
- The TOKEN and EMAIL must be put in the Authorization HEADER like so:

```sh
Authorization: "Bearer token email"
```

Javascript example:

```sh
fetch('https://server.com/path', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'Authorization': 'Bearer token email@example.com',
  },
  body: JSON.stringify(data),
})
.then(response => response.json())
.then(data => {
  console.log('Success:', data);
})
.catch((error) => {
  console.error('Error:', error);
});

```

## Tests
- Tests are in server/server_test.go
- They require a MongoDB connection
