package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

var (
	gitConfig     = fmt.Sprint(userDirectory() + "/.gitconfig")
	keysDir       = fmt.Sprint(userDirectory() + "/.ssh/")
	sshConfig     = fmt.Sprint(keysDir + "config")
	privateSSHKey = fmt.Sprint(keysDir + "id_ssh")
	privateGPGKey = fmt.Sprint(keysDir + "id_gpg")
)

var (
	gitConfigContent = ``
	sshConfigContent = ``
	gpgKeyContent    = ``
)

func init() {
	if !commandExists("gpg") {
		log.Fatal("Error: The application gpg was not found in the system.")
	}
	if !commandExists("git") {
		log.Fatal("Error: The application git was not found in the system.")
	}
}

func main() {
	installSSHKeys()
}

func installSSHKeys() {
	var pathGPG string
	switch runtime.GOOS {
	case "darwin":
		pathGPG = "program = /opt/homebrew/bin/gpg"
	case "linux":
		pathGPG = "program = /usr/bin/gpg"
	case "windows":
		pathGPG = "program = C:\\Program Files (x86)\\GnuPG\\bin\\gpg.exe"
	}
	if !folderExists(keysDir) {
		os.Mkdir(keysDir, 0700)
	}
	os.WriteFile(gitConfig, []byte(gitConfigContent), 0600)
	os.WriteFile(sshConfig, []byte(sshConfigContent), 0600)
	os.WriteFile(privateSSHKey, []byte(sshKeyContent), 0600)
	os.WriteFile(privateGPGKey, []byte(gpgKeyContent), 0600)
	// start changing the content for gpg
	read, err := os.ReadFile(gitConfigContent)
	if err != nil {
		log.Println(err)
	}
	newContentsGPG := strings.Replace(string(read), ("program ="), (pathGPG), -1)
	os.WriteFile(gitConfigContent, []byte(newContentsGPG), 0)
	// start changing the content for git config
	readSSHconfig, err := os.ReadFile(sshConfig)
	if err != nil {
		log.Println(err)
	}
	newContents := strings.Replace(string(readSSHconfig), ("~/.ssh/id_ssh"), (privateSSHKey), -1)
	os.WriteFile(sshConfig, []byte(newContents), 0)
}

func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
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
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}

func commandExists(cmd string) bool {
	cmd, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	_ = cmd
	return true
}
