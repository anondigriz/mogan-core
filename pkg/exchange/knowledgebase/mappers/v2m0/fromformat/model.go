package fromformat

import (
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ff FromFormat) processModel(model formatV2M0.Model, ws workspaceHandler) error {
	err := ff.processRelations(model.Relations.Relations, ws)
	if err != nil {
		ff.lg.Error(errMsgs.ParsingRelationsFromXMLFail, zap.Error(err))
		return err
	}

	rootGroup, err := ff.processClass(model.Class, ws)
	if err != nil {
		ff.lg.Error(errMsgs.ParsingClassFromXMLFail, zap.Error(err))
		return err
	}

	for _, v := range rootGroup.Groups {
		ws.AddGroup(v)
	}

	return nil
}
