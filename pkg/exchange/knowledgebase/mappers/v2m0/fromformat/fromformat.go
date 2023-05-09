package fromformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type FromFormat struct {
	lg     *zap.Logger
	ws     workspaceHandler
	kbUUID string
	model  *formatV2M0.Model
}

func New(lg *zap.Logger, kbUUID string, model *formatV2M0.Model) *FromFormat {
	vm := &FromFormat{
		lg:     lg,
		kbUUID: kbUUID,
		model:  model,
	}
	return vm
}

func (ff *FromFormat) Map() (kbEnt.Container, error) {
	ff.ws = newWorkspace()
	err := ff.processModel()
	if err != nil {
		ff.lg.Error(errMsgs.MapKnowledgeBaseFail, zap.Error(err))
		return kbEnt.Container{}, err
	}

	return ff.ws.GetContainer(), nil
}
