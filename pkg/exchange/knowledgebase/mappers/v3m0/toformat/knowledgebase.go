package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (tf ToFormat) processContainer(knowledgeBase *formatV3M0.KnowledgeBase, cont kbEnt.Container, ws workspaceHandler) error {
	tf.prepareKnowledgeBase(knowledgeBase, cont, ws)

	gs, err := tf.mapToGroups(cont.Groups, ws)
	if err != nil {
		tf.lg.Error(errMsgs.MappingGroupsFail, zap.Error(err))
		return nil
	}
	knowledgeBase.Groups.Groups = gs

	ps, err := tf.mapToParameters(cont.Parameters, ws)
	if err != nil {
		tf.lg.Error(errMsgs.MappingParametersFail, zap.Error(err))
		return nil
	}
	knowledgeBase.Parameters.Parameters = ps

	pts, err := tf.mapToPatterns(cont.Patterns, ws)
	if err != nil {
		tf.lg.Error(errMsgs.MappingPatternsFail, zap.Error(err))
		return nil
	}
	knowledgeBase.Patterns.Patterns = pts

	rs, err := tf.mapToRules(cont, ws)
	if err != nil {
		tf.lg.Error(errMsgs.MappingRulesFail, zap.Error(err))
		return nil
	}
	knowledgeBase.Rules.Rules = rs

	return nil
}

func (tf ToFormat) prepareKnowledgeBase(knowledgeBase *formatV3M0.KnowledgeBase, cont kbEnt.Container, ws workspaceHandler) {
	knowledgeBase.BaseInfo = formatV3M0.BaseInfo{
		ID:           cont.KnowledgeBase.BaseInfo.ID,
		ShortName:    cont.KnowledgeBase.BaseInfo.ShortName,
		Description:  cont.KnowledgeBase.BaseInfo.Description,
		CreatedDate:  cont.KnowledgeBase.CreatedDate.UTC().Unix(),
		ModifiedDate: cont.KnowledgeBase.ModifiedDate.UTC().Unix(),
	}
	knowledgeBase.FormatXMLVersion = string(formats.VersionV3M0)
}
