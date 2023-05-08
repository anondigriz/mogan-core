package toformat

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

type processGroupsArgs struct {
	groups     map[string]kbEnt.Group
	parameters map[string]kbEnt.Parameter
	rules      map[string]kbEnt.Rule
	model      *formatV2M0.Class
	ws         workspaceHandler
}

func (tf ToFormat) processGroups(args processGroupsArgs) error {
	return nil
}
