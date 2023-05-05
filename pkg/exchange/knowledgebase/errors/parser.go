package errors

import (
	"fmt"
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
		Message: "xml exchange document file version could not be detected",
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewUnsupportedFormatXMLVersionErr(version string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    UnsupportedFormatXMLVersion,
		Message: fmt.Sprintf("xml exchange document file version '%s' is not supported", version),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewUnexpectedJobExecutionFailErr(err error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    UnexpectedJobExecutionFail,
		Message: "unexpected job execution error occurred",
		Err:     nil,
		Dt:      map[string]string{},
	}
}
