/*
[Working on it]
DockerSwarmSetter:
  This program will configure docker swarm on the remote hosts.

* Provide a configuration file to this program mentioning master and agent nodes listing, somthing like
    docker_master:
      - node 1
      - node 2
      - node 3
    docker_agent:
      - node 4
      - node 5
      - node 6

* block of code which will validate that file is provided or not , if not provided then error out
* block of code which will validate that docker nodees are mentioned properly , if not then error out.
* function which will test connectivity to enlisted host , and disply connected hosts.
* function for master to execute docker swarm master set up
  -> log in to master nodes and execute docker swarm master related command
  -> log docker swarm join token provided during swarm init and save it in memory
* function for agent to execute docker swarm agent set up
  -> log in to agent nodes and execute docker swarm agent related command
* function for validating swarm cluster
* if all functions returned success , then dispaly success message
  -> Docker swarm cluster is ready to use.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func SwarmMaster() {

}

func SwarmAgent() {

}

func SwarmValidate() {

}

func Connector() {

}

func main() {

}
