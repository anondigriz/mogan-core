package errors

const (
	ReadingXMLFail   = "ReadingXMLFail"
	XMLUnmarshalFail = "XMLUnmarshalFail"
)

func NewReadingXMLFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ReadingXMLFail,
		Message: "error reading XML file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewXMLUnmarshalFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    XMLUnmarshalFail,
		Message: "fail to unmarshal the xml file",
		Err:     e,
		Dt:      map[string]string{},
	}
}
