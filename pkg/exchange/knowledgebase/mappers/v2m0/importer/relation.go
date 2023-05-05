package importer

import (
	"fmt"
	"strings"
	"time"

	uuidGen "github.com/google/uuid"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (vm Importer) parseRelation(relation formatV2M0.Relation, cont *kbEnt.Container, ids *ids) error {
	rn := kbEnt.Pattern{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        uuidGen.NewString(),
			ID:          relation.ID,
			ShortName:   relation.ShortName,
			CreatedDate: time.Now(),
		},
		ExtraData: kbEnt.ExtraDataPattern{
			Description:      relation.Description,
			Language:         types.JS,
			Script:           relation.Script,
			InputParameters:  []kbEnt.ParameterPattern{},
			OutputParameters: []kbEnt.ParameterPattern{},
		},
	}
	rn.ModifiedDate = rn.CreatedDate

	t, err := vm.convertTypePattern(relation.RelationType)
	if err != nil {
		return err
	}
	rn.Type = t

	inObj, err := vm.splitParameters(relation.InObjects)
	if err != nil {
		return err
	}

	for k, v := range inObj {
		paramType, err := vm.convertTypeParameter(v)
		if err != nil {
			return err
		}

		rn.ExtraData.InputParameters = append(rn.ExtraData.InputParameters, kbEnt.ParameterPattern{
			ShortName: k,
			Type:      paramType,
		})
	}

	outObj, err := vm.splitParameters(relation.OutObjects)
	if err != nil {
		return err
	}

	for k, v := range outObj {
		paramType, err := vm.convertTypeParameter(v)
		if err != nil {
			return err
		}

		rn.ExtraData.OutputParameters = append(rn.ExtraData.OutputParameters, kbEnt.ParameterPattern{
			ShortName: k,
			Type:      paramType,
		})
	}

	ids.Patterns[rn.ID] = rn.UUID
	cont.Patterns[rn.UUID] = rn
	return nil

}

func (vm Importer) convertTypePattern(base string) (types.PatternType, error) {
	switch base {
	case "constr":
		return types.Constraint, nil
	case "ifclause":
		return types.IfThenElse, nil
	case "prog":
		return types.Program, nil
	case "simple":
		return types.Formula, nil
	default:
		return types.Program, errors.NewParsingXMLFailErr(
			fmt.Sprintf("unknown pattern type from the XML file %s", base),
			nil)
	}
}

func (vm Importer) splitParameters(base string) (map[string]string, error) {
	params := map[string]string{}
	if base == "" {
		return params, nil
	}

	pair := strings.Split(base, ";")
	for _, v := range pair {
		keyValue := strings.Split(v, ":")
		if len(keyValue) != 2 {
			return nil, errors.NewParsingXMLFailErr(
				fmt.Sprintf("'%s' is not a key-value pair", v),
				nil)
		}
		params[keyValue[0]] = keyValue[1]
	}
	return params, nil
}
