package dataloader

import (
	"github.com/H15Z/gptsummary/domain/models"
)

type DataLoader struct {
	Loader Loader
}

func NewDataLoader(loader Loader) *DataLoader {
	return &DataLoader{
		Loader: loader,
	}
}

func (d DataLoader) StreamData(callback func(models.Article)) {
	d.Loader.StreamData(callback)
}
