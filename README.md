# Go Web Application

This is a simple Go web application that serves as a starting point for building web applications using the Go programming language.

## Features

- Basic web application structure in Go.
- Web framework using the [Gin](https://github.com/gin-gonic/gin) router.
- PostgreSQL database integration using the `github.com/lib/pq` driver.
- Session management using [Gorilla Sessions](https://github.com/gorilla/sessions).
- User authentication and authorization.


## Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

- Go (1.20 or later)
- PostgreSQL (15.2)


### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/dituanson/go-webapp.git
   ```

2. Change to the project directory:

   ```bash
   cd go-webapp
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Set up your PostgreSQL database and update the database configuration in `config/db.go`.

5. Build and run the application:

   ```bash
   go run main.go
   ```

The web application should now be running locally at `http://localhost:8080`.

## Usage

The application is a simple web app with user authentication and a basic user management system. You can use it as a starting point for your own web projects.

## Contributing

If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes.
4. Commit your changes and push them to your fork.
5. Create a pull request to this repository's `main` branch.

Please ensure your code follows the project's coding standards and includes tests when necessary.

## Contact

If you have any questions or issues, please feel free to open an issue on the GitHub repository or contact the project owner at [dotuansondn@gmail.com].
