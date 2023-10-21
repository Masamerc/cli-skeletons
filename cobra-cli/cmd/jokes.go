package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

const endpoint = "https://icanhazdadjoke.com/"

type JokeResp struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// get a random joke from  https://icanhazdadjoke.com/api
func getJoke() (JokeResp, error) {
	client := &http.Client{}
	request, err := http.NewRequest(
		"GET",
		endpoint,
		nil,
	)
	request.Header.Set("Accept", "application/json")

	if err != nil {
		fmt.Println("Error making request")
		return JokeResp{}, err
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error getting response")
		return JokeResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return JokeResp{}, err
	}

	var jokeResp JokeResp
	err = json.Unmarshal(body, &jokeResp)

	if err != nil {
		fmt.Println("Error unmarshalling response")
		return JokeResp{}, err
	}

	return jokeResp, nil
}

func printJoke(cmd *cobra.Command, args []string) {
	jokeResp, err := getJoke()
	if err != nil {
		fmt.Println("Error getting joke")
		return
	}

	joke := jokeResp.Joke

	var uppercase bool
	if cmd.Flags().Changed("upper") {
		uppercase, _ = cmd.Flags().GetBool("upper")
	}

	if uppercase {
		joke = strings.ToUpper(joke)
	}

	fmt.Println(joke)
}

var JokesCmd = &cobra.Command{
	Use:   "jokes",
	Short: "generate jokes",
	Long: `generate jokes from https://icanhazdadjoke.com/api
and displays them in the terminal`,
	Run: printJoke,
}
