package main

import (
	"github.com/H15Z/gptsummary/cmd"
)

func main() {
	cmd.Execute()
	// configs.InitConfig()
	// chatgpt.QueryGPT("Translate Hellow World to Polish")
	// csvloader.LoadCSV()

	// gpt := chatgpt.NewGPT()
	// enricher := enrich.NewEnricher(gpt)

	// test_art := models.Article{
	// 	Title: "Title",
	// 	Text: `When it comes to training your brain, your sense of smell is possibly the last thing you’d think could strengthen your neural pathways. Learning a new language or reading more books (and fewer social media posts) — sure. But your nose?

	// 	That’s because the olfactory system is one of the most plastic systems in your brain. Neuroplasticity describes how the brain flexibly adapts to changes in the environment or when exposed to neural damage. Stimulating the brain strengthens existing neural structures and further adds fuel to the brain’s capacity to remain adaptive, thereby keeping it young. And your smell system is particularly adept at repair and renewal. (Olfactory cells have recently been used in human transplant therapy to treat spinal cord injury, for example.)

	// 	One reason for the olfactory system’s adaptive responsiveness is that it undergoes adult neurogenesis. Humans grow new olfactory neurons every three to four weeks throughout their entire life, not just during child development. (These sensory neurons sit in the mucous of your nose, where they pick up airborne chemicals and send activity signals straight to the core of the brain.) If it weren’t for this ongoing regeneration of sensory cells in your nose, we would stop detecting smells after our first few colds.
	// 	`,
	// }

	// enricher.Summarize(&test_art)

	// csvloader.LoadCSV(PrintArticleTitle)

	// dl := dataloader.NewDataLoader(
	// 	csvloader.NewCSVLoader(),
	// )

	// dl.StreamData(PrintArticleTitle)
}
