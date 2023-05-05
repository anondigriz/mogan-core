package errors

import (
	"strconv"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
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
		Message: errMsgs.OpeningXSDFail,
		Err:     e,
		Dt:      map[string]string{},
	}
}
func NewReadingXSDFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ReadingXSDFail,
		Message: errMsgs.ReadingXSDFail,
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewParsingXSDFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ParsingXSDFail,
		Message: errMsgs.ParsingXSDFail,
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewValidationXMLFailErr(baseError error, errors []error) error {
	e := ExchangeKnowledgeBaseErr{
		Stat:    ValidationXMLFail,
		Message: errMsgs.ValidationXMLFail,
		Err:     baseError,
		Dt:      map[string]string{},
	}

	for i, err := range errors {
		e.Dt[strconv.Itoa(i)] = err.Error()
	}
	return e
}
