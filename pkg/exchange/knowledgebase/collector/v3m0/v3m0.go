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
}

func New(lg *zap.Logger) *V3M0 {
	vm := &V3M0{
		lg: lg,
	}
	return vm
}

func (vm V3M0) CollectXML(cont *kbEnt.Container, prefix, indent string) ([]byte, error) {
	tf := toformatV3M0.New(vm.lg, cont)
	model, err := tf.Map()
	if err != nil {
		vm.lg.Error(errMsgs.MapKnowledgeBaseFail, zap.Error(err))
		return []byte{}, err
	}

	b, err := xml.MarshalIndent(model, prefix, indent)
	if err != nil {
		vm.lg.Error(errMsgs.XMLMarshalFail, zap.Error(err))
		return []byte{}, errors.NewXMLMarshalFailErr(err)
	}

	return b, nil
}
