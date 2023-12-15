package main

import (
	"fmt"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers"
)

func main() {

	TestInstructionsAndTesInstructionContainersAndAllowedUsers.GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_OnPremDemo()
	TestInstructionsAndTesInstructionContainersAndAllowedUsers.GenerateAndVerifyRPCMessages()

	fmt.Println(TestInstructionsAndTesInstructionContainersAndAllowedUsers.TestInstructionsAndTestInstructionContainersAndAllowedUsers_OnPremDemo)
	fmt.Println("Success when generating and verifying all messages for TestInstructions, TestInstructionContainers and AllowedUsers")
}
