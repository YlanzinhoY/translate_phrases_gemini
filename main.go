package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Erro ao carregar o arquivo .env: %s", err)
		return
	}
	apikey := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))

	if err != nil {
		fmt.Println(err)
	}

	defer client.Close()
	model := client.GenerativeModel("gemini-pro")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Choice target language for translate (input name language in english) \n")

	t, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("enter phrase to translate")
	p, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("erro")
		return
	}
	t = strings.TrimSpace(t)
	p = strings.TrimSpace(p)

	prompt := []genai.Part{
		genai.Text(fmt.Sprintf("Translate this pharase in this language %s ", t)),
		genai.Text(p),
	}
	res, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		fmt.Println(res)
		return
	}

	for i, k := range res.Candidates {
		fmt.Println(k.Content.Parts[i])
	}

}
