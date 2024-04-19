package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/swxctx/pot/library/gopot"
)

var (
	potClient *gopot.PotClient
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of Pot CLI:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nCommands:")
		fmt.Fprintln(os.Stderr, "  set <key> <value> [expiration] - Set the value of a key")
		fmt.Fprintln(os.Stderr, "  get <key> - Get the value of a key")
		fmt.Fprintln(os.Stderr, "  del <key> - Delete a key")
		fmt.Fprintln(os.Stderr, "  exists <key> - Check if a key exists")
		fmt.Fprintln(os.Stderr, "  ttl <key> - Get the TTL of a key")
		fmt.Fprintln(os.Stderr, "  expire <key> <seconds> - Set the expiration of a key")
	}

	// command
	versionFlag := flag.Bool("version", false, "Print version information and exit")
	host := flag.String("h", "127.0.0.1", "host address of the server")
	port := flag.String("p", "9577", "port of the server")

	flag.Parse()

	if *versionFlag {
		fmt.Println("Pot CLI Version 1.0.0")
		return
	}

	address := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Println("Welcome to Pot CLI")
	fmt.Printf("Connecting to %s\n", address)
	fmt.Println("Type 'exit' to quit.")

	var (
		err error
	)
	// init pot client
	potClient, err = gopot.NewClient(&gopot.Config{
		Address: address,
	})
	if err != nil {
		fmt.Printf("connect pot server err-> %v\n", err)
		return
	}
	fmt.Println("Connecting pot server success.")

	// read input
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s> ", address)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading command:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Exiting Pot CLI...")
			break
		}

		// input cmd logic
		handleCommand(input)
	}
}

// handleCommand
func handleCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}
	command := parts[0]
	args := parts[1:]

	switch command {
	case "set":
		if len(args) < 2 {
			fmt.Println("Usage: set <key> <value>")
			return
		}

		expiration := time.Duration(0) * time.Second
		if len(args) == 3 {
			inputSeconds, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				fmt.Println("Usage: expiration is second(number).")
				return
			}
			if inputSeconds < 0 {
				fmt.Println("Usage: input expiration > 0.")
				return
			}
			expiration = time.Duration(inputSeconds) * time.Second
		}

		// set
		setCmd := potClient.Set(args[0], args[1], expiration)
		if setCmd.Err() != nil {
			fmt.Println("set err-> ", setCmd.Err())
			return
		}
		fmt.Println(setCmd.Success())
	case "get":
		if len(args) != 1 {
			fmt.Println("Usage: get <key>")
			return
		}
		getCmd := potClient.Get(args[0])
		if getCmd.Err() != nil {
			fmt.Println("get err-> ", getCmd.Err())
			return
		}
		fmt.Println(getCmd.String())
	case "del":
		if len(args) != 1 {
			fmt.Println("Usage: del <key>")
			return
		}
		delCmd := potClient.Del(args[0])
		if delCmd.Err() != nil {
			fmt.Println("del err-> ", delCmd.Err())
			return
		}
		fmt.Println(delCmd.Success())
	case "exists":
		if len(args) != 1 {
			fmt.Println("Usage: exists <key>")
			return
		}
		existsCmd := potClient.Exists(args[0])
		if existsCmd.Err() != nil {
			fmt.Println("exists err-> ", existsCmd.Err())
			return
		}
		fmt.Println(existsCmd.Result())
	case "ttl":
		if len(args) != 1 {
			fmt.Println("Usage: ttl <key>")
			return
		}
		ttlCmd := potClient.TTL(args[0])
		if ttlCmd.Err() != nil {
			fmt.Println("ttl err-> ", ttlCmd.Err())
			return
		}
		fmt.Println(ttlCmd.Result())
	case "expire":
		if len(args) != 2 {
			fmt.Println("Usage: expire <key> <expiration>(seconds)")
			return
		}
		inputSeconds, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println("Usage: expiration is second(number).")
			return
		}
		if inputSeconds < 0 {
			fmt.Println("Usage: input expiration > 0.")
			return
		}
		expiration := time.Duration(inputSeconds) * time.Second

		expireCmd := potClient.Expire(args[0], expiration)
		if expireCmd.Err() != nil {
			fmt.Println("expire err-> ", expireCmd.Err())
			return
		}
		fmt.Println(expireCmd.Success())
	default:
		fmt.Println("Unsupported command:", command)
	}
}
