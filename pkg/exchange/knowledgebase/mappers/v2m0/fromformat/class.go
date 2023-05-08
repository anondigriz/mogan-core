package fromformat

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ff FromFormat) processClasses(classes []formatV2M0.Class, parentGroup *kbEnt.Group, ws workspaceHandler) error {
	for _, v := range classes {
		childGroup, err := ff.processClass(v, ws)
		if err != nil {
			ff.lg.Error(errMsgs.ParsingClassFromXMLFail, zap.Error(err))
			return err
		}
		parentGroup.Groups = append(parentGroup.Groups, childGroup)
	}

	return nil
}

func (ff FromFormat) processClass(class formatV2M0.Class, ws workspaceHandler) (kbEnt.Group, error) {
	group := ff.extractGroup(class, ws)

	if err := ff.processParameters(class.Parameters.Parameters, &group, ws); err != nil {
		ff.lg.Error(errMsgs.ParsingParametersFromXMLFail, zap.Error(err))
		return kbEnt.Group{}, err
	}

	if err := ff.processClasses(class.Classes.Classes, &group, ws); err != nil {
		ff.lg.Error(errMsgs.ParsingClassesFromXMLFail, zap.Error(err))
		return kbEnt.Group{}, err
	}

	if err := ff.processRules(class.Rules.Rules, &group, ws); err != nil {
		ff.lg.Error(errMsgs.ParsingRulesFromXMLFail, zap.Error(err))
		return kbEnt.Group{}, err
	}

	if err := ff.processRules(class.Constraints.Constraints, &group, ws); err != nil {
		ff.lg.Error(errMsgs.ParsingConstraintsFromXMLFail, zap.Error(err))
		return kbEnt.Group{}, err
	}

	return group, nil
}

func (ff FromFormat) extractGroup(class formatV2M0.Class, ws workspaceHandler) kbEnt.Group {
	now := time.Now()
	g := kbEnt.Group{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ws.CreateGroupUUID(),
			ID:           class.ID,
			ShortName:    class.ShortName,
			Description:  class.Description,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}
	return g
}
