package exporter

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (ex Exporter) processClasses(classes []formatV2M0.Class, gh *kbEnt.GroupsHierarchy, ws workspaceHandler) error {
	for _, v := range classes {
		ghNested, err := ex.processClass(v, ws)
		if err != nil {
			ex.lg.Error(errMsgs.ParsingClassFromXMLFail, zap.Error(err))
			return err
		}
		gh.Contains = append(gh.Contains, ghNested)
	}

	return nil
}

func (ex Exporter) processClass(class formatV2M0.Class, ws workspaceHandler) (kbEnt.GroupsHierarchy, error) {
	g := ex.extractClass(class, ws)
	ws.AddGroup(g)

	if err := ex.processParameters(class.Parameters.Parameters, g.UUID, ws); err != nil {
		ex.lg.Error(errMsgs.ParsingParametersFromXMLFail, zap.Error(err))
		return kbEnt.GroupsHierarchy{}, err
	}

	gh := kbEnt.GroupsHierarchy{
		GroupUUID: g.UUID,
		Contains:  []kbEnt.GroupsHierarchy{},
	}
	if err := ex.processClasses(class.Classes.Classes, &gh, ws); err != nil {
		ex.lg.Error(errMsgs.ParsingClassesFromXMLFail, zap.Error(err))
		return kbEnt.GroupsHierarchy{}, err
	}

	if err := ex.processRules(class.Rules.Rules, ws); err != nil {
		ex.lg.Error(errMsgs.ParsingRulesFromXMLFail, zap.Error(err))
		return kbEnt.GroupsHierarchy{}, err
	}

	if err := ex.processRules(class.Constraints.Constraints, ws); err != nil {
		ex.lg.Error(errMsgs.ParsingConstraintsFromXMLFail, zap.Error(err))
		return kbEnt.GroupsHierarchy{}, err
	}

	return gh, nil
}

func (ex Exporter) extractClass(class formatV2M0.Class, ws workspaceHandler) kbEnt.Group {
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
