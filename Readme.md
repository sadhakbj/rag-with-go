### Code structure

```
your-app/
│── cmd/ # Main application entry point
│ └── main.go # Main Go file to run the application

│── internal/ # Application-specific code, internal logic
│ ├── app/ # Core business logic for the app
│ │ └── app.go # Application setup and main logic
│ ├── di/ # Dependency Injection container
│ │ └── container.go # DI management and registration
│ ├── services/ # Business-specific services
│ │ ├── github/ # GitHub-specific service
│ │ └── ollama/ # Ollama-specific service
│ └── http/ # Internal HTTP-related functionality
│ └── client.go # HTTP client or related utilities

│── pkg/ # Reusable packages across different applications
│ └── httpclient/ # HTTP client package
│ └── client.go # HTTP client implementation

│── config/ # Configuration files and logic
│ ├── config.go # Load and manage configuration settings
│ └── config.yaml # Configuration file (can be YAML, JSON, etc.)

│── utils/ # Utility functions and helpers
│ ├── logger/ # Logger utility
│ │ └── logger.go # Logger implementation
│ └── helpers/ # Generic helper functions
│ └── utils.go # Utility functions like date formatting, etc.

│── go.mod # Go module file
│── go.sum # Go checksum file for dependencies
└── README.md # Project documentation
```
