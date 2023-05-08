package fromformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
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

func (ff FromFormat) Map(kbUUID string, kb formatV3M0.KnowledgeBase) (kbEnt.Container, error) {
	ws := newWorkspace()
	err := ff.processKnowledgeBase(kbUUID, kb, ws)
	if err != nil {
		ff.lg.Error(errMsgs.MapKnowledgeBaseFail, zap.Error(err))
		return kbEnt.Container{}, err
	}

	return ws.cont, nil

}
