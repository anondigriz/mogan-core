package errors

import (
	"fmt"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	MappingFail = "MappingFail"
)

func NewRelationNotFoundForRuleErr(ruleID string, relationID string) error {
	return ExchangeKnowledgeBaseErr{
		Stat: MappingFail,
		Message: fmt.Sprintf("%s. Rule ID: '%s', specified relation ID in the body of the rule that was not found: '%s'",
			errMsgs.RelationNotFoundForRule, relationID, ruleID),
		Err: nil,
		Dt:  map[string]string{},
	}
}
