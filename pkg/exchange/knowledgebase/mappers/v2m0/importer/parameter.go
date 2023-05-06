package importer

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (vm Importer) processParameters(parameters []formatV2M0.Parameter, groupUUID string, ws workspaceHandler) error {
	for _, v := range parameters {
		parameter, err := vm.extractParameter(v, groupUUID, ws)
		if err != nil {
			vm.lg.Error(errMsgs.ParsingParameterFromXMLFail, zap.Error(err))
			return err
		}
		ws.AddParameter(parameter)
	}
	return nil
}

func (im Importer) extractParameter(parameter formatV2M0.Parameter, groupUUID string, ws workspaceHandler) (kbEnt.Parameter, error) {
	now := time.Now()
	p := kbEnt.Parameter{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ws.GetOrCreateParameterUUID(parameter.ID),
			ID:           parameter.ID,
			ShortName:    parameter.ShortName,
			Description:  parameter.Description,
			CreatedDate:  now,
			ModifiedDate: now,
		},
		GroupUUID:    groupUUID,
		DefaultValue: parameter.DefaultValue,
	}

	t, err := im.extractParameterType(parameter.Type)
	if err != nil {
		im.lg.Error(errMsgs.ExtractParameterTypeFail, zap.Error(err))
		return kbEnt.Parameter{}, err
	}
	p.Type = t

	return p, nil
}
