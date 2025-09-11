package ping

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

//Global Variables
// Why global? Because Cobra's flag binding system (StringVarP, DurationVarP)
// needs persistent variables to store flag values
var (
	url string
	timeout time.Duration
)


var pingCmd = &cobra.Command{
	Use: "ping",
	Short: "Ping an HTTP endpoint",
	Long: `Send HTTP GET request to check endpoint availability and response time`,
	Run: runPing,
	// Run takes any function that accepts (cmd *cobra.Command, args []string)
}

func init(){
	pingCmd.Flags().StringVarP(&url, "url", "u", "", "URL to ping (required)")
	pingCmd.Flags().DurationVarP(&timeout, "timeout", "t", 5*time.Second, "request timeout")
	pingCmd.MarkFlagRequired("url")
}

func runPing(cmd *cobra.Command, args []string){
	verbose, _ := cmd.Flags().GetBool("verbose")

	if verbose{
		fmt.Printf("Pinging %s with timeout %v ... \n",url, timeout)
	}

	//Create HTTP client with timeout
	client := &http.Client{
		Timeout: timeout,
	}

	//Record start time
	start := time.Now()

	//Make the request
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("❌ Oops!, Failed to reach %s, %v\n", url, err)
		os.Exit(2) // exit with code 2 = misuse/network error
	}
	defer resp.Body.Close()

	//Check status and print result
	if resp.StatusCode == 200 {
		fmt.Printf("✅ %s responded with %d in %v\n", url, resp.StatusCode, duration)
		os.Exit(0)
	} else {
		fmt.Printf("⚠️ %s responded with %d in %v\n", url, resp.StatusCode,duration)
		os.Exit(1)
	}
}
