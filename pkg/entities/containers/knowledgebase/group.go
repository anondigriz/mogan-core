package knowledgebase

type Group struct {
	BaseInfo
	Groups     map[string]Group
	Parameters []string
	Rules      []string
}
