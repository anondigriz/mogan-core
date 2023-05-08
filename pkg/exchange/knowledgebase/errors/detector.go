package errors

import (
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	FormatXMLVersionNotFound = "FormatXMLVersionNotFound"
)

func NewFormatXMLVersionNotFoundErr() error {
	return ExchangeKnowledgeBaseErr{
		Stat:    FormatXMLVersionNotFound,
		Message: errMsgs.FormatXMLVersionNotFound,
		Err:     nil,
		Dt:      map[string]string{},
	}
}
