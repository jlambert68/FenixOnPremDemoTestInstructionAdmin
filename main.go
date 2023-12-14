package main

import (
	"FenixOnPremDemoTestInstructionAdmin/SubCustody"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
	"log"
	"os"
)

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func init() {

	// *********************** Extract environment variables for 'This Application' ***********************

	// Get Environment variable for relative path us json file with allowed users for this connector
	shared_code.RelativePathToAllowedUsersList = mustGetenv("RelativePathToAllowedUsersList")

}

func main() {

	SubCustody.GenerateTestInstructions_SC()

}
