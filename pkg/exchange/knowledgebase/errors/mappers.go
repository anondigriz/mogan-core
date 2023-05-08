package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

func NewUnknownPatternTypeErr(t string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    MappingFail,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.UnknownPatternType, t),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewUnknownParameterTypeErr(t string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    MappingFail,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.UnknownParameterType, t),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewUnknownScriptLanguageTypeErr(t string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    MappingFail,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.UnknownScriptLanguageType, t),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewBrokenDictionaryInStringErr(value string, keyValue string) error {
	return ExchangeKnowledgeBaseErr{
		Stat: MappingFail,
		Message: fmt.Sprintf("%s. String: '%s'. Broken key-value pair: '%s'",
			errMsgs.BrokenDictionaryInString, value, keyValue),
		Err: nil,
		Dt:  map[string]string{},
	}
}
