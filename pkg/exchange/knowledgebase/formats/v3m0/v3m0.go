package v3m0

import (
	"encoding/xml"
)

type (
	ParameterType      string
	PatternType        string
	ScriptLanguageType string
)

const (
	String  ParameterType = "string"
	Double  ParameterType = "double"
	Boolean ParameterType = "boolean"

	Program    PatternType = "program"
	Constraint PatternType = "constraint"

	JS  ScriptLanguageType = "js"
	Lua ScriptLanguageType = "lua"
)

type Description struct {
	Text string `xml:",cdata"`
}

type BaseInfo struct {
	ID           string      `xml:"id"`
	ShortName    string      `xml:"shortName"`
	Description  Description `xml:"description"`
	CreatedDate  int64       `xml:"createdDate"`
	ModifiedDate int64       `xml:"modifiedDate"`
}

type KnowledgeBase struct {
	BaseInfo
	XMLName          xml.Name `xml:"knowledgeBase"`
	FormatXMLVersion string   `xml:"formatXmlVersion,attr"`
	Parameters       struct {
		Parameters []Parameter `xml:"parameter"`
	} `xml:"parameters"`
	Patterns struct {
		Patterns []Pattern `xml:"pattern"`
	} `xml:"patterns"`
	Rules struct {
		Rules []Rule `xml:"rule"`
	} `xml:"rules"`
	Groups struct {
		Groups []Group `xml:"group"`
	} `xml:"groups"`
}

type Group struct {
	BaseInfo
	Groups struct {
		Groups []Group `xml:"group"`
	} `xml:"groups"`
	Parameters struct {
		Parameters []string `xml:"parameter"`
	} `xml:"parameters"`
	Rules struct {
		Rules []string `xml:"rule"`
	} `xml:"rules"`
}

type Parameter struct {
	BaseInfo
	Type         string `xml:"type"`
	DefaultValue string `xml:"defaultValue"`
}

type Script struct {
	Text string `xml:",cdata"`
}

type Pattern struct {
	BaseInfo
	Type            string `xml:"type"`
	Language        string `xml:"language"`
	Script          Script `xml:"script"`
	InputParameters struct {
		InputParameters []ParameterPattern `xml:"parameter"`
	} `xml:"inputParameters"`
	OutputParameters struct {
		OutputParameters []ParameterPattern `xml:"parameter"`
	} `xml:"outputParameters"`
}

type ParameterPattern struct {
	ShortName string `xml:"shortName"`
	Type      string `xml:"type"`
}

type Rule struct {
	BaseInfo
	PatternID       string `xml:"patternID"`
	InputParameters struct {
		InputParameters []ParameterRule `xml:"parameter"`
	} `xml:"inputParameters"`
	OutputParameters struct {
		OutputParameters []ParameterRule `xml:"parameter"`
	} `xml:"outputParameters"`
}

type ParameterRule struct {
	ShortName   string `xml:"shortName"`
	ParameterID string `xml:"parameterID"`
}
