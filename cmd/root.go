package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ara-o/doc-find/utils"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
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
		err := godotenv.Load("/home/ara/go/bin/.env")

		if err != nil {
			log.Fatal("Error loading environment variables", err)
		}

		pterm.DefaultBox.
			WithRightPadding(10).
			WithLeftPadding(10).
			WithTopPadding(1).
			WithBottomPadding(1).
			WithHorizontalString("═").
			WithVerticalString("║").
			WithBottomLeftCornerString("╗").
			WithBottomRightCornerString("╔").
			WithTopLeftCornerString("╝").
			WithTopRightCornerString("╚").
			Println("Welcome to Doc-Find!")

		parsingURLSpinner, _ := pterm.DefaultSpinner.Start("Parsing URL Data...")

		body, err := utils.ParseURL(url, comprehensive)

		//Showing spinner + details
		if err != nil {
			parsingURLSpinner.Fail("Error parsing URL")
			log.Fatal("Information: ", err)
		} else {
			parsingURLSpinner.Success("URL parsed successfully")
		}

		e, err := openai.NewOpenAI()

		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()

		doc := []schema.Document{{
			PageContent: body,
			Metadata: map[string]any{
				"source": url,
			},
		}}

		//Setting up text splitter

		textSplitter := textsplitter.NewRecursiveCharacter()
		textSplitter.ChunkOverlap = 0
		textSplitter.ChunkSize = 500

		splitDocs, err := textsplitter.SplitDocuments(textSplitter, doc)

		if err != nil {
			log.Fatal(err)
		}

		creatingStoreSpinner, _ := pterm.DefaultSpinner.Start("Connecting to Vectore Store...")

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
			creatingStoreSpinner.Fail("Connecting to vector store failed")
		} else {
			creatingStoreSpinner.Success("Successfully connected to vector store")
		}

		addingDocumentsSpinner, _ := pterm.DefaultSpinner.Start("Adding documents, this may take a while...")

		err = store.AddDocuments(ctx, splitDocs)

		if err != nil {
			log.Fatal(err)
			addingDocumentsSpinner.Fail("Error adding documents")
		} else {
			addingDocumentsSpinner.Success("Successfully added documents")
		}

		llm, err := l.New()

		if err != nil {
			fmt.Println(err)
		}

		qa := chains.NewRetrievalQAFromLLM(llm, vectorstores.ToRetriever(store, 10))
		_ = qa

		pterm.Println()

		question, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Input the question you would like answered ").WithMultiLine(false).Show()

		pterm.Println()

		searchingSpinner, _ := pterm.DefaultSpinner.Start("Looking for answer...")

		ans, err := qa.Call(context.Background(), map[string]any{
			"query": question,
		})

		if err != nil {
			searchingSpinner.Fail("Error in looking for answer")
			log.Fatal(err)
		} else {
			searchingSpinner.Success(ans["text"])
		}

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
