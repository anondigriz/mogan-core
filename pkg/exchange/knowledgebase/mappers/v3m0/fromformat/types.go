package fromformat

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

func (ff FromFormat) mapToPatternType(base string) (types.PatternType, error) {
	switch base {
	case "Constraint":
		return types.Constraint, nil
	case "IfThenElse":
		return types.IfThenElse, nil
	case "Program":
		return types.Program, nil
	case "Formula":
		return types.Formula, nil
	default:
		err := errors.NewUnknownPatternTypeErr(base)
		ff.lg.Error(errMsgs.UnknownPatternType, zap.Error(err))
		return 0, err
	}
}

func (ff FromFormat) mapToParameterType(base string) (types.ParameterType, error) {
	switch base {
	case "Double":
		return types.Double, nil
	case "String":
		return types.String, nil
	default:
		err := errors.NewUnknownParameterTypeErr(base)
		ff.lg.Error(errMsgs.UnknownParameterType, zap.Error(err))
		return 0, err
	}
}

func (ff FromFormat) mapToScriptLanguageType(base string) (types.ScriptLanguageType, error) {
	switch base {
	case "JS":
		return types.JS, nil
	case "Lua":
		return types.Lua, nil
	default:
		err := errors.NewUnknownScriptLanguageTypeErr(base)
		ff.lg.Error(errMsgs.UnknownScriptLanguageType, zap.Error(err))
		return 0, err
	}
}
