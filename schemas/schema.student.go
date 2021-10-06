package schemas

type Student struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Npm       uint64 `json:"npm"`
	Fak       string `json:"fak"`
	Bid       string `json:"bid"`
	TeacherId uint   `json:"teacher_id"`
}
