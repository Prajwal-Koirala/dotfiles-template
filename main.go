package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

var (
	// Git
	gitConfigPath        = fmt.Sprint(userDirectory() + "/.gitconfig")
	gitConfigPathContent []byte
	// SSH
	sshKeysPath          = fmt.Sprint(userDirectory() + "/.ssh")
	sshConfigPath        = fmt.Sprint(sshKeysPath + "/config")
	sshConfigPathContent []byte
	privateSSHKey        = fmt.Sprint(sshKeysPath + "/id_ssh")
	privateSSHKeyContent []byte
	// GPG
	privateGPGKey        = fmt.Sprint(sshKeysPath + "/id_gpg")
	privateGPGKeyContent []byte
	// Handle error
	err error
)

func init() {
	// Check software requirements
	commandExists("git")
	commandExists("gpg")
	// Read the content files
	/* Git */
	gitConfigPathContent, err := os.ReadFile(gitConfigPath)
	handleErrors(err)
	/* SSH */
	sshConfigPathContent, err := os.ReadFile(sshConfigPath)
	handleErrors(err)
	privateSSHKeyContent, err := os.ReadFile(privateSSHKey)
	handleErrors(err)
	/* GPG */
	privateGPGKeyContent, err := os.ReadFile(privateGPGKey)
	handleErrors(err)
}

func main() {
	installSSHKeys()
}

func installSSHKeys() {
	// Make sure we have the ssh folder
	if !folderExists(sshKeysPath) {
		err = os.Mkdir(sshKeysPath, 0700)
		handleErrors(err)
	}
	// Git
	if len(gitConfigPathContent) > 1 {
		if fileExists(gitConfigPath) {
			os.Remove(gitConfigPath)
		}
		err = os.WriteFile(gitConfigPath, []byte(gitConfigPathContent), 0600)
		handleErrors(err)
	}
	// SSH Config
	if len(sshConfigPathContent) > 1 {
		if fileExists(sshConfigPath) {
			os.Remove(sshConfigPath)
		}
		err = os.WriteFile(sshConfigPath, []byte(sshConfigPathContent), 0600)
		handleErrors(err)
	}
	// SSH Key
	if len(privateSSHKeyContent) > 1 {
		if fileExists(privateSSHKey) {
			os.Remove(privateSSHKey)
		}
		err = os.WriteFile(privateSSHKey, []byte(privateSSHKeyContent), 0600)
		handleErrors(err)
	}
	// GPG
	if len(privateGPGKeyContent) > 1 {
		if fileExists(privateGPGKey) {
			os.Remove(privateGPGKey)
		}
		err = os.WriteFile(privateGPGKey, []byte(privateGPGKeyContent), 0600)
		handleErrors(err)
	}
}

func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func userDirectory() string {
	user, err := user.Current()
	handleErrors(err)
	return user.HomeDir
}

func commandExists(cmd string) {
	cmd, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatalf("Error: The application %s was not found in the system.\n", cmd)
	}
}

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
