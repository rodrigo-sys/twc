package browser

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type Browser struct {
	Cmd exec.Cmd
}

