package toformat

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type processPatternsArgs struct {
	patterns map[string]kbEnt.Pattern
	model    *formatV2M0.Model
	ws       workspaceHandler
}

func (tf ToFormat) processPatterns(args processPatternsArgs) error {
	for _, v := range args.patterns {
		r, err := tf.mapToRelation(v, args.ws)
		if err != nil {
			tf.lg.Error(errMsgs.MappingRelationFail, zap.Error(err))
			return err
		}
		args.model.Relations.Relations = append(args.model.Relations.Relations, r)
	}
	return nil
}

func (tf ToFormat) mapToRelation(pattern kbEnt.Pattern, ws workspaceHandler) (formatV2M0.Relation, error) {
	if err := ws.CheckAndRememberPattern(pattern); err != nil {
		tf.lg.Error(errMsgs.MappingPatternFail, zap.Error(err))
		return formatV2M0.Relation{}, err
	}

	if err := tf.isAllowedScriptLanguageType(pattern.ScriptLanguage); err != nil {
		tf.lg.Error(errMsgs.TypeIsNotSupportedByFormat, zap.Error(err))
		return formatV2M0.Relation{}, err
	}

	r := formatV2M0.Relation{
		BaseInfo: formatV2M0.BaseInfo{
			ID:          pattern.ID,
			ShortName:   pattern.ShortName,
			Description: pattern.Description,
		},
		Script: pattern.Script,
	}

	relationType, err := tf.mapToRelationType(pattern.Type)
	if err != nil {
		tf.lg.Error(errMsgs.MappingRelationTypeFail, zap.Error(err))
		return formatV2M0.Relation{}, err
	}
	r.RelationType = relationType

	inObjects, err := tf.mapToRelationParameters(pattern.InputParameters)
	if err != nil {
		tf.lg.Error(errMsgs.MappingRelationParametersFail, zap.Error(err))
		return formatV2M0.Relation{}, err
	}
	r.InObjects = inObjects

	outObjects, err := tf.mapToRelationParameters(pattern.OutputParameters)
	if err != nil {
		tf.lg.Error(errMsgs.MappingRelationParametersFail, zap.Error(err))
		return formatV2M0.Relation{}, err
	}
	r.OutObjects = outObjects

	return r, nil

}

func (tf ToFormat) mapToRelationParameters(parameters []kbEnt.ParameterPattern) (string, error) {
	keys := make([]string, 0, len(parameters))
	for _, v := range parameters {
		parameterType, err := tf.mapToParameterType(v.Type)
		if err != nil {
			tf.lg.Error(errMsgs.MappingParameterTypeFail, zap.Error(err))
			return "", err
		}
		keys = append(keys, fmt.Sprintf("%s:%s", v.ShortName, parameterType))
	}
	return strings.Join(keys, ";"), nil
}
