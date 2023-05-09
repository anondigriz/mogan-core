package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

func (ff *FromFormat) processKnowledgeBase() error {
	ff.ws.AddKnowledgeBase(ff.mapToKnowledgeBase())

	if err := ff.processParameters(ff.kb.Parameters.Parameters); err != nil {
		ff.lg.Error(errMsgs.MappingParametersFail, zap.Error(err))
		return err
	}

	if err := ff.processPatterns(ff.kb.Patterns.Patterns); err != nil {
		ff.lg.Error(errMsgs.MappingRelationsFail, zap.Error(err))
		return err
	}

	if err := ff.processRules(ff.kb.Rules.Rules); err != nil {
		ff.lg.Error(errMsgs.MappingRulesFail, zap.Error(err))
		return err
	}

	ff.processGroups(ff.kb.Groups.Groups)

	return nil
}

func (ff *FromFormat) mapToKnowledgeBase() kbEnt.KnowledgeBase {
	k := kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ff.kbUUID,
			ID:           ff.kb.ID,
			ShortName:    ff.kb.ShortName,
			Description:  ff.kb.Description,
			CreatedDate:  time.Unix(ff.kb.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(ff.kb.ModifiedDate, 0).UTC(),
		},
	}

	return k
}
