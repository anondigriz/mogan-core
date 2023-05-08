package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ff FromFormat) processRules(rules []formatV2M0.Rule, parentGroup *kbEnt.Group, ws workspaceHandler) error {
	for _, v := range rules {
		rule, err := ff.extractRule(v, ws)
		if err != nil {
			ff.lg.Error(errMsgs.ParsingRulesFromXMLFail, zap.Error(err))
			return err
		}
		parentGroup.Rules = append(parentGroup.Rules, rule.UUID)
		ws.AddRule(rule)
	}

	return nil
}

func (ff FromFormat) extractRule(rule formatV2M0.Rule, ws workspaceHandler) (kbEnt.Rule, error) {
	now := time.Now()
	r := kbEnt.Rule{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ws.CreateRuleUUID(),
			ID:           rule.ID,
			ShortName:    rule.ShortName,
			Description:  rule.Description,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}

	patternUUID, ok := ws.GetPatternUUID(rule.RelationID)
	if !ok {
		err := errors.NewRelationNotFoundForRuleErr(rule.ID, rule.RelationID)
		ff.lg.Error(errMsgs.RelationNotFoundForRule, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	r.PatternUUID = patternUUID

	inputParameters, err := ff.extractRuleParameters(rule.InitIDs, ws)
	if err != nil {
		ff.lg.Error(errMsgs.ParsingRuleParametersFromXMLFail, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	r.InputParameters = inputParameters

	outputParameters, err := ff.extractRuleParameters(rule.ResultIDs, ws)
	if err != nil {
		ff.lg.Error(errMsgs.ParsingRuleParametersFromXMLFail, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	r.OutputParameters = outputParameters

	return r, nil
}

func (ff FromFormat) extractRuleParameters(attribute string, ws workspaceHandler) ([]kbEnt.ParameterRule, error) {
	var parameters []kbEnt.ParameterRule
	dict, err := ff.extractDictionaryFromAttribute(attribute)
	if err != nil {
		return []kbEnt.ParameterRule{}, err
	}

	for k, v := range dict {
		parameters = append(parameters, kbEnt.ParameterRule{
			ShortName:     k,
			ParameterUUID: ws.GetOrCreateParameterUUID(v),
		})
	}
	return parameters, nil
}