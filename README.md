# zero-trust-network

This project demonstrates a simple implementation of a Zero Trust Network model using the Go programming language. It provides an HTTP server with basic authentication and access control based on user roles.

# Features

Basic Authentication: Users must provide a username and password to access the resources.
Access Control: Users are granted access to specific resources based on predefined policies.
Resource Endpoints:

    /resource1: Accessible to user1 and user2.
    /resource2: Accessible to user1 and user2, with role-based restrictions.

# Components

Authentication Middleware: Validates user credentials using a simple in-memory store.
Authorization Middleware: Checks if the authenticated user has permission to access the requested resource based on predefined policies.
HTTP Server: Configures and serves the application endpoints with integrated middleware for security.

# Getting Started

Clone the Repository:

git clone https://github.com/Daxcellcn/zero-trust-network.git
cd zero-trust-network

Build the Project:

go build -o server main.go

Run the Server:

./server

The server will start on port 8080. Access the resources using your web browser or tools like curl with basic authentication.

# Example Usage

To access /resource1 or /resource2, you need to provide valid credentials. For example, using curl:

curl -u user1:password1 http://localhost:8080/resource1

# License

This project is licensed under the MIT License. See the LICENSE file for details.
