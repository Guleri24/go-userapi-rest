# User Information RESTful API on Go: Fiber, MongoDB Atlas

## Description
REST API for a User Information application in which we create new users, view them, update & delete their information. Here I used a NoSQL database MongoDB for storing the data.

# Hosted on Heroku
https://go-userapi-rest.herokuapp.com/swagger/index.html

# Requirements

* Go `1.18`

    Rest of requirements are pulled by go itself. 
    Database is already configured but can be changed.

# Installation

1. Clone the repo

    `https://github.com/Guleri24/go-userapi-rest.git`

2. Open `go-userapi-rest`

    `cd go-userapi-rest`

5. Install project dependencies

    `go get ./...`

4. Optional: Swagger Docs

    `swag init`

5.  Run

    `go run main.go`


## User
### Create a new user

<details>
    <summary><code>POST</code> <code><b>/user</b></code> <code>(Creates a new user)</code></summary>

##### Parameters
> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | Body      |  required | object (JSON)   | User obj |


##### Example cURL

> ```javascript
>  curl -X 'POST' 'https://go-userapi-rest.herokuapp.com/api/v1/user' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"address": "Guler, Himachal Pradesh","description": "Student@NITH","dob": "24-12-2000","name": "Guleri"}'
> ```

</details>

_____________________________________________________________

### Get user by given ID

<details>
    <summary><code>GET</code> <code><b>/user/{id}</b></code> <code>(Get user by given ID)</code></summary>

##### Parameters
> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | string  | User ID |
 
##### Example cURL
> ```javascript
>  curl -X 'GET' 'https://go-userapi-rest.herokuapp.com/api/v1/user/63023e32606a56bdd637c202' -H 'accept: application/json'
> ```

</details>

______________________________________________________________________________________

### Delete user by given ID

<details>
    <summary><code>DELETE</code> <code><b>/user/{id}</b></code> <code>(Deletes a user by given ID)</code></summary>

##### Parameters
> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | string | User ID |


##### Example cURL

> ```javascript
>  curl -X 'DELETE' 'https://go-userapi-rest.herokuapp.com/api/v1/user/63023e32606a56bdd637c202' -H 'accept: application/json'
> ```

</details>

______________________________________________________________________________________

### Edit/Update user 

<details>
    <summary><code>PATCH</code> <code><b>/user/{id}</b></code> <code>(Updates a user by given ID)</code></summary>

##### Parameters
> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | id      |  required | string | User Id|
> | Body      |  required | object(json) | User obj |


##### Example cURL

> ```javascript
>  curl -X 'PATCH' 'https://go-userapi-rest.herokuapp.com/api/v1/user/63000c7aa8079cb65b7478e7' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"address": "Hostel@DBH","description": "Student@NITH","dob": "24-12-2000","name": "Guleri24"}'
> ```

</details>

______________________________________________________________________________________

## Users
### Get all exists users 

<details>
    <summary><code>GEt</code> <code><b>/users</b></code> <code>(Get all exists users)</code></summary>

##### Parameters
> none


##### Example cURL

> ```javascript
>  curl -X 'GET' 'https://go-userapi-rest.herokuapp.com/api/v1/users' -H 'accept: application/json'
> ```

</details>