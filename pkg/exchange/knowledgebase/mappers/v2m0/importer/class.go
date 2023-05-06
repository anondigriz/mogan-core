package importer

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (im Importer) processClasses(classes []formatV2M0.Class, gh *kbEnt.GroupsHierarchy, ws workspaceHandler) error {
	for _, v := range classes {
		ghNested, err := im.processClass(v, ws)
		if err != nil {
			im.lg.Error(errMsgs.ParsingClassFromXMLFail, zap.Error(err))
			return err
		}
		gh.Contains = append(gh.Contains, ghNested)
	}

	return nil
}

func (im Importer) processClass(class formatV2M0.Class, ws workspaceHandler) (kbEnt.GroupsHierarchy, error) {
	g := im.extractClass(class, ws)
	ws.AddGroup(g)

	if err := im.processParameters(class.Parameters.Parameters, g.UUID, ws); err != nil {
		im.lg.Error(errMsgs.ParsingParametersFromXMLFail, zap.Error(err))
		return kbEnt.GroupsHierarchy{}, err
	}

	gh := kbEnt.GroupsHierarchy{
		GroupUUID: g.UUID,
		Contains:  []kbEnt.GroupsHierarchy{},
	}
	if err := im.processClasses(class.Classes.Classes, &gh, ws); err != nil {
		im.lg.Error(errMsgs.ParsingClassesFromXMLFail, zap.Error(err))
		return kbEnt.GroupsHierarchy{}, err
	}

	err := im.processRules(class.Rules.Rules, ws)
	if err != nil {
		im.lg.Error(errMsgs.ParsingRulesFromXMLFail, zap.Error(err))
		return kbEnt.GroupsHierarchy{}, err
	}

	err = im.processRules(class.Constraints.Constraints, ws)
	if err != nil {
		im.lg.Error(errMsgs.ParsingConstraintsFromXMLFail, zap.Error(err))
		return kbEnt.GroupsHierarchy{}, err
	}

	return gh, nil
}

func (im Importer) extractClass(class formatV2M0.Class, ws workspaceHandler) kbEnt.Group {
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
