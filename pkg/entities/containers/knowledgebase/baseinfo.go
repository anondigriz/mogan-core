package knowledgebase

import "time"

type BaseInfo struct {
	UUID         string
	ID           string
	ShortName    string
	Description  string
	CreatedDate  time.Time
	ModifiedDate time.Time
}
