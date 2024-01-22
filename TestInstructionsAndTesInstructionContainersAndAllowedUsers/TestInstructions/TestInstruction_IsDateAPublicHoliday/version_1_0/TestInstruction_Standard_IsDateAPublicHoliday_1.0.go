package version_1_0

import (
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/LocalExecutionMethods"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/LocalExecutionMethods/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/DomainData"
	"github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions"
	fixedValuesOverVersions "github.com/jlambert68/FenixOnPremDemoTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_IsDateAPublicHoliday"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/Domains"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

const (

	// *************************************
	// *** TestInstruction *** 'IsDateAPublicHoliday'
	TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday               TypeAndStructs.OriginalElementUUIDType = fixedValuesOverVersions.TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday
	TestInstructionName_OnPremDemo_IsDateAPublicHoliday               TypeAndStructs.TestInstructionNameType = "IsDateAPublicHoliday"
	TestInstructionTypeUUID_OnPremDemo_IsDateAPublicHoliday                                                  = TestInstructions.TestInstructionTypeUUID_OnPremDemo_Standard
	TestInstructionTypeName_OnPremDemo_IsDateAPublicHoliday                                                  = TestInstructions.TestInstructionTypeName_OnPremDemo_Standard
	TestInstructionDescription_OnPremDemo_IsDateAPublicHoliday        string                                 = "This TestInstruction validates if a date is public holiday or not"
	TestInstructionMouseOverText_OnPremDemo_IsDateAPublicHoliday      string                                 = "This TestInstruction validates if a date is public holiday or not"
	TestInstructionDeprecated_OnPremDemo_IsDateAPublicHoliday         bool                                   = false
	TestInstructionEnabled_OnPremDemo_IsDateAPublicHoliday            bool                                   = true
	TestInstructionMajorVersionNumber_OnPremDemo_IsDateAPublicHoliday int                                    = 1
	TestInstructionMinorVersionNumber_OnPremDemo_IsDateAPublicHoliday int                                    = 0
	TestInstructionColor_OnPremDemo_IsDateAPublicHoliday              TypeAndStructs.ColorType               = "#ffff00AA"
	TCRuleDeletion_OnPremDemo_IsDateAPublicHoliday                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_OnPremDemo_IsDateAPublicHoliday                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"
	TestInstructionCreatingTimeStamp                                  TypeAndStructs.UpdatedTimeStampType    = "2024-01-08 15:00:00"

	// *** DropZone *** 'IsDateAPublicHoliday_ExpectsToSucceed'
	TestInstructionDropZoneUUID_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "dc713785-9bed-4e6e-aeab-964a5dbf648a"
	TestInstructionDropZoneName_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "IsDateAPublicHoliday_ExpectsToSucceed"
	TestInstructionDropZoneDescription_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed       TypeAndStructs.ColorType        = "#ffff00AA"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = TestInstructions.TestInstructionAttributeUUID_OnPremDemo_ExpectedToBePassed // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeName_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructions.TestInstructionAttributeName_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeType_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructions.TestInstructionAttributeType_OnPremDemo_ExpectedToBePassed
	TestInstructionAttributeActionCommand_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	// Attribute - 'HolidayDateToCheck'
	TestInstructionAttributeUUID_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck          TypeAndStructs.TestInstructionAttributeUUIDType = "ebfeb406-75fc-4b5d-8e83-28d02d0d69d6" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_OnPremDemo_HolidayDateToCheck
	TestInstructionAttributeName_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck          TypeAndStructs.TestInstructionAttributeNameType = "Holiday Date To Check"
	TestInstructionAttributeType_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck   string                                          = "The date that should be checked if it is a Public Holiday, use format 'YYYY-MM-DD'"
	TestInstructionAttributeMouseOverText_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck string                                          = "The date that should be checked if it is a Public Holiday, use format 'YYYY-MM-DD'"

	// Attribute - 'CountryCode'
	TestInstructionAttributeUUID_OnPremDemo_IsDateAPublicHoliday_CountryCode          TypeAndStructs.TestInstructionAttributeUUIDType = "371349ff-787d-4031-9e2e-4a937703cc0c" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_OnPremDemo_CountryCode
	TestInstructionAttributeName_OnPremDemo_IsDateAPublicHoliday_CountryCode          TypeAndStructs.TestInstructionAttributeNameType = "Country Code"
	TestInstructionAttributeType_OnPremDemo_IsDateAPublicHoliday_CountryCode          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_OnPremDemo_IsDateAPublicHoliday_CountryCode   string                                          = "The Country for which to check the Public Holiday, use format 'XX'"
	TestInstructionAttributeMouseOverText_OnPremDemo_IsDateAPublicHoliday_CountryCode string                                          = "The Country for which to check the Public Holiday, use format 'XX'"
)

// TestInstruction_OnPremDemo_IsDateAPublicHoliday
// Variable that holds the data for the TestInstruction
var TestInstruction_OnPremDemo_IsDateAPublicHoliday *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct

// Initate_TestInstruction_OnPremDemo_IsDateAPublicHoliday
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_OnPremDemo_IsDateAPublicHoliday() *TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_OnPremDemo_IsDateAPublicHoliday = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
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

	// TestInstruction - IsDateAPublicHoliday
	TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstruction = &TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_OnPremDemo_InGCP,
		ExecutionDomainName:          DomainData.ExecutionDomainName_OnPremDemo_InGCP,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionName:          TestInstructionName_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionTypeName:      TestInstructionTypeName_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionDescription:   TestInstructionDescription_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionMouseOverText: TestInstructionMouseOverText_OnPremDemo_IsDateAPublicHoliday,
		Deprecated:                   TestInstructionDeprecated_OnPremDemo_IsDateAPublicHoliday,
		Enabled:                      TestInstructionEnabled_OnPremDemo_IsDateAPublicHoliday,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_OnPremDemo_IsDateAPublicHoliday,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_OnPremDemo_IsDateAPublicHoliday,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
	}

	// BasicTestInstructionInformation - IsDateAPublicHoliday
	TestInstruction_OnPremDemo_IsDateAPublicHoliday.BasicTestInstructionInformation = &TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		ExecutionDomainUUID:          DomainData.ExecutionDomainUUID_OnPremDemo_InGCP,
		ExecutionDomainName:          DomainData.ExecutionDomainName_OnPremDemo_InGCP,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionName:          TestInstructionName_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionTypeName:      TestInstructionTypeName_OnPremDemo_IsDateAPublicHoliday,
		Deprecated:                   TestInstructionDeprecated_OnPremDemo_IsDateAPublicHoliday,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_OnPremDemo_IsDateAPublicHoliday,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_OnPremDemo_IsDateAPublicHoliday,
		UpdatedTimeStamp:             TestInstructionCreatingTimeStamp,
		TestInstructionColor:         TestInstructionColor_OnPremDemo_IsDateAPublicHoliday,
		TCRuleDeletion:               TCRuleDeletion_OnPremDemo_IsDateAPublicHoliday,
		TCRuleSwap:                   TCRuleSwap_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionDescription:   TestInstructionDescription_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionMouseOverText: TestInstructionMouseOverText_OnPremDemo_IsDateAPublicHoliday,
		Enabled:                      TestInstructionEnabled_OnPremDemo_IsDateAPublicHoliday,
	}

	// DropZone 'IsDateAPublicHoliday_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: IsDateAPublicHoliday_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed *TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   DomainData.DomainUUID_OnPremDemo,
		DomainName:                   DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:          TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionName:          TestInstructionName_OnPremDemo_IsDateAPublicHoliday,
		DropZoneUUID:                 TestInstructionDropZoneUUID_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_OnPremDemo_IsDateAPublicHoliday_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
	}
	TestInstruction_OnPremDemo_IsDateAPublicHoliday.ImmatureTestInstructionInformation = append(
		TestInstruction_OnPremDemo_IsDateAPublicHoliday.ImmatureTestInstructionInformation,
		TestInstruction_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_OnPremDemo,
		DomainName:                                    DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:                           TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionName:                           TestInstructionName_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed,
	}
	TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstructionAttribute = append(
		TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstructionAttribute,
		TestInstructionAttribute_OnPremDemo_IsDateAPublicHoliday_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// TestInstruction Attribute - 'HolidayDateToCheck'
	var TestInstructionAttribute_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_OnPremDemo,
		DomainName:                                    DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:                           TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionName:                           TestInstructionName_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck,
		TestInstructionAttributeName:                  TestInstructionAttributeName_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck,
	}
	TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstructionAttribute = append(
		TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstructionAttribute,
		TestInstructionAttribute_OnPremDemo_IsDateAPublicHoliday_HolidayDateToCheck)

	// TestInstruction Attribute - 'CountryCode'
	var TestInstructionAttribute_OnPremDemo_IsDateAPublicHoliday_CountryCode *TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_OnPremDemo_IsDateAPublicHoliday_CountryCode = &TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    DomainData.DomainUUID_OnPremDemo,
		DomainName:                                    DomainData.DomainName_OnPremDemo,
		TestInstructionUUID:                           TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionName:                           TestInstructionName_OnPremDemo_IsDateAPublicHoliday,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_OnPremDemo_IsDateAPublicHoliday_CountryCode,
		TestInstructionAttributeName:                  TestInstructionAttributeName_OnPremDemo_IsDateAPublicHoliday_CountryCode,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_OnPremDemo_IsDateAPublicHoliday_CountryCode,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_OnPremDemo_IsDateAPublicHoliday_CountryCode,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_OnPremDemo_IsDateAPublicHoliday_CountryCode,
	}
	TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstructionAttribute = append(
		TestInstruction_OnPremDemo_IsDateAPublicHoliday.TestInstructionAttribute,
		TestInstructionAttribute_OnPremDemo_IsDateAPublicHoliday_CountryCode)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	// Nothing here

	// ImmatureElementModel - IsDateAPublicHoliday
	var TestInstructionImmatureElementModel_OnPremDemo_IsDateAPublicHoliday *TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_OnPremDemo_IsDateAPublicHoliday = &TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               DomainData.DomainUUID_OnPremDemo,
		DomainName:               DomainData.DomainName_OnPremDemo,
		ImmatureElementUUID:      TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_OnPremDemo_IsDateAPublicHoliday),
		PreviousElementUUID:      TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		NextElementUUID:          TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		FirstChildElementUUID:    TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		ParentElementUUID:        TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		TopImmatureElementUUID:   TestInstructionUUID_OnPremDemo_IsDateAPublicHoliday,
		IsTopElement:             true,
	}
	TestInstruction_OnPremDemo_IsDateAPublicHoliday.ImmatureElementModel = append(
		TestInstruction_OnPremDemo_IsDateAPublicHoliday.ImmatureElementModel,
		TestInstructionImmatureElementModel_OnPremDemo_IsDateAPublicHoliday)

	return TestInstruction_OnPremDemo_IsDateAPublicHoliday
}
