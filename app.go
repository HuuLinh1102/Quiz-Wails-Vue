package main

import (
	"context"
	"embed"
	"fmt"
	"strings"

	"changeme/app/model"
)

//go:embed questions.txt
var questionsFile embed.FS

// App struct
type App struct {
	ctx       context.Context
	questions []model.Question // Dữ liệu câu hỏi
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// LoadQuestions loads questions from questions.txt file
func (a *App) LoadQuestions() ([]model.Question, error) {
	data, err := questionsFile.ReadFile("questions.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var questions []model.Question
	var q model.Question
	for _, line := range lines {
		if strings.HasPrefix(line, "Đáp án:") {
			q.Answer = strings.TrimSpace(line[8:])
			questions = append(questions, q)
			q = model.Question{}
		} else if len(line) > 0 && line[1] == '.' {
			choice := model.Choice{Content: line[3:]}
			q.Choices = append(q.Choices, choice)
		} else {
			q.Content = line
		}
	}

	return questions, nil
}

// GetQuestions returns the loaded questions
func (a *App) GetQuestions() []model.Question {
	return a.questions
}

// SetQuestions sets the loaded questions
func (a *App) SetQuestions(questions []model.Question) {
	a.questions = questions
}
