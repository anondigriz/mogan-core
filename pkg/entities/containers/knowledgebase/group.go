package knowledgebase

type Group struct {
	BaseInfo
	Groups     []Group
	Parameters []string
	Rules      []string
}
