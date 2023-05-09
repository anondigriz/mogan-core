package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type ToFormat struct {
	lg   *zap.Logger
	cont *kbEnt.Container
	ws   workspaceHandler
}

func New(lg *zap.Logger, cont *kbEnt.Container) *ToFormat {
	vm := &ToFormat{
		lg:   lg,
		cont: cont,
	}
	return vm
}

func (tf *ToFormat) Map() (formatV2M0.Model, error) {
	tf.ws = newWorkspace()
	model := &formatV2M0.Model{}

	err := tf.processContainer(model)
	if err != nil {
		tf.lg.Error(errMsgs.MapKnowledgeBaseFail, zap.Error(err))
		return formatV2M0.Model{}, err
	}

	return *model, nil
}
