package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (tf *ToFormat) mapToParameters(parameters map[string]kbEnt.Parameter) ([]formatV3M0.Parameter, error) {
	ps := []formatV3M0.Parameter{}
	for _, v := range parameters {
		p, err := tf.mapToParameter(v)
		if err != nil {
			tf.lg.Error(errMsgs.MappingParameterFail, zap.Error(err))
			return []formatV3M0.Parameter{}, err
		}
		ps = append(ps, p)
	}

	return ps, nil
}

func (tf *ToFormat) mapToParameter(parameter kbEnt.Parameter) (formatV3M0.Parameter, error) {
	if err := tf.ws.CheckAndRememberParameter(parameter); err != nil {
		tf.lg.Error(errMsgs.MappingParameterFail, zap.Error(err))
		return formatV3M0.Parameter{}, err
	}

	p := formatV3M0.Parameter{
		BaseInfo: formatV3M0.BaseInfo{
			ID:           parameter.BaseInfo.ID,
			ShortName:    parameter.BaseInfo.ShortName,
			Description:  parameter.BaseInfo.Description,
			CreatedDate:  parameter.CreatedDate.UTC().Unix(),
			ModifiedDate: parameter.ModifiedDate.UTC().Unix(),
		},
		DefaultValue: parameter.DefaultValue,
	}
	parameterType, err := tf.mapToParameterType(parameter.Type)
	if err != nil {
		tf.lg.Error(errMsgs.MappingParameterTypeFail, zap.Error(err))
		return formatV3M0.Parameter{}, err
	}
	p.Type = parameterType

	return p, nil
}
