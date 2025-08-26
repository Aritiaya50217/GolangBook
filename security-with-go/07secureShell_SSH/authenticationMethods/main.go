package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	host := "192.168.1.100:22" // IP หรือ Host ของ SSH server
	user := "user"
	password := "password"

	// Config ใช้ password authentication อย่างเดียว
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatalf("SSH connect fail: %s", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("New session failed: %s", err)
	}
	defer session.Close()

	// Run command
	output, err := session.CombinedOutput("hostname && whoami")
	if err != nil {
		log.Fatalf("Command failed: %s", err)
	}

	fmt.Println(string(output))
}
