package errors

import (
	"strconv"
)

const (
	OpeningXSDFail    = "OpeningXSDFail"
	ReadingXSDFail    = "ReadingXSDFail"
	ParsingXSDFail    = "ParsingXSDFail"
	ValidationXMLFail = "ValidationXMLFail"
)

func NewOpeningXSDFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    OpeningXSDFail,
		Message: "error opening XSD file",
		Err:     e,
		Dt:      map[string]string{},
	}
}
func NewReadingXSDFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ReadingXSDFail,
		Message: "error reading XSD file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewParsingXSDFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ParsingXSDFail,
		Message: "error parsing XSD file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewValidationXMLFailErr(baseError error, errors []error) error {
	e := ExchangeKnowledgeBaseErr{
		Stat:    ValidationXMLFail,
		Message: "xml document validation error",
		Err:     baseError,
		Dt:      map[string]string{},
	}

	for i, err := range errors {
		e.Dt[strconv.Itoa(i)] = err.Error()
	}
	return e
}
