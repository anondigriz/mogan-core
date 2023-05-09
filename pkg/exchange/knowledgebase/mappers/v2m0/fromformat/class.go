package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ff *FromFormat) processClasses(classes []formatV2M0.Class, parentGroup *kbEnt.Group) error {
	for _, v := range classes {
		childGroup, err := ff.processClass(v)
		if err != nil {
			ff.lg.Error(errMsgs.MappingClassFail, zap.Error(err))
			return err
		}
		parentGroup.Groups[childGroup.UUID] = childGroup
	}

	return nil
}

func (ff *FromFormat) processClass(class formatV2M0.Class) (kbEnt.Group, error) {
	group := ff.mapToGroup(class)

	if err := ff.processParameters(class.Parameters.Parameters, &group); err != nil {
		ff.lg.Error(errMsgs.MappingParametersFail, zap.Error(err))
		return kbEnt.Group{}, err
	}

	if err := ff.processClasses(class.Classes.Classes, &group); err != nil {
		ff.lg.Error(errMsgs.MappingClassesFail, zap.Error(err))
		return kbEnt.Group{}, err
	}

	if err := ff.processRules(class.Rules.Rules, &group); err != nil {
		ff.lg.Error(errMsgs.MappingRulesFail, zap.Error(err))
		return kbEnt.Group{}, err
	}

	if err := ff.processRules(class.Constraints.Constraints, &group); err != nil {
		ff.lg.Error(errMsgs.MappingConstraintsFail, zap.Error(err))
		return kbEnt.Group{}, err
	}

	return group, nil
}

func (ff *FromFormat) mapToGroup(class formatV2M0.Class) kbEnt.Group {
	now := time.Now()
	g := kbEnt.Group{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ff.ws.CreateGroupUUID(),
			ID:           class.ID,
			ShortName:    class.ShortName,
			Description:  class.Description,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}
	g.Groups = map[string]kbEnt.Group{}
	return g
}
