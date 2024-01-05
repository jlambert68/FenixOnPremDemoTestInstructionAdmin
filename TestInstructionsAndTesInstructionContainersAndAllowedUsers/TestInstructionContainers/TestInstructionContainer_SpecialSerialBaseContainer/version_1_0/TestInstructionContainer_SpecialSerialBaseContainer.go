package version_1_0

import (
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructionContainers/TestInstructionContainer_SpecialSerialBaseContainer"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp"
	TI_TestCaseSetUp "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseSetUp/version_1_1"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown"
	TI_TestCaseTearDown "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_GeneralSetupTearDown_TestCaseTearDown/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Bonds"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (
	// *** TestInstructionContainer 'SpecialSerialBaseContainer'
	TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer                        TypeAndStructs.OriginalElementUUIDType          = TestInstructionContainer_SpecialSerialBaseContainer.TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer
	TestInstructionContainerName_OnPremDemo_SpecialSerialBaseContainer                        TypeAndStructs.TestInstructionContainerNameType = "Serial Special TestInstructionsContainer"
	TestInstructionContainerTypeUUID_OnPremDemo_SpecialSerialBaseContainer                                                                    = TestInstructionContainers.TestInstructionContainerTypeUUID_OnPremDemo_BaseContainers
	TestInstructionContainerTypeName_OnPremDemo_SpecialSerialBaseContainer                                                                    = TestInstructionContainers.TestInstructionContainerTypeNameType_OnPremDemo_BaseContainers
	TestInstructionContainerDescription_OnPremDemo_SpecialSerialBaseContainer                 string                                          = "Children of this special container is processed in serial"
	TestInstructionContainerMouseOverText_OnPremDemo_SpecialSerialBaseContainer               string                                          = "Children of this special container is processed in serial"
	TestInstructionContainerDeprecated_OnPremDemo_SpecialSerialBaseContainer                  bool                                            = false
	TestInstructionContainerEnabled_OnPremDemo_SpecialSerialBaseContainer                     bool                                            = true
	TestInstructionContainerMajorVersionNumber_OnPremDemo_SpecialSerialBaseContainer          int                                             = 1
	TestInstructionContainerMinorVersionNumber_OnPremDemo_SpecialSerialBaseContainer          int                                             = 0
	TestInstructionContainerChildrenIsParallelProcessed_OnPremDemo_SpecialSerialBaseContainer bool                                            = false
	TestInstructionContainerColor_OnPremDemo_SpecialSerialBaseContainer                       TypeAndStructs.ColorType                        = "#AAAAAA"
	TCRuleDeletion_OnPremDemo_SpecialSerialBaseContainer                                      TypeAndStructs.TCRuleDeletionType               = "TCRuleDeletion011"
	TCRuleSwap_OnPremDemo_SpecialSerialBaseContainer                                          TypeAndStructs.TCRuleSwapType                   = "TCRuleSwap012"
	TestInstructionContainerCreatingTimeStamp                                                 TypeAndStructs.UpdatedTimeStampType             = "2023-11-28 14:00:00"

	// *** DropZone *** 'TestInstructionsInTIC_ExpectsToSucceed'
	TestInstructionDropZoneUUID_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed         TypeAndStructs.DropZoneUUIDType = "fdfa0f19-16df-4ce2-88bf-88acbf1db48d"
	TestInstructionDropZoneName_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed         TypeAndStructs.DropZoneNameType = "AllTestInstructionsInTIC_ExpectsToSucceed"
	TestInstructionDropZoneDescription_OnPremDemo_TTestInstructionsInTIC_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed    string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed        TypeAndStructs.ColorType        = "#FF0000FF"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeActionCommand_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType     = Domains.TestInstructionAttributeValueUUID_TRUE
)

// TestInstructionContainer_OnPremDemo_SpecialSerialBase
// Variable that holds the data for the TestInstructionContainer
var TestInstructionContainer_OnPremDemo_SpecialSerialBase *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct

// Initiate_TestInstructionContainer_OnPremDemo_Serial
// Function that creates all data for the TestInstructionContainer
func Initiate_TestInstructionContainer_OnPremDemo_Serial(testInstructionsAndTestInstructionContainersMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct) {

	// Initiate the full structure
	TestInstructionContainer_OnPremDemo_SpecialSerialBase = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct{}

	// TestInstructionContainer - 'SpecialSerialBaseContainer'
	TestInstructionContainer_OnPremDemo_SpecialSerialBase.TestInstructionContainer = &TypeAndStructs.TestInstructionContainerStruct{
		DomainUUID:                            DomainData.DomainUUID_OnPremDemo,
		DomainName:                            DomainData.DomainName_OnPremDemo,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerName:          TestInstructionContainerName_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_OnPremDemo_SpecialSerialBaseContainer,
		Deprecated:                            TestInstructionContainerDeprecated_OnPremDemo_SpecialSerialBaseContainer,
		Enabled:                               TestInstructionContainerEnabled_OnPremDemo_SpecialSerialBaseContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_OnPremDemo_SpecialSerialBaseContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_OnPremDemo_SpecialSerialBaseContainer,
		UpdatedTimeStamp:                      TestInstructionContainerCreatingTimeStamp,
		ChildrenIsParallelProcessed:           TestInstructionContainerChildrenIsParallelProcessed_OnPremDemo_SpecialSerialBaseContainer,
	}

	// BasicTestInstructionContainerInformation - 'SpecialSerialBaseContainer'
	TestInstructionContainer_OnPremDemo_SpecialSerialBase.BasicTestInstructionContainerInformation = &TypeAndStructs.BasicTestInstructionContainerInformationStruct{
		DomainUUID:                            DomainData.DomainUUID_OnPremDemo,
		DomainName:                            DomainData.DomainName_OnPremDemo,
		TestInstructionContainerUUID:          TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerName:          TestInstructionContainerName_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerTypeUUID:      TestInstructionContainerTypeUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerTypeName:      TestInstructionContainerTypeName_OnPremDemo_SpecialSerialBaseContainer,
		Deprecated:                            TestInstructionContainerDeprecated_OnPremDemo_SpecialSerialBaseContainer,
		MajorVersionNumber:                    TestInstructionContainerMajorVersionNumber_OnPremDemo_SpecialSerialBaseContainer,
		MinorVersionNumber:                    TestInstructionContainerMinorVersionNumber_OnPremDemo_SpecialSerialBaseContainer,
		UpdatedTimeStamp:                      TestInstructionContainerCreatingTimeStamp,
		TestInstructionContainerColor:         TestInstructionContainerColor_OnPremDemo_SpecialSerialBaseContainer,
		TCRuleDeletion:                        TCRuleDeletion_OnPremDemo_SpecialSerialBaseContainer,
		TCRuleSwap:                            TCRuleSwap_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerDescription:   TestInstructionContainerDescription_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerMouseOverText: TestInstructionContainerMouseOverText_OnPremDemo_SpecialSerialBaseContainer,
		Enabled:                               TestInstructionContainerEnabled_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerExecutionType: Domains.TestInstructionContainerExecutionType_SERIAL_PROCESSED,
	}

	// ImmatureTestInstructionContainerMessage - 'SpecialSerialBaseContainer'
	// *** DropZone *** 'TestInstructionsInTIC_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseSetUp_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_OnPremDemo_TestCaseSetUp_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionContainerMessageStruct
	TestInstruction_OnPremDemo_TestCaseSetUp_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionContainerMessageStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionContainerUUID: TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerName: TestInstructionContainerName_OnPremDemo_SpecialSerialBaseContainer,
		DropZoneUUID:                 TestInstructionDropZoneUUID_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_OnPremDemo_TTestInstructionsInTIC_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed,
		TestInstructionAttributeType: TI_TestCaseSetUp.TestInstructionAttributeType_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeUUID: TI_TestCaseSetUp.TestInstructionAttributeUUID_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName: TI_TestCaseSetUp.TestInstructionAttributeName_OnPremDemo_TestCaseSetUp_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed,
		FirstImmatureElementUUID:     TI_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed,
	}

	TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureTestInstructionContainer = append(
		TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureTestInstructionContainer,
		TestInstruction_OnPremDemo_TestCaseSetUp_ExpectedToBePassed)

	// DropZone 'TestInstructionsInTIC_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestInstructionsInTIC_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionContainerMessageStruct
	TestInstruction_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionContainerMessageStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionContainerUUID: TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestInstructionContainerName: TestInstructionContainerName_OnPremDemo_SpecialSerialBaseContainer,
		DropZoneUUID:                 TestInstructionDropZoneUUID_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_OnPremDemo_TTestInstructionsInTIC_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_OnPremDemo_TestInstructionsInTIC_ExpectsToSucceed,
		TestInstructionAttributeType: TI_TestCaseTearDown.TestInstructionAttributeType_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeUUID: TI_TestCaseTearDown.TestInstructionAttributeUUID_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName: TI_TestCaseTearDown.TestInstructionAttributeName_OnPremDemo_TestCaseTearDown_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed,
		FirstImmatureElementUUID:     TI_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed,
	}

	TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureTestInstructionContainer = append(
		TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureTestInstructionContainer,
		TestInstruction_OnPremDemo_TestInstructionsInTIC_ExpectedToBePassed)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'TIC' in 'TIC(B11fx-TI-B12-TI-B11lx)'
	var ImmatureElementModel_TIC *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TIC = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_OnPremDemo,
		DomainName:               DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID:      TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionContainerName_OnPremDemo_SpecialSerialBaseContainer),
		PreviousElementUUID:      TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		NextElementUUID:          TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		FirstChildElementUUID:    Bonds.Bond_B11fx_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TIC,
		OriginalElementUUID:      TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TopImmatureElementUUID:   TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		IsTopElement:             true,
	}
	TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_TIC)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'B11fx_1' in 'TIC(B11fx_1-TIx_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_B11fx_1 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B11fx_1 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:          DomainData.DomainUUID_OnPremDemo,
		DomainName:          DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID: Bonds.Bond_B11fx_BondUuid,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(Bonds.Bond_B11fx_BondName),
		PreviousElementUUID: Bonds.Bond_B11fx_BondUuid,
		NextElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		FirstChildElementUUID:    Bonds.Bond_B11fx_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B11fx,
		OriginalElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		TopImmatureElementUUID: TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		IsTopElement:           false,
	}
	TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_B11fx_1)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'TI_1' in 'TIC(B11fx_1-TIx_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_TIx_1 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TIx_1 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID: DomainData.DomainUUID_OnPremDemo,
		DomainName: DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionName),
		PreviousElementUUID: Bonds.Bond_B11fx_BondUuid,
		NextElementUUID:     Bonds.Bond_B12_BondUuid,
		FirstChildElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ParentElementUUID:        TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TIx,
		OriginalElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		TopImmatureElementUUID: TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		IsTopElement:           false,
	}
	TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_TIx_1)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'B12' in 'TIC(B11fx_1-TIx_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_B12 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B12 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:          DomainData.DomainUUID_OnPremDemo,
		DomainName:          DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID: Bonds.Bond_B12_BondUuid,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(Bonds.Bond_B12_BondName),
		PreviousElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseSetUp.TestInstructionUUID_OnPremDemo_TestCaseSetUp].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		NextElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		FirstChildElementUUID:    Bonds.Bond_B12_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B12,
		OriginalElementUUID:      Bonds.Bond_B12_BondUuid,
		TopImmatureElementUUID:   TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		IsTopElement:             false,
	}
	TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_B12)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'TI_2' in 'TIC(B11fx_1-TIx_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_TI_2 *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_TI_2 = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID: DomainData.DomainUUID_OnPremDemo,
		DomainName: DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionName),
		PreviousElementUUID: Bonds.Bond_B12_BondUuid,
		NextElementUUID:     Bonds.Bond_B11lx_BondUuid,
		FirstChildElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		ParentElementUUID:        TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		TopImmatureElementUUID: TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		IsTopElement:           false,
	}
	TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_TI_2)

	// ImmatureElementModelMessage - 'SpecialSerialBaseContainer' - 'B10' in 'TIC(B11fx_1-TIx_1-B12-TI_2-B11lx)'
	var ImmatureElementModel_B11lx *TypeAndStructs.ImmatureElementModelMessageStruct
	ImmatureElementModel_B11lx = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:          DomainData.DomainUUID_OnPremDemo,
		DomainName:          DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID: Bonds.Bond_B11lx_BondUuid,
		ImmatureElementName: TypeAndStructs.OriginalElementNameType(Bonds.Bond_B11lx_BondName),
		PreviousElementUUID: testInstructionsAndTestInstructionContainersMessage.TestInstructions.
			TestInstructionsMap[TestInstruction_GeneralSetupTearDown_TestCaseTearDown.TestInstructionUUID_OnPremDemo_TestCaseTearDown].
			TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
		NextElementUUID:          Bonds.Bond_B11lx_BondUuid,
		FirstChildElementUUID:    Bonds.Bond_B11lx_BondUuid,
		ParentElementUUID:        TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_B11lx,
		OriginalElementUUID:      Bonds.Bond_B11lx_BondUuid,
		TopImmatureElementUUID:   TestInstructionContainerUUID_OnPremDemo_SpecialSerialBaseContainer,
		IsTopElement:             false,
	}
	TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel = append(TestInstructionContainer_OnPremDemo_SpecialSerialBase.ImmatureElementModel, ImmatureElementModel_B11lx)

}
