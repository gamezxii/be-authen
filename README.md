# be-authen

be-authen is a learning project focused on exploring the Go programming language (Go 1.21.1) and applying the principles of Clean Architecture. This project integrates with a PostgreSQL database and leverages modern Go libraries and frameworks to build a maintainable and scalable backend system. A REST API is implemented using the Gin framework for its performance, simplicity, and middleware support.

Features
Authentication System: Handle user registration, login, and authentication using JWT (JSON Web Tokens).
Clean Architecture: Separation of concerns into layers such as Domain, Usecase, and Infrastructure for better scalability and maintainability.
REST API: Exposes endpoints for user management and authentication, designed with scalability and maintainability in mind.
PostgreSQL: Integrated with PostgreSQL database using GORM for Object Relational Mapping (ORM).
Gin Framework: Used for handling HTTP requests and middleware to enable fast and flexible web applications.
JWT Authentication: Secure user authentication using JSON Web Tokens for safe access control.
Technologies Used
Go 1.21.1: The programming language used to build the backend system.
Gin: A high-performance web framework for Go, used to build REST APIs with scalability in mind.
GORM: An ORM library used for interacting with PostgreSQL.
PostgreSQL: Relational database for storing user data and authentication information.
JWT (JSON Web Token): Secure mechanism for user authentication.
