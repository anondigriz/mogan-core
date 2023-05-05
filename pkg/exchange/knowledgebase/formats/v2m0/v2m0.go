package v2m0

import "encoding/xml"

type BaseInfo struct {
	ID          string `xml:"id,attr"`
	ShortName   string `xml:"shortName,attr"`
	Description string `xml:"description,attr"`
}

type Model struct {
	BaseInfo
	XMLName          xml.Name `xml:"model"`
	FormatXMLVersion string   `xml:"formatXmlVersion,attr"`
	Class            Class    `xml:"class"`
	Relations        struct {
		Relations []Relation `xml:"relation"`
	} `xml:"relations"`
}

type Class struct {
	BaseInfo
	Classes struct {
		Classes []Class `xml:"class"`
	} `xml:"classes"`
	Rules struct {
		Rules []Rule `xml:"rule"`
	} `xml:"rules"`
	Constraints struct {
		Constraints []Rule `xml:"constraint"`
	} `xml:"constraints,omitempty"`
	Parameters struct {
		Parameters []Parameter `xml:"parameter"`
	} `xml:"parameters"`
}

type Rule struct {
	BaseInfo
	RelationID string `xml:"relation,attr"`
	InitIDs    string `xml:"initId,attr"`
	ResultIDs  string `xml:"resultId,attr"`
}

type Relation struct {
	BaseInfo
	OutObjects   string `xml:"outObj,attr"`
	InObjects    string `xml:"inObj,attr"`
	RelationType string `xml:"relationType,attr"`
	Script       string `xml:",innerxml"`
}

type Parameter struct {
	BaseInfo
	Type         string `xml:"type,attr"`
	DefaultValue string `xml:"defaultValue,attr"`
}
