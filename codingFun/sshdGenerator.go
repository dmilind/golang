/*
sshdGenerator to keep sshd_config file updated always. This program parses all files provided under /etc/ssh/allowed_groups, collects all group, and all those groups will be added
AllowGroups in sshd_config file. Also this will make sure not to add duplicate value in sshd_config.
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
	destDir           = "/etc/ssh/allowed_groups/"
	serviceFiles      []string
	allowedGroups     []string
	groups            []string
	allowedGroupsLine string
)

const (
	sshdFile = "/etc/ssh/sshd_config"
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
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Pass flag 'update' for updating sshd_config")
		os.Exit(1)
	}
	if args[1] == "update" {
		_, err := os.Stat(destDir)
		if err != nil {
			fmt.Println("Directory missing -->", destDir)
			os.Exit(1)
		}
		allFiles, err := ioutil.ReadDir(destDir)
		if err != nil {
			fmt.Println("Failed to read the directory -->", destDir)
		}
		for _, f := range allFiles {
			serviceFiles = append(serviceFiles, f.Name())
		}
		if len(serviceFiles) == 0 {
			fmt.Println("Updating sshd_config with all default values!")
			os.Exit(0)
		}

		fmt.Println("Available files coming from services are -->", serviceFiles)
		for _, file := range serviceFiles {
			allowedGroups = Groups(destDir + file)
		}
		uniqueAllowGroups := make([]string, 0)
		for _, val := range allowedGroups {
			if !isInSlice(val, uniqueAllowGroups) {
				uniqueAllowGroups = append(uniqueAllowGroups, val)
			}
		}
		fmt.Println("Required sshd groups to be injected in sshd_config -->", uniqueAllowGroups)
		UpdateConf(uniqueAllowGroups)
		_, err = exec.Command("systemctl", "restart", "sshd").Output()
		if err != nil {
			fmt.Println("Error while starting sshd daemon")
			os.Exit(1)
		}
		fmt.Println("Updated sshd_config and restarted sshd service successfully!")
	} else {
		fmt.Println("Expecting 'update' flag to update sshd_config")
	}
}
