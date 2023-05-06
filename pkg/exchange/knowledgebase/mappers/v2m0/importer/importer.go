package importer

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type Importer struct {
	lg *zap.Logger
}

func New(lg *zap.Logger) *Importer {
	vm := &Importer{
		lg: lg,
	}
	return vm
}

func (im Importer) Import(kbUUID string, model *formatV2M0.Model) (kbEnt.Container, error) {
	ws := newWorkspace()
	ws.AddKnowledgeBase(im.extractKnowledgeBase(model, kbUUID))
	err := im.processModel(model, ws)
	if err != nil {
		im.lg.Error(errMsgs.ParsingModelFromXMLFail, zap.Error(err))
		return kbEnt.Container{}, err
	}

	return ws.cont, nil

}

func (vm Importer) extractKnowledgeBase(model *formatV2M0.Model, kbUUID string) kbEnt.KnowledgeBase {
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
