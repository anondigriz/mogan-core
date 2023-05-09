package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

func (ff *FromFormat) processModel() error {
	ff.ws.AddKnowledgeBase(ff.mapToKnowledgeBase())

	err := ff.processRelations(ff.model.Relations.Relations)
	if err != nil {
		ff.lg.Error(errMsgs.MappingRelationsFail, zap.Error(err))
		return err
	}

	rootGroup, err := ff.processClass(ff.model.Class)
	if err != nil {
		ff.lg.Error(errMsgs.MappingClassFail, zap.Error(err))
		return err
	}

	for _, v := range rootGroup.Groups {
		ff.ws.AddGroup(v)
	}

	return nil
}

func (ff *FromFormat) mapToKnowledgeBase() kbEnt.KnowledgeBase {
	now := time.Now()
	return kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ff.kbUUID,
			ID:           ff.model.ID,
			ShortName:    ff.model.ShortName,
			Description:  ff.model.Description,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}
}
