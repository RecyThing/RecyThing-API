package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"recything/features/recybot/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"
	"strconv"

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

func (rb *recybotService) FindAllData(page, category, limit string) ([]entity.RecybotCore, pagination.PageInfo, int, error) {
	log.Println("page service sebelum validasi", page)
	// pageInt, limitInt, err := validation.ValidateTypePaginationParameters(limit, page)
	var limitInt int
	var pageInt int
	var err error
	if limit == "" {
		limitInt = 10
	}
	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return nil, pagination.PageInfo{}, 0, errors.New("limit harus berupa angka")
		}
	}

	if page == "" {
		pageInt = 1
	}
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return nil, pagination.PageInfo{}, 0, errors.New("page harus berupa angka")
		}
	}

	if pageInt <= 0 {
		pageInt = 1
	}

	maxLimit := 10

	if limitInt <= 0 || limitInt > maxLimit {
		limitInt = maxLimit
	}

	log.Println("limitservice : ", limitInt)
	log.Println("pageservice :", pageInt)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	// validPage, validLimit := validation.ValidatePaginationParameters(pageInt, limitInt)
	result, pagnationInfo, count, err := rb.recybotRepository.FindAll(pageInt, limitInt, category)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}
	return result, pagnationInfo, count, nil
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

	validCategory, errCategory := validation.CheckEqualData(data.Category, constanta.Category)
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

	var prompt string
	for category, questions := range output {
		prompt += "\n" + fmt.Sprintf("kategori %s:\n", category)
		for _, question := range questions {
			prompt += fmt.Sprintf("%s\n", question)
		}
	}

	log.Println(prompt)
	ctx := context.Background()
	client := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: prompt,
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
