package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"recything/features/recybot/entity"
	"recything/utils/validation"
	"strings"
	"recything/utils/constanta"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type recybotService struct {
	recybotRepository entity.RecybotRepositoryInterface
}

func NewRecybotService(recybot entity.RecybotRepositoryInterface) entity.RecybotServiceInterface {
	return &recybotService{
		recybotRepository: recybot,
	}
}

// CreateData implements entity.RecybotServiceInterface.
func (rb *recybotService) CreateData(data entity.RecybotCore) (entity.RecybotCore, error) {

	errEmpty := validation.CheckDataEmpty(data.Category, data.Question)
	if errEmpty != nil {
		return entity.RecybotCore{}, errEmpty
	}

	validCategory, errCategory := validation.CheckEqualData(data.Category, constanta.Category)
	if errCategory != nil {
		return entity.RecybotCore{}, errCategory
	}

	data.Category = validCategory
	result, err := rb.recybotRepository.Create(data)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (rb *recybotService) GetAllData() ([]entity.RecybotCore, error) {
	result, err := rb.recybotRepository.GetAll()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (rb *recybotService) GetById(idData string) (entity.RecybotCore, error) {
	result, err := rb.recybotRepository.GetById(idData)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Delete implements entity.RecybotServiceInterface.
func (rb *recybotService) DeleteData(idData string) error {

	err := rb.recybotRepository.Delete(idData)
	if err != nil {
		return err
	}
	return nil
}

// UpdateData implements entity.RecybotServiceInterface.
func (rb *recybotService) UpdateData(idData string, data entity.RecybotCore) (entity.RecybotCore, error) {

	errEmpty := validation.CheckDataEmpty(data.Category, data.Question)
	if errEmpty != nil {
		return entity.RecybotCore{}, errEmpty
	}

	validCategory,errCategory := validation.CheckEqualData(data.Category, constanta.Category)
	if errCategory != nil {
		return entity.RecybotCore{}, errCategory
	}

	data.Category = validCategory
	result, err := rb.recybotRepository.Update(idData, data)
	if err != nil {
		return result, err
	}
	result.ID = idData
	return result, nil
}

func (rb *recybotService) GetPrompt(question string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataRecybot, err := rb.recybotRepository.GetAll()
	if err != nil {
		return "", err
	}

	output := make(map[string][]string)
	for _, item := range dataRecybot {
		output[item.Category] = append(output[item.Category], item.Question)
	}

	var prompt strings.Builder
	for category, questions := range output {
		prompt.WriteString(fmt.Sprintln(" "))
		prompt.WriteString(fmt.Sprintf("kategori %s:\n", category))
		for _, question := range questions {
			prompt.WriteString(fmt.Sprintf("%s\n", question))
		}
	}
	log.Println(prompt.String())
	ctx := context.Background()
	client := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: prompt.String(),
		},
		{
			Role:    "user",
			Content: question,
		},
	}

	response, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	if err != nil {
		return "", err
	}

	answer := response.Choices[0].Message.Content
	return answer, nil
}
