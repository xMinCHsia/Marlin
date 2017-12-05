
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

