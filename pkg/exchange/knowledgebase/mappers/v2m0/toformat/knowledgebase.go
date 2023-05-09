package toformat

import (
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (tf ToFormat) processContainer(model *formatV2M0.Model, cont kbEnt.Container, ws workspaceHandler) error {
	tf.prepareModel(model, cont, ws)
	tf.warmUpWorkspace(cont, ws)

	err := tf.processPatterns(processPatternsArgs{
		patterns: cont.Patterns,
		model:    model,
		ws:       ws,
	})
	if err != nil {
		tf.lg.Error(errMsgs.MappingPatternFail, zap.Error(err))
		return err
	}

	err = tf.processGroups(processGroupsArgs{
		cont:        cont,
		childGroups: cont.Groups,
		parentClass: &model.Class,
		ws:          ws,
	})

	if err != nil {
		tf.lg.Error(errMsgs.MappingGroupsFail, zap.Error(err))
		return err
	}

	return tf.processRootClass(cont, model, ws)
}

func (tf ToFormat) prepareModel(model *formatV2M0.Model, cont kbEnt.Container, ws workspaceHandler) {
	model.BaseInfo = formatV2M0.BaseInfo{
		ID:          cont.KnowledgeBase.BaseInfo.ID,
		ShortName:   cont.KnowledgeBase.BaseInfo.ShortName,
		Description: cont.KnowledgeBase.BaseInfo.Description,
	}
	model.FormatXMLVersion = string(formats.VersionV2M0)
	model.Relations.Relations = []formatV2M0.Relation{}

	model.Class = formatV2M0.Class{
		BaseInfo: formatV2M0.BaseInfo{
			ID:          uuidGen.NewString(),
			ShortName:   cont.KnowledgeBase.ShortName,
			Description: cont.KnowledgeBase.Description,
		},
	}
}

func (ToFormat) warmUpWorkspace(cont kbEnt.Container, ws workspaceHandler) {
	unprocessedParameters := []string{}
	for k := range cont.Parameters {
		unprocessedParameters = append(unprocessedParameters, k)
	}
	ws.SaveUnprocessedParameters(unprocessedParameters)

	unprocessedRules := []string{}
	for k := range cont.Rules {
		unprocessedRules = append(unprocessedRules, k)
	}
	ws.SaveUnprocessedRules(unprocessedRules)
}

func (tf ToFormat) processRootClass(cont kbEnt.Container, model *formatV2M0.Model, ws workspaceHandler) error {
	rootGroup := kbEnt.Group{}
	rootGroup.Parameters = append(rootGroup.Parameters, ws.GetUnprocessedParameters()...)
	rootGroup.Rules = append(rootGroup.Rules, ws.GetUnprocessedRules()...)

	err := tf.processRules(processRulesArgs{
		cont:        cont,
		parentGroup: rootGroup,
		parentClass: &model.Class,
		ws:          ws,
	})
	if err != nil {
		tf.lg.Error(errMsgs.MappingRulesFail, zap.Error(err))
		return err
	}

	err = tf.processParameters(processParametersArgs{
		cont:        cont,
		parentGroup: rootGroup,
		parentClass: &model.Class,
		ws:          ws,
	})
	if err != nil {
		tf.lg.Error(errMsgs.MappingParametersFail, zap.Error(err))
		return err
	}

	return nil
}
