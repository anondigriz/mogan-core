package types

type ParameterType int

const (
	String  ParameterType = 0
	Double  ParameterType = 1
	Boolean ParameterType = 2
)

type PatternType int

const (
	Program    PatternType = 0
	Constraint PatternType = 1
)

type ScriptLanguageType int

const (
	JS  ScriptLanguageType = 0
	Lua ScriptLanguageType = 1
)
