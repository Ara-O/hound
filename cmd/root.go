package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ara-o/doc-find/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
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

		body, err := utils.ParseURL(url, comprehensive)

		if err != nil {
			log.Fatal(err)
		}

		doc := []schema.Document{schema.Document{
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

		llm, err := openai.New()

		if err != nil {
			log.Fatal(err)
		}

		stuffQAChain := chains.LoadStuffQA(llm)

		answer, err := chains.Call(context.Background(), stuffQAChain, map[string]any{
			"input_documents": splitDocs,
			"question":        "How can i conditionally render in react?",
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(answer)
		// retrievalQA.Call(context.TODO(), nil, chains.ChainCallOption{})
		// fmt.Printf("%+v", splitBody[0])
		// fmt.Println(body)
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
