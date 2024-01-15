package version_1_0

import (
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/LocalExecutionMethods/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'TestCaseTearDown'
	TestInstructionUUID_OnPremDemo_TestCaseTearDown               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_OnPremDemo_TestCaseTearDown
	TestInstructionName_OnPremDemo_TestCaseTearDown               TypeAndStructs.TestInstructionNameType = "TestCaseTearDown"
	TestInstructionTypeUUID_OnPremDemo_TestCaseTearDown                                                  = TestInstructions.TestInstructionTypeUUID_OnPremDemo_GeneralSetUpTearDown
	TestInstructionTypeName_OnPremDemo_TestCaseTearDown                                                  = TestInstructions.TestInstructionTypeName_OnPremDemo_GeneralSetUpTearDown
	TestInstructionDescription_OnPremDemo_TestCaseTearDown        string                                 = "TestCaseTearDown, runs last for every TestCase for OnPrem-demo"
	TestInstructionMouseOverText_OnPremDemo_TestCaseTearDown      string                                 = "TestCaseTearDown, runs last for every TestCase for OnPrem-demo"
	TestInstructionDeprecated_OnPremDemo_TestCaseTearDown         bool                                   = false
	TestInstructionEnabled_OnPremDemo_TestCaseTearDown            bool                                   = false
	TestInstructionMajorVersionNumber_OnPremDemo_TestCaseTearDown int                                    = 1
	TestInstructionMinorVersionNumber_OnPremDemo_TestCaseTearDown int                                    = 0
	TestInstructionColor_OnPremDemo_TestCaseTearDown              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_OnPremDemo_TestCaseTearDown                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_OnPremDemo_TestCaseTearDown                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                              TypeAndStructs.UpdatedTimeStampType    = "2023-11-28 14:00:00"

	// *** DropZone *** 'TestCaseTearDown_ExpectsToSucceed'
	TestInstructionDropZoneUUID_OnPremDemo_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "bdf1b383-a4c5-4fb2-a9a1-17a7bea75efb"
	TestInstructionDropZoneName_OnPremDemo_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseTearDown_ExpectsToSucceed"
	TestInstructionDropZoneDescription_OnPremDemo_TestCaseTearDown_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_OnPremDemo_TestCaseTearDown_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_OnPremDemo_TestCaseTearDown_ExpectsToSucceed       TypeAndStructs.ColorType        = "#000000AA"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_OnPremDemo_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "737b77d1-f468-4f07-acef-bdc087d71fe0" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeName_OnPremDemo_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeType_OnPremDemo_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeActionCommand_OnPremDemo_TestCaseTearDown_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_OnPremDemo_TestCaseTearDown_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_OnPremDemo_TestCaseTearDown_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_OnPremDemo_TestCaseTearDown_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_OnPremDemo_TestCaseTearDown_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	FangEngine_Class_Name_OnPremDemo_TestCaseTearDown = "GeneralSetupTearDown"
)

// TestInstruction_OnPremDemo_TestCaseTearDown
// Variable that holds the data for the TestInstruction
var TestInstruction_OnPremDemo_TestCaseTearDown *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_OnPremDemo_TestCaseTearDown
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_OnPremDemo_TestCaseTearDown() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	var fangMethodAttributeMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineClassesAndMethods.FangEngineAttributesStruct
	fangMethodAttributeMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineClassesAndMethods.FangEngineAttributesStruct)

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_OnPremDemo_TestCaseTearDown = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
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
				FangEngineClassesMethodsAttributes: &FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
					TestInstructionOriginalUUID: TestInstructionUUID_OnPremDemo_TestCaseTearDown,
					TestInstructionName:         TestInstructionName_OnPremDemo_TestCaseTearDown,
					FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_OnPremDemo_GeneralSetupTearDown,
					FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_OnPremDemo_GeneralSetupTearDown,
					FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_OnPremDemo_GeneralSetupTearDown_TearDown,
					FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_OnPremDemo_GeneralSetupTearDown_TearDown,
					Attributes:                  fangMethodAttributeMap,
				},
			},
		},
	}

	// TestInstruction - TestCaseTearDown
	TestInstruction_OnPremDemo_TestCaseTearDown.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_OnPremDemo_TestCaseTearDown,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_OnPremDemo_TestCaseTearDown,
		TestInstructionTypeName:      TestInstructionTypeName_OnPremDemo_TestCaseTearDown,
		TestInstructionDescription:   TestInstructionDescription_OnPremDemo_TestCaseTearDown,
		TestInstructionMouseOverText: TestInstructionMouseOverText_OnPremDemo_TestCaseTearDown,
		Deprecated:                   TestInstructionDeprecated_OnPremDemo_TestCaseTearDown,
		Enabled:                      TestInstructionEnabled_OnPremDemo_TestCaseTearDown,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_OnPremDemo_TestCaseTearDown,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_OnPremDemo_TestCaseTearDown,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseTearDown
	TestInstruction_OnPremDemo_TestCaseTearDown.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_OnPremDemo_TestCaseTearDown,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_OnPremDemo_TestCaseTearDown,
		TestInstructionTypeName:      TestInstructionTypeName_OnPremDemo_TestCaseTearDown,
		Deprecated:                   TestInstructionDeprecated_OnPremDemo_TestCaseTearDown,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_OnPremDemo_TestCaseTearDown,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_OnPremDemo_TestCaseTearDown,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_OnPremDemo_TestCaseTearDown,
		TCRuleDeletion:               TCRuleDeletion_OnPremDemo_TestCaseTearDown,
		TCRuleSwap:                   TCRuleSwap_OnPremDemo_TestCaseTearDown,
		TestInstructionDescription:   TestInstructionDescription_OnPremDemo_TestCaseTearDown,
		TestInstructionMouseOverText: TestInstructionMouseOverText_OnPremDemo_TestCaseTearDown,
		Enabled:                      TestInstructionEnabled_OnPremDemo_TestCaseTearDown,
	}

	// DropZone 'TestCaseTearDown_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseTearDown_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_OnPremDemo_TestCaseTearDown_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_OnPremDemo_TestCaseTearDown_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_OnPremDemo_TestCaseTearDown,
		DropZoneUUID:                 TestInstructionDropZoneUUID_OnPremDemo_TestCaseTearDown_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_OnPremDemo_TestCaseTearDown_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_OnPremDemo_TestCaseTearDown_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_OnPremDemo_TestCaseTearDown_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_OnPremDemo_TestCaseTearDown_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
	}
	TestInstruction_OnPremDemo_TestCaseTearDown.ImmatureTestInstructionInformation = append(
		TestInstruction_OnPremDemo_TestCaseTearDown.ImmatureTestInstructionInformation,
		TestInstruction_OnPremDemo_TestCaseTearDown_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_OnPremDemo_TestCaseTearDown_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_OnPremDemo_TestCaseTearDown_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_OnPremDemo,
		DomainName:                                    DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:                           TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		TestInstructionName:                           TestInstructionName_OnPremDemo_TestCaseTearDown,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
	}
	TestInstruction_OnPremDemo_TestCaseTearDown.TestInstructionAttribute = append(
		TestInstruction_OnPremDemo_TestCaseTearDown.TestInstructionAttribute,
		TestInstructionAttribute_OnPremDemo_TestCaseTearDown_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed *FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = &FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructions.TestInstructionAttributeTypeUUID_OnPremDemo_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_OnPremDemo_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_OnPremDemo_GeneralAttribute_ExpectedToBePassed,
	}
	fangMethodAttributeMap[TestInstructionAttributeUUID_OnPremDemo_TestCaseTearDown_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed
	//TestInstruction_OnPremDemo_TestCaseTearDown.LocalExecutionMethods.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_OnPremDemo_TestCaseTearDown_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseTearDown
	var TestInstructionImmatureElementModel_OnPremDemo_TestCaseTearDown *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_OnPremDemo_TestCaseTearDown = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_OnPremDemo,
		DomainName:               DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID:      TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_OnPremDemo_TestCaseTearDown),
		PreviousElementUUID:      TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		NextElementUUID:          TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		FirstChildElementUUID:    TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		ParentElementUUID:        TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		TopImmatureElementUUID:   TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		IsTopElement:             true,
	}
	TestInstruction_OnPremDemo_TestCaseTearDown.ImmatureElementModel = append(
		TestInstruction_OnPremDemo_TestCaseTearDown.ImmatureElementModel,
		TestInstructionImmatureElementModel_OnPremDemo_TestCaseTearDown)

	return TestInstruction_OnPremDemo_TestCaseTearDown
}
