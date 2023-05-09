package fromformat

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (ff *FromFormat) mapToPatternType(base string) (types.PatternType, error) {
	switch base {
	case string(formatV3M0.Program):
		return types.Program, nil
	case string(formatV3M0.Constraint):
		return types.Constraint, nil
	default:
		err := errors.NewUnknownPatternTypeErr(base)
		ff.lg.Error(errMsgs.UnknownPatternType, zap.Error(err))
		return 0, err
	}
}

func (ff *FromFormat) mapToParameterType(base string) (types.ParameterType, error) {
	switch base {
	case string(formatV3M0.String):
		return types.String, nil
	case string(formatV3M0.Double):
		return types.Double, nil
	case string(formatV3M0.Boolean):
		return types.Boolean, nil
	default:
		err := errors.NewUnknownParameterTypeErr(base)
		ff.lg.Error(errMsgs.UnknownParameterType, zap.Error(err))
		return 0, err
	}
}

func (ff *FromFormat) mapToScriptLanguageType(base string) (types.ScriptLanguageType, error) {
	switch base {
	case string(formatV3M0.JS):
		return types.JS, nil
	case string(formatV3M0.Lua):
		return types.Lua, nil
	default:
		err := errors.NewUnknownScriptLanguageTypeErr(base)
		ff.lg.Error(errMsgs.UnknownScriptLanguageType, zap.Error(err))
		return 0, err
	}
}
