package actors

import (
	"fmt"
	"log"

	"github.com/H15Z/gptsummary/domain/enrich"
	"github.com/H15Z/gptsummary/domain/models"
	"github.com/H15Z/gptsummary/pkg/chatgpt"
)

type EnrichMsg struct {
	Article models.Article
}

type EnricherActor struct {
	Enricher enrich.Enricher
	*Actor
}

func NewEnricherActor(super *Supervisor, threads int) *EnricherActor {
	gpt := chatgpt.NewGPT()
	enricher := enrich.NewEnricher(gpt)

	a := &EnricherActor{
		Enricher: *enricher,
		Actor: &Actor{
			ID: "ENRICHER ACTOR",
		},
	}

	a.Init(super, a.Recieve, threads)
	return a
}

// concurrent actor loop
func (l *EnricherActor) Recieve(m ActorMsg) {

	enrich_msg, ok := m.Msg.(EnrichMsg)

	if !ok {
		log.Println("MESSAGE TYPE NOT EnrichMsg SKIPPING")
		return
	}

	l.Enricher.Summarize(&enrich_msg.Article)

	fmt.Println("=======================")
	fmt.Println("Title:", enrich_msg.Article.Title)
	fmt.Println("Summary:", enrich_msg.Article.Summary)
	fmt.Println("Sentiment:", enrich_msg.Article.Sentiment)
	fmt.Println("Category:", enrich_msg.Article.Category)

	// TODO implement an reciever actor that handles enriched articles
	// such as adding them to a db, REST request, or writing to file/bucket

}
