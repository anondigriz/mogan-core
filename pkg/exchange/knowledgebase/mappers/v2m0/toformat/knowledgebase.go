package toformat

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
	"go.uber.org/zap"
)

func (tf ToFormat) processContainer(cont kbEnt.Container, model *formatV2M0.Model, ws workspaceHandler) error {
	model.BaseInfo = formatV2M0.BaseInfo{
		ID:          cont.KnowledgeBase.BaseInfo.ID,
		ShortName:   cont.KnowledgeBase.BaseInfo.ShortName,
		Description: cont.KnowledgeBase.BaseInfo.Description,
	}
	model.FormatXMLVersion = string(formats.VersionV2M0)
	model.Class = formatV2M0.Class{}
	model.Relations.Relations = []formatV2M0.Relation{}

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

	err := tf.processPatterns(processPatternsArgs{
		patterns: cont.Patterns,
		model:    model,
		ws:       ws,
	})
	if err != nil {
		tf.lg.Error(errMsgs.MappingPatternFail, zap.Error(err))
		return err
	}

	rootGroup := kbEnt.Group{}
	rootGroup.Parameters = append(rootGroup.Parameters, ws.GetUnprocessedParameters()...)
	rootGroup.Rules = append(rootGroup.Rules, ws.GetUnprocessedRules()...)

	return nil
}
