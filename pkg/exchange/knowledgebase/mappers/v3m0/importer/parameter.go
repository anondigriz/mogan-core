package importer

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (vm Importer) processParameters(parameters []formatV3M0.Parameter, ws workspaceHandler) error {
	for _, v := range parameters {
		parameter, err := vm.extractParameter(v, ws)
		if err != nil {
			vm.lg.Error(errMsgs.ParsingParameterFromXMLFail, zap.Error(err))
			return err
		}
		ws.AddParameter(parameter)
	}
	return nil
}

func (im Importer) extractParameter(parameter formatV3M0.Parameter, ws workspaceHandler) (kbEnt.Parameter, error) {
	p := kbEnt.Parameter{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ws.GetOrCreateParameterUUID(parameter.ID),
			ID:           parameter.ID,
			ShortName:    parameter.ShortName,
			Description:  parameter.Description,
			CreatedDate:  time.Unix(parameter.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(parameter.ModifiedDate, 0).UTC(),
		},
		GroupUUID:    ws.GetOrCreateGroupUUID(parameter.GroupUUID),
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
