package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ff FromFormat) processParameters(parameters []formatV2M0.Parameter, parentGroup *kbEnt.Group, ws workspaceHandler) error {
	for _, v := range parameters {
		parameter, err := ff.mapToParameter(v, ws)
		if err != nil {
			ff.lg.Error(errMsgs.MappingParameterFail, zap.Error(err))
			return err
		}
		parentGroup.Parameters = append(parentGroup.Parameters, parameter.UUID)
		ws.AddParameter(parameter)
	}
	return nil
}

func (ff FromFormat) mapToParameter(parameter formatV2M0.Parameter, ws workspaceHandler) (kbEnt.Parameter, error) {
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
		DefaultValue: parameter.DefaultValue,
	}

	t, err := ff.mapToParameterType(parameter.Type)
	if err != nil {
		ff.lg.Error(errMsgs.MappingParameterTypeFail, zap.Error(err))
		return kbEnt.Parameter{}, err
	}
	p.Type = t

	return p, nil
}
