package FangEngineClassesAndMethods

import "github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"

// Types for FangEngine
type FangEngine_ClassName_UUID_OnPremDemo_Type string
type FangEngine_ClassName_Name_OnPremDemo_Type string
type FangEngine_MethodName_UUID_OnPremDemo_Type string
type FangEngine_MethodName_Name_OnPremDemo_Type string
type FangEngine_AttributeName_UUID_OnPremDemo_Type string
type FangEngine_AttributeName_Name_OnPremDemo_Type string

// Type this is used for specifying Classes, Methods and Attributes for FangEngine which is used by _OnPremDemos TestAutomation
type FangEngineClassesMethodsAttributesStruct struct {
	TestInstructionOriginalUUID TypeAndStructs.OriginalElementUUIDType                                          `json:"TestInstructionOriginalUUID"`
	TestInstructionName         TypeAndStructs.TestInstructionNameType                                          `json:"TestInstructionName"`
	FangEngineClassNameUUID     FangEngine_ClassName_UUID_OnPremDemo_Type                                       `json:"FangEngineClassNameUUID"`
	FangEngineClassNameNAME     FangEngine_ClassName_Name_OnPremDemo_Type                                       `json:"FangEngineClassNameNAME"`
	FangEngineMethodNameUUID    FangEngine_MethodName_UUID_OnPremDemo_Type                                      `json:"FangEngineMethodNameUUID"`
	FangEngineMethodNameNAME    FangEngine_MethodName_Name_OnPremDemo_Type                                      `json:"FangEngineMethodNameNAME"`
	Attributes                  map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineAttributesStruct `json:"Attributes"`
}

type FangEngineAttributesStruct struct {
	TestInstructionAttributeUUID     TypeAndStructs.TestInstructionAttributeUUIDType     `json:"TestInstructionAttributeUUID"`
	TestInstructionAttributeName     TypeAndStructs.TestInstructionAttributeNameType     `json:"TestInstructionAttributeName"`
	TestInstructionAttributeTypeUUID TypeAndStructs.TestInstructionAttributeTypeUUIDType `json:"TestInstructionAttributeTypeUUID"`
	FangEngineAttributeNameUUID      FangEngine_AttributeName_UUID_OnPremDemo_Type       `json:"FangEngineAttributeNameUUID"`
	FangEngineAttributeNameName      FangEngine_AttributeName_Name_OnPremDemo_Type       `json:"FangEngineAttributeNameName"`
}

// Classes, Methods and their Parameters in FangEngine for _OnPremDemo
const (

	// General Attribute - ''
	FangEngine_ClassName_UUID_OnPremDemo_GeneralAttribute_ExpectedToBePassed FangEngine_AttributeName_UUID_OnPremDemo_Type = "9b9e4ca8-e9a3-4939-b9dc-184b4e44f60e"
	FangEngine_ClassName_Name_OnPremDemo_GeneralAttribute_ExpectedToBePassed FangEngine_AttributeName_Name_OnPremDemo_Type = "expectedToBePassed"

	// ClassName - ***** 'GeneralSetupTearDown' *****
	FangEngine_ClassName_UUID_OnPremDemo_GeneralSetupTearDown FangEngine_ClassName_UUID_OnPremDemo_Type = "85373d2b-30ec-49ee-823a-0d8a0b2d5599"
	FangEngine_ClassName_Name_OnPremDemo_GeneralSetupTearDown FangEngine_ClassName_Name_OnPremDemo_Type = "GeneralSetupTearDown"

	// ClassName: 'GeneralSetupTearDown' - MethodName: 'Setup'
	FangEngine_MethodName_UUID_OnPremDemo_GeneralSetupTearDown_Setup FangEngine_MethodName_UUID_OnPremDemo_Type = "093cfe88-0970-427e-9548-82568bfede8c"
	FangEngine_MethodName_Name_OnPremDemo_GeneralSetupTearDown_Setup FangEngine_MethodName_Name_OnPremDemo_Type = "Setup"

	// ClassName: 'GeneralSetupTearDown' - MethodName: 'TearDown'
	FangEngine_MethodName_UUID_OnPremDemo_GeneralSetupTearDown_TearDown FangEngine_MethodName_UUID_OnPremDemo_Type = "0db4a61c-5a85-49a2-b039-4f411de0edd9"
	FangEngine_MethodName_Name_OnPremDemo_GeneralSetupTearDown_TearDown FangEngine_MethodName_Name_OnPremDemo_Type = "TearDown"
)
