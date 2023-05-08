package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/entities/types"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ff FromFormat) processRelations(relations []formatV2M0.Relation, ws workspaceHandler) error {
	for _, v := range relations {
		pattern, err := ff.extractPattern(v, ws)
		if err != nil {
			ff.lg.Error(errMsgs.ParsingPatternFromXMLFail, zap.Error(err))
			return err
		}
		ws.AddPattern(pattern)
	}
	return nil
}

func (ff FromFormat) extractPattern(relation formatV2M0.Relation, ws workspaceHandler) (kbEnt.Pattern, error) {
	now := time.Now()
	p := kbEnt.Pattern{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ws.CreatePatternUUID(relation.ID),
			ID:           relation.ID,
			ShortName:    relation.ShortName,
			Description:  relation.Description,
			CreatedDate:  now,
			ModifiedDate: now,
		},
		ScriptLanguage: types.JS,
		Script:         relation.Script,
	}

	patternType, err := ff.extractPatternType(relation.RelationType)
	if err != nil {
		ff.lg.Error(errMsgs.ExtractPatternTypeFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.Type = patternType

	inputParameters, err := ff.extractPatternParameters(relation.InObjects)
	if err != nil {
		ff.lg.Error(errMsgs.ParsingRelationParametersFromXMLFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.InputParameters = inputParameters

	outputParameters, err := ff.extractPatternParameters(relation.OutObjects)
	if err != nil {
		ff.lg.Error(errMsgs.ParsingRelationParametersFromXMLFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.OutputParameters = outputParameters

	return p, nil

}

func (ff FromFormat) extractPatternParameters(attribute string) ([]kbEnt.ParameterPattern, error) {
	var parameters []kbEnt.ParameterPattern
	dict, err := ff.extractDictionaryFromAttribute(attribute)
	if err != nil {
		return []kbEnt.ParameterPattern{}, err
	}

	for k, v := range dict {
		parameterType, err := ff.extractParameterType(v)
		if err != nil {
			ff.lg.Error(errMsgs.ExtractParameterTypeFail, zap.Error(err))
			return []kbEnt.ParameterPattern{}, err
		}

		parameters = append(parameters, kbEnt.ParameterPattern{
			ShortName: k,
			Type:      parameterType,
		})
	}
	return parameters, nil
}
