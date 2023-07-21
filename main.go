package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/caseymrm/menuet"
)

type Config struct {
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	Address     string       `json:"address"`
	Port        int          `json:"port"`
	Forwardings []Forwarding `json:"forwardings"`
}

type Forwarding struct {
	BindPort int  `json:"bindPort"`
	Enabled  bool `json:"enabled"`
}

type TunnelStatus struct {
	Active bool
	Host   string
}

var (
	tunnelStatus      = make(map[string]*TunnelStatus)
	tunnelStatusMutex = sync.RWMutex{}
)

func main() {

	go initMenu()
	go createTunnels()
	menuet.App().RunApplication()

}

func createTunnels() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Unmarshal JSON data
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse JSON data: %v", err)
	}

	// Loop over all profiles and establish SSH tunnels
	for _, profile := range config.Profiles {
		for _, forwarding := range profile.Forwardings {
			if forwarding.Enabled {
				// Start a goroutine to establish the SSH tunnel
				go establishSSHTunnel(profile.Address, profile.Port, forwarding.BindPort)
			}
		}
	}

	// Prevent the program from exiting immediately
	for {
		time.Sleep(time.Minute)
	}
}

func establishSSHTunnel(address string, port, bindPort int) {
	key := fmt.Sprintf("%s:%d", address, bindPort)
	tunnelStatusMutex.Lock()
	status, ok := tunnelStatus[key]
	if !ok {
		tunnelStatus[key] = &TunnelStatus{
			Active: false,
		}
		status = tunnelStatus[key]
	}
	tunnelStatusMutex.Unlock()
	for {
		sshCmd := fmt.Sprintf("ssh -o BatchMode=yes -o ConnectTimeout=5 -N -D localhost:%d %s", bindPort, address)
		fmt.Println("Running command:", sshCmd)

		// Establish SSH tunnel
		cmd := exec.Command("ssh", "-N", "-D", fmt.Sprintf("localhost:%d", bindPort), fmt.Sprintf("%s", address))
		err := cmd.Start()
		if err != nil {
			log.Printf("Failed to establish SSH tunnel: %v", err)
			status.Active = false
		} else {
			log.Printf("SSH tunnel established for %s", address)
			status.Active = true
		}
		// Wait for SSH tunnel to finish setup (replace with your own condition)
		err = cmd.Wait()
		if err != nil {
			log.Printf("SSH tunnel closed unexpectedly: %v", err)
		}
		status.Active = false

		// If the command finishes (SSH tunnel is closed), wait for a minute before the next check
		time.Sleep(time.Second * 5)
	}
}

// Helper function to check if a string is in a slice
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
