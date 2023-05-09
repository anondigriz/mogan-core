package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ff *FromFormat) processRules(rules []formatV2M0.Rule, parentGroup *kbEnt.Group) error {
	for _, v := range rules {
		rule, err := ff.mapToRule(v)
		if err != nil {
			ff.lg.Error(errMsgs.MappingRulesFail, zap.Error(err))
			return err
		}
		parentGroup.Rules = append(parentGroup.Rules, rule.UUID)
		ff.ws.AddRule(rule)
	}

	return nil
}

func (ff *FromFormat) mapToRule(rule formatV2M0.Rule) (kbEnt.Rule, error) {
	now := time.Now()
	r := kbEnt.Rule{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ff.ws.CreateRuleUUID(),
			ID:           rule.ID,
			ShortName:    rule.ShortName,
			Description:  rule.Description,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}

	patternUUID, ok := ff.ws.GetPatternUUID(rule.RelationID)
	if !ok {
		err := errors.NewRelationNotFoundForRuleErr(rule.ID, rule.RelationID)
		ff.lg.Error(errMsgs.RelationNotFoundForRule, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	r.PatternUUID = patternUUID

	inputParameters, err := ff.mapToRuleParameters(rule.InitIDs)
	if err != nil {
		ff.lg.Error(errMsgs.MappingRuleParametersFail, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	r.InputParameters = inputParameters

	outputParameters, err := ff.mapToRuleParameters(rule.ResultIDs)
	if err != nil {
		ff.lg.Error(errMsgs.MappingRuleParametersFail, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	r.OutputParameters = outputParameters

	return r, nil
}

func (ff *FromFormat) mapToRuleParameters(attribute string) ([]kbEnt.ParameterRule, error) {
	var parameters []kbEnt.ParameterRule
	dict, err := ff.mapToDictionary(attribute)
	if err != nil {
		return []kbEnt.ParameterRule{}, err
	}

	for k, v := range dict {
		parameters = append(parameters, kbEnt.ParameterRule{
			ShortName:     k,
			ParameterUUID: ff.ws.GetOrCreateParameterUUID(v),
		})
	}
	return parameters, nil
}
