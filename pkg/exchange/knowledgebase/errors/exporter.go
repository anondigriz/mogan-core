package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	TypeIsNotSupportedByFormat = "TypeIsNotSupportedByFormat"
)

func NewTypeIsNotSupportedByFormatErr(t string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    TypeIsNotSupportedByFormat,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.TypeIsNotSupportedByFormat, t),
		Err:     nil,
		Dt:      map[string]string{},
	}
}
