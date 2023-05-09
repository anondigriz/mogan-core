package toformat

import (
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (tf *ToFormat) processContainer(knowledgeBase *formatV3M0.KnowledgeBase) error {
	tf.prepareKnowledgeBase(knowledgeBase)

	gs, err := tf.mapToGroups(tf.cont.Groups)
	if err != nil {
		tf.lg.Error(errMsgs.MappingGroupsFail, zap.Error(err))
		return nil
	}
	knowledgeBase.Groups.Groups = gs

	ps, err := tf.mapToParameters(tf.cont.Parameters)
	if err != nil {
		tf.lg.Error(errMsgs.MappingParametersFail, zap.Error(err))
		return nil
	}
	knowledgeBase.Parameters.Parameters = ps

	pts, err := tf.mapToPatterns(tf.cont.Patterns)
	if err != nil {
		tf.lg.Error(errMsgs.MappingPatternsFail, zap.Error(err))
		return nil
	}
	knowledgeBase.Patterns.Patterns = pts

	rs, err := tf.mapToRules()
	if err != nil {
		tf.lg.Error(errMsgs.MappingRulesFail, zap.Error(err))
		return nil
	}
	knowledgeBase.Rules.Rules = rs

	return nil
}

func (tf *ToFormat) prepareKnowledgeBase(knowledgeBase *formatV3M0.KnowledgeBase) {
	knowledgeBase.BaseInfo = formatV3M0.BaseInfo{
		ID:           tf.cont.KnowledgeBase.BaseInfo.ID,
		ShortName:    tf.cont.KnowledgeBase.BaseInfo.ShortName,
		Description:  formatV3M0.Description{Text: tf.cont.KnowledgeBase.BaseInfo.Description},
		CreatedDate:  tf.cont.KnowledgeBase.CreatedDate.UTC().Unix(),
		ModifiedDate: tf.cont.KnowledgeBase.ModifiedDate.UTC().Unix(),
	}
	knowledgeBase.FormatXMLVersion = string(formats.VersionV3M0)
}
