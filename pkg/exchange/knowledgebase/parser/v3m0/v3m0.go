package v3m0

import (
	"encoding/xml"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
	fromFormatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/mappers/v3m0/fromformat"
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

func (vm V3M0) ParseXML(kbUUID string, content []byte) (kbEnt.Container, error) {
	knowledgeBase := formatV3M0.KnowledgeBase{}
	if err := xml.Unmarshal(content, &knowledgeBase); err != nil {
		vm.lg.Error(errMsgs.XMLUnmarshalFail, zap.Error(err))
		return kbEnt.Container{}, errors.NewXMLUnmarshalFailErr(err)
	}

	ff := fromFormatV3M0.New(vm.lg, kbUUID, &knowledgeBase)
	cont, err := ff.Map()
	if err != nil {
		vm.lg.Error(errMsgs.MapKnowledgeBaseFail, zap.Error(err))
		return kbEnt.Container{}, err
	}

	return cont, nil
}
