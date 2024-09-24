package processor

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func resetTorCircuit() error {
	conn, err := net.Dial("tcp", "127.0.0.1:9051")
	if err != nil {
		return fmt.Errorf("failed to connect to Tor: %v. Make sure Tor is running", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "AUTHENTICATE \"\"\n")
	status, err := readResponse(conn)

	if err != nil || !strings.Contains(status, "250 OK") {
		return fmt.Errorf("tor authentication failed: %v", err)
	}

	fmt.Fprintf(conn, "signal NEWNYM\n")
	status, err = readResponse(conn)
	if err != nil || !strings.Contains(status, "250 OK") {
		return fmt.Errorf("tor NEWNYM signal failed: %v", err)
	}

	return nil
}

func readResponse(conn net.Conn) (string, error) {
	buffer := make([]byte, 512)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}
