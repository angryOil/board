package res

import "board/internal/controller/res"

type ListTotalDto struct {
	Content []res.Info `json:"content"`
	Total   int        `json:"total"`
}
