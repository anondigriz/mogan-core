package importer

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
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

func (im Importer) Import(kbUUID string, kb *formatV3M0.KnowledgeBase) (kbEnt.Container, error) {
	ws := newWorkspace()
	err := im.processKnowledgeBase(kbUUID, kb, ws)
	if err != nil {
		im.lg.Error(errMsgs.ParsingModelFromXMLFail, zap.Error(err))
		return kbEnt.Container{}, err
	}

	return ws.cont, nil

}
