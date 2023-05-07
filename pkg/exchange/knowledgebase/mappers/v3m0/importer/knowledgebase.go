package importer

import (
	"time"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (im Importer) processKnowledgeBase(kbUUID string, knowledgeBase formatV3M0.KnowledgeBase, ws workspaceHandler) error {
	im.processGroups(knowledgeBase.Groups.Groups, ws)

	ws.AddKnowledgeBase(im.extractKnowledgeBase(kbUUID, knowledgeBase, ws))

	if err := im.processParameters(knowledgeBase.Parameters.Parameters, ws); err != nil {
		im.lg.Error(errMsgs.ParsingParametersFromXMLFail, zap.Error(err))
		return err
	}

	if err := im.processPatterns(knowledgeBase.Patterns.Patterns, ws); err != nil {
		im.lg.Error(errMsgs.ParsingRelationsFromXMLFail, zap.Error(err))
		return err
	}

	if err := im.processRules(knowledgeBase.Rules.Rules, ws); err != nil {
		im.lg.Error(errMsgs.ParsingRulesFromXMLFail, zap.Error(err))
		return err
	}

	return nil
}

func (im Importer) extractKnowledgeBase(kbUUID string, knowledgeBase formatV3M0.KnowledgeBase, ws workspaceHandler) kbEnt.KnowledgeBase {
	k := kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         kbUUID,
			ID:           knowledgeBase.ID,
			ShortName:    knowledgeBase.ShortName,
			Description:  knowledgeBase.Description,
			CreatedDate:  time.Unix(knowledgeBase.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(knowledgeBase.ModifiedDate, 0).UTC(),
		},
		GroupsHierarchy: im.extractGroupsHierarchy(knowledgeBase.GroupsHierarchy, ws),
	}

	return k
}

func (im Importer) extractGroupsHierarchy(gh formatV3M0.GroupsHierarchy, ws workspaceHandler) kbEnt.GroupsHierarchy {
	g := kbEnt.GroupsHierarchy{
		GroupUUID: ws.GetOrCreateGroupUUID(gh.GroupID),
		Contains:  []kbEnt.GroupsHierarchy{},
	}

	for _, v := range gh.Contains.GroupsHierarchies {
		g.Contains = append(g.Contains, im.extractGroupsHierarchy(v, ws))
	}

	return g
}