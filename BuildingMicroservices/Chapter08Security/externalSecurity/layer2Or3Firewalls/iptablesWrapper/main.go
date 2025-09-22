package main

import (
	"fmt"
	"log"
	"os/exec"
)

// runCmd runs a command and returns its output or an error
func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cmd %s %v failed: %v\n%s", name, args, err, string(out))
	}
	if len(out) > 0 {
		log.Printf("output : %s\n", string(out))
	}
	return nil
}

func allowHTTPSFromSubnet(subnet string) error {
	// ตัวอย่าง: อนุญาต TCP 443 จาก subnet
	return runCmd("iptables", "-A", "INPUT", "-p", "tcp", "-s", subnet, "--dport", "443", "-j", "ACCEPT")
}

func dropAllFromIP(ip string) error {
	return runCmd("iptables", "-A", "INPUT", "-s", ip, "-j", "DROP")
}

func listRules() error {
	return runCmd("iptables", "-L", "-n", "-v")
}

func main() {
	// ต้องรันด้วย sudo / root
	if err := allowHTTPSFromSubnet("192.168.1.0/24"); err != nil {
		log.Fatalf("allow error: %v", err)
	}
	if err := dropAllFromIP("1.2.3.4"); err != nil {
		log.Fatalf("drop error: %v", err)
	}

	if err := listRules(); err != nil {
		log.Fatalf("list error : %v", err)
	}
}
