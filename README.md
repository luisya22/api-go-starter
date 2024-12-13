# api-go-starter
API Go Starter is a CLI tool to generate a pre-defined Go project structure for building APIs. It simplifies scaffolding a new API project with a consistent and organized structure.

# Features
- Generates a structured API project template in Go.
- Customizable repository name, description, and repository path.
- Prepares common files and directories like server, api, users, and more.
= Built-in templates for faster setup.
- Future updates will include additional project templates and customization options.

# Installation
Follow these steps to clone, build, and install the API Go Starter tool on an Ubuntu system.

## Prerequisites
Ensure following are installed:
- Go
- Git

## Clone the Repository
```bash
git clone https://github.com/yourusername/api-go-starter.git
cd api-go-starter
```

## Build and Install
Build the binary and install it globally to use the tool from anywhere:
```bash
go build -o api-go-starter
sudo mv api-go-starter /usr/local/bin
echo 'export PATH=$PATH:/path/to/your/binary' >> ~/.bashrc
source ~/.bashrc
```
You can now run `api-go-starter` globally.

# Usage
Run the tool with the following options:

```bash
api-go-starter -repo <RepoName> -desc <Description> -repo-path <RepoPath>
```
## Example
To generate a project named MyAPI:

```bash
api-go-starter -repo MyAPI -desc "A Go-based API starter project" -repo-path github.com/yourusername/myapi
```
This will create a new project structure like:

```bash
MyAPI/
├── server/
│   ├── start.go
│   ├── server.go
│   ├── routes.go
│   ├── helpers.go
├── api/
│   ├── api.go
│   ├── config.go
│   ├── context.go
│   ├── errors.go
│   ├── helpers.go
│   ├── middleware.go
│   ├── server.go
├── internal/tests/
│   ├── assert/
│   │   └── assert.go
│   ├── testdata.go
│   ├── testutils.go
├── validator/
│   └── validator.go
├── users/
│   ├── user_model.go
│   ├── user_service.go
│   ├── user_store.go
├── cmd/MyAPI/
│   └── main.go
├── Makefile
└── .envrc
```

# File Structure
- server/: Contains server-related files like routes and helpers.
- api/: Handles API configurations, middleware, and server logic.
- internal/tests/: Utilities and helpers for testing.
- validator/: Custom validation logic.
- users/: Models, services, and storage for user management.
- cmd/: Entry point for the application.
- Makefile: Build and automation commands.
- .envrc: Environment configuration file.

# License
This project is licensed under the MIT License. See the LICENSE file for more information.
