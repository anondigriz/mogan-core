package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (ff *FromFormat) processPatterns(patterns []formatV3M0.Pattern) error {
	for _, v := range patterns {
		pattern, err := ff.mapToPattern(v)
		if err != nil {
			ff.lg.Error(errMsgs.MappingPatternFail, zap.Error(err))
			return err
		}
		ff.ws.AddPattern(pattern)
	}
	return nil
}

func (ff *FromFormat) mapToPattern(pattern formatV3M0.Pattern) (kbEnt.Pattern, error) {
	p := kbEnt.Pattern{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ff.ws.CreatePatternUUID(pattern.ID),
			ID:           pattern.ID,
			ShortName:    pattern.ShortName,
			Description:  pattern.Description.Text,
			CreatedDate:  time.Unix(pattern.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(pattern.ModifiedDate, 0).UTC(),
		},
		Script: pattern.Script.Text,
	}

	patternType, err := ff.mapToPatternType(pattern.Type)
	if err != nil {
		ff.lg.Error(errMsgs.MappingPatternTypeFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.Type = patternType

	scriptLanguageType, err := ff.mapToScriptLanguageType(pattern.Language)
	if err != nil {
		ff.lg.Error(errMsgs.MappingScriptLanguageTypeFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.ScriptLanguage = scriptLanguageType

	inputParameters, err := ff.mapToPatternParameters(pattern.InputParameters.InputParameters)
	if err != nil {
		ff.lg.Error(errMsgs.MappingRelationParametersFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.InputParameters = inputParameters

	outputParameters, err := ff.mapToPatternParameters(pattern.OutputParameters.OutputParameters)
	if err != nil {
		ff.lg.Error(errMsgs.MappingRelationParametersFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.OutputParameters = outputParameters

	return p, nil

}

func (ff *FromFormat) mapToPatternParameters(parameters []formatV3M0.ParameterPattern) ([]kbEnt.ParameterPattern, error) {
	var ps []kbEnt.ParameterPattern
	for _, v := range parameters {
		parameterType, err := ff.mapToParameterType(v.Type)
		if err != nil {
			ff.lg.Error(errMsgs.MappingParameterTypeFail, zap.Error(err))
			return []kbEnt.ParameterPattern{}, err
		}
		ps = append(ps, kbEnt.ParameterPattern{
			ShortName: v.ShortName,
			Type:      parameterType,
		})
	}
	return ps, nil
}
