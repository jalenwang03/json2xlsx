package json2xlsx

import ()

type completion struct {
	Is_completed int `json:"is_completed"`
}
type due_date struct {
	Time     string  `json:"time"`
	Due_date float64 `json:"date"`
}
type attachments struct {
	Id  string `json:"_id"`
	Url string `json:"url"`
}
type task_interface struct {
	Id                    string        `json:"_id"`
	Team                  string        `json:"team"`
	Project               string        `json:"project"`
	Parent                string        `json:"entry"`
	Identifier            int           `json:"identifier"`
	Title                 string        `json:"title"`
	Description           string        `json:"description"`
	Assignment            string        `json:"assignment"`
	Watchers              string        `json:"watcher"`
	Due_date              due_date      `json:"due_date"`
	Tags                  []string      `json:"tags"`
	Attachments           []attachments `json:"attachments"`
	Completion            completion    `json:"completion"`
	Visibility            string        `json:"visibility"`
	Extensions            []string      `json:"extensions"`
	Is_archived           int           `json:"is_archived"`
	Priority              string        `json:"priority"`
	Is_deleted            int           `json:"is_deleted"`
	Created_at            float64       `json:"created_at"`
	Created_by            string        `json:"created_by"`
	Updated_at            float64       `json:"updated_at"`
	Updated_by            string        `json:"updated_by"`
	Is_cascading_deleted  int           `json:"is_cascading_deleted"`
	Is_cascading_archived int           `json:"is_cascading_archived"`
	Comments              []string      `json:"comments"`
}
type from struct {
	Message_type string `json:"type"`
	Uid          string `json:"uid"`
}
type body struct {
	Content           string `json:"content"`
	Markdown          int    `json:"markdown"`
	Style             int    `json:"style"`
	Attachment        string `json:"attachment"`
	Inline_attachment inline `json:"inline_attachment"`
}
type inline struct {
	Fields      []field `json:"fields"`
	Text        string  `json:"text"`
	Title_link  string  `json:"title_link"`
	Title       string  `json:"title"`
	Author_icon string  `json:"author_icon"`
	Author_link string  `json:"author_link"`
	Author_name string  `json:"author_name"`
	Pretext     string  `json:"pretext"`
	Color       string  `json:"color"`
	Fallback    string  `json:"fallback"`
}
type field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short int    `json:"short"`
}
type message_interface struct {
	Message_type string  `json:"type"`
	Created_at   float64 `json:"created_at"`
	Updated_at   float64 `json:"updated_at"`
	From         from    `json:"from"`
	Body         body    `json:"body"`
}

type user_interface struct {
	Uid          string  `json:"uid"`
	Name         string  `json:"name"`
	Display_name string  `json:"display_name"`
	Email        string  `json:"email"`
	Avatar       string  `json:"avatar"`
	Short_code   int     `json:"short_code"`
	Desc         string  `json:"desc"`
	Location     string  `json:"localtion"`
	Mobile       int     `json:"mobile"`
	Created_at   float64 `json:"created_at"`
	Updated_at   float64 `json:"updated_at"`
	Role         string  `json:"role"`
	Status       string  `json:"status"`
}

type members_interface struct {
	Uid        string `json:"uid"`
	ID         string `json:"_id"`
	Permission int    `json:"permission"`
}

type project_interface struct {
	Team           string              `json:"team"`
	Name           string              `json:"name"`
	Description    string              `json:"description"`
	Color          string              `json:"color"`
	Mambers        []members_interface `json:"members"`
	Visibility     int                 `json:"visibility"`
	Permission     int                 `json:"permission"`
	Is_archived    int                 `json:"is_archived"`
	Is_deleted     int                 `json:"is_deleted"`
	Created_at     float64             `json:"created_at"`
	Created_by     string              `json:"created_by"`
	Updated_at     float64             `json:"updated_at"`
	Updated_by     string              `json:"updated_by"`
	Archived_at    float64             `json:"archived_at"`
	Status_history string              `json:"status_history"`
	Entries        []string            `json:"entries"`
}

type groups_interface struct {
	Id         string   `json:"_id"`
	Name       string   `json:"name"`
	Desc       string   `json:"desc"`
	Color      string   `json:"color"`
	Is_system  int      `json:"is_system"`
	Created_at float64  `json:"created_at"`
	Created_by string   `json:"created_by"`
	Updated_at float64  `json:"updated_at"`
	Updated_by string   `json:"updated_by"`
	Members    []string `json:"members"`
	Visibility string   `json:"visibility"`
	Status     string   `json:"status"`
}
type scope_interface struct {
	Visibility string              `json:"visibility"`
	Permission string              `json:"permission"`
	Members    []members_interface `json:"members"`
}
type addition_interface struct {
	Ext  string `json:"ext"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}
type drivers_interface struct {
	Id           string             `json:"_id"`
	Title        string             `json:"title"`
	Created_at   float64            `json:"created_at"`
	Created_by   string             `json:"created_by"`
	Updated_at   float64            `json:"updated_at"`
	Updated_by   string             `json:"updated_by"`
	Drivers_type string             `json:"type"`
	Scope        scope_interface    `json:"scope"`
	Addition     addition_interface `json:"addition"`
	Tags         []string           `json:"tags"`
}
