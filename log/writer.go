package log

import (
	"io"
	"os"
)

// Writer log writer
var Writer io.Writer = os.Stdout
