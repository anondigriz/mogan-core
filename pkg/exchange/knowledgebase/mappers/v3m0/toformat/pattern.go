package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (tf *ToFormat) mapToPatterns(patterns map[string]kbEnt.Pattern) ([]formatV3M0.Pattern, error) {
	ps := []formatV3M0.Pattern{}
	for _, v := range patterns {
		p, err := tf.mapToPattern(v)
		if err != nil {
			tf.lg.Error(errMsgs.MappingPatternFail, zap.Error(err))
			return []formatV3M0.Pattern{}, err
		}
		ps = append(ps, p)
	}

	return ps, nil
}

func (tf *ToFormat) mapToPattern(pattern kbEnt.Pattern) (formatV3M0.Pattern, error) {
	if err := tf.ws.CheckAndRememberPattern(pattern); err != nil {
		tf.lg.Error(errMsgs.MappingPatternFail, zap.Error(err))
		return formatV3M0.Pattern{}, err
	}

	p := formatV3M0.Pattern{
		BaseInfo: formatV3M0.BaseInfo{
			ID:           pattern.BaseInfo.ID,
			ShortName:    pattern.BaseInfo.ShortName,
			Description:  formatV3M0.Description{Text: pattern.BaseInfo.Description},
			CreatedDate:  pattern.CreatedDate.UTC().Unix(),
			ModifiedDate: pattern.ModifiedDate.UTC().Unix(),
		},
		Script: formatV3M0.Script{Text: pattern.Script},
	}

	scriptLanguageType, err := tf.mapToScriptLanguageType(pattern.ScriptLanguage)
	if err != nil {
		tf.lg.Error(errMsgs.TypeIsNotSupportedByFormat, zap.Error(err))
		return formatV3M0.Pattern{}, err
	}
	p.Language = scriptLanguageType

	patternType, err := tf.mapToPatternType(pattern.Type)
	if err != nil {
		tf.lg.Error(errMsgs.MappingPatternTypeFail, zap.Error(err))
		return formatV3M0.Pattern{}, err
	}
	p.Type = patternType

	inputParameters, err := tf.mapToPatternParameters(pattern.InputParameters)
	if err != nil {
		tf.lg.Error(errMsgs.MappingPatternParametersFail, zap.Error(err))
		return formatV3M0.Pattern{}, err
	}
	p.InputParameters.InputParameters = inputParameters

	outputParameters, err := tf.mapToPatternParameters(pattern.OutputParameters)
	if err != nil {
		tf.lg.Error(errMsgs.MappingPatternParametersFail, zap.Error(err))
		return formatV3M0.Pattern{}, err
	}
	p.OutputParameters.OutputParameters = outputParameters

	return p, nil
}

func (tf *ToFormat) mapToPatternParameters(parameters []kbEnt.ParameterPattern) ([]formatV3M0.ParameterPattern, error) {
	ps := []formatV3M0.ParameterPattern{}
	for _, v := range parameters {
		parameterType, err := tf.mapToParameterType(v.Type)
		if err != nil {
			tf.lg.Error(errMsgs.MappingParameterTypeFail, zap.Error(err))
			return []formatV3M0.ParameterPattern{}, err
		}
		ps = append(ps, formatV3M0.ParameterPattern{
			ShortName: v.ShortName,
			Type:      parameterType,
		})
	}
	return ps, nil
}
