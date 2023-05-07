package exporter

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type Exporter struct {
	lg *zap.Logger
}

func New(lg *zap.Logger) *Exporter {
	vm := &Exporter{
		lg: lg,
	}
	return vm
}

func (ex Exporter) Export(knowledgeBase kbEnt.Container) (kbEnt.Container, error) {
	err := ex.processModel(model, ws)
	if err != nil {
		ex.lg.Error(errMsgs.ParsingModelFromXMLFail, zap.Error(err))
		return kbEnt.Container{}, err
	}

	return ws.cont, nil

}

func (ex Exporter) extractKnowledgeBase(model *, kbUUID string) formatV2M0.Model {
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
