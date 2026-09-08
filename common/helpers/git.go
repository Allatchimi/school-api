package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"api/common/utils"
	"api/config"

	"go.uber.org/zap"
)

const (
	mRepoDir            = ".repository"
	mDeployedSchoolsDir = "deployed"
	mDeletedSchoolsDir  = "deleted"
)

func newGitCommand(dir string, args ...string) *exec.Cmd {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(),
		`GIT_SSH_COMMAND=ssh -i /root/.ssh/id_ed25519 -o IdentitiesOnly=yes -o StrictHostKeyChecking=yes`)
	return cmd
}

// gitPush pushes to the repository.
func gitPush(repoDir string, commitMessage string, branch string) (err error) {
	Logger.Info("Configuring Git user...")
	configNameCmd := newGitCommand(repoDir, "config", "user.name", "Bot")
	if output, errConfig := configNameCmd.CombinedOutput(); errConfig != nil {
		return fmt.Errorf("Failed to set git user.name: %v, output: %s", errConfig, string(output))
	}
	configEmailCmd := newGitCommand(repoDir, "config", "user.email", "github-actions[bot]@users.noreply.github.com")
	if output, errConfig := configEmailCmd.CombinedOutput(); errConfig != nil {
		return fmt.Errorf("Failed to set git user.email: %v, output: %s", errConfig, string(output))
	}

	Logger.Info("Pushing changes to GitHub...")
	commands := [][]string{
		{"add", "."},
		{"commit", "-m", commitMessage},
		{"push", "origin", branch},
	}
	for _, args := range commands {
		commitCmd := newGitCommand(repoDir, args...)
		output, errCommit := commitCmd.CombinedOutput()
		if errCommit != nil {
			errMsg := "Failed to push changes to GitHub!"
			err = fmt.Errorf("%s: %s\nGit output:\n%s", errMsg, errCommit.Error(), string(output))
			return
		}
	}
	return
}

// gitClone clones the repository into a directory.
func gitClone(repoDir string, branch string, withSubmodules bool) (err error) {
	Logger.Info("Cloning repository...")

	parentDir := filepath.Dir(repoDir)
	if err = os.MkdirAll(parentDir, 0755); err != nil {
		return fmt.Errorf("Failed to create parent directory: %w", err)
	}

	var cloneCmd *exec.Cmd
	if withSubmodules {
		cloneCmd = newGitCommand(parentDir, "clone", "--recurse-submodules", config.Env.GitRepoSshUrl, repoDir)
	} else {
		cloneCmd = newGitCommand(parentDir, "clone", config.Env.GitRepoSshUrl, repoDir)
	}
	output, errClone := cloneCmd.CombinedOutput()
	if errClone != nil {
		errMsg := "Failed to clone repo!"
		err = fmt.Errorf("%s: %s %s\nGit output:\n%s", errMsg, repoDir, errClone.Error(), string(output))
		return
	}

	if branch == "" {
		branch = "main"
	}

	Logger.Info("Checking out repository...")
	checkoutCmd := newGitCommand(repoDir, "checkout", branch)
	outCheckout, errCheckout := checkoutCmd.CombinedOutput()
	if errCheckout != nil {
		Logger.Warn("Failed to switch to git branch! Creating branch and switching to branch instead...", zap.String("gitOutput", string(outCheckout)))
		createBranchCmd := newGitCommand(repoDir, "checkout", "-b", branch)
		outCreate, errCreate := createBranchCmd.CombinedOutput()
		if errCreate != nil {
			errMsg := "Failed to create and switch to git branch!"
			Logger.Error(errMsg, zap.String("git output", string(outCreate)))
			err = fmt.Errorf("%s: %s\nGit output:\n%s", errMsg, branch, string(outCreate))
			return
		}
	}
	return
}

// gitPull pulls latest changes from remote.
func gitPull(repoDir string, branch string) (err error) {
	Logger.Info("Pulling repository...")
	pullCmd := newGitCommand(repoDir, "pull", "origin", branch)
	outPull, errPull := pullCmd.CombinedOutput()
	if errPull != nil {
		errMsg := "Failed to pull git changes!"
		Logger.Error(errMsg, zap.String("git output", string(outPull)))
		err = fmt.Errorf("%s: %s\nGit output:\n%s", errMsg, branch, string(outPull))
		return
	}
	return
}

// GitPushSchoolDeployment pushes the generated files to the repository.
func GitPushSchoolDeployment(schoolID string, baseDir string, filesDir string) (err error) {
	repoDir := filepath.Join(baseDir, mRepoDir)
	os.RemoveAll(repoDir)

	// Clone the repo into a directory
	if err = config.GitDistributedLock(func() error {
		return gitClone(repoDir, config.Env.GitRepoBranch, false)
	}); err != nil {
		Logger.Error("Failed to clone repo!", zap.Error(err))
		return
	}

	// Pull the changes from main branch
	if err = config.GitDistributedLock(func() error {
		return gitPull(repoDir, "main")
	}); err != nil {
		Logger.Warn("Failed to pull git changes! Skipping...")
	}

	// Define source and destination directories for the school deployment files
	srcDir := filepath.Join(filesDir)
	dstDir := filepath.Join(repoDir, mDeployedSchoolsDir)

	// Ensure the parent directory exists
	if err = os.MkdirAll(dstDir, os.ModePerm); err != nil {
		errMsg := "Failed to create schools directory!"
		err = fmt.Errorf("%s: %s %w", errMsg, dstDir, err)
		return
	}

	// Remove existing school files to have a clean directory
	os.RemoveAll(filepath.Join(dstDir, schoolID))
	// Remove school from deleted folder
	os.RemoveAll(filepath.Join(repoDir, mDeletedSchoolsDir, schoolID))

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
		err = gitPush(repoDir, fmt.Sprintf("Deploy school %s", schoolID), config.Env.GitRepoBranch)
		return err
	})

	return
}

// GitPushDeletedSchoolDeployment pushes the deleted school files to the repository.
func GitPushDeletedSchoolDeployment(schoolID string, baseDir string, folderToAdd string) (err error) {
	repoDir := filepath.Join(baseDir, mRepoDir)

	// Clone the repo into a directory
	if err = config.GitDistributedLock(func() error {
		return gitClone(repoDir, config.Env.GitRepoBranch, false)
	}); err != nil {
		return
	}

	// Pull the changes from GitHub
	if err = config.GitDistributedLock(func() error {
		return gitPull(repoDir, config.Env.GitRepoBranch)
	}); err != nil {
		Logger.Warn("Failed to pull git changes! Skipping...")
	}

	// Add school to deleted folder by creating a directory and adding a .gitkeep file
	deletedSchoolDir := filepath.Join(repoDir, mDeletedSchoolsDir, schoolID)
	os.MkdirAll(deletedSchoolDir, os.ModePerm)
	if err = os.WriteFile(filepath.Join(deletedSchoolDir, ".gitkeep"), nil, os.ModePerm); err != nil {
		errMsg := "Failed to create .gitkeep file!"
		err = fmt.Errorf("%s: %s %w", errMsg, filepath.Join(deletedSchoolDir, ".gitkeep"), err)
		return
	}

	// Copy the generated deployment files into the cloned repo
	if len(folderToAdd) > 0 {
		if err = utils.CopyDir(folderToAdd, deletedSchoolDir); err != nil {
			errMsg := "Failed to copy deployment files!"
			err = fmt.Errorf("%s: %s %s %s %w", errMsg, folderToAdd, repoDir, err.Error(), err)
			return
		}
	}

	// Remove existing school from deploys folder
	deployedSchoolDir := filepath.Join(repoDir, mDeployedSchoolsDir, schoolID)
	os.RemoveAll(deployedSchoolDir)

	// Change working directory to the cloned repo for git operations
	if err = os.Chdir(repoDir); err != nil {
		errMsg := "Failed to change directory to repo!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, repoDir, err.Error(), err)
		return
	}

	// Push the changes to GitHub
	err = config.GitDistributedLock(func() error {
		err = gitPush(repoDir, fmt.Sprintf("Deploy school %s", schoolID), config.Env.GitRepoBranch)
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

	keyPathTmp := filepath.Join(sshDir, "id_ed25519.tmp")
	keyPath := filepath.Join(sshDir, "id_ed25519")

	// Write the private key to ~/.ssh/id_ed25519 with restrictive permissions
	if err := os.WriteFile(keyPathTmp, []byte(privateKey), 0600); err != nil {
		errMsg := "Failed to write private SSH key!"
		return fmt.Errorf("%s: %s %s %w", errMsg, keyPathTmp, err.Error(), err)
	}
	// Cleanup the private key
	cmdCleanup := exec.Command("awk", "NF", keyPathTmp)
	output, err := cmdCleanup.Output()
	if err != nil {
		return fmt.Errorf("Failed to cleanup GitHub SSH key: %w", err)
	}
	err = os.WriteFile(keyPath, output, 0600)
	if err != nil {
		return fmt.Errorf("Failed to write cleaned key: %w", err)
	}
	os.Remove(keyPathTmp)

	// Preload GitHub's SSH host key to known_hosts to avoid interactive prompts
	if err := GitPreloadGitHubSSHKey(); err != nil {
		return err
	}

	return nil
}

// GitPreloadGitHubSSHKey adds GitHub's SSH host key to known_hosts to prevent prompt on first connection.
func GitPreloadGitHubSSHKey() error {
	// Restrict the scan to GitHub host key algorithms supported by the image.
	cmdScan := exec.Command("ssh-keyscan", "-T", "10", "-t", "rsa,ecdsa,ed25519", "github.com")
	outputScan, errScan := cmdScan.Output()
	if errScan != nil {
		errMsg := "Failed to scan GitHub SSH key!"
		return fmt.Errorf("%s: %s %w", errMsg, "github.com", errScan)
	}

	// Append GitHub's SSH key to known_hosts file (creating it if necessary)
	knownHostsPath := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
	file, err := os.OpenFile(knownHostsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		errMsg := "Failed to open known_hosts file!"
		return fmt.Errorf("%s: %s %s %w", errMsg, knownHostsPath, err.Error(), err)
	}
	defer file.Close()

	if _, err := file.Write(outputScan); err != nil {
		errMsg := "Failed to write to known_hosts!"
		return fmt.Errorf("%s: %s %s %w", errMsg, knownHostsPath, err.Error(), err)
	}

	return nil
}
