package model

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Content string   `json:"content"` // Nội dung câu hỏi
	Choices []Choice `json:"choices"` // Các lựa chọn trả lời
	Answer  string   `json:"answer"`  // Đáp án đúng
}

type Choice struct {
	gorm.Model
	QuestionID uint   `json:"question_id"` // ID của câu hỏi mà lựa chọn này thuộc về
	Content    string `json:"content"`     // Nội dung lựa chọn
}
