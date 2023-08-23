package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ara-o/doc-find/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
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

		textSplitter := textsplitter.NewRecursiveCharacter()
		textSplitter.ChunkSize = 500
		textSplitter.ChunkOverlap = 0

		var test []schema.Document

		metadata := make(map[string]string)

		textsplitter.CreateDocuments(textSplitter, _, _)

		test, err = textsplitter.SplitDocuments(textSplitter, test)

		fmt.Println(test[0])

		// test, err := textSplitter.SplitText(body)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(test)
		//Splitting the docs
		// splitBody, err := textsplitter.SplitDocuments(textSplitter, body)

		// openaiVar, err := openai.New()
		// openaiVar.CreateEmbedding(context.Background(), splitBody)

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
