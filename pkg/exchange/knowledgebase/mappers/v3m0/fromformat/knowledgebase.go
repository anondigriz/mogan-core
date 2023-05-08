package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (ff FromFormat) processKnowledgeBase(kbUUID string, knowledgeBase formatV3M0.KnowledgeBase, ws workspaceHandler) error {
	ws.AddKnowledgeBase(ff.mapToKnowledgeBase(kbUUID, knowledgeBase, ws))

	ff.processGroups(knowledgeBase.Groups.Groups, ws)

	if err := ff.processParameters(knowledgeBase.Parameters.Parameters, ws); err != nil {
		ff.lg.Error(errMsgs.MappingParametersFail, zap.Error(err))
		return err
	}

	if err := ff.processPatterns(knowledgeBase.Patterns.Patterns, ws); err != nil {
		ff.lg.Error(errMsgs.MappingRelationsFail, zap.Error(err))
		return err
	}

	if err := ff.processRules(knowledgeBase.Rules.Rules, ws); err != nil {
		ff.lg.Error(errMsgs.MappingRulesFail, zap.Error(err))
		return err
	}

	return nil
}

func (ff FromFormat) mapToKnowledgeBase(kbUUID string, knowledgeBase formatV3M0.KnowledgeBase, ws workspaceHandler) kbEnt.KnowledgeBase {
	k := kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         kbUUID,
			ID:           knowledgeBase.ID,
			ShortName:    knowledgeBase.ShortName,
			Description:  knowledgeBase.Description,
			CreatedDate:  time.Unix(knowledgeBase.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(knowledgeBase.ModifiedDate, 0).UTC(),
		},
	}

	return k
}
