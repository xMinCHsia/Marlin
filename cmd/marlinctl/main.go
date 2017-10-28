
// Command marlinctl is the operator CLI for Marlin daemons.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", "localhost:8590", "daemon address")
	cmd := flag.String("cmd", "status", "status | plans | drafts | tune")
	flag.Parse()

	var body any
	switch *cmd {
	case "status", "plans", "drafts":
		body = nil
	case "tune":
		body = map[string]any{"force": true}
	default:
		fail("unknown command: " + *cmd)
	}

	url := "http://" + *addr + "/v1/" + *cmd
	client := &http.Client{Timeout: 30 * time.Second}
	var resp *http.Response
	var err error
	if body == nil {
		resp, err = client.Get(url)
	} else {
		raw, _ := json.Marshal(body)
		resp, err = client.Post(url, "application/json", bytes.NewReader(raw))
	}
	if err != nil {
		fail("request: " + err.Error())
	}
	defer resp.Body.Close()
	out, _ := io.ReadAll(resp.Body)
	fmt.Printf("%d %s\n", resp.StatusCode, pretty(out))
}

