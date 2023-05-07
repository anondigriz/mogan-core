package importer

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

func (im Importer) extractPatternType(base string) (types.PatternType, error) {
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
		err := errors.NewUnknownPatternTypeInXMLErr(base)
		im.lg.Error(errMsgs.UnknownPatternTypeInXML, zap.Error(err))
		return 0, err
	}
}

func (im Importer) extractParameterType(base string) (types.ParameterType, error) {
	switch base {
	case "Double":
		return types.Double, nil
	case "String":
		return types.String, nil
	default:
		err := errors.NewUnknownParameterTypeInXMLErr(base)
		im.lg.Error(errMsgs.UnknownParameterTypeInXML, zap.Error(err))
		return 0, err
	}
}

func (im Importer) extractScriptLanguageType(base string) (types.ScriptLanguageType, error) {
	switch base {
	case "JS":
		return types.JS, nil
	case "Lua":
		return types.Lua, nil
	default:
		err := errors.NewUnknownScriptLanguageTypeInXMLErr(base)
		im.lg.Error(errMsgs.UnknownScriptLanguageTypeInXML, zap.Error(err))
		return 0, err
	}
}
