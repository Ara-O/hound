package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ara-o/doc-find/utils"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/embeddings/openai"
	l "github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pinecone"
)

var url string
var comprehensive bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doc-find",
	Short: "Easily query through long pages of documentation!",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading environment variables")
		}

		e, err := openai.NewOpenAI()

		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()

		body, err := utils.ParseURL(url, comprehensive)

		if err != nil {
			log.Fatal(err)
		}

		doc := []schema.Document{{
			PageContent: body,
			Metadata: map[string]any{
				"source": url,
			},
		}}

		textSplitter := textsplitter.NewRecursiveCharacter()
		textSplitter.ChunkOverlap = 0
		textSplitter.ChunkSize = 500

		splitDocs, err := textsplitter.SplitDocuments(textSplitter, doc)

		if err != nil {
			log.Fatal(err)
		}

		// Create a new Pinecone vector store.
		store, err := pinecone.New(
			ctx,
			pinecone.WithProjectName(os.Getenv("PINECONE_PROJECT_NAME")),
			pinecone.WithIndexName(os.Getenv("PINECONE_INDEX_NAME")),
			pinecone.WithEnvironment(os.Getenv("PINECONE_ENVIRONMENT_NAME")),
			pinecone.WithEmbedder(e),
			pinecone.WithAPIKey(os.Getenv("PINECONE_API_KEY")),
			pinecone.WithNameSpace(uuid.New().String()),
		)

		if err != nil {
			log.Fatal(err)
		}

		err = store.AddDocuments(ctx, splitDocs)

		if err != nil {
			log.Fatal(err)
		}

		llm, err := l.New()

		if err != nil {
			fmt.Println(err)
		}

		qa := chains.NewRetrievalQAFromLLM(llm, vectorstores.ToRetriever(store, 10))

		ans, err := qa.Call(context.Background(), map[string]any{
			"query": "Imagine this is a question",
		})

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("ans", ans)

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Set up flags
func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "Enter the URL of the documentation page you want to query")
	rootCmd.Flags().BoolVarP(&comprehensive, "comprehensive", "c", false, "Select true if you want to search through the entire site's documentation")
	rootCmd.MarkFlagRequired("url")
}
