package importer

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (vm Importer) processPatterns(patterns []formatV3M0.Pattern, ws workspaceHandler) error {
	for _, v := range patterns {
		pattern, err := vm.extractPattern(v, ws)
		if err != nil {
			vm.lg.Error(errMsgs.ParsingPatternFromXMLFail, zap.Error(err))
			return err
		}
		ws.AddPattern(pattern)
	}
	return nil
}

func (im Importer) extractPattern(pattern formatV3M0.Pattern, ws workspaceHandler) (kbEnt.Pattern, error) {
	p := kbEnt.Pattern{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ws.CreatePatternUUID(pattern.ID),
			ID:           pattern.ID,
			ShortName:    pattern.ShortName,
			Description:  pattern.Description,
			CreatedDate:  time.Unix(pattern.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(pattern.ModifiedDate, 0).UTC(),
		},
		Script: pattern.Script,
	}

	patternType, err := im.extractPatternType(pattern.Type)
	if err != nil {
		im.lg.Error(errMsgs.ExtractPatternTypeFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.Type = patternType

	scriptLanguageType, err := im.extractScriptLanguageType(pattern.Type)
	if err != nil {
		im.lg.Error(errMsgs.ExtractScriptLanguageTypeInXML, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.ScriptLanguage = scriptLanguageType

	inputParameters, err := im.extractPatternParameters(pattern.InputParameters.InputParameters)
	if err != nil {
		im.lg.Error(errMsgs.ParsingRelationParametersFromXMLFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.InputParameters = inputParameters

	outputParameters, err := im.extractPatternParameters(pattern.OutputParameters.OutputParameters)
	if err != nil {
		im.lg.Error(errMsgs.ParsingRelationParametersFromXMLFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.OutputParameters = outputParameters

	return p, nil

}

func (im Importer) extractPatternParameters(parameters []formatV3M0.ParameterPattern) ([]kbEnt.ParameterPattern, error) {
	var ps []kbEnt.ParameterPattern
	for _, v := range parameters {
		parameterType, err := im.extractParameterType(v.Type)
		if err != nil {
			im.lg.Error(errMsgs.ExtractParameterTypeFail, zap.Error(err))
			return []kbEnt.ParameterPattern{}, err
		}
		ps = append(ps, kbEnt.ParameterPattern{
			ShortName: v.ShortName,
			Type:      parameterType,
		})
	}
	return ps, nil
}
