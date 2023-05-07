package exporter

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ex Exporter) processParameters(parameters []formatV2M0.Parameter, groupUUID string, ws workspaceHandler) error {
	for _, v := range parameters {
		parameter, err := ex.extractParameter(v, groupUUID, ws)
		if err != nil {
			ex.lg.Error(errMsgs.ParsingParameterFromXMLFail, zap.Error(err))
			return err
		}
		ws.AddParameter(parameter)
	}
	return nil
}

func (ex Exporter) extractParameter(parameter formatV2M0.Parameter, groupUUID string, ws workspaceHandler) (kbEnt.Parameter, error) {
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

	t, err := ex.extractParameterType(parameter.Type)
	if err != nil {
		ex.lg.Error(errMsgs.ExtractParameterTypeFail, zap.Error(err))
		return kbEnt.Parameter{}, err
	}
	p.Type = t

	return p, nil
}
