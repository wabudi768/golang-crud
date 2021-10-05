package controllers

type Teacher struct {
	ID []int `json:"teacher_id"`
}

type InputStudent struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Npm       uint64 `json:"npm"`
	Fak       string `json:"fak"`
	Bid       string `json:"bid"`
	TeacherID []Teacher
}
