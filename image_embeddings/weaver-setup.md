
# Weaver Setup Guide

This document outlines the steps taken to configure and deploy a backend application using Service Weaver. The setup involves creating a suitable project structure, configuring the `weaver.toml`, and setting up the `main.go` for proper service deployment.

## Project Structure

The project is organized into separate directories for each service and a shared utility directory. This structure helps in maintaining clear separation of concerns and modularity.

```
project-root/
├── admin-service/
│   └── cmd/
│       └── main.go
├── user-service/
│   └── cmd/
│       └── main.go
|__ main.go
└── go.mod
```

- **Admin and User Services**: Each service has its own directory containing source files specific to that service.
- **Shared Directory**: Contains shared utilities like the Weaver setup and configuration code.
- **`go.mod`**: Manages dependencies for the entire project.

## Weaver Configuration (`weaver.toml`)

The `weaver.toml` file specifies how Service Weaver should deploy the services, including the binary to use and the ports for the listeners.

```toml
[serviceweaver]
binary = "./main.exe"

[multi]
listeners.UserService = { address = "localhost:8081" }
listeners.AdminService = { address = "localhost:8082" }

# Other global configurations
[global]
log_level = "info"
```

- **Binary**: Points to the compiled executable of the project.
- **Listeners**: Configures specific ports for each service, ensuring that each listens on its intended port.

## Main Go File (`main.go`)

The `main.go` in the project root is set up to use Service Weaver for managing services. It initializes the services and uses Weaver's capabilities to handle network listeners and service lifecycle.

```go
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/ServiceWeaver/weaver"
    "github.com/gofiber/adaptor/v2"
    "github.com/gofiber/fiber/v2"

    admin "admin-service/cmd"
    user "user-service/cmd"
)

type App struct {
    weaver.Implements[weaver.Main]
    UserService  weaver.Listener
    AdminService weaver.Listener
}

func serve(ctx context.Context, a *App) error {
    // Service initialization and HTTP server setup for both user and admin services
}

func main() {
    // Weaver run command to manage the application lifecycle
}
```

- **Type `App`**: Defines the structure required by Weaver, including listeners for each service.
- **`serve` Function**: Manages the initialization and running of HTTP servers for each service.
- **`main` Function**: Starts the Weaver management of the app.

## Conclusion

