package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	ParsingXMLFail              = "ParsingXMLFail"
	UnsupportedFormatXMLVersion = "UnsupportedFormatXMLVersion"
)

func NewParsingXMLFailErr(msg string, e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ParsingXMLFail,
		Message: msg,
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewUnsupportedFormatXMLVersionErr(version string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    UnsupportedFormatXMLVersion,
		Message: fmt.Sprintf("%s. Unsupported version: '%s'", errMsgs.UnsupportedFormatXMLVersion, version),
		Err:     nil,
		Dt:      map[string]string{},
	}
}
