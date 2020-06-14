package main

import(
	"flag"
    "fmt"
	"os"
    "strconv"
    "net/http"
)

type Args struct {
	Port    int
	Dir     string
}

var args Args

func main() {
	// first we gonne parse the args
	parseArgs()
    fmt.Println("Starting gss on port '",args.Port,"' serving  '",args.Dir)
    
    // than we gonne start the webserver
    var portStr string
    portStr += ":"
    portStr += strconv.Itoa(args.Port)
    http.ListenAndServe(portStr, http.FileServer(http.Dir(args.Dir)))
}

func parseArgs() {
	// first we check for the help flag
	if 1 < len(os.Args) {
		if ok := os.Args[1]; ok == "help" {
			printHelpText()
			os.Exit(1)
		}
	}

	// delimiter to be used for custom generator input
	var dir string
	flag.StringVar(&dir, "dir", "./", "Directory to serve (recursive)")

	// thread amount to be used
	var port int
	flag.IntVar(&port, "port", 8091, "Port to serve the static files on");

	// parse the flags
	flag.Parse()

	args = Args{
		Port: port,
		Dir : dir,
	}

}

func printHelpText() {
	helpText := "GSS Help:\n" +
		"> GSS serve's static files from a given directory(recursivly) via http\n" + 
        "  on a defined port .\n" +
		"  For examples check https://github.com/voodooEntity/gss readme.\n\n" +
		"  Args: \n" +
		"    -port INT           | Specifys the port number used to server the files, default 8091\n" +
		"    -dir STRING         | Path to directory to server the files from\n";
	fmt.Println(helpText)
}
