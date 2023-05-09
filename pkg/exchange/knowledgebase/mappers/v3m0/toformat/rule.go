package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (tf ToFormat) mapToRules(cont kbEnt.Container, ws workspaceHandler) ([]formatV3M0.Rule, error) {
	ps := []formatV3M0.Rule{}
	for _, v := range cont.Rules {
		p, err := tf.mapToRule(v, cont, ws)
		if err != nil {
			tf.lg.Error(errMsgs.MappingRuleFail, zap.Error(err))
			return []formatV3M0.Rule{}, err
		}
		ps = append(ps, p)
	}

	return ps, nil
}

func (tf ToFormat) mapToRule(rule kbEnt.Rule, cont kbEnt.Container, ws workspaceHandler) (formatV3M0.Rule, error) {
	if err := ws.CheckAndRememberRule(rule); err != nil {
		tf.lg.Error(errMsgs.MappingRuleFail, zap.Error(err))
		return formatV3M0.Rule{}, err
	}
	r := formatV3M0.Rule{
		BaseInfo: formatV3M0.BaseInfo{
			ID:           rule.BaseInfo.ID,
			ShortName:    rule.BaseInfo.ShortName,
			Description:  rule.BaseInfo.Description,
			CreatedDate:  rule.CreatedDate.UTC().Unix(),
			ModifiedDate: rule.ModifiedDate.UTC().Unix(),
		},
	}

	pattern, ok := cont.Patterns[rule.PatternUUID]
	if !ok {
		err := errors.NewPatternNotFoundForRuleErr(rule.UUID, rule.ID, rule.PatternUUID)
		tf.lg.Error(errMsgs.PatternNotFoundForRule, zap.Error(err))
		return formatV3M0.Rule{}, err
	}
	r.PatternID = pattern.ID

	inputParameters, err := tf.mapToRuleParameters(rule.InputParameters, rule, cont)
	if err != nil {
		tf.lg.Error(errMsgs.MappingRuleParametersFail, zap.Error(err))
		return formatV3M0.Rule{}, err
	}
	r.InputParameters.InputParameters = inputParameters

	outputParameters, err := tf.mapToRuleParameters(rule.OutputParameters, rule, cont)
	if err != nil {
		tf.lg.Error(errMsgs.MappingRuleParametersFail, zap.Error(err))
		return formatV3M0.Rule{}, err
	}
	r.OutputParameters.OutputParameters = outputParameters

	return r, nil
}

func (tf ToFormat) mapToRuleParameters(parameters []kbEnt.ParameterRule, rule kbEnt.Rule, cont kbEnt.Container) ([]formatV3M0.ParameterRule, error) {
	ps := []formatV3M0.ParameterRule{}
	for _, v := range parameters {
		parameter, ok := cont.Parameters[v.ParameterUUID]
		if !ok {
			err := errors.NewParameterNotFoundForRuleErr(rule.UUID, rule.ID, rule.PatternUUID)
			tf.lg.Error(errMsgs.ParameterNotFoundForRule, zap.Error(err))
			return []formatV3M0.ParameterRule{}, err
		}
		ps = append(ps, formatV3M0.ParameterRule{
			ShortName:   v.ShortName,
			ParameterID: parameter.ID,
		})
	}
	return ps, nil
}
