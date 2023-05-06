package importer

import (
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (im Importer) processModel(model *formatV2M0.Model, ws workspaceHandler) error {
	err := im.processRelations(model.Relations.Relations, ws)
	if err != nil {
		im.lg.Error(errMsgs.ParsingRelationsFromXMLFail, zap.Error(err))
		return err
	}

	gh, err := im.processClass(model.Class, ws)
	if err != nil {
		im.lg.Error(errMsgs.ParsingClassFromXMLFail, zap.Error(err))
		return err
	}
	ws.AddGroupHierarchy(gh)

	return nil
}
