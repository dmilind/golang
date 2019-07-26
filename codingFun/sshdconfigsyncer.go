/*
sshdConfigSyncer to keep sshd_config file updated always. This program parses all files provided under /etc/ssh/allowed_groups, collects all group, and all those groups will be added
AllowGroups in sshd_config file. Also no duplicate value will be kept in sshd_config file.
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

var (
	destDir = "/etc/ssh/allowed_groups/"
	//destDir           = "/Users/mdhoke/Documents/kachara/ansible/allowed_groups/" // remove this line later
	serviceFiles      []string
	allowedGroups     []string
	groups            []string
	allowedGroupsLine string
)

const (
	sshdFile = "/etc/ssh/sshd_config"
	//sshdFile = "/Users/mdhoke/Documents/kachara/ansible/sshd_config" // change the value of this param later
)

// isInSlice function to remove all duplicate sshd groups from provided allowed groups slice
func isInSlice(str string, rawGroup []string) bool {
	for _, val := range rawGroup {
		if val == str {
			return true
		}
	}
	return false
}

// Groups  function to read a service file line by line containing sshd group. After traversing all files, alphabetically sorted slice of all required sshd groups is created.
func Groups(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error while opening a file -->", fileName)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	defer file.Close()
	for _, line := range fileLines {
		groups = append(groups, line)
		sort.Strings(groups)
	}
	return groups
}

// UpdateConf function updates the sshd_config file in order to place all allowed groups.Already sorted slice of sshd groups parsed here.
func UpdateConf(allowedGroups []string) {
	allowedGroupsLine = "AllowGroups"
	for _, g := range allowedGroups {
		allowedGroupsLine += " " + g
	}
	file, err := ioutil.ReadFile(sshdFile)
	if err != nil {
		fmt.Println("Error while reading a file: -->", sshdFile)
	}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		if strings.Contains(line, "AllowGroups") {
			lines[i] = allowedGroupsLine
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(sshdFile, []byte(output), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	_, err := os.Stat(destDir)
	if err != nil {
		fmt.Println("Directory missing", destDir)
		os.Exit(1)
	}
	allFiles, err := ioutil.ReadDir(destDir)
	if err != nil {
		fmt.Println("Failed to read the directory")
	}
	for _, f := range allFiles {
		serviceFiles = append(serviceFiles, f.Name())
	}
	if len(serviceFiles) == 0 {
		fmt.Println("No service's allowed group file found, keeping default sshd groups!")
		os.Exit(0)
	}

	fmt.Println("Available service's allowed_groups files are -->", serviceFiles)
	for _, file := range serviceFiles {
		allowedGroups = Groups(destDir + file)
	}
	uniqueAllowGroups := make([]string, 0)
	for _, val := range allowedGroups {
		if !isInSlice(val, uniqueAllowGroups) {
			uniqueAllowGroups = append(uniqueAllowGroups, val)
		}
	}

	fmt.Println("Required sshd groups to be injected:-->", uniqueAllowGroups)
	UpdateConf(uniqueAllowGroups)
	_, err = exec.Command("systemctl", "restart", "sshd").Output()
	if err != nil {
		fmt.Println("Error while starting sshd daemon")
		os.Exit(1)
	}
}
