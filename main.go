package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Check if you can run the go get

// Embed the `templates` directory and its contents into the binary.
//
//go:embed templates/*
var templates embed.FS

// Define the project file structure and the corresponding template paths in the `templates` folder.
var fileStructure = map[string]string{
	"/server/start.go":                 "templates/start.txt",
	"/server/server.go":                "templates/server.txt",
	"/server/routes.go":                "templates/routes.txt",
	"/server/helpers.go":               "templates/helpers.txt",
	"/api/api.go":                      "templates/api.txt",
	"/api/config.go":                   "templates/config.txt",
	"/api/context.go":                  "templates/context.txt",
	"/api/errors.go":                   "templates/errors.txt",
	"/api/helpers.go":                  "templates/api_helpers.txt",
	"/api/middleware.go":               "templates/middleware.txt",
	"/api/server.go":                   "templates/api_server.txt",
	"/internal/tests/assert/assert.go": "templates/assert.txt",
	"/internal/tests/testdata.go":      "templates/testdata.txt",
	"/internal/tests/testutils.go":     "templates/testutils.txt",
	"/validator/validator.go":          "templates/validator.txt",
	"/users/user_model.go":             "templates/user_model.txt",
	"/users/user_service.go":           "templates/user_service.txt",
	"/users/user_store.go":             "templates/user_store.txt",
	"/cmd/{projectName}/main.go":       "templates/main.txt",
	"/Makefile":                        "templates/makefile.txt",
	"/.envrc":                          "templates/env.txt",
}

func main() {
	// Define CLI flags for the repository name and description.
	repoName := flag.String("repo", "MyProject", "The name of the repository")
	description := flag.String("desc", "A sample project structure created by Go CLI", "The description of the project")
	repoPath := flag.String("repo-path", "github.com/yourusername/yourrepo", "The path/url of the repository")
	flag.Parse()

	// Replace placeholders in each file template with the provided repo name and description.
	for targetPath, templatePath := range fileStructure {
		fullTargetPath := "." + strings.ReplaceAll(targetPath, "{projectName}", *repoName)
		dir := filepath.Dir(fullTargetPath)

		// Create directories if they don’t exist.
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", dir, err)
			continue
		}

		// Read content from the embedded template file.
		content, err := templates.ReadFile(templatePath)
		if err != nil {
			fmt.Printf("Failed to read template file %s: %v\n", templatePath, err)
			continue
		}

		// Replace placeholders with actual values.
		updatedContent := strings.ReplaceAll(string(content), "{{RepoName}}", *repoName)
		updatedContent = strings.ReplaceAll(updatedContent, "{{Description}}", *description)
		updatedContent = strings.ReplaceAll(updatedContent, "{{RepoPath}}", *repoPath)

		// Write the updated content to the target file.
		if err := os.WriteFile(fullTargetPath, []byte(updatedContent), 0644); err != nil {
			fmt.Printf("Failed to create file %s: %v\n", fullTargetPath, err)
			continue
		}

		fmt.Printf("Created file %s with content from %s.\n", fullTargetPath, templatePath)
	}

	fmt.Printf("Project %s created successfully!\n", *repoName)
}
