//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

var (
	owner   = "rallinator7"
	baseDir = getMageDir()
)

func gitCommit() string {
	s, e := sh.Output("git", "rev-parse", "--short", "HEAD")
	if e != nil {
		fmt.Printf("Failed to get GIT version: %s\n", e)
		return ""
	}
	return s
}

func getMageDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	return dir
}

type DockerImage struct {
	RelativePath string
	Name         string
	Tag          string
}

// builds the akita-poker server image
func Build() error {
	// get current dir mage is running in
	dir := baseDir

	builds := []DockerImage{
		{RelativePath: "..", Name: "ghcr.io/rallinator7/greeter-app", Tag: gitCommit()},
	}

	for _, build := range builds {
		os.Chdir(filepath.Join(dir, build.RelativePath))

		err := sh.Run("docker", "build", "-f", filepath.Join("app", "docker", "Dockerfile"), "-t", fmt.Sprintf("%s:%s", build.Name, build.Tag), "-t", fmt.Sprintf("%s:%s", build.Name, "latest"), ".")
		if err != nil {
			return fmt.Errorf("could not build docker image: %s", err)
		}

		err = sh.Run("rm", "akita-poker-server")
		if err != nil {
			return fmt.Errorf("could not remove binary: %s", err)
		}
	}

	return nil
}

// pushes images to GitHub
func PushImages() error {
	images := []string{"greeter-app"}
	commit := gitCommit()
	for _, image := range images {
		err := sh.Run("docker", "push", fmt.Sprintf("ghcr.io/%s/%s:%s", owner, image, commit))
		if err != nil {
			return fmt.Errorf("could not push docker image: %s", err)
		}
		err = sh.Run("docker", "push", fmt.Sprintf("ghcr.io/%s/%s:latest", owner, image))
		if err != nil {
			return fmt.Errorf("could not push docker image: %s", err)
		}
	}

	return nil
}
