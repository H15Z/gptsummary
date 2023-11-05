package dataloader

import "github.com/H15Z/gptsummary/domain/models"

type Loader interface {
	StreamData(callback func(models.Article))
}
