# Bookstore Api

Bookstore api made using GoLang and MySql as database. JWT for token based security.

## Table of Contents

- [Bookstore Api](#bookstore-api)
  - [Table of Contents](#table-of-contents)
  - [Demo](#demo)
  - [Installation](#installation)
  - [Testing](#testing)

## Demo
[Demo](https://github.com/Sid435/go-bookstore.git)


## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Sid435/go-bookstore.git
   ```
2. **Navigate to the project directory:**
   ```bash
   cd go-bookstore
   ```
3. **Docker:**
   ```bash
    docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=rootpassword -e MYSQL_DATABASE=simplerest -e MYSQL_USER=siddharth -e MYSQL_PASSWORD=siddharth -p 3306:3306 -d mysql:latest
    ```

## Usage
1.  **Navigate to main folder:**
    ```bash
    cd cmd/main
    ```
3. **Start the server:**
   ```bash
   cd go run main.go
   ```

## Testing

1. **Go to postman:**
    Use the url :- <b>http://localhost:9010/</b>
    1. Auth :- /auth/login -> POST Method
    2. Signup :- /auth/sigup -> POST Method
    3. Get all the books :- /books -> GET Method
    4. Get a specific book :- /books/{id} -> GET Method
    5. Post a book :- /books/ -> POST Method
    6. Update a book :- /books/{id} -> PUT Method
    7. Delete a book :- /books/{id} -> DELETE Method