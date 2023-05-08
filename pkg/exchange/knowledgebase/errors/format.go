package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	UnsupportedFormatXMLVersion = "UnsupportedFormatXMLVersion"
)

func NewUnsupportedFormatXMLVersionErr(version string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    UnsupportedFormatXMLVersion,
		Message: fmt.Sprintf("%s. Unsupported version: '%s'", errMsgs.UnsupportedFormatXMLVersion, version),
		Err:     nil,
		Dt:      map[string]string{},
	}
}
