package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type processGroupsArgs struct {
	childGroups map[string]kbEnt.Group
	parentClass *formatV2M0.Class
}

type processGroupArgs struct {
	group       kbEnt.Group
	parentClass *formatV2M0.Class
}

func (tf *ToFormat) processGroups(args processGroupsArgs) error {
	for _, v := range args.childGroups {
		err := tf.processGroup(processGroupArgs{
			group:       v,
			parentClass: args.parentClass,
		})
		if err != nil {
			tf.lg.Error(errMsgs.MappingGroupFail, zap.Error(err))
			return err
		}
	}

	return nil
}

func (tf *ToFormat) processGroup(args processGroupArgs) error {
	c, err := tf.mapToClass(args.group)
	if err != nil {
		tf.lg.Error(errMsgs.MappingClassFail, zap.Error(err))
		return err
	}

	err = tf.processParameters(processParametersArgs{
		parentGroup: args.group,
		parentClass: &c,
	})
	if err != nil {
		tf.lg.Error(errMsgs.MappingParametersFail, zap.Error(err))
		return err
	}

	err = tf.processGroups(processGroupsArgs{
		childGroups: args.group.Groups,
		parentClass: &c,
	})
	if err != nil {
		tf.lg.Error(errMsgs.MappingGroupsFail, zap.Error(err))
		return err
	}

	args.parentClass.Classes.Classes = append(args.parentClass.Classes.Classes, c)

	return nil
}

func (tf *ToFormat) mapToClass(group kbEnt.Group) (formatV2M0.Class, error) {
	if err := tf.ws.CheckAndRememberGroup(group); err != nil {
		tf.lg.Error(errMsgs.MappingGroupFail, zap.Error(err))
		return formatV2M0.Class{}, err
	}

	c := formatV2M0.Class{
		BaseInfo: formatV2M0.BaseInfo{
			ID:          group.ID,
			ShortName:   group.ShortName,
			Description: group.Description,
		},
	}

	return c, nil
}
