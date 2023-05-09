package fromformat

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ff FromFormat) mapToPatternType(base string) (types.PatternType, error) {
	switch base {
	case string(formatV2M0.Prog):
		return types.Program, nil
	case string(formatV2M0.Constr):
		return types.Constraint, nil
	case string(formatV2M0.Simple):
		return types.Program, nil
	case string(formatV2M0.Ifclause):
		return types.Program, nil
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
