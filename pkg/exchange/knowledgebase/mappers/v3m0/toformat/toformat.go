package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

type ToFormat struct {
	lg *zap.Logger
}

func New(lg *zap.Logger) *ToFormat {
	vm := &ToFormat{
		lg: lg,
	}
	return vm
}

func (tf ToFormat) Map(cont kbEnt.Container) (formatV3M0.KnowledgeBase, error) {
	knowledgeBase := &formatV3M0.KnowledgeBase{}
	ws := newWorkspace()

	err := tf.processContainer(knowledgeBase, cont, ws)
	if err != nil {
		tf.lg.Error(errMsgs.MapKnowledgeBaseFail, zap.Error(err))
		return formatV3M0.KnowledgeBase{}, err
	}

	return *knowledgeBase, nil
}
