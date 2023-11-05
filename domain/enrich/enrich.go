package enrich

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/H15Z/gptsummary/domain/models"
)

// CORE STRUCTS FOR ENRICH ACTOR
type Enricher struct {
	GPT GPT
}

func NewEnricher(gpt GPT) *Enricher {
	return &Enricher{
		GPT: gpt,
	}
}

func (e Enricher) Summarize(article *models.Article) {

	prompt := e.GeneratePrompt(article.Text)
	reponse := e.GPT.QueryGPT(prompt)

	r := e.ProcessResponse(reponse)

	article.Summary = r.Summary
	article.Sentiment = r.Sentiment
	article.Category = r.Category
}

// Generate prompt from article text
func (e Enricher) GeneratePrompt(article_text string) string {
	return `Summorize the  text  to max 30 words, provide sentiment analysis (values: positive, negative or netural), assign a bussiness category,respond in the JSON format {"summary":"text summary", "sentiment":"value","category":"value"} : \n` + article_text
}

func (e Enricher) ProcessResponse(reponse models.Response) models.ResponseContent {
	var data models.ResponseContent

	content := reponse.Choices[0].Message.Content
	content = strings.Replace(content, "\\", "", -1)

	if err := json.Unmarshal([]byte(content), &data); err != nil {
		log.Printf("Error unmarshaling JSON: %v\n", err)
		return models.ResponseContent{}
	}

	return data
}
