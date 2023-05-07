package v3m0

import (
	"encoding/xml"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/mappers/v3m0/importer"

	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

type V3M0 struct {
	lg *zap.Logger
	im *importer.Importer
}

func New(lg *zap.Logger) *V3M0 {
	vm := &V3M0{
		lg: lg,
		im: importer.New(lg),
	}
	return vm
}

func (vm V3M0) ParseXML(kbUUID string, content []byte) (kbEnt.Container, error) {
	knowledgeBase := formatV3M0.KnowledgeBase{}
	if err := xml.Unmarshal(content, &knowledgeBase); err != nil {
		vm.lg.Error("fail to unmarshal the xml file", zap.Error(err))
		return kbEnt.Container{}, errors.NewXMLUnmarshalFailErr(err)
	}

	cont, err := vm.im.Import(kbUUID, knowledgeBase)
	if err != nil {
		vm.lg.Error("mapping xml entities to app entities fail", zap.Error(err))
		return kbEnt.Container{}, err
	}

	return cont, nil
}
