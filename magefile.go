//go:build mage

package main

import (
    "github.com/magefile/mage/sh"
)

// Запустить все сервисы
func Up() error {
    return sh.Run("docker-compose", "-f", "docker/docker-compose.yml", "up", "-d")
}

// Остановить все сервисы
func Down() error {
    return sh.Run("docker-compose", "-f", "docker/docker-compose.yml", "down")
}

// Запустить миграции
func Migrate() error {
    return sh.Run("docker-compose", "-f", "docker/docker-compose.yml", "run", "--rm", "app", "go", "run", "./cmd/migrations/main.go")
}

// Пересобрать и запустить
func Build() error {
    return sh.Run("docker-compose", "-f", "docker/docker-compose.yml", "up", "-d", "--build")
}