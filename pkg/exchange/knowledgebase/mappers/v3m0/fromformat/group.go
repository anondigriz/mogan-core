package fromformat

import (
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (ff FromFormat) processGroups(groups []formatV3M0.Group, ws workspaceHandler) {
	for _, v := range groups {
		ws.AddGroup(ff.mapToGroup(v, ws))
	}
}

func (ff FromFormat) mapToGroup(group formatV3M0.Group, ws workspaceHandler) kbEnt.Group {
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
	g.Groups = map[string]kbEnt.Group{}

	for _, v := range group.Groups.Groups {
		childGroup := ff.mapToGroup(v, ws)
		g.Groups[childGroup.UUID] = childGroup
	}

	g.Parameters = append(g.Parameters, group.Parameters.Parameters...)
	g.Rules = append(g.Rules, group.Rules.Rules...)

	return g
}
