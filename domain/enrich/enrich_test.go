package enrich

import (
	"testing"

	"github.com/H15Z/gptsummary/domain/models"
	"github.com/go-playground/assert/v2"
)

func TestResponseParser(t *testing.T) {

	response := models.Response{
		Choices: []models.Choice{
			{
				Message: models.Message{
					Content: `{\"summary\": \"Training your brain through smell can strengthen neural pathways. Olfactory system is plastic, repairs and renews. Ongoing regeneration of sensory cells enables smell detection.\", \"sentiment\": \"positive\", \"category\": \"Brain Training\"}`,
				},
			},
		},
	}

	gpt := MockGPT{}
	enricher := NewEnricher(gpt)

	r := enricher.ProcessResponse(response)
	assert.Equal(t, r.Summary, "Training your brain through smell can strengthen neural pathways. Olfactory system is plastic, repairs and renews. Ongoing regeneration of sensory cells enables smell detection.")
	assert.Equal(t, r.Sentiment, "positive")
	assert.Equal(t, r.Category, "Brain Training")

}

func TestArticleEnrichment(t *testing.T) {

	gpt := MockGPT{}
	enricher := NewEnricher(gpt)

	test_art := models.Article{
		Title: "Title",
		Text:  `When it comes to training your brain, your sense of smell is possibly the last thing you’d think could strengthen your neural pathways. Learning a new language or reading more books (and fewer social media posts) — sure. But your nose?`,
	}

	enricher.Summarize(&test_art)

	assert.Equal(t, test_art.Summary, "Training your brain through smell can strengthen neural pathways. Olfactory system is plastic, repairs and renews. Ongoing regeneration of sensory cells enables smell detection.")
	assert.Equal(t, test_art.Sentiment, "positive")
	assert.Equal(t, test_art.Category, "Brain Training")

}

type MockGPT struct {
}

func (m MockGPT) QueryGPT(prompt string) models.Response {
	return models.Response{
		Choices: []models.Choice{
			{
				Message: models.Message{
					Content: `{\"summary\": \"Training your brain through smell can strengthen neural pathways. Olfactory system is plastic, repairs and renews. Ongoing regeneration of sensory cells enables smell detection.\", \"sentiment\": \"positive\", \"category\": \"Brain Training\"}`,
				},
			},
		},
	}
}
