package knowledgebase

import "github.com/anondigriz/mogan-core/pkg/entities/types"

type Pattern struct {
	BaseInfo
	Type      types.PatternType
	ExtraData ExtraDataPattern
}

type ExtraDataPattern struct {
	Description      string
	Language         types.ScriptLanguageType
	Script           string
	InputParameters  []ParameterPattern
	OutputParameters []ParameterPattern
}

type ParameterPattern struct {
	ShortName string
	Type      types.ParameterType
}
