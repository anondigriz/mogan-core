package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	TypeIsNotSupportedByFormat = "TypeIsNotSupportedByFormat"
	NonUniqueIDFound           = "NonUniqueIDFound"
)

func NewTypeIsNotSupportedByFormatErr(t string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    TypeIsNotSupportedByFormat,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.TypeIsNotSupportedByFormat, t),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

type NonUniqueIDFoundArgs struct {
	ID              string
	FirstUUID       string
	FirstShortName  string
	SecondUUID      string
	SecondShortName string
}

func NewNonUniqueIDFoundErr(args NonUniqueIDFoundArgs) error {
	return ExchangeKnowledgeBaseErr{
		Stat: NonUniqueIDFound,
		Message: fmt.Sprintf("%s. ID: '%s'. First object: UUID: '%s', short name: '%s'. Second object: UUID: '%s', short name: '%s'",
			errMsgs.NonUniqueIDFound,
			args.ID, args.FirstUUID, args.FirstShortName, args.SecondUUID, args.SecondShortName),
		Err: nil,
		Dt:  map[string]string{},
	}
}
