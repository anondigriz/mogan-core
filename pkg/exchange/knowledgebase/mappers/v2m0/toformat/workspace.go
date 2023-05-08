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

type processingMarks struct {
	rules      map[string]bool
	parameters map[string]bool
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
	SaveUnprocessedParameters(parameters []string)
	SaveUnprocessedRules(rules []string)
	RemoveProcessedParameters(uuid string)
	RemoveProcessedRules(uuid string)
	GetUnprocessedParameters() []string
	GetUnprocessedRules() []string
}

type workspace struct {
	mapping idToUUIDmap
	marks   processingMarks
}

func newWorkspace() *workspace {
	ws := &workspace{}

	ws.mapping.groups = map[string]identifyInfo{}
	ws.mapping.parameters = map[string]identifyInfo{}
	ws.mapping.patterns = map[string]identifyInfo{}
	ws.mapping.rules = map[string]identifyInfo{}

	ws.marks.parameters = map[string]bool{}
	ws.marks.rules = map[string]bool{}

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

func (ws *workspace) SaveUnprocessedParameters(parameters []string) {
	for _, v := range parameters {
		ws.marks.parameters[v] = true
	}
}

func (ws *workspace) SaveUnprocessedRules(rules []string) {
	for _, v := range rules {
		ws.marks.parameters[v] = true
	}
}

func (ws *workspace) RemoveProcessedParameters(uuid string) {
	ws.marks.parameters[uuid] = true
}

func (ws *workspace) RemoveProcessedRules(uuid string) {
	ws.marks.rules[uuid] = true
}

func (ws workspace) GetUnprocessedParameters() []string {
	result := []string{}
	for k, v := range ws.marks.parameters {
		if v {
			result = append(result, k)
		}
	}
	return result
}

func (ws workspace) GetUnprocessedRules() []string {
	result := []string{}
	for k, v := range ws.marks.rules {
		if v {
			result = append(result, k)
		}
	}
	return result
}
