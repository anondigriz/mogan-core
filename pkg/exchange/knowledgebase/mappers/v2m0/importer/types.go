package importer

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

func (im Importer) extractPatternType(base string) (types.PatternType, error) {
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
		err := errors.NewUnknownPatternTypeInXMLErr(base)
		im.lg.Error(errMsgs.UnknownPatternTypeInXML, zap.Error(err))
		return 0, err
	}
}

func (im Importer) extractParameterType(base string) (types.ParameterType, error) {
	switch base {
	case "double":
		return types.Double, nil
	case "string":
		return types.String, nil
	default:
		err := errors.NewUnknownParameterTypeInXMLErr(base)
		im.lg.Error(errMsgs.UnknownParameterTypeInXML, zap.Error(err))
		return 0, err
	}
}
