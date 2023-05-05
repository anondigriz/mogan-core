package importer

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type Importer struct {
	lg *zap.Logger
}

type ids struct {
	Parameters map[string]string
	Patterns   map[string]string
}

func New(lg *zap.Logger) *Importer {
	vm := &Importer{
		lg: lg,
	}
	return vm
}

func (vm Importer) Import(kbUUID string, model *formatV2M0.Model) (kbEnt.Container, error) {
	cont := &kbEnt.Container{
		Groups:     map[string]kbEnt.Group{},
		Parameters: map[string]kbEnt.Parameter{},
		Patterns:   map[string]kbEnt.Pattern{},
		Rules:      map[string]kbEnt.Rule{},
	}

	cont.KnowledgeBase = kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        kbUUID,
			ID:          model.ID,
			ShortName:   model.ShortName,
			CreatedDate: time.Now(),
		},
		ExtraData: kbEnt.ExtraDataKnowledgeBase{
			Description: model.Description,
		},
	}
	cont.KnowledgeBase.ModifiedDate = cont.KnowledgeBase.CreatedDate

	mapIDs := &ids{
		Parameters: map[string]string{},
		Patterns:   map[string]string{},
	}

	for _, v := range model.Relations.Relations {
		err := vm.parseRelation(v, cont, mapIDs)
		if err != nil {
			vm.lg.Error("parsing of the rule ended with an error", zap.Error(err))
			return kbEnt.Container{}, err
		}
	}

	gh, err := vm.parseClass(model.Class, cont, mapIDs)
	if err != nil {
		vm.lg.Error("parsing of the main class ended with an error", zap.Error(err))
		return kbEnt.Container{}, err
	}
	cont.KnowledgeBase.ExtraData.Groups = gh

	return *cont, nil
}
