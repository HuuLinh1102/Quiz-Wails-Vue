package model

import (
	"gorm.io/gorm"
)

// Model Đề thi
type ExamQues struct {
	gorm.Model
	ExamID    uint `json:"exam_id"`    // Mã đề
	NameExam  string `json:"name_exam"` // Tên đề
	ExamTime  uint `json:"exam_time"`  // Thời gian làm bài
	TimeStart uint `json:"time_start"` // Thời gian bắt đầu làm bài
	TimeEnd   uint `json:"time_end"`   // Thời gian kết thúc làm bài
	Score     uint `json:"score"`      // Điểm số đạt được
	EasyQuest uint `json:"easy_quest"` // Số câu hỏi dễ
	MediumQuest uint `json:"medium_quest"` // Số câu hỏi trung bình
	HardQuest uint `json:"hard_quest"` // Số câu hỏi khó
}

// Model Câu hỏi
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
