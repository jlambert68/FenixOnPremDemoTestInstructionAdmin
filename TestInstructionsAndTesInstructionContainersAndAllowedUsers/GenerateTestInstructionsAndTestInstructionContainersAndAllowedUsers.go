package TestInstructionsAndTesInstructionContainersAndAllowedUsers

import (
	"fmt"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	testInstructionContainer_SpecialSerialBaseContainer "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer"
	testInstructionContainer_SpecialSerialBaseContainer_1_0 "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer/version_1_0"
	generalSetupTearDown_TestCaseSetUp "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	generalSetupTearDown_TestCaseSetUp_1_0 "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/version_1_0"
	generalSetupTearDown_TestCaseSetUp_1_1 "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/version_1_1"
	generalSetupTearDown_TestCaseTearDown "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	generalSetupTearDown_TestCaseTearDown_1_0 "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown/version_1_0"
	isDateAPublicHoliday "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsDateAPublicHoliday"
	isDateAPublicHoliday_1_0 "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsDateAPublicHoliday/version_1_0"
	isServerAlive "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsServerAlive"
	isServerAlive_v_1_0 "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsServerAlive/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/shared_code"
	"os"
	"time"
)

var TestInstructionsAndTestInstructionContainersAndAllowedUsers_OnPremDemo *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct

func GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_OnPremDemo(allowedUsers []byte) {

	var err error

	// Load Allowed users from json-file
	err = shared_code.ParseAllowedUsersFromEmbeddedFile(allowedUsers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set AllUsersAuthorizationRights and initial Hash-value
	shared_code.AllowedUsersLoadFromJsonFile.AllUsersAuthorizationRights = &TestInstructionAndTestInstuctionContainerTypes.AllUsersAuthorizationRightsStruct{
		AllUsersCanListAndViewTestCaseHavingTIandTICFromThisDomain:  true,
		AllUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomain: true,
	}
	shared_code.AllowedUsersLoadFromJsonFile.AllowedUsersHash = shared_code.InitialValueBeforeHashed

	// Generate TestInstructions
	// GeneralSetupTearDown::TestCaseSetUp
	generalSetupTearDown_TestCaseSetUp_1_1.Initate_TestInstruction_OnPremDemo_TestCaseSetUp()
	generalSetupTearDown_TestCaseSetUp_1_0.Initate_TestInstruction_OnPremDemo_TestCaseSetUp()

	// GeneralSetupTearDown::TestCaseTearDown
	generalSetupTearDown_TestCaseTearDown_1_0.Initate_TestInstruction_OnPremDemo_TestCaseTearDown()

	// Standard::IsDateAPublicHoliday
	isDateAPublicHoliday_1_0.Initate_TestInstruction_OnPremDemo_IsDateAPublicHoliday()

	// Standard::IsServerAlive
	isServerAlive_v_1_0.Initate_TestInstruction_OnPremDemo_IsServerAlive()

	// Build structure for all TestInstructions & TestInstructionContainers to be sent over gRPC to Fenix Backend
	TestInstructionsAndTestInstructionContainersAndAllowedUsers_OnPremDemo = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct{

		// TestInstructions
		TestInstructions: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructionsMap: map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{

				//TestInstruction 'generalSetupTearDown_TestCaseSetUp'
				generalSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						//Version 'generalSetupTearDown_TestCaseSetUp_1.1'
						{
							TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_OnPremDemo_TestCaseSetUp,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_OnPremDemo_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_OnPremDemo_TestCaseSetUp.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_OnPremDemo_TestCaseSetUp.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseSetUp_1_1.TestInstruction_OnPremDemo_TestCaseSetUp.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},

						// Version 'generalSetupTearDown_TestCaseSetUp_1.0'
						{
							TestInstructionInstance:             generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_OnPremDemo_TestCaseSetUp,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_OnPremDemo_TestCaseSetUp.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_OnPremDemo_TestCaseSetUp.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_OnPremDemo_TestCaseSetUp.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseSetUp_1_0.TestInstruction_OnPremDemo_TestCaseSetUp.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},

				// TestInstruction 'generalSetupTearDown_TestCaseSetUp'
				generalSetupTearDown_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown: {
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'generalSetupTearDown_TestCaseSetUp_1.0'
						{
							TestInstructionInstance:             generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_OnPremDemo_TestCaseTearDown,
							TestInstructionInstanceMajorVersion: generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_OnPremDemo_TestCaseTearDown.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_OnPremDemo_TestCaseTearDown.TestInstruction.MinorVersionNumber,
							Deprecated:                          generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_OnPremDemo_TestCaseTearDown.TestInstruction.Deprecated,
							Enabled:                             generalSetupTearDown_TestCaseTearDown_1_0.TestInstruction_OnPremDemo_TestCaseTearDown.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},

				// TestInstruction 'IsDateAPublicHoliday'
				isDateAPublicHoliday.TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday: {
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'isDateAPublicHoliday_1_0'
						{
							TestInstructionInstance:             isDateAPublicHoliday_1_0.TestInstruction_OnPremDemo_IsDateAPublicHoliday,
							TestInstructionInstanceMajorVersion: isDateAPublicHoliday_1_0.TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: isDateAPublicHoliday_1_0.TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstruction.MinorVersionNumber,
							Deprecated:                          isDateAPublicHoliday_1_0.TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstruction.Deprecated,
							Enabled:                             isDateAPublicHoliday_1_0.TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},

				// TestInstruction 'IsServerAlive'
				isServerAlive.TestInstructionUUID_OnPremDemo_IsServerAlive: {
					TestInstructionVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{

						// Version 'IsServerAlive_1_0'
						{
							TestInstructionInstance:             isServerAlive_v_1_0.TestInstruction_OnPremDemo_IsServerAlive,
							TestInstructionInstanceMajorVersion: isServerAlive_v_1_0.TestInstruction_OnPremDemo_IsServerAlive.TestInstruction.MajorVersionNumber,
							TestInstructionInstanceMinorVersion: isServerAlive_v_1_0.TestInstruction_OnPremDemo_IsServerAlive.TestInstruction.MinorVersionNumber,
							Deprecated:                          isServerAlive_v_1_0.TestInstruction_OnPremDemo_IsServerAlive.TestInstruction.Deprecated,
							Enabled:                             isServerAlive_v_1_0.TestInstruction_OnPremDemo_IsServerAlive.TestInstruction.Enabled,
							TestInstructionInstanceVersionHash:  shared_code.InitialValueBeforeHashed,
						},
					},
					TestInstructionVersionsHash: shared_code.InitialValueBeforeHashed,
				},
			},
			TestInstructionsHash: shared_code.InitialValueBeforeHashed,
		},

		// TestInstructionContainers
		TestInstructionContainers: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{},
		AllowedUsers:              shared_code.AllowedUsersLoadFromJsonFile,
		TestInstructionsAndTestInstructionsContainersAndUsersMessageHash: shared_code.InitialValueBeforeHashed,
		MessageCreationTimeStamp: time.Now(),
		ForceNewBaseLineForTestInstructionsAndTestInstructionContainers: false,
		ConnectorsDomain: TestInstructionAndTestInstuctionContainerTypes.ConnectorsDomainStruct{
			ConnectorsDomainUUID: DomainData.DomainUUID_OnPremDemo,
			ConnectorsDomainName: DomainData.DomainName_OnPremDemo,
			ConnectorsDomainHash: shared_code.InitialValueBeforeHashed,
		},
	}

	// Generate TestInstructionContainers
	// testInstructionContainer_SpecialSerialBaseContainer
	testInstructionContainer_SpecialSerialBaseContainer_1_0.Initiate_TestInstructionContainer_OnPremDemo_Serial(TestInstructionsAndTestInstructionContainersAndAllowedUsers_OnPremDemo)

	// TestInstructionContainers
	var testInstructionContainers *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct
	testInstructionContainers = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{
		TestInstructionContainersMap: map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{

			// 'SC_SpecialSerialBaseContainer'
			testInstructionContainer_SpecialSerialBaseContainer.TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer: {
				TestInstructionContainerVersions: []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct{

					//Version 'TestInstructionContainer_SpecialSerialBaseContainer_1.0'
					{
						TestInstructionContainerInstance:             testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_OnPremDemo_SpecialSerialBase,
						TestInstructionContainerInstanceMajorVersion: testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_OnPremDemo_SpecialSerialBase.TestInstructionContainer.MajorVersionNumber,
						TestInstructionContainerInstanceMinorVersion: testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_OnPremDemo_SpecialSerialBase.TestInstructionContainer.MinorVersionNumber,
						Deprecated:                           testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_OnPremDemo_SpecialSerialBase.TestInstructionContainer.Deprecated,
						Enabled:                              testInstructionContainer_SpecialSerialBaseContainer_1_0.TestInstructionContainer_OnPremDemo_SpecialSerialBase.TestInstructionContainer.Enabled,
						TestInstructionContainerInstanceHash: shared_code.InitialValueBeforeHashed,
					},
				},
				TestInstructionContainerVersionsHash: shared_code.InitialValueBeforeHashed,
			},
		},

		TestInstructionContainersHash: shared_code.InitialValueBeforeHashed,
	}
	TestInstructionsAndTestInstructionContainersAndAllowedUsers_OnPremDemo.TestInstructionContainers = testInstructionContainers

	// Define temporary storage for 'LocalExecutionMethods.MethodsForLocalExecutionsStruct'
	/*
		var tempStorageLocalExecutionMethodsObject *TestInstructionAndTestInstuctionContainerTypes.AnyType


		var PushToTempStoreFunction = func(incomingLocalExecutionMethodsObject *TestInstructionAndTestInstuctionContainerTypes.AnyType) {
			tempStorageLocalExecutionMethodsObject = incomingLocalExecutionMethodsObject

		}

		var PullFromTempStoreFunction = func() (outgoingLocalExecutionMethodsObject *TestInstructionAndTestInstuctionContainerTypes.AnyType) {
			return tempStorageLocalExecutionMethodsObject
		}

	*/
	//type PushToTempStoreFunctionType[ T, any] func(T)
	//var PushToTempStore = shared_code.PushToTempStoreFunctionType[*TestInstructionAndTestInstuctionContainerTypes.AnyType](PushToTempStoreFunction)
	//var PullFromTempStore = shared_code.PullFromTempStoreFunctionType[*TestInstructionAndTestInstuctionContainerTypes.AnyType](PullFromTempStoreFunction)

	// Calculate hashes that is included in the Supported TestInstructions, TestInstructionContainers and Allowed Users message
	err = shared_code.CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes(
		TestInstructionsAndTestInstructionContainersAndAllowedUsers_OnPremDemo)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
