package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ara-o/hound/db"
	"github.com/ara-o/hound/utils"
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
var forceDocumentStorage bool
var searchDepth int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hound",
	Short: "Easily query through long pages of documentation!",
	Run: func(cmd *cobra.Command, args []string) {
		path := utils.GetEnvironmentVariablePath()

		err := godotenv.Load(path)

		if err != nil {
			fmt.Println("Error: Error loading environment variables")
			pterm.Println("Run the", pterm.Sprint(pterm.LightCyan("hound setup")), "command to set up environment variables")
			return
		}

		defer db.Close()

		utils.ViewWelcomeLogo()

		//Parse URL
		parsingURLSpinner, _ := pterm.DefaultSpinner.Start("Parsing URL Data...")

		body, err := utils.ParseURL(url, comprehensive, searchDepth)

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

		// Use the namespace of URL to keep documentation separate
		// If url is not in database, add documents to the store related to that url
		// If url is in database, no need to add documents, the documents would've been added
		// The first time, under that URLs namespace
		store, err := pinecone.New(
			ctx,
			pinecone.WithProjectName(os.Getenv("PINECONE_PROJECT_ID")),
			pinecone.WithIndexName(os.Getenv("PINECONE_INDEX_NAME")),
			pinecone.WithEnvironment(os.Getenv("PINECONE_ENVIRONMENT_NAME")),
			pinecone.WithEmbedder(e),
			pinecone.WithAPIKey(os.Getenv("PINECONE_API_KEY")),
			pinecone.WithNameSpace(url),
		)

		if err != nil {
			log.Fatal(err)
			creatingStoreSpinner.Fail("Connecting to vector store failed")
		} else {
			creatingStoreSpinner.Success("Successfully connected to vector store")
		}

		// check if url has already been indexed
		indexExistsInDB := true
		addingDocumentsSpinner, _ := pterm.DefaultSpinner.Start("Adding documents, this may take a while...")

		if !db.IndexExistsInDB(url) {
			indexExistsInDB = false
			err := db.AddIndex(url)

			if err != nil {
				log.Fatal(err)
			}
		}

		//If the index does not exist, add the index
		if !indexExistsInDB || forceDocumentStorage {
			err = store.AddDocuments(ctx, splitDocs)
			if err != nil {
				addingDocumentsSpinner.Fail("Error adding documents")
				log.Fatal(err)

			}
		}

		addingDocumentsSpinner.Success("Successfully added documents")

		llm, err := l.New()

		if err != nil {
			fmt.Println(err)
		}

		qa := chains.NewRetrievalQAFromLLM(llm, vectorstores.ToRetriever(store, 10))
		_ = qa

		for {
			pterm.Println()

			pterm.Yellow("Note: Type 'end' ( without the quotes ) to end conversation")
			pterm.Println("Enter the question you would like answered [Enter " + pterm.Red("end") + " to end]")
			question, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Input").WithMultiLine(false).Show()

			if strings.ToLower(strings.TrimSpace(question)) == "end" {
				break
			}

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
	rootCmd.Flags().BoolVarP(&forceDocumentStorage, "force-storage", "f", false, "Parsed URL data are cached for later re-use. Use this if you want to force the re-storage of already parsed URL sites")
	rootCmd.Flags().BoolVarP(&comprehensive, "comprehensive", "c", false, "Use if you want to search through the entire site's documentation [Tentative: Goes through a maximum of 10 links]")
	rootCmd.Flags().IntVarP(&searchDepth, "depth", "d", 10, "Select how many links to parse when doing a comprehensive search ( Warning: A higher number will result in longer parse times )")
	rootCmd.MarkFlagRequired("url")
}
