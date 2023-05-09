package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (ff *FromFormat) processParameters(parameters []formatV3M0.Parameter) error {
	for _, v := range parameters {
		parameter, err := ff.mapToParameter(v)
		if err != nil {
			ff.lg.Error(errMsgs.MappingParameterFail, zap.Error(err))
			return err
		}
		ff.ws.AddParameter(parameter)
	}
	return nil
}

func (ff *FromFormat) mapToParameter(parameter formatV3M0.Parameter) (kbEnt.Parameter, error) {
	p := kbEnt.Parameter{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ff.ws.GetOrCreateParameterUUID(parameter.ID),
			ID:           parameter.ID,
			ShortName:    parameter.ShortName,
			Description:  parameter.Description,
			CreatedDate:  time.Unix(parameter.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(parameter.ModifiedDate, 0).UTC(),
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
