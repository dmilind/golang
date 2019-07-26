/*
Custom ansible module written in golang. This program can be used directly either
by using go tool (go run) or exeuting binary built by go tool (go build)
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// global variable declaration
var (
	returnResponse []byte
	response ModuleResponse

)

// ModuleArgs to collect the state provided in argument file, this file will be generated by ansible when module is invoked.
type ModuleArgs struct {
	State string
}

// ModuleResponse struct to provide json formatted data to ansible, ansbile expects json data in return
type ModuleResponse struct {
	Msg     string `json:"msg"`
	Failed  bool   `json:"failed"`
	Changed bool   `json:"changed"`
}

// ReturnResponse function prints json data , data is collected from struct ModuleResponse
func ReturnResponse(response ModuleResponse) {
	retVal, _ := json.Marshal(response)
	fmt.Println(string(retVal))
}

// FileEdit function to edit the contents of a provided file. This function take 3 arguments "filename", "oldline", "newline" all of string type.
// "filename" is the file which will be edited, "oldline" is the line to be edited, "newline" is the line to be added
func FileEdit(filename, oldLine, newLine string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while reading a file -->", filename)
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, oldLine) {
			lines[i] = newLine
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error while writing a file -->", filename)
		os.Exit(1)
	}
}

// FileCopy function copies a content of source file to destination file. Function accepts 2 arguments src, and dest both of string type.
func FileCopy(src, dest string) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error while opening a file", src)
		os.Exit(1)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		fmt.Println("Error while creating a file -->", dest)
		os.Exit(1)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error while copying a file -->", src)
		os.Exit(1)
	}
}

// Executor function to execute system command on the machine.
func Executor(cmd string) {
	output, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println("Error while running command", cmd)
		os.Exit(1)
	}
	fmt.Println(string(output)) // change later
}

// EnableGoogleAuth function will enable google authentication.
func EnableGoogleAuth() {
	output, err := exec.Command("/bin/google-authenticator", "-t", "-d", "-f", "-r", "3", "-R", "30", "-W").Output()
	if err != nil {
		fmt.Println("Error while running command --> google-authenticator")
		os.Exit(1)
	}
	google_secrets := string(output)
	// fmt.Println("copying file a file --> /etc/pam.d/sshd")
	FileCopy("/etc/pam.d/sshd", "/home/vagrant/sshd.back")

	// fmt.Println("copying file 2")
	FileCopy("/etc/ssh/sshd_config", "/home/vagrant/sshd_config.back")

	// fmt.Println("editing file 1")
	FileEdit("/etc/pam.d/sshd", "auth       substack     password-auth", "#auth       substack     password-auth")

	// fmt.Println("adding a line in a file --> /etc/pam.d/sshd")
	LineInPam := "auth required pam_google_authenticator.so nullok"
	file, err := os.OpenFile("/etc/pam.d/sshd", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error while opening a file --> /etc/pam.d/sshd")
		os.Exit(1)
	}

	defer file.Close()
	_, err = file.WriteString(LineInPam)
	if err != nil {
		fmt.Println("Error while writing into a file --> /etc/pam.d/sshd")
		os.Exit(1)
	}

	// fmt.Println("Editing a file --> /etc/ssh/sshd_config")
	FileEdit("/etc/ssh/sshd_config", "#ChallengeResponseAuthentication yes", "ChallengeResponseAuthentication yes")

	// fmt.Println("Editing a file --> /etc/ssh/sshd_config")
	FileEdit("/etc/ssh/sshd_config", "ChallengeResponseAuthentication no", "#ChallengeResponseAuthentication no")

  LineInSSH := "ClientAliveInterval 120\n"
  file, err = os.OpenFile("/etc/ssh/sshd_config", os.O_APPEND|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println("Error while opening a file --> /etc/ssh/sshd_config")
    os.Exit(1)
  }

  defer file.Close()
  _, err = file.WriteString(LineInSSH)
  if err != nil {
    fmt.Println("Error while writing a file --> /etc/ssh/sshd_config")
    os.Exit(1)
  }

  LineInSSH = "ClientAliveCountMax 2\n"
  file, err = os.OpenFile("/etc/ssh/sshd_config", os.O_APPEND|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println("Error while opening a file --> /etc/ssh/sshd_config")
    os.Exit(1)
  }

  defer file.Close()
  _, err = file.WriteString(LineInSSH)
  if err != nil {
    fmt.Println("Error while writing a file --> /etc/ssh/sshd_config")
    os.Exit(1)
  }

	// fmt.Println("Adding a line in a file --> /etc/ssh/sshd_config")
	LineInSSH = "AuthenticationMethods publickey,keyboard-interactive\n"
	file, err = os.OpenFile("/etc/ssh/sshd_config", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error while opening a file --> /etc/ssh/sshd_config")
		os.Exit(1)
	}

	defer file.Close()
	_, err = file.WriteString(LineInSSH)
	if err != nil {
		fmt.Println("Error while writing a file --> /etc/ssh/sshd_config")
		os.Exit(1)
	}

	fmt.Println("Restarting sshd service")
	_, err = exec.Command("systemctl", "restart", "sshd").Output()
	if err != nil {
		fmt.Println("Error while starting sshd daemon")
		os.Exit(1)
	}
	response.Msg = google_secrets
	response.Failed = false
	response.Changed = true
	ReturnResponse(response)
}

// DisableGoogleAuth function disables google authentication also removes all configurations of google authenticator
func DisableGoogleAuth() {
	// copy back up files and restart sshd :) simple
	FileCopy("/home/vagrant/sshd.back", "/etc/pam.d/sshd")
	FileCopy("/home/vagrant/sshd_config.back", "/etc/ssh/sshd_config")

	_, err := exec.Command("systemctl", "restart", "sshd").Output()
	if err != nil {
		fmt.Println("Error while starting sshd daemon")
		os.Exit(1)
	}
	response.Msg = "Google two factor authentication is disabled"
	response.Failed = false
	response.Changed = true
	ReturnResponse(response)
}
func main() {
	// check length of the arguments
	if len(os.Args) != 2 {
		response.Msg = "Missing argument file"
		response.Failed = true
		response.Changed = false
		ReturnResponse(response)
		os.Exit(1)
	}

	argFile := os.Args[1]
	// this is file provide by the ansible at ansible's execution runtime
	// now we need to read this file to get the state which is mentioned in ansible play.
	State, err := ioutil.ReadFile(argFile)
	if err != nil {
		response.Msg = "could not read argument file"
		response.Failed = true
		response.Changed = false
		ReturnResponse(response)
	}
	// state contains byte of data which should be converted into actual string.
	// this data is coming from json file, for that we need to use unmarshal function to
	// get this data into a struct.
	var modArg ModuleArgs
	err = json.Unmarshal(State, &modArg)
	if err != nil {
		response.Msg = "error occured while unmarshling"
		response.Failed = true
		response.Changed = false
		ReturnResponse(response)
	}
	// use switch statement to call appropriate function to enable or disable TFA
	switch modArg.State {
	case "enable":
		// check google authentication status
		// if enabled: --> msg that it is already enabled exit 0
		// calling EnableGoogleAuth function
		EnableGoogleAuth()

	case "disable":
		// if disable: --> msg that it is already disabled exit 0
		// calling DisableGoogleAuth Function
		DisableGoogleAuth()

	default:
		response.Msg = "state attribute takes either enable or disable value"
		response.Failed = true
		response.Changed = false
		ReturnResponse(response)
		os.Exit(1)
  }
}
