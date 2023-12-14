package LocalExecutionMethods

import (
	"FenixOnPremDemoTestInstructionAdmin/LocalExecutionMethods/FangEngineClassesAndMethods"
)

// MethodsForLocalExecutionsStruct
// Struct for holding all information of how to execute a TestInstruction
// Even when there are no information about the methods this struct must exist
type MethodsForLocalExecutionsStruct struct {
	LocalParametersUsedInRunTime       *LocalParametersUsedInRunTimeStruct                                   `json:"LocalParametersUsedInRunTime"`
	FangEngineClassesMethodsAttributes *FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct `json:"FangEngineClassesMethodsAttributes"`
}

// LocalParametersUsedInRunTimeStruct
// Struct for holding the local parameters used during runtime
type LocalParametersUsedInRunTimeStruct struct {
	ExpectedTestInstructionExecutionDurationInSeconds int64 `json:"TimeOutTimeInSeconds"`
}
