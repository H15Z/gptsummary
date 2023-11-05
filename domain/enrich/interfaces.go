package enrich

import "github.com/H15Z/gptsummary/domain/models"

type GPT interface {
	QueryGPT(prompt string) models.Response
}
