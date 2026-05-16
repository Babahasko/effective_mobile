//go:build mage
package main

import (
    "github.com/magefile/mage/sh"
)
const composeFile = "docker/docker-compose.yml"
const envFile = ".env"

func Up() error {
    return sh.Run("docker", "compose", "--env-file", envFile, "-f", composeFile, "up", "-d")
}

func Down() error {
    return sh.Run("docker", "compose", "--env-file", envFile, "-f", composeFile, "down")
}

func Build() error {
    return sh.Run("docker", "compose", "--env-file", envFile, "-f", composeFile, "up", "-d", "--build")
}

func Migrate() error {
    return sh.Run("docker", "compose", "--env-file", envFile, "-f", composeFile, "run", "--rm", "app", "go", "run", "./cmd/migrations/main.go")
}