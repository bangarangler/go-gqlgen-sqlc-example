// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func SQLCGen() error {
	fmt.Println("sqlc generating queries...")
	system := runtime.GOOS
	switch system {
	case "windows":
		println("No Thank You, Switch to Linux ; )")
	case "darwin":
		println("Running on mac")
		return sh.Run("docker", "run", "--rm", "-v", "/Users/jonathanpalacio/Desktop/go-gqlgen-sqlc-example:/src", "-w", "/src", "kjconroy/sqlc", "generate")
	case "linux":
		println("Linux ; )")
		return sh.Run("docker", "run", "--rm", "-v", "/home/jonathan/Desktop/go-gqlgen-sqlc-example:/src", "-w", "/src", "kjconroy/sqlc", "generate")
	}
	return nil
}

func GQLGen() error {
	fmt.Println("gql generating queries...")
	return sh.Run("go", "run", "gqlgen", "generate")
}

func hackPermissions() error {
	fmt.Println("sqlc generating queries...")
	// TODO: worked on mac no problem. only on linux it locked me out and i needed
	// this
	// TODO: add sudo chmod o+w ./pg/*
	return nil
}

func Server() error {
	fmt.Println("running server ...")
	return sh.Run("go", "run", "cmd/go-gqlgen-sqlc-example/main.go")
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return cmd.Run()
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./MyApp", "/usr/bin/MyApp")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}
