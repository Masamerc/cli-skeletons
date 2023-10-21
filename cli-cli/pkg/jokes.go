package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/urfave/cli/v2"
)

const endpoint = "https://icanhazdadjoke.com/"

type JokeResp struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getJoke() (JokeResp, error) {
	client := &http.Client{}

	request, err := http.NewRequest(
		"GET",
		endpoint,
		nil,
	)
	if err != nil {
		fmt.Println("Error while creating request")
		return JokeResp{}, err
	}

	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error while sending request")
		return JokeResp{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while reading response")
		return JokeResp{}, err
	}

	var joke JokeResp
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Println("Error while unmarshalling response")
		return JokeResp{}, err
	}

	return joke, nil
}

func PrintJoke(c *cli.Context) error {
	joke, err := getJoke()
	jokeText := joke.Joke
	if err != nil {
		fmt.Println("Error while getting joke")
		return err
	}

	if c.Bool("upper") {
		jokeText = strings.ToUpper(jokeText)
	}

	fmt.Println(jokeText)

	return nil
}
