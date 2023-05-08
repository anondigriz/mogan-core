package v2m0

import (
	"encoding/xml"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/mappers/v2m0/fromformat"

	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type V2M0 struct {
	lg *zap.Logger
	im *fromformat.FromFormat
}

func New(lg *zap.Logger) *V2M0 {
	vm := &V2M0{
		lg: lg,
		im: fromformat.New(lg),
	}
	return vm
}

func (vm V2M0) ParseXML(kbUUID string, content []byte) (kbEnt.Container, error) {
	model := formatV2M0.Model{}
	if err := xml.Unmarshal(content, &model); err != nil {
		vm.lg.Error("fail to unmarshal the xml file", zap.Error(err))
		return kbEnt.Container{}, errors.NewXMLUnmarshalFailErr(err)
	}

	cont, err := vm.im.Import(kbUUID, model)
	if err != nil {
		vm.lg.Error("mapping xml entities to app entities fail", zap.Error(err))
		return kbEnt.Container{}, err
	}

	return cont, nil
}
