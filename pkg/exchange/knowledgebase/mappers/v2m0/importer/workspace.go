package importer

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	uuidGen "github.com/google/uuid"
)

type idToUUIDmap struct {
	parameters map[string]string
	patterns   map[string]string
}

type workspaceHandler interface {
	AddKnowledgeBase(knowledgeBase kbEnt.KnowledgeBase)
	AddParameter(parameter kbEnt.Parameter)
	AddPattern(pattern kbEnt.Pattern)
	AddGroup(group kbEnt.Group)
	AddRule(rule kbEnt.Rule)
	AddGroupHierarchy(gh kbEnt.GroupsHierarchy)
	CreateGroupUUID() string
	CreatePatternUUID(id string) string
	CreateRuleUUID() string
	GetOrCreateParameterUUID(id string) string
	GetPatternUUID(relationID string) (string, bool)
}

type workspace struct {
	mapping idToUUIDmap
	cont    kbEnt.Container
}

func newWorkspace() *workspace {
	ws := &workspace{}
	ws.cont.Groups = map[string]kbEnt.Group{}
	ws.cont.Parameters = map[string]kbEnt.Parameter{}
	ws.cont.Patterns = map[string]kbEnt.Pattern{}
	ws.cont.Rules = map[string]kbEnt.Rule{}
	ws.mapping.parameters = map[string]string{}
	ws.mapping.patterns = map[string]string{}
	return ws
}

func (ws *workspace) AddKnowledgeBase(knowledgeBase kbEnt.KnowledgeBase) {
	ws.cont.KnowledgeBase = knowledgeBase
}

func (ws *workspace) AddParameter(parameter kbEnt.Parameter) {
	ws.cont.Parameters[parameter.UUID] = parameter
}

func (ws *workspace) AddPattern(pattern kbEnt.Pattern) {
	ws.cont.Patterns[pattern.UUID] = pattern
}

func (ws *workspace) AddGroup(group kbEnt.Group) {
	ws.cont.Groups[group.UUID] = group
}

func (ws *workspace) AddRule(rule kbEnt.Rule) {
	ws.cont.Rules[rule.UUID] = rule
}

func (ws *workspace) AddGroupHierarchy(gh kbEnt.GroupsHierarchy) {
	ws.cont.KnowledgeBase.GroupsHierarchy = gh
}

func (ws *workspace) CreateGroupUUID() string {
	uuid := uuidGen.NewString()
	return uuid
}

func (ws *workspace) CreatePatternUUID(id string) string {
	uuid := uuidGen.NewString()
	ws.mapping.patterns[id] = uuid
	return uuid
}

func (ws *workspace) CreateRuleUUID() string {
	uuid := uuidGen.NewString()
	return uuid
}
func (ws *workspace) GetOrCreateParameterUUID(id string) string {
	if uuid, ok := ws.mapping.parameters[id]; ok {
		return uuid
	}

	uuid := uuidGen.NewString()
	ws.mapping.parameters[id] = uuid
	return uuid
}

func (ws *workspace) GetPatternUUID(relationID string) (string, bool) {
	uuid, ok := ws.mapping.patterns[relationID]
	return uuid, ok
}
