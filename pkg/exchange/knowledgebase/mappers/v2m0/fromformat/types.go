package fromformat

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

func (ff FromFormat) mapToPatternType(base string) (types.PatternType, error) {
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
		err := errors.NewUnknownPatternTypeErr(base)
		ff.lg.Error(errMsgs.UnknownPatternType, zap.Error(err))
		return 0, err
	}
}

func (ff FromFormat) mapToParameterType(base string) (types.ParameterType, error) {
	switch base {
	case "double":
		return types.Double, nil
	case "string":
		return types.String, nil
	default:
		err := errors.NewUnknownParameterTypeErr(base)
		ff.lg.Error(errMsgs.UnknownParameterType, zap.Error(err))
		return 0, err
	}
}
