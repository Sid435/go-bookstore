# Bookstore Api

Bookstore api made using GoLang and MySql as backend

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
    Use the url :- <b>http://localhost:9010/book/</b>
    1. Get all the books :- /book -> GET Method
    2. Get a specific book :- /book/{id} -> GET Method
    3. Post a book :- /book/ -> POST Method
    4. Update a book :- /book/{id} -> PUT Method
    5. Delete a book :- /book/{id} -> DELETE Method