package main

import (
	"changeme/app/model"
	"changeme/app/modules/auth"

	"bufio"
	"context"
	"embed"
	"fmt"
	"os"
	"strings"
)

//go:embed questions.txt
var questionsFile embed.FS

// App struct
type App struct {
	ctx       context.Context
	questions []model.Question
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

func (a *App) Login(idnum string) (res model.LoginResponse) {
	ls := auth.LoginService{}

	user, err := ls.Login(idnum)
	if err != nil {
		res = model.LoginResponse{
			Message: err.Error(),
			Status:  false,
		}
		return
	}
	res = model.LoginResponse{
		Message: "Đăng nhập thành công",
		Status:  true,
		User:    &user,
	}
	return
}

// LoadQuestions loads questions from questions.txt file
func (a *App) LoadQuestions() ([]model.Question, error) {
	file, err := os.Open("./questions.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	questions := []model.Question{}
	question := model.Question{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Đáp án:") {
			question.Answer = strings.TrimSpace(line[7:])
			questions = append(questions, question)
			question = model.Question{}
		} else if strings.HasPrefix(line, "A.") || strings.HasPrefix(line, "B.") || strings.HasPrefix(line, "C.") || strings.HasPrefix(line, "D.") {
			question.Choices = append(question.Choices, model.Choice{Content: strings.TrimSpace(line[2:])})
		} else {
			question.Content = line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

// CheckAnswer checks the answer of the given question from questions.txt file
func (a *App) CheckAnswers(questions []model.Question, answers []string) int {
    correctCount := 0
    for i, question := range questions {
        if question.Answer == answers[i] {
            correctCount++
        }
    }
    return correctCount
}

// GetQuestions returns the loaded questions
func (a *App) GetQuestions() []model.Question {
	return a.questions
}

// SetQuestions sets the loaded questions
func (a *App) SetQuestions(questions []model.Question) {
	a.questions = questions
}
