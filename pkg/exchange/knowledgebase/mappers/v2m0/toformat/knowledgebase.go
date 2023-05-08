package toformat

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
	"go.uber.org/zap"
)

func (tf ToFormat) processKnowledgeBase(knowledgeBase kbEnt.Container, model *formatV2M0.Model, ws workspaceHandler) error {
	model.BaseInfo = formatV2M0.BaseInfo{
		ID:          model.ID,
		ShortName:   model.ShortName,
		Description: model.Description,
	}
	model.FormatXMLVersion = "2.0"
	model.Class = formatV2M0.Class{}
	model.Relations.Relations = []formatV2M0.Relation{}

	unprocessedParameters := []string{}
	for k := range knowledgeBase.Parameters {
		unprocessedParameters = append(unprocessedParameters, k)
	}
	ws.SaveUnprocessedParameters(unprocessedParameters)

	unprocessedRules := []string{}
	for k := range knowledgeBase.Rules {
		unprocessedRules = append(unprocessedRules, k)
	}
	ws.SaveUnprocessedRules(unprocessedRules)

	pArgs := processPatternsArgs{
		patterns: knowledgeBase.Patterns,
		model:    model,
		ws:       ws,
	}
	err := tf.processPatterns(pArgs)
	if err != nil {
		tf.lg.Error(errMsgs.MappingPatternFail, zap.Error(err))
		return err
	}

	rootGroup := kbEnt.Group{}
	rootGroup.Parameters = append(rootGroup.Parameters, ws.GetUnprocessedParameters()...)
	rootGroup.Rules = append(rootGroup.Rules, ws.GetUnprocessedRules()...)

	return nil
}
