package importer

import (
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (im Importer) processGroups(groups []formatV3M0.Group, ws workspaceHandler) {
	for _, v := range groups {
		ws.AddGroup(im.extractGroup(v, ws))
	}
}

func (im Importer) extractGroup(group formatV3M0.Group, ws workspaceHandler) kbEnt.Group {
	g := kbEnt.Group{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ws.CreateGroupUUID(),
			ID:           group.ID,
			ShortName:    group.ShortName,
			Description:  group.Description,
			CreatedDate:  time.Unix(group.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(group.ModifiedDate, 0).UTC(),
		},
	}

	for _, v := range group.Groups.Groups {
		g.Groups = append(g.Groups, im.extractGroup(v, ws))
	}

	for _, v := range group.Parameters.Parameters {
		g.Parameters = append(g.Parameters, v)
	}

	for _, v := range group.Rules.Rules {
		g.Rules = append(g.Rules, v)
	}
	return g
}
