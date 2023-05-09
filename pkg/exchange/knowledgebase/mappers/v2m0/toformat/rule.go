package toformat

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type processRulesArgs struct {
	parentGroup kbEnt.Group
	parentClass *formatV2M0.Class
}

func (tf *ToFormat) processRules(args processRulesArgs) error {
	for _, v := range args.parentGroup.Rules {
		rule, ok := tf.cont.Rules[v]
		if !ok {
			continue
		}

		pattern, ok := tf.cont.Patterns[rule.PatternUUID]
		if !ok {
			err := errors.NewPatternNotFoundForRuleErr(rule.UUID, rule.ID, rule.PatternUUID)
			tf.lg.Error(errMsgs.PatternNotFoundForRule, zap.Error(err))
			return err
		}

		r, err := tf.mapToRule(rule, pattern)
		if err != nil {
			tf.lg.Error(errMsgs.MappingRelationFail, zap.Error(err))
			return err
		}

		if pattern.Type == types.Constraint {
			args.parentClass.Constraints.Constraints = append(args.parentClass.Constraints.Constraints, r)
		} else {
			args.parentClass.Rules.Rules = append(args.parentClass.Rules.Rules, r)
		}
	}
	return nil
}

func (tf *ToFormat) mapToRule(rule kbEnt.Rule, pattern kbEnt.Pattern) (formatV2M0.Rule, error) {
	if err := tf.ws.CheckAndRememberRule(rule); err != nil {
		tf.lg.Error(errMsgs.MappingRuleFail, zap.Error(err))
		return formatV2M0.Rule{}, err
	}

	r := formatV2M0.Rule{
		BaseInfo: formatV2M0.BaseInfo{
			ID:          rule.ID,
			ShortName:   rule.ShortName,
			Description: rule.Description,
		},
		RelationID: pattern.ID,
	}

	initIDs, err := tf.mapToRuleParameters(rule.InputParameters, rule)
	if err != nil {
		tf.lg.Error(errMsgs.MappingRuleParametersFail, zap.Error(err))
		return formatV2M0.Rule{}, err
	}
	r.InitIDs = initIDs

	resultIDs, err := tf.mapToRuleParameters(rule.OutputParameters, rule)
	if err != nil {
		tf.lg.Error(errMsgs.MappingRuleParametersFail, zap.Error(err))
		return formatV2M0.Rule{}, err
	}
	r.ResultIDs = resultIDs

	return r, nil
}

func (tf *ToFormat) mapToRuleParameters(parameters []kbEnt.ParameterRule, rule kbEnt.Rule) (string, error) {
	keys := make([]string, 0, len(parameters))
	for _, v := range parameters {
		parameter, ok := tf.cont.Parameters[v.ParameterUUID]
		if !ok {
			err := errors.NewParameterNotFoundForRuleErr(rule.UUID, rule.ID, rule.PatternUUID)
			tf.lg.Error(errMsgs.ParameterNotFoundForRule, zap.Error(err))
			return "", err
		}

		keys = append(keys, fmt.Sprintf("%s:%s", v.ShortName, parameter.ID))
	}
	return strings.Join(keys, ";"), nil
}
