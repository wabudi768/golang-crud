package schemas

type Student struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Npm      uint64   `json:"npm"`
	Fak      string   `json:"fak"`
	Bid      string   `json:"bid"`
	Teachers []string `json:teacher_id`
}
