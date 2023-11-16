package req

type Patch struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Requester int    `json:"requester_id"` // 요청하는사람 (요청하는 사람과 작성자가 같아야 함)
}
