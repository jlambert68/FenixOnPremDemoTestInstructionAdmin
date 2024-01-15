package version_1_0

import (
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/LocalExecutionMethods/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsServerAlive"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'IsServerAlive'
	TestInstructionUUID_OnPremDemo_IsServerAlive               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_OnPremDemo_IsServerAlive
	TestInstructionName_OnPremDemo_IsServerAlive               TypeAndStructs.TestInstructionNameType = "IsServerAlive"
	TestInstructionTypeUUID_OnPremDemo_IsServerAlive                                                  = TestInstructions.TestInstructionTypeUUID_OnPremDemo_Standard
	TestInstructionTypeName_OnPremDemo_IsServerAlive                                                  = TestInstructions.TestInstructionTypeName_OnPremDemo_Standard
	TestInstructionDescription_OnPremDemo_IsServerAlive        string                                 = "This TestInstruction Checks if server 'X' is alive and responding "
	TestInstructionMouseOverText_OnPremDemo_IsServerAlive      string                                 = "This TestInstruction Checks if server 'X' is alive and responding"
	TestInstructionDeprecated_OnPremDemo_IsServerAlive         bool                                   = false
	TestInstructionEnabled_OnPremDemo_IsServerAlive            bool                                   = true
	TestInstructionMajorVersionNumber_OnPremDemo_IsServerAlive int                                    = 1
	TestInstructionMinorVersionNumber_OnPremDemo_IsServerAlive int                                    = 0
	TestInstructionColor_OnPremDemo_IsServerAlive              TypeAndStructs.ColorType               = "#fff000AA"
	TCRuleDeletion_OnPremDemo_IsServerAlive                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_OnPremDemo_IsServerAlive                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                           TypeAndStructs.UpdatedTimeStampType    = "2024-01-15 20:00:00"

	// *** DropZone *** 'IsServerAlive_ExpectsToSucceed'
	TestInstructionDropZoneUUID_OnPremDemo_IsServerAlive_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "280b8ccd-3ed8-4c88-985a-3190da72016d"
	TestInstructionDropZoneName_OnPremDemo_IsServerAlive_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "IsServerAlive_ExpectsToSucceed"
	TestInstructionDropZoneDescription_OnPremDemo_IsServerAlive_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_OnPremDemo_IsServerAlive_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_OnPremDemo_IsServerAlive_ExpectsToSucceed       TypeAndStructs.ColorType        = "#ffff00AA"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_OnPremDemo_IsServerAlive_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = TestInstructions.TestInstructionAttributeUUID_OnPremDemo_ExpectedToBePassed // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeName_OnPremDemo_IsServerAlive_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeType_OnPremDemo_IsServerAlive_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeActionCommand_OnPremDemo_IsServerAlive_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_OnPremDemo_IsServerAlive_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_OnPremDemo_IsServerAlive_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_OnPremDemo_IsServerAlive_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_OnPremDemo_IsServerAlive_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	// Attribute - 'ServerServerIpAddress'
	TestInstructionAttributeUUID_OnPremDemo_IsServerAlive_ServerIpAddress          TypeAndStructs.TestInstructionAttributeUUIDType = "7cd1932e-1842-48ee-b4f7-c5782d6e7453" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_OnPremDemo_ServerIpAddress
	TestInstructionAttributeName_OnPremDemo_IsServerAlive_ServerIpAddress          TypeAndStructs.TestInstructionAttributeNameType = "The IP-address of the server"
	TestInstructionAttributeType_OnPremDemo_IsServerAlive_ServerIpAddress          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_OnPremDemo_IsServerAlive_ServerIpAddress   string                                          = "The IP-address to the server to be checked"
	TestInstructionAttributeMouseOverText_OnPremDemo_IsServerAlive_ServerIpAddress string                                          = "The IP-address to the server to be checked"
)

// TestInstruction_OnPremDemo_IsServerAlive
// Variable that holds the data for the TestInstruction
var TestInstruction_OnPremDemo_IsServerAlive *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_OnPremDemo_IsServerAlive
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_OnPremDemo_IsServerAlive() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_OnPremDemo_IsServerAlive = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				LocalParametersUsedInRunTime: &LocalExecutionMethods.LocalParametersUsedInRunTimeStruct{
					ExpectedTestInstructionExecutionDurationInSeconds: 360,
				},
				FangEngineClassesMethodsAttributes: &FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{},
			},
		},
	}

	// TestInstruction - IsServerAlive
	TestInstruction_OnPremDemo_IsServerAlive.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_IsServerAlive,
		TestInstructionName:          TestInstructionName_OnPremDemo_IsServerAlive,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_OnPremDemo_IsServerAlive,
		TestInstructionTypeName:      TestInstructionTypeName_OnPremDemo_IsServerAlive,
		TestInstructionDescription:   TestInstructionDescription_OnPremDemo_IsServerAlive,
		TestInstructionMouseOverText: TestInstructionMouseOverText_OnPremDemo_IsServerAlive,
		Deprecated:                   TestInstructionDeprecated_OnPremDemo_IsServerAlive,
		Enabled:                      TestInstructionEnabled_OnPremDemo_IsServerAlive,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_OnPremDemo_IsServerAlive,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_OnPremDemo_IsServerAlive,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - IsServerAlive
	TestInstruction_OnPremDemo_IsServerAlive.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_IsServerAlive,
		TestInstructionName:          TestInstructionName_OnPremDemo_IsServerAlive,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_OnPremDemo_IsServerAlive,
		TestInstructionTypeName:      TestInstructionTypeName_OnPremDemo_IsServerAlive,
		Deprecated:                   TestInstructionDeprecated_OnPremDemo_IsServerAlive,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_OnPremDemo_IsServerAlive,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_OnPremDemo_IsServerAlive,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_OnPremDemo_IsServerAlive,
		TCRuleDeletion:               TCRuleDeletion_OnPremDemo_IsServerAlive,
		TCRuleSwap:                   TCRuleSwap_OnPremDemo_IsServerAlive,
		TestInstructionDescription:   TestInstructionDescription_OnPremDemo_IsServerAlive,
		TestInstructionMouseOverText: TestInstructionMouseOverText_OnPremDemo_IsServerAlive,
		Enabled:                      TestInstructionEnabled_OnPremDemo_IsServerAlive,
	}

	// DropZone 'IsServerAlive_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: IsServerAlive_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_OnPremDemo_IsServerAlive_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_OnPremDemo_IsServerAlive_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_IsServerAlive,
		TestInstructionName:          TestInstructionName_OnPremDemo_IsServerAlive,
		DropZoneUUID:                 TestInstructionDropZoneUUID_OnPremDemo_IsServerAlive_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_OnPremDemo_IsServerAlive_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_OnPremDemo_IsServerAlive_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_OnPremDemo_IsServerAlive_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_OnPremDemo_IsServerAlive_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_OnPremDemo_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_OnPremDemo_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_OnPremDemo_IsServerAlive_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_OnPremDemo_IsServerAlive_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_OnPremDemo_IsServerAlive_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_OnPremDemo_IsServerAlive,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_OnPremDemo_IsServerAlive_ExpectedToBePassed,
	}
	TestInstruction_OnPremDemo_IsServerAlive.ImmatureTestInstructionInformation = append(
		TestInstruction_OnPremDemo_IsServerAlive.ImmatureTestInstructionInformation,
		TestInstruction_OnPremDemo_IsServerAlive_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_OnPremDemo_IsServerAlive_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_OnPremDemo_IsServerAlive_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_OnPremDemo,
		DomainName:                                    DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:                           TestInstructionUUID_OnPremDemo_IsServerAlive,
		TestInstructionName:                           TestInstructionName_OnPremDemo_IsServerAlive,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_OnPremDemo_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_OnPremDemo_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_OnPremDemo_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_OnPremDemo_IsServerAlive_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:              TestInstructions.TestInstructionAttributeTypeUUID_OnPremDemo_ExpectedToPass,
		TestInstructionAttributeTypeName:              TestInstructions.TestInstructionAttributeTypeName_OnPremDemo_ExpectedToPass,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_OnPremDemo_IsServerAlive_ExpectedToBePassed,
	}
	TestInstruction_OnPremDemo_IsServerAlive.TestInstructionAttribute = append(
		TestInstruction_OnPremDemo_IsServerAlive.TestInstructionAttribute,
		TestInstructionAttribute_OnPremDemo_IsServerAlive_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// TestInstruction Attribute - 'ServerIpAddress'
	var TestInstructionAttribute_OnPremDemo_IsServerAlive_ServerIpAddress *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_OnPremDemo_IsServerAlive_ServerIpAddress = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_OnPremDemo,
		DomainName:                                    DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:                           TestInstructionUUID_OnPremDemo_IsServerAlive,
		TestInstructionName:                           TestInstructionName_OnPremDemo_IsServerAlive,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_OnPremDemo_IsServerAlive_ServerIpAddress,
		TestInstructionAttributeName:                  TestInstructionAttributeName_OnPremDemo_IsServerAlive_ServerIpAddress,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_OnPremDemo_IsServerAlive_ServerIpAddress,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_OnPremDemo_IsServerAlive_ServerIpAddress,
		TestInstructionAttributeTypeUUID:              TestInstructions.TestInstructionAttributeTypeUUID_OnPremDemo_Standard,
		TestInstructionAttributeTypeName:              TestInstructions.TestInstructionAttributeTypeName_OnPremDemo_Standard,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_OnPremDemo_IsServerAlive_ServerIpAddress,
	}
	TestInstruction_OnPremDemo_IsServerAlive.TestInstructionAttribute = append(
		TestInstruction_OnPremDemo_IsServerAlive.TestInstructionAttribute,
		TestInstructionAttribute_OnPremDemo_IsServerAlive_ServerIpAddress)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// ImmatureElementModel - IsServerAlive
	var TestInstructionImmatureElementModel_OnPremDemo_IsServerAlive *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_OnPremDemo_IsServerAlive = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_OnPremDemo,
		DomainName:               DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID:      TestInstructionUUID_OnPremDemo_IsServerAlive,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_OnPremDemo_IsServerAlive),
		PreviousElementUUID:      TestInstructionUUID_OnPremDemo_IsServerAlive,
		NextElementUUID:          TestInstructionUUID_OnPremDemo_IsServerAlive,
		FirstChildElementUUID:    TestInstructionUUID_OnPremDemo_IsServerAlive,
		ParentElementUUID:        TestInstructionUUID_OnPremDemo_IsServerAlive,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_OnPremDemo_IsServerAlive,
		TopImmatureElementUUID:   TestInstructionUUID_OnPremDemo_IsServerAlive,
		IsTopElement:             true,
	}
	TestInstruction_OnPremDemo_IsServerAlive.ImmatureElementModel = append(
		TestInstruction_OnPremDemo_IsServerAlive.ImmatureElementModel,
		TestInstructionImmatureElementModel_OnPremDemo_IsServerAlive)

	return TestInstruction_OnPremDemo_IsServerAlive
}
