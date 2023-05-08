package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	TypeIsNotSupportedByFormat = "TypeIsNotSupportedByFormat"
	NonUniqueIDFound           = "NonUniqueIDFound"
	PatternNotFoundForRule     = "PatternNotFoundForRule"
	ParameterNotFoundForRule   = "ParameterNotFoundForRule"
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

func NewPatternNotFoundForRuleErr(ruleUUID string, ruleID string, patternUUID string) error {
	return ExchangeKnowledgeBaseErr{
		Stat: PatternNotFoundForRule,
		Message: fmt.Sprintf("%s. Rule with UUID '%s' and ID '%s', specified pattern UUID in the body of the rule that was not found: '%s'",
			errMsgs.PatternNotFoundForRule, ruleID, ruleUUID, patternUUID),
		Err: nil,
		Dt:  map[string]string{},
	}
}

func NewParameterNotFoundForRuleErr(ruleUUID string, ruleID string, parameterUUID string) error {
	return ExchangeKnowledgeBaseErr{
		Stat: ParameterNotFoundForRule,
		Message: fmt.Sprintf("%s. Rule with UUID '%s' and ID '%s', specified parameter UUID in the body of the rule that was not found: '%s'",
			errMsgs.ParameterNotFoundForRule, ruleID, ruleUUID, parameterUUID),
		Err: nil,
		Dt:  map[string]string{},
	}
}
