package fromformat

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/entities/types"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ff *FromFormat) processRelations(relations []formatV2M0.Relation) error {
	for _, v := range relations {
		pattern, err := ff.mapToPattern(v)
		if err != nil {
			ff.lg.Error(errMsgs.MappingPatternFail, zap.Error(err))
			return err
		}
		ff.ws.AddPattern(pattern)
	}
	return nil
}

func (ff *FromFormat) mapToPattern(relation formatV2M0.Relation) (kbEnt.Pattern, error) {
	now := time.Now()
	p := kbEnt.Pattern{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ff.ws.CreatePatternUUID(relation.ID),
			ID:           relation.ID,
			ShortName:    relation.ShortName,
			Description:  relation.Description,
			CreatedDate:  now,
			ModifiedDate: now,
		},
		ScriptLanguage: types.JS,
		Script:         relation.Script,
	}

	patternType, err := ff.mapToPatternType(relation.RelationType)
	if err != nil {
		ff.lg.Error(errMsgs.MappingPatternTypeFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.Type = patternType

	inputParameters, err := ff.mapToPatternParameters(relation.InObjects)
	if err != nil {
		ff.lg.Error(errMsgs.MappingRelationParametersFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.InputParameters = inputParameters

	outputParameters, err := ff.mapToPatternParameters(relation.OutObjects)
	if err != nil {
		ff.lg.Error(errMsgs.MappingRelationParametersFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	p.OutputParameters = outputParameters

	if relation.RelationType == string(formatV2M0.Simple) || relation.RelationType == string(formatV2M0.Ifclause) {
		p.Script = ff.convertScriptToProgramType(p.Script, inputParameters, outputParameters)
	}

	return p, nil
}

func (ff *FromFormat) convertScriptToProgramType(script string, input []kbEnt.ParameterPattern, output []kbEnt.ParameterPattern) string {
	vars := make([]string, 0, len(input)+len(output))
	for _, v := range input {
		vars = append(vars, v.ShortName)
	}
	for _, v := range output {
		vars = append(vars, v.ShortName)
	}
	result := fmt.Sprintf("var %s;\n%s", strings.Join(vars, ", "), script)
	return result
}

func (ff *FromFormat) mapToPatternParameters(attribute string) ([]kbEnt.ParameterPattern, error) {
	var parameters []kbEnt.ParameterPattern
	dict, err := ff.mapToDictionary(attribute)
	if err != nil {
		return []kbEnt.ParameterPattern{}, err
	}

	for k, v := range dict {
		parameterType, err := ff.mapToParameterType(v)
		if err != nil {
			ff.lg.Error(errMsgs.MappingParameterTypeFail, zap.Error(err))
			return []kbEnt.ParameterPattern{}, err
		}

		parameters = append(parameters, kbEnt.ParameterPattern{
			ShortName: k,
			Type:      parameterType,
		})
	}
	return parameters, nil
}
