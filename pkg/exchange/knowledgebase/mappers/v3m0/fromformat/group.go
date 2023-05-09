package fromformat

import (
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (ff *FromFormat) processGroups(groups []formatV3M0.Group) {
	for _, v := range groups {
		ff.ws.AddGroup(ff.mapToGroup(v))
	}
}

func (ff *FromFormat) mapToGroup(group formatV3M0.Group) kbEnt.Group {
	g := kbEnt.Group{
		BaseInfo: kbEnt.BaseInfo{
			UUID:         ff.ws.CreateGroupUUID(),
			ID:           group.ID,
			ShortName:    group.ShortName,
			Description:  group.Description,
			CreatedDate:  time.Unix(group.CreatedDate, 0).UTC(),
			ModifiedDate: time.Unix(group.ModifiedDate, 0).UTC(),
		},
	}
	g.Groups = map[string]kbEnt.Group{}

	for _, v := range group.Groups.Groups {
		childGroup := ff.mapToGroup(v)
		g.Groups[childGroup.UUID] = childGroup
	}

	for _, v := range group.Parameters.Parameters {
		uuid, ok := ff.ws.GetParameterUUID(v)
		if !ok {
			continue
		}
		g.Parameters = append(g.Parameters, uuid)
	}

	for _, v := range group.Rules.Rules {
		uuid, ok := ff.ws.GetRuleUUID(v)
		if !ok {
			continue
		}
		g.Rules = append(g.Rules, uuid)
	}

	return g
}
