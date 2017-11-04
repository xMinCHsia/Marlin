
// Command marlind runs the speculative decoding orchestrator.
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xMinCHsia/marlin/internal/admission"
	"github.com/xMinCHsia/marlin/internal/config"
	"github.com/xMinCHsia/marlin/internal/kvcache"
	"github.com/xMinCHsia/marlin/internal/tuning"
)

