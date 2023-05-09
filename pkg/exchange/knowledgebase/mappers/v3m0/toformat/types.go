package toformat

import (
	"strconv"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (tf *ToFormat) mapToPatternType(base types.PatternType) (string, error) {
	switch base {
	case types.Program:
		return string(formatV3M0.Program), nil
	case types.Constraint:
		return string(formatV3M0.Constraint), nil
	default:
		err := errors.NewTypeIsNotSupportedByFormatErr(strconv.Itoa(int(base)))
		tf.lg.Error(errMsgs.TypeIsNotSupportedByFormat, zap.Error(err))
		return "", err
	}
}

func (tf *ToFormat) mapToParameterType(base types.ParameterType) (string, error) {
	switch base {
	case types.Double:
		return string(formatV3M0.Double), nil
	case types.String:
		return string(formatV3M0.String), nil
	case types.Boolean:
		return string(formatV3M0.Boolean), nil
	default:
		err := errors.NewTypeIsNotSupportedByFormatErr(strconv.Itoa(int(base)))
		tf.lg.Error(errMsgs.TypeIsNotSupportedByFormat, zap.Error(err))
		return "", err
	}
}

func (tf *ToFormat) mapToScriptLanguageType(base types.ScriptLanguageType) (string, error) {
	switch base {
	case types.JS:
		return string(formatV3M0.JS), nil
	case types.Lua:
		return string(formatV3M0.Lua), nil
	default:
		err := errors.NewTypeIsNotSupportedByFormatErr(strconv.Itoa(int(base)))
		tf.lg.Error(errMsgs.TypeIsNotSupportedByFormat, zap.Error(err))
		return "", err

	}
}
