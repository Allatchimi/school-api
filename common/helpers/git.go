package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"api/common/utils"
	"api/config"
)

const (
	repoDir            = ".repository"
	deployedSchoolsDir = "deployed"
	deletedSchoolsDir  = "deleted"
)

// gitPush pushes to the repository.
func gitPush(commitMessage string, branch string) (ok bool, err error) {
	// Check for changes
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		errMsg := "Failed to check git status!"
		err = fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
		return
	}
	if len(output) < 1 {
		return
	}

	// Commit and push the changes to GitHub
	commands := [][]string{
		{"git", "add", "."},
		{"git", "commit", "-m", commitMessage},
		{"git", "push", "origin", branch},
	}

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Run(); err != nil {
			errMsg := "Failed to push changes to GitHub!"
			err = fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
			return
		}
	}
	ok = true
	return
}

// gitClone clones the repository into a directory.
func gitClone(outputDir string, branch string, withSubmodules bool) (err error) {
	if err != nil {
		if err = exec.Command(
			"git", "clone", "--recurse-submodules", config.Env.GitRepoSshUrl, outputDir,
		).Run(); err != nil {
			errMsg := "Failed to clone repo!"
			err = fmt.Errorf("%s: %s %s %w", errMsg, outputDir, err.Error(), err)
			return
		}
	} else {
		if err = exec.Command(
			"git", "clone", config.Env.GitRepoSshUrl, outputDir,
		).Run(); err != nil {
			errMsg := "Failed to clone repo!"
			err = fmt.Errorf("%s: %s %s %w", errMsg, outputDir, err.Error(), err)
			return
		}
	}

	// Switch to the specified branch
	if len(branch) < 1 {
		branch = "main" // Default branch if not specified
	}
	if err = exec.Command("git", "checkout", branch).Run(); err != nil {
		errMsg := "Failed to switch to git branch!"
		err = fmt.Errorf("%s: %s %w", errMsg, branch, err)
		return
	}
	return
}

// GitPushSchoolDeployment pushes the generated files to the repository.
func GitPushSchoolDeployment(schoolID string, baseDir string, filesDir string) (ok bool, err error) {
	repoDir := filepath.Join(baseDir, repoDir)
	os.RemoveAll(repoDir)

	// Clone the repo into a directory
	if err = config.GitDistributedLock(func() error {
		return gitClone(repoDir, config.Env.GitRepoBranch, false)
	}); err != nil {
		return
	}

	// Define source and destination directories for the school deployment files
	srcDir := filepath.Join(filesDir)
	dstDir := filepath.Join(repoDir, deployedSchoolsDir)

	// Ensure the parent directory exists
	if err = os.MkdirAll(dstDir, os.ModePerm); err != nil {
		errMsg := "Failed to create schools directory!"
		err = fmt.Errorf("%s: %s %w", errMsg, dstDir, err)
		return
	}

	// Remove existing school files to have a clean directory
	os.RemoveAll(filepath.Join(dstDir, schoolID))
	// Remove school from deleted folder
	os.RemoveAll(filepath.Join(repoDir, deletedSchoolsDir, schoolID))

	// Copy the generated deployment files into the cloned repo
	if err = utils.CopyDir(srcDir, dstDir); err != nil {
		errMsg := "Failed to copy deployment files!"
		err = fmt.Errorf("%s: %s %s %s %w", errMsg, srcDir, dstDir, err.Error(), err)
		return
	}

	// Change working directory to the cloned repo for git operations
	if err = os.Chdir(repoDir); err != nil {
		errMsg := "Failed to change directory to repo!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, repoDir, err.Error(), err)
		return
	}

	// Push the changes to GitHub
	err = config.GitDistributedLock(func() error {
		ok, err = gitPush(fmt.Sprintf("Deploy school %s", schoolID), config.Env.GitRepoBranch)
		return err
	})

	return
}

// GitPushDeletedSchoolDeployment pushes the deleted school files to the repository.
func GitPushDeletedSchoolDeployment(schoolID string, baseDir string) (ok bool, err error) {
	repoDir := filepath.Join(baseDir, repoDir)

	// Clone the repo into a directory
	if err = config.GitDistributedLock(func() error {
		return gitClone(repoDir, config.Env.GitRepoBranch, false)
	}); err != nil {
		return
	}

	// Add school to deleted folder by creating a directory and adding a .gitkeep file
	deletedSchoolDir := filepath.Join(repoDir, deletedSchoolsDir, schoolID)
	os.MkdirAll(deletedSchoolDir, os.ModePerm)
	if err = os.WriteFile(filepath.Join(deletedSchoolDir, ".gitkeep"), nil, os.ModePerm); err != nil {
		errMsg := "Failed to create .gitkeep file!"
		err = fmt.Errorf("%s: %s %w", errMsg, filepath.Join(deletedSchoolDir, ".gitkeep"), err)
		return
	}

	// Remove existing school from deploys folder
	deployedSchoolDir := filepath.Join(repoDir, deployedSchoolsDir, schoolID)
	os.RemoveAll(deployedSchoolDir)

	// Change working directory to the cloned repo for git operations
	if err = os.Chdir(repoDir); err != nil {
		errMsg := "Failed to change directory to repo!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, repoDir, err.Error(), err)
		return
	}

	// Push the changes to GitHub
	err = config.GitDistributedLock(func() error {
		ok, err = gitPush(fmt.Sprintf("Deploy school %s", schoolID), config.Env.GitRepoBranch)
		return err
	})
	return
}

// GitSetupSSHKey writes the private SSH key to the user's .ssh directory for GitHub access.
func GitSetupSSHKey() error {
	// Retrieve the private key from environment variable or fallback config
	privateKey := strings.ReplaceAll(os.Getenv("GIT_DEPLOY_REPO_SSH_ED25519_PRIVATE_KEY"), `\n`, "\n")
	if len(privateKey) < 1 {
		privateKey = strings.ReplaceAll(config.Env.GitRepoSshEd25519PrivateKey, `\n`, "\n")
		if len(privateKey) < 1 {
			return fmt.Errorf("GIT_DEPLOY_REPO_SSH_ED25519_PRIVATE_KEY is not set")
		}
	}

	// Ensure the ~/.ssh directory exists with correct permissions
	sshDir := filepath.Join(os.Getenv("HOME"), ".ssh")
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		errMsg := "Failed to create .ssh directory!"
		return fmt.Errorf("%s: %s %s %w", errMsg, sshDir, err.Error(), err)
	}

	// Write the private key to ~/.ssh/id_ed25519 with restrictive permissions
	keyPath := filepath.Join(sshDir, "id_ed25519")
	if err := os.WriteFile(keyPath, []byte(privateKey), 0600); err != nil {
		errMsg := "Failed to write private SSH key!"
		return fmt.Errorf("%s: %s %s %w", errMsg, keyPath, err.Error(), err)
	}

	// Preload GitHub's SSH host key to known_hosts to avoid interactive prompts
	if err := GitPreloadGitHubSSHKey(); err != nil {
		return err
	}

	return nil
}

// GitPreloadGitHubSSHKey adds GitHub's SSH host key to known_hosts to prevent prompt on first connection.
func GitPreloadGitHubSSHKey() error {
	// Use ssh-keyscan to fetch GitHub's SSH public key fingerprint
	cmd := exec.Command("ssh-keyscan", "github.com")
	output, err := cmd.Output()
	if err != nil {
		errMsg := "Failed to scan GitHub SSH key!"
		return fmt.Errorf("%s: %s %s %w", errMsg, "github.com", err.Error(), err)
	}

	// Append GitHub's SSH key to known_hosts file (creating it if necessary)
	knownHostsPath := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
	file, err := os.OpenFile(knownHostsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		errMsg := "Failed to open known_hosts file!"
		return fmt.Errorf("%s: %s %s %w", errMsg, knownHostsPath, err.Error(), err)
	}
	defer file.Close()

	if _, err := file.Write(output); err != nil {
		errMsg := "Failed to write to known_hosts!"
		return fmt.Errorf("%s: %s %s %w", errMsg, knownHostsPath, err.Error(), err)
	}

	return nil
}
