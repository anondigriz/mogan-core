package fromformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

type FromFormat struct {
	lg     *zap.Logger
	ws     workspaceHandler
	kbUUID string
	kb     *formatV3M0.KnowledgeBase
}

func New(lg *zap.Logger, kbUUID string, kb *formatV3M0.KnowledgeBase) *FromFormat {
	vm := &FromFormat{
		lg:     lg,
		kbUUID: kbUUID,
		kb:     kb,
	}
	return vm
}

func (ff *FromFormat) Map() (kbEnt.Container, error) {
	ff.ws = newWorkspace()
	err := ff.processKnowledgeBase()
	if err != nil {
		ff.lg.Error(errMsgs.MapKnowledgeBaseFail, zap.Error(err))
		return kbEnt.Container{}, err
	}

	return ff.ws.GetContainer(), nil
}
