package exporter

import (
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ex Exporter) processModel(model *formatV2M0.Model, ws workspaceHandler) error {
	err := ex.processRelations(model.Relations.Relations, ws)
	if err != nil {
		ex.lg.Error(errMsgs.ParsingRelationsFromXMLFail, zap.Error(err))
		return err
	}

	gh, err := ex.processClass(model.Class, ws)
	if err != nil {
		ex.lg.Error(errMsgs.ParsingClassFromXMLFail, zap.Error(err))
		return err
	}
	ws.AddGroupHierarchy(gh)

	return nil
}
