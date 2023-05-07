package importer

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (im Importer) processRules(rules []formatV3M0.Rule, ws workspaceHandler) error {
	for _, v := range rules {
		rule, err := im.extractRule(v, ws)
		if err != nil {
			im.lg.Error(errMsgs.ParsingRulesFromXMLFail, zap.Error(err))
			return err
		}
		ws.AddRule(rule)
	}

	return nil
}

func (im Importer) extractRule(rule formatV3M0.Rule, ws workspaceHandler) (kbEnt.Rule, error) {
	r := kbEnt.Rule{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ws.CreateRuleUUID(),
			ID:           rule.ID,
			ShortName:    rule.ShortName,
			Description:  rule.Description,
			CreatedDate:  time.Unix(rule.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(rule.ModifiedDate, 0).UTC(),
		},
	}

	patternUUID, ok := ws.GetPatternUUID(rule.PatternID)
	if !ok {
		err := errors.NewRelationNotFoundForRuleErr(rule.ID, rule.PatternID)
		im.lg.Error(errMsgs.RelationNotFoundForRule, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	r.PatternUUID = patternUUID

	r.InputParameters = im.extractRuleParameters(rule.InputParameters.InputParameters, ws)
	r.OutputParameters = im.extractRuleParameters(rule.OutputParameters.OutputParameters, ws)

	return r, nil
}

func (im Importer) extractRuleParameters(parameters []formatV3M0.ParameterRule, ws workspaceHandler) []kbEnt.ParameterRule {
	var ps []kbEnt.ParameterRule
	for _, v := range parameters {
		ps = append(ps, kbEnt.ParameterRule{
			ShortName:     v.ShortName,
			ParameterUUID: ws.GetOrCreateParameterUUID(v.ParameterID),
		})
	}
	return ps
}
