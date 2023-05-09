package toformat

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
)

type identifyInfo struct {
	UUID      string
	ID        string
	ShortName string
}

type idToUUIDmap struct {
	groups     map[string]identifyInfo
	parameters map[string]identifyInfo
	patterns   map[string]identifyInfo
	rules      map[string]identifyInfo
}

type workspaceHandler interface {
	CheckAndRememberGroup(group kbEnt.Group) error
	CheckAndRememberParameter(parameter kbEnt.Parameter) error
	CheckAndRememberPattern(pattern kbEnt.Pattern) error
	CheckAndRememberRule(rule kbEnt.Rule) error
}

type workspace struct {
	mapping idToUUIDmap
}

func newWorkspace() *workspace {
	ws := &workspace{}

	ws.mapping.groups = map[string]identifyInfo{}
	ws.mapping.parameters = map[string]identifyInfo{}
	ws.mapping.patterns = map[string]identifyInfo{}
	ws.mapping.rules = map[string]identifyInfo{}

	return ws
}

func (ws *workspace) CheckAndRememberGroup(group kbEnt.Group) error {
	if v, ok := ws.mapping.groups[group.ID]; ok {
		return errors.NewNonUniqueIDFoundErr(
			errors.NonUniqueIDFoundArgs{
				ID:              group.ID,
				FirstUUID:       v.UUID,
				FirstShortName:  v.ShortName,
				SecondUUID:      group.ID,
				SecondShortName: group.ShortName,
			})
	}

	ws.mapping.groups[group.ID] = identifyInfo{
		ID:        group.ID,
		UUID:      group.UUID,
		ShortName: group.ShortName,
	}
	return nil
}

func (ws *workspace) CheckAndRememberParameter(parameter kbEnt.Parameter) error {
	if v, ok := ws.mapping.parameters[parameter.ID]; ok {
		return errors.NewNonUniqueIDFoundErr(
			errors.NonUniqueIDFoundArgs{
				ID:              parameter.ID,
				FirstUUID:       v.UUID,
				FirstShortName:  v.ShortName,
				SecondUUID:      parameter.ID,
				SecondShortName: parameter.ShortName,
			})
	}

	ws.mapping.parameters[parameter.ID] = identifyInfo{
		ID:        parameter.ID,
		UUID:      parameter.UUID,
		ShortName: parameter.ShortName,
	}
	return nil
}

func (ws *workspace) CheckAndRememberPattern(pattern kbEnt.Pattern) error {
	if v, ok := ws.mapping.patterns[pattern.ID]; ok {
		return errors.NewNonUniqueIDFoundErr(
			errors.NonUniqueIDFoundArgs{
				ID:              pattern.ID,
				FirstUUID:       v.UUID,
				FirstShortName:  v.ShortName,
				SecondUUID:      pattern.ID,
				SecondShortName: pattern.ShortName,
			})
	}

	ws.mapping.patterns[pattern.ID] = identifyInfo{
		ID:        pattern.ID,
		UUID:      pattern.UUID,
		ShortName: pattern.ShortName,
	}
	return nil
}

func (ws *workspace) CheckAndRememberRule(rule kbEnt.Rule) error {
	if v, ok := ws.mapping.rules[rule.ID]; ok {
		return errors.NewNonUniqueIDFoundErr(
			errors.NonUniqueIDFoundArgs{
				ID:              rule.ID,
				FirstUUID:       v.UUID,
				FirstShortName:  v.ShortName,
				SecondUUID:      rule.ID,
				SecondShortName: rule.ShortName,
			})
	}

	ws.mapping.rules[rule.ID] = identifyInfo{
		ID:        rule.ID,
		UUID:      rule.UUID,
		ShortName: rule.ShortName,
	}
	return nil
}
