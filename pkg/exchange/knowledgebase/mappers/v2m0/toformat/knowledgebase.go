package toformat

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ex ToFormat) processKnowledgeBase(knowledgeBase kbEnt.Container, model *formatV2M0.Model, ws workspaceHandler) error {
	model.BaseInfo = formatV2M0.BaseInfo{
		ID:          model.ID,
		ShortName:   model.ShortName,
		Description: model.Description,
	}

	unprocessedParameters := []string{}
	for k, _ := range knowledgeBase.Parameters {
		unprocessedParameters = append(unprocessedParameters, k)
	}
	ws.SaveUnprocessedParameters(unprocessedParameters)

	unprocessedRules := []string{}
	for k, _ := range knowledgeBase.Rules {
		unprocessedRules = append(unprocessedRules, k)
	}
	ws.SaveUnprocessedRules(unprocessedRules)

	rootGroup := formatV2M0.Class{}
	_ = rootGroup

	return nil
}
