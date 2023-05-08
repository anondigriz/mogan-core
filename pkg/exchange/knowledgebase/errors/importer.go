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

func NewRelationNotFoundForRuleErr(ruleID string, relationID string) error {
	return ExchangeKnowledgeBaseErr{
		Stat: ParsingXMLFail,
		Message: fmt.Sprintf("%s. Rule ID: '%s', specified relation ID in the body of the rule that was not found: '%s'",
			errMsgs.RelationNotFoundForRule, relationID, ruleID),
		Err: nil,
		Dt:  map[string]string{},
	}
}

func NewUnknownPatternTypeInXMLErr(t string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ParsingXMLFail,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.UnknownPatternTypeInXML, t),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewUnknownParameterTypeInXMLErr(t string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ParsingXMLFail,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.UnknownParameterTypeInXML, t),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewUnknownScriptLanguageTypeInXMLErr(t string) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ParsingXMLFail,
		Message: fmt.Sprintf("%s. Type: '%s'", errMsgs.UnknownScriptLanguageTypeInXML, t),
		Err:     nil,
		Dt:      map[string]string{},
	}
}

func NewBrokenDictionaryInXMLAttributeErr(value string, keyValue string) error {
	return ExchangeKnowledgeBaseErr{
		Stat: ParsingXMLFail,
		Message: fmt.Sprintf("%s. Attribute value: '%s'. Broken key-value pair: '%s'",
			errMsgs.BrokenDictionaryInXMLAttribute, value, keyValue),
		Err: nil,
		Dt:  map[string]string{},
	}
}
