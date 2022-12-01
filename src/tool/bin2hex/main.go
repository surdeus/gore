package bin2hex

import (
	"os"
	"io"
	"encoding/hex"
	"fmt"
	"bufio"
	"flag"
	//"strings"
)

var (
	blockSize = 512
	verbose bool
)

func notVerboseHandle() {
	rd := bufio.NewReader(os.Stdin)
	buf := make([]byte, blockSize)
	for {
		n, err := rd.Read(buf)
		if err == io.EOF {
			break
		}
		out := hex.EncodeToString(buf[:n])
		fmt.Print(out)
	}
}

func verboseHandle() {
}

func Run(args []string) {
	arg0 := args[0]
	args = args[1:]

	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.BoolVar(&verbose, "v", false, "verbose mode")
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options]\n", arg0)
		flagSet.PrintDefaults()
	}
	flagSet.Parse(args)

	if !verbose {
		notVerboseHandle()
	} else {
		verboseHandle()
	}

}
