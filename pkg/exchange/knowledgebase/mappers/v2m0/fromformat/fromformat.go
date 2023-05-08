package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type FromFormat struct {
	lg *zap.Logger
}

func New(lg *zap.Logger) *FromFormat {
	vm := &FromFormat{
		lg: lg,
	}
	return vm
}

func (ff FromFormat) Map(kbUUID string, model formatV2M0.Model) (kbEnt.Container, error) {
	ws := newWorkspace()
	ws.AddKnowledgeBase(ff.mapToKnowledgeBase(model, kbUUID))
	err := ff.processModel(model, ws)
	if err != nil {
		ff.lg.Error(errMsgs.MapKnowledgeBaseFail, zap.Error(err))
		return kbEnt.Container{}, err
	}

	return ws.cont, nil
}

func (ff FromFormat) mapToKnowledgeBase(model formatV2M0.Model, kbUUID string) kbEnt.KnowledgeBase {
	now := time.Now()
	return kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         kbUUID,
			ID:           model.ID,
			ShortName:    model.ShortName,
			Description:  model.Description,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}
}
