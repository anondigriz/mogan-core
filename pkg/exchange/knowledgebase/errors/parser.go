package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	ParsingXMLFail              = "ParsingXMLFail"
	FormatXMLVersionNotFound    = "FormatXMLVersionNotFound"
	UnsupportedFormatXMLVersion = "UnsupportedFormatXMLVersion"
	UnexpectedJobExecutionFail  = "UnexpectedJobExecutionFail"
)

func NewParsingXMLFailErr(msg string, e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ParsingXMLFail,
		Message: msg,
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewFormatXMLVersionNotFoundErr() error {
	return ExchangeKnowledgeBaseErr{
		Stat:    FormatXMLVersionNotFound,
		Message: errMsgs.FormatXMLVersionNotFound,
		Err:     nil,
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

func NewUnexpectedJobExecutionFailErr(err error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    UnexpectedJobExecutionFail,
		Message: errMsgs.UnexpectedJobExecution,
		Err:     nil,
		Dt:      map[string]string{},
	}
}
