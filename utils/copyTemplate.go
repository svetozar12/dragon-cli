package utils

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
)

func CopyTemplateFromRepo(templateToCopy string, destinationDir string, copyFolderContent bool) error {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user's home directory:", err)
		return err
	}

	templateDir := path.Join(usr.HomeDir, "dragon-cli-tmp/dragon-cli", templateToCopy)
	if copyFolderContent {
		templateDir = templateDir + "/."
	}

	if err := copyTemplate(templateDir, destinationDir); err != nil {
		fmt.Println("Error removing unnecessary files:", err)
		return err
	}

	return nil
}

func CloneTemplateRepo(tmpRepoDir string) error {
	repoURL := "https://github.com/svetozar12/dragon-cli"

	if err := os.MkdirAll(tmpRepoDir, os.ModePerm); err != nil {
		fmt.Println("Error creating temporary directory:", err)
		return err
	}
	cmd := exec.Command("git", "clone", "--depth", "1", "--branch", "master", repoURL)
	cmd.Dir = tmpRepoDir

	if err := cmd.Run(); err != nil {
		fmt.Println("Error cloning repository:", err)
		return err
	}
	return nil
}

func DeleteTemplateRepo(tmpRepoDir string) error {
	cmd := exec.Command("rm", "-r", tmpRepoDir)
	cmd.Dir = tmpRepoDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error deleting repository:", err)
		return err
	}
	return nil
}

func copyTemplate(templateDir, outputDIr string) error {
	cmd := exec.Command("cp", "-a", templateDir, outputDIr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error cloning repository:", err)
		return err
	}
	return nil
}
