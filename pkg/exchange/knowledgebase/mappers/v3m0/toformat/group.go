package toformat

import (
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	formatV3M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v3m0"
)

func (tf ToFormat) mapToGroups(groups map[string]kbEnt.Group, ws workspaceHandler) ([]formatV3M0.Group, error) {
	gs := []formatV3M0.Group{}
	for _, v := range groups {
		g, err := tf.mapToGroup(v, ws)
		if err != nil {
			tf.lg.Error(errMsgs.MappingGroupFail, zap.Error(err))
			return []formatV3M0.Group{}, err
		}
		gs = append(gs, g)
	}

	return gs, nil
}

func (tf ToFormat) mapToGroup(group kbEnt.Group, ws workspaceHandler) (formatV3M0.Group, error) {
	if err := ws.CheckAndRememberGroup(group); err != nil {
		tf.lg.Error(errMsgs.MappingGroupFail, zap.Error(err))
		return formatV3M0.Group{}, err
	}

	g := formatV3M0.Group{
		BaseInfo: formatV3M0.BaseInfo{
			ID:           group.BaseInfo.ID,
			ShortName:    group.BaseInfo.ShortName,
			Description:  group.BaseInfo.Description,
			CreatedDate:  group.CreatedDate.UTC().Unix(),
			ModifiedDate: group.ModifiedDate.UTC().Unix(),
		},
	}
	g.Parameters.Parameters = append(g.Parameters.Parameters, group.Parameters...)
	g.Rules.Rules = append(g.Rules.Rules, group.Rules...)

	gs, err := tf.mapToGroups(group.Groups, ws)
	if err != nil {
		tf.lg.Error(errMsgs.MappingGroupsFail, zap.Error(err))
		return formatV3M0.Group{}, err
	}
	g.Groups.Groups = gs

	return g, nil
}
