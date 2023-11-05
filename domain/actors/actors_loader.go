package actors

import (
	"fmt"
	"log"
	"time"

	"github.com/H15Z/gptsummary/domain/dataloader"
	"github.com/H15Z/gptsummary/domain/models"
	"github.com/H15Z/gptsummary/pkg/csvloader"
)

type LoaderMsg struct {
	Count   int
	Threads int
}

type LoaderActor struct {
	DataEnricher *EnricherActor
	LinkName     string
	Count        int
	Limit        int
	*Actor
}

func NewLoaderActor(super *Supervisor) *LoaderActor {
	a := &LoaderActor{
		Actor: &Actor{
			ID: "DATA LOADER ACTOR",
		},
		Count: 0,
	}

	a.Init(super, a.Recieve, 1)
	return a
}

// concurrent actor loop
func (l *LoaderActor) Recieve(m ActorMsg) {
	defer l.ActorStop()

	start := time.Now()
	log.Println("LOADER MSG RECIEVED :", m)

	load_msg, ok := m.Msg.(LoaderMsg)

	if !ok {
		log.Println("MESSAGE TYPE NOT LoaderMsg QUITING")
		l.ActorStop()
		return
	}

	log.Println("STARTING DATA LOADER ACTOR - Count: ", load_msg.Count)

	l.Limit = load_msg.Count

	// spawn Enricher Actors
	l.DataEnricher = NewEnricherActor(l.Supervisor, load_msg.Threads)

	dl := dataloader.NewDataLoader(
		csvloader.NewCSVLoader(),
	)

	dl.StreamData(l.SendForEnriching) // Pass Callback into data streamer
	l.Monitor()

	log.Println("DATA LOADER ACTOR FINISHED:", time.Since(start))
}

func (l *LoaderActor) Monitor() {
	// Monitor monitor and stop enrichers
	for {
		//  CHECK IF FINISHED PROCESSING AND UPDATING
		if l.DataEnricher.MessageCount() == 0 && !l.DataEnricher.IsActive() {
			l.DataEnricher.ActorStop()
			break
		}
	}
}

func (l *LoaderActor) SendForEnriching(article models.Article) {

	if l.Count > l.Limit {
		return
	}

	PrintArticleTitle(article)

	l.Count += 1

	l.Actor.SendMsg(l.DataEnricher, ActorMsg{
		Msg: EnrichMsg{
			Article: article,
		},
	})
}

// Use for testing
func PrintArticleTitle(article models.Article) {
	fmt.Println(article.Title)
	fmt.Println("---------------------------------------")

}