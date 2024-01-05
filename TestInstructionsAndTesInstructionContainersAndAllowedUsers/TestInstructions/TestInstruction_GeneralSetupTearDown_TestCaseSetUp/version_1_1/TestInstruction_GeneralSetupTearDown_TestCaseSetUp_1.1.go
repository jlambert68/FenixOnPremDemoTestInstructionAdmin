package version_1_1

import (
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/LocalExecutionMethods/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'TestCaseSetUp'
	TestInstructionUUID_OnPremDemo_TestCaseSetUp               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_OnPremDemo_TestCaseSetUp
	TestInstructionName_OnPremDemo_TestCaseSetUp               TypeAndStructs.TestInstructionNameType = "TestCaseSetUp"
	TestInstructionTypeUUID_OnPremDemo_TestCaseSetUp                                                  = TestInstructions.TestInstructionTypeUUID_OnPremDemo_GeneralSetUpTearDown
	TestInstructionTypeName_OnPremDemo_TestCaseSetUp                                                  = TestInstructions.TestInstructionTypeName_OnPremDemo_GeneralSetUpTearDown
	TestInstructionDescription_OnPremDemo_TestCaseSetUp        string                                 = "Initiate _OnPremDemos execution engine to be able to execute TestInstructionsMap"
	TestInstructionMouseOverText_OnPremDemo_TestCaseSetUp      string                                 = "Initiate _OnPremDemos execution engine to be able to execute TestInstructionsMap"
	TestInstructionDeprecated_OnPremDemo_TestCaseSetUp         bool                                   = false
	TestInstructionEnabled_OnPremDemo_TestCaseSetUp            bool                                   = false // Is not shown up in tree of available building blocks, but can be used in pre-create TestInstructionContainers
	TestInstructionMajorVersionNumber_OnPremDemo_TestCaseSetUp int                                    = 1
	TestInstructionMinorVersionNumber_OnPremDemo_TestCaseSetUp int                                    = 1
	TestInstructionColor_OnPremDemo_TestCaseSetUp              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_OnPremDemo_TestCaseSetUp                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_OnPremDemo_TestCaseSetUp                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                           TypeAndStructs.UpdatedTimeStampType    = "2023-11-28 14:00:00"

	// *** DropZone *** 'TestCaseSetUp_ExpectsToSucceed'
	TestInstructionDropZoneUUID_OnPremDemo_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "e17c1058-67d9-47af-8a38-f28979a50eec"
	TestInstructionDropZoneName_OnPremDemo_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseSetUp_ExpectsToSucceed"
	TestInstructionDropZoneDescription_OnPremDemo_TestCaseSetUp_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_OnPremDemo_TestCaseSetUp_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_OnPremDemo_TestCaseSetUp_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_OnPremDemo_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "818218d0-3fd4-4dd1-8033-7e7c9ff3143f" //TestInstructionAttributeUUID_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeName_OnPremDemo_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeType_OnPremDemo_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeActionCommand_OnPremDemo_TestCaseSetUp_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_OnPremDemo_TestCaseSetUp_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_OnPremDemo_TestCaseSetUp_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_OnPremDemo_TestCaseSetUp_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_OnPremDemo_TestCaseSetUp_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

// TestInstruction_OnPremDemo_TestCaseSetUp
// Variable that holds the data for the TestInstruction
var TestInstruction_OnPremDemo_TestCaseSetUp *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_OnPremDemo_TestCaseSetUp
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_OnPremDemo_TestCaseSetUp() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Initiate Map for FangMethodAttributes
	var fangMethodAttributeMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineClassesAndMethods.FangEngineAttributesStruct
	fangMethodAttributeMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineClassesAndMethods.FangEngineAttributesStruct)

	// Initiate variable to store all TestInstruction data
	TestInstruction_OnPremDemo_TestCaseSetUp = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
		TestInstruction:                    &TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    &TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,

		// Local Execution Methods are specified here
		LocalExecutionMethods: TestInstructionAndTestInstuctionContainerTypes.AnyType{
			&LocalExecutionMethods.MethodsForLocalExecutionsStruct{
				FangEngineClassesMethodsAttributes: &FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
					TestInstructionOriginalUUID: TestInstructionUUID_OnPremDemo_TestCaseSetUp,
					TestInstructionName:         TestInstructionName_OnPremDemo_TestCaseSetUp,
					FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_OnPremDemo_GeneralSetupTearDown,
					FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_OnPremDemo_GeneralSetupTearDown,
					FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_OnPremDemo_GeneralSetupTearDown_Setup,
					FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_OnPremDemo_GeneralSetupTearDown_Setup,
					Attributes:                  fangMethodAttributeMap,
				},
			},
		},
	}

	// TestInstruction - TestCaseSetUp
	TestInstruction_OnPremDemo_TestCaseSetUp.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_OnPremDemo_TestCaseSetUp,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_OnPremDemo_TestCaseSetUp,
		TestInstructionTypeName:      TestInstructionTypeName_OnPremDemo_TestCaseSetUp,
		TestInstructionDescription:   TestInstructionDescription_OnPremDemo_TestCaseSetUp,
		TestInstructionMouseOverText: TestInstructionMouseOverText_OnPremDemo_TestCaseSetUp,
		Deprecated:                   TestInstructionDeprecated_OnPremDemo_TestCaseSetUp,
		Enabled:                      TestInstructionEnabled_OnPremDemo_TestCaseSetUp,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_OnPremDemo_TestCaseSetUp,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_OnPremDemo_TestCaseSetUp,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseSetUp
	TestInstruction_OnPremDemo_TestCaseSetUp.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_OnPremDemo_TestCaseSetUp,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_OnPremDemo_TestCaseSetUp,
		TestInstructionTypeName:      TestInstructionTypeName_OnPremDemo_TestCaseSetUp,
		Deprecated:                   TestInstructionDeprecated_OnPremDemo_TestCaseSetUp,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_OnPremDemo_TestCaseSetUp,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_OnPremDemo_TestCaseSetUp,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_OnPremDemo_TestCaseSetUp,
		TCRuleDeletion:               TCRuleDeletion_OnPremDemo_TestCaseSetUp,
		TCRuleSwap:                   TCRuleSwap_OnPremDemo_TestCaseSetUp,
		TestInstructionDescription:   TestInstructionDescription_OnPremDemo_TestCaseSetUp,
		TestInstructionMouseOverText: TestInstructionMouseOverText_OnPremDemo_TestCaseSetUp,
		Enabled:                      TestInstructionEnabled_OnPremDemo_TestCaseSetUp,
	}

	// DropZone 'TestCaseSetUp_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseSetUp_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_OnPremDemo_TestCaseSetUp_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_OnPremDemo_TestCaseSetUp_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_OnPremDemo_TestCaseSetUp,
		DropZoneUUID:                 TestInstructionDropZoneUUID_OnPremDemo_TestCaseSetUp_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_OnPremDemo_TestCaseSetUp_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_OnPremDemo_TestCaseSetUp_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_OnPremDemo_TestCaseSetUp_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_OnPremDemo_TestCaseSetUp_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
	}
	TestInstruction_OnPremDemo_TestCaseSetUp.ImmatureTestInstructionInformation = append(
		TestInstruction_OnPremDemo_TestCaseSetUp.ImmatureTestInstructionInformation,
		TestInstruction_OnPremDemo_TestCaseSetUp_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_OnPremDemo_TestCaseSetUp_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_OnPremDemo_TestCaseSetUp_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_OnPremDemo,
		DomainName:                                    DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:                           TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		TestInstructionName:                           TestInstructionName_OnPremDemo_TestCaseSetUp,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
	}
	TestInstruction_OnPremDemo_TestCaseSetUp.TestInstructionAttribute = append(
		TestInstruction_OnPremDemo_TestCaseSetUp.TestInstructionAttribute,
		TestInstructionAttribute_OnPremDemo_TestCaseSetUp_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed *FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = &FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructions.TestInstructionAttributeTypeUUID_OnPremDemo_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_OnPremDemo_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_OnPremDemo_GeneralAttribute_ExpectedToBePassed,
	}
	fangMethodAttributeMap[TestInstructionAttributeUUID_OnPremDemo_TestCaseSetUp_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed
	//TestInstruction_OnPremDemo_TestCaseSetUp.LocalExecutionMethods.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_OnPremDemo_TestCaseSetUp_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseSetUp
	var TestInstructionImmatureElementModel_OnPremDemo_TestCaseSetUp *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_OnPremDemo_TestCaseSetUp = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_OnPremDemo,
		DomainName:               DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID:      TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_OnPremDemo_TestCaseSetUp),
		PreviousElementUUID:      TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		NextElementUUID:          TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		FirstChildElementUUID:    TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		ParentElementUUID:        TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		TopImmatureElementUUID:   TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		IsTopElement:             true,
	}
	TestInstruction_OnPremDemo_TestCaseSetUp.ImmatureElementModel = append(
		TestInstruction_OnPremDemo_TestCaseSetUp.ImmatureElementModel,
		TestInstructionImmatureElementModel_OnPremDemo_TestCaseSetUp)

	return TestInstruction_OnPremDemo_TestCaseSetUp
}
