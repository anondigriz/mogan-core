package knowledgebase

type KnowledgeBase struct {
	BaseInfo
	GroupsHierarchy GroupsHierarchy
}

type GroupsHierarchy struct {
	GroupUUID string
	Contains  []GroupsHierarchy
}
