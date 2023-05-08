package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
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

func (tf ToFormat) Export(knowledgeBase kbEnt.Container) (formatV2M0.Model, error) {
	model := &formatV2M0.Model{}
	ws := newWorkspace()

	err := tf.processKnowledgeBase(knowledgeBase, model, ws)
	if err != nil {
		tf.lg.Error(errMsgs.ExportKnowledgeBaseToXMLFail, zap.Error(err))
		return formatV2M0.Model{}, err
	}

	return *model, nil

}
