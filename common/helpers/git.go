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

// gitPush pushes to the repository.
func gitPush(repoDir string, commitMessage string, branch string) (err error) {
	Logger.Info("Pushing changes to GitHub...")
	commands := [][]string{
		{"git", "add", "."},
		{"git", "commit", "-m", commitMessage},
		{"git", "push", "origin", branch},
	}
	for _, args := range commands {
		commitCmd := exec.Command(args[0], args[1:]...)
		commitCmd.Dir = repoDir
		if errCommit := commitCmd.Run(); errCommit != nil {
			errMsg := "Failed to push changes to GitHub!"
			err = fmt.Errorf("%s: %s %w", errMsg, errCommit.Error(), errCommit)
			return
		}
	}
	return
}

// gitClone clones the repository into a directory.
func gitClone(repoDir string, branch string, withSubmodules bool) (err error) {
	Logger.Info("Cloning repository...")
	if withSubmodules {
		if err = exec.Command(
			"git", "clone", "--recurse-submodules", config.Env.GitRepoSshUrl, repoDir,
		).Run(); err != nil {
			errMsg := "Failed to clone repo!"
			err = fmt.Errorf("%s: %s %s %w", errMsg, repoDir, err.Error(), err)
			return
		}
	} else {
		if err = exec.Command(
			"git", "clone", config.Env.GitRepoSshUrl, repoDir,
		).Run(); err != nil {
			errMsg := "Failed to clone repo!"
			err = fmt.Errorf("%s: %s %s %w", errMsg, repoDir, err.Error(), err)
			return
		}
	}

	// Switch to the specified branch
	if len(branch) < 1 {
		branch = "main" // Default branch if not specified
	}
	Logger.Info("Checking out repository...")
	checkoutCmd := exec.Command("git", "checkout", branch)
	checkoutCmd.Dir = repoDir
	if outCheckout, errCheckout := checkoutCmd.CombinedOutput(); errCheckout != nil {
		Logger.Warn("Failed to switch to git branch! Creating branch and switching to branch instead...", zap.String("gitOutput", string(outCheckout)))
		createBranchCmd := exec.Command("git", "checkout", "-b", branch)
		createBranchCmd.Dir = repoDir
		if outCreate, errCreate := createBranchCmd.CombinedOutput(); errCreate != nil {
			errMsg := "Failed to create and switch to git branch!"
			Logger.Error(errMsg, zap.String("git output", string(outCreate)))
			err = fmt.Errorf("%s: %s\nGit output: %s", errMsg, branch, string(outCreate))
			return
		}
	}
	return
}

// gitClone clones the repository into a directory.
func gitPull(repoDir string, branch string) (err error) {
	Logger.Info("Pulling repository...")
	pullCmd := exec.Command("git", "pull", "origin", branch)
	pullCmd.Dir = repoDir
	if outPull, errPull := pullCmd.CombinedOutput(); errPull != nil {
		errMsg := "Failed to pull git changes!"
		Logger.Error(errMsg, zap.String("git output", string(outPull)))
		err = fmt.Errorf("%s: %s\nGit output: %s", errMsg, branch, string(outPull))
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
	// Evaluate and add GitHub's SSH key
	cmdEval := exec.Command("sh", "-c", "eval", "\"$(ssh-agent -s)\"")
	_, errEval := cmdEval.Output()
	if errEval != nil {
		errMsg := "Failed to evaluate GitHub SSH key!"
		return fmt.Errorf("%s: %s %s %w", errMsg, "github.com", errEval.Error(), errEval)
	}
	cmdAgent := exec.Command("ssh-add", "/root/.ssh/id_ed25519")
	_, errAgent := cmdAgent.Output()
	if errAgent != nil {
		errMsg := "Failed to add GitHub SSH key to ssh-agent!"
		return fmt.Errorf("%s: %s %s %w", errMsg, "github.com", errAgent.Error(), errAgent)
	}

	// Use ssh-keyscan to fetch GitHub's SSH public key fingerprint
	cmdScan := exec.Command("ssh-keyscan", "github.com")
	outputScan, errScan := cmdScan.Output()
	if errScan != nil {
		errMsg := "Failed to scan GitHub SSH key!"
		return fmt.Errorf("%s: %s %s %w", errMsg, "github.com", errScan.Error(), errScan)
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
