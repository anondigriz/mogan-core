package exporter

import (
	"strconv"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

func (ex Exporter) extractRelationType(base types.PatternType) (string, error) {
	switch base {
	case types.Constraint:
		return "constr", nil
	case types.IfThenElse:
		return "ifclause", nil
	case types.Formula:
		return "simple", nil
	default:
		err := errors.NewTypeIsNotSupportedByFormatErr(strconv.Itoa(int(base)))
		ex.lg.Error(errMsgs.UnknownParameterTypeInXML, zap.Error(err))
		return "", err
	}
}

func (ex Exporter) extractParameterType(base types.ParameterType) (string, error) {
	switch base {
	case types.Double:
		return "double", nil
	case types.String:
		return "string", nil
	case types.Boolean:
		err := errors.NewTypeIsNotSupportedByFormatErr("Boolean")
		ex.lg.Error(errMsgs.UnknownParameterTypeInXML, zap.Error(err))
		return "", err
	default:
		err := errors.NewTypeIsNotSupportedByFormatErr("BigInteger")
		ex.lg.Error(errMsgs.UnknownParameterTypeInXML, zap.Error(err))
		return "", err
	}
}

func (ex Exporter) isAllowedScriptLanguageType(base types.ScriptLanguageType) error {
	switch base {
	case types.JS:
		return nil
	default:
		err := errors.NewTypeIsNotSupportedByFormatErr("Lua")
		ex.lg.Error(errMsgs.UnknownParameterTypeInXML, zap.Error(err))
		return err

	}
}
