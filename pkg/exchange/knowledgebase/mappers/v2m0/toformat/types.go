package toformat

import (
	"strconv"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (tf ToFormat) mapToRelationType(base types.PatternType) (string, error) {
	switch base {
	case types.Program:
		return string(formatV2M0.Prog), nil
	case types.Constraint:
		return string(formatV2M0.Constr), nil
	case types.Formula:
		return string(formatV2M0.Simple), nil
	case types.IfThenElse:
		return string(formatV2M0.Ifclause), nil
	default:
		err := errors.NewTypeIsNotSupportedByFormatErr(strconv.Itoa(int(base)))
		tf.lg.Error(errMsgs.TypeIsNotSupportedByFormat, zap.Error(err))
		return "", err
	}
}

func (tf ToFormat) mapToParameterType(base types.ParameterType) (string, error) {
	switch base {
	case types.Double:
		return string(formatV2M0.Double), nil
	case types.String:
		return string(formatV2M0.String), nil
	case types.Boolean:
		err := errors.NewTypeIsNotSupportedByFormatErr("Boolean")
		tf.lg.Error(errMsgs.TypeIsNotSupportedByFormat, zap.Error(err))
		return "", err
	default:
		err := errors.NewTypeIsNotSupportedByFormatErr("BigInteger")
		tf.lg.Error(errMsgs.TypeIsNotSupportedByFormat, zap.Error(err))
		return "", err
	}
}

func (ex ToFormat) isAllowedScriptLanguageType(base types.ScriptLanguageType) error {
	switch base {
	case types.JS:
		return nil
	default:
		err := errors.NewTypeIsNotSupportedByFormatErr("Lua")
		ex.lg.Error(errMsgs.TypeIsNotSupportedByFormat, zap.Error(err))
		return err

	}
}
