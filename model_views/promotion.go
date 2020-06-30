package model_views

type Promotion struct {
	Id           int       `json:"id"`
	Name         string    `json:"name" description:"名称"`
	Title        string    `json:"title" description:"标题"`
	DeleteFlag   int8      `json:"delete_flag" description:"删除标记"`
}
