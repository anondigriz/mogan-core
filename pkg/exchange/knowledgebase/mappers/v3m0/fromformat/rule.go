package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (ff *FromFormat) processRules(rules []formatV3M0.Rule) error {
	for _, v := range rules {
		rule, err := ff.mapToRule(v)
		if err != nil {
			ff.lg.Error(errMsgs.MappingRulesFail, zap.Error(err))
			return err
		}
		ff.ws.AddRule(rule)
	}

	return nil
}

func (ff *FromFormat) mapToRule(rule formatV3M0.Rule) (kbEnt.Rule, error) {
	r := kbEnt.Rule{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ff.ws.CreateRuleUUID(rule.ID),
			ID:           rule.ID,
			ShortName:    rule.ShortName,
			Description:  rule.Description.Text,
			CreatedDate:  time.Unix(rule.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(rule.ModifiedDate, 0).UTC(),
		},
	}

	patternUUID, ok := ff.ws.GetPatternUUID(rule.PatternID)
	if !ok {
		err := errors.NewRelationNotFoundForRuleErr(rule.ID, rule.PatternID)
		ff.lg.Error(errMsgs.RelationNotFoundForRule, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	r.PatternUUID = patternUUID

	r.InputParameters = ff.mapToRuleParameters(rule.InputParameters.InputParameters)
	r.OutputParameters = ff.mapToRuleParameters(rule.OutputParameters.OutputParameters)

	return r, nil
}

func (ff *FromFormat) mapToRuleParameters(parameters []formatV3M0.ParameterRule) []kbEnt.ParameterRule {
	var ps []kbEnt.ParameterRule
	for _, v := range parameters {
		ps = append(ps, kbEnt.ParameterRule{
			ShortName:     v.ShortName,
			ParameterUUID: ff.ws.GetOrCreateParameterUUID(v.ParameterID),
		})
	}
	return ps
}
