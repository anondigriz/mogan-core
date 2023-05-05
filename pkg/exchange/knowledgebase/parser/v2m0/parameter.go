package v2m0

import (
	"fmt"
	"time"

	uuidGen "github.com/google/uuid"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	entV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/v2m0"
)

func (vm V2M0) parseParameter(parameter entV2M0.Parameter, gr kbEnt.Group, cont *kbEnt.Container, ids *ids) error {
	pr := kbEnt.Parameter{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        vm.getOrCreateParameterUUID(parameter.ID, ids),
			ID:          parameter.ID,
			ShortName:   parameter.ShortName,
			CreatedDate: time.Now(),
		},
		ExtraData: kbEnt.ExtraDataParameter{
			Description:  parameter.Description,
			DefaultValue: parameter.DefaultValue,
		},
	}
	pr.ModifiedDate = pr.CreatedDate

	t, err := vm.convertTypeParameter(parameter.Type)
	if err != nil {
		return err
	}
	pr.Type = t
	pr.GroupUUID = gr.UUID

	ids.Parameters[pr.ID] = pr.UUID
	cont.Parameters[pr.UUID] = pr
	return nil
}

func (vm V2M0) convertTypeParameter(base string) (types.ParameterType, error) {
	switch base {
	case "double":
		return types.Double, nil
	case "string":
		return types.String, nil
	default:
		return types.String, errors.NewParsingXMLFailErr(
			fmt.Sprintf("unknown parameter type from the XML file %s", base),
			nil)
	}
}

func (vm V2M0) getOrCreateParameterUUID(id string, ids *ids) string {
	prUUID, ok := ids.Parameters[id]
	if !ok {
		prUUID = uuidGen.NewString()
	}
	return prUUID
}
