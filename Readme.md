# task-5-pbi-btpns-JoseBagusRamadhan

## Overview
This repository contains the implementation of a Full Stack Developer task for VIX, focusing on creating APIs for user registration, login, and photo management. The API includes endpoints for user actions such as registration, login, update, and deletion, as well as photo-related actions like uploading, retrieving, updating, and deleting.

## Project Structure
- **app**: Holds the struct definition, including the user struct for data and authentication purposes.
- **controllers**: Contains logic related to the database, including models and queries.
- **database**: Holds the database configuration and is used for establishing database connections and migrations.
- **helpers**: Includes various functions that can be used throughout the project, such as JWT, bcrypt, and header value handling.
- **middlewares**: Consists of functions used for JWT authentication to protect the API.
- **models**: Contains the models used for defining database relationships.
- **router**: Configures the routing and endpoints for accessing the API.
- **go mod**: Manages packages and dependencies, including the jwt-go library.

## API Endpoints

### User Endpoint
- **POST** `/users/register`
  - Attributes:
    - ID (primary key, required)
    - Username (required)
    - Email (unique & required)
    - Password (required & minlength 6)
    - Created At (timestamp)
    - Updated At (timestamp)
  - Relationship with Photo model (Constraint cascade)

- **GET** `/users/login`
  - Using email & password (required)

- **PUT** `/users/:userId`
  - Update User

- **DELETE** `/users/:userId`
  - Delete User

### Photos Endpoint
- **POST** `/photos`
  - Attributes:
    - ID
    - Title
    - Caption
    - PhotoUrl
    - UserID
  - Relationship with User model

- **GET** `/photos`

- **PUT** `/photos/:photoId`
  - Update Photo

- **DELETE** `/photos/:photoId`

## Authentication
Authorization is implemented using the Go JWT library ([jwt-go](https://github.com/dgrijalva/jwt-go)). Only the user who created a photo is allowed to modify or delete that photo.

## Database
Use PostgreSQL as the database.

## Frameworks and Libraries Used
- [Gin Gonic Framework](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/index.html)
- [JWT Go](https://github.com/dgrijalva/jwt-go)
- [Go Validator](http://github.com/asaskevich/govalidator)

## How to Run
1. Set up the database and configure the connection in the `database` folder.
2. Configure environment variables if needed.
3. Run the necessary migrations.
4. Start the server using the provided configurations in the `router` folder.

## Dependencies
Ensure that you have the required dependencies specified in the `go mod` file.

Feel free to reach out for any further clarification or assistance!
