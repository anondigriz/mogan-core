package knowledgebase

type Rule struct {
	BaseInfo
	PatternUUID      string
	InputParameters  []ParameterRule
	OutputParameters []ParameterRule
}

type ParameterRule struct {
	ShortName     string
	ParameterUUID string
}
