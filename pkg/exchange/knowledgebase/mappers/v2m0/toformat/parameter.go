package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type processParametersArgs struct {
	parentGroup kbEnt.Group
	parentClass *formatV2M0.Class
}

func (tf *ToFormat) processParameters(args processParametersArgs) error {
	for _, v := range args.parentGroup.Parameters {
		parameter, ok := tf.cont.Parameters[v]
		if !ok || tf.ws.IsProcessedParameters(parameter.UUID) {
			continue
		}

		p, err := tf.mapToParameter(parameter)
		if err != nil {
			tf.lg.Error(errMsgs.MappingRelationFail, zap.Error(err))
			return err
		}

		args.parentClass.Parameters.Parameters = append(args.parentClass.Parameters.Parameters, p)

		tf.ws.SetProcessedParameters(parameter.UUID)
	}
	return nil
}

func (tf *ToFormat) mapToParameter(parameter kbEnt.Parameter) (formatV2M0.Parameter, error) {
	if err := tf.ws.CheckAndRememberParameter(parameter); err != nil {
		tf.lg.Error(errMsgs.MappingParameterFail, zap.Error(err))
		return formatV2M0.Parameter{}, err
	}

	p := formatV2M0.Parameter{
		BaseInfo: formatV2M0.BaseInfo{
			ID:          parameter.ID,
			ShortName:   parameter.ShortName,
			Description: parameter.Description,
		},
		DefaultValue: parameter.DefaultValue,
	}
	parameterType, err := tf.mapToParameterType(parameter.Type)
	if err != nil {
		tf.lg.Error(errMsgs.MappingParameterTypeFail, zap.Error(err))
		return formatV2M0.Parameter{}, err
	}
	p.Type = parameterType

	return p, nil
}
