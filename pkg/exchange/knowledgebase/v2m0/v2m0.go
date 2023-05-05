package v2m0

import "encoding/xml"

type Model struct {
	XMLName          xml.Name `xml:"model"`
	ID               string   `xml:"id,attr"`
	ShortName        string   `xml:"shortName,attr"`
	FormatXMLVersion string   `xml:"formatXmlVersion,attr"`
	Description      string   `xml:"description,attr"`
	Class            Class    `xml:"class"`
	Relations        struct {
		Relations []Relation `xml:"relation"`
	} `xml:"relations"`
}

type Class struct {
	ID          string `xml:"id,attr"`
	ShortName   string `xml:"shortName,attr"`
	Description string `xml:"description,attr"`
	Classes     struct {
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
	ID          string `xml:"id,attr"`
	ShortName   string `xml:"shortName,attr"`
	Description string `xml:"description,attr"`
	RelationID  string `xml:"relation,attr"`
	InitIDs     string `xml:"initId,attr"`
	ResultIDs   string `xml:"resultId,attr"`
}

type Relation struct {
	ID           string `xml:"id,attr"`
	ShortName    string `xml:"shortName,attr"`
	Description  string `xml:"description,attr"`
	OutObjects   string `xml:"outObj,attr"`
	InObjects    string `xml:"inObj,attr"`
	RelationType string `xml:"relationType,attr"`
	Script       string `xml:",innerxml"`
}

type Parameter struct {
	ID           string `xml:"id,attr"`
	ShortName    string `xml:"shortName,attr"`
	Description  string `xml:"description,attr"`
	Type         string `xml:"type,attr"`
	DefaultValue string `xml:"defaultValue,attr"`
}
