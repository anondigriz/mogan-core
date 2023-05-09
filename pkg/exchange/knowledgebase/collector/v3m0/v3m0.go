package v3m0

import (
	"encoding/xml"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	toformatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/mappers/v3m0/toformat"
)

type V3M0 struct {
	lg *zap.Logger
	tf *toformatV3M0.ToFormat
}

func New(lg *zap.Logger) *V3M0 {
	vm := &V3M0{
		lg: lg,
		tf: toformatV3M0.New(lg),
	}
	return vm
}

func (vm V3M0) CollectXML(KnowledgeBase kbEnt.Container) ([]byte, error) {
	model, err := vm.tf.Map(KnowledgeBase)
	if err != nil {
		vm.lg.Error(errMsgs.MapKnowledgeBaseFail, zap.Error(err))
		return []byte{}, err
	}

	b, err := xml.MarshalIndent(model, "", "  ")
	if err != nil {
		vm.lg.Error(errMsgs.XMLMarshalFail, zap.Error(err))
		return []byte{}, errors.NewXMLMarshalFailErr(err)
	}

	return b, nil
}
