package pokeapi

// "github.com/mpedone/pokedexcli"

const (
	baseURL = "https://pokeapi.co/api/v2"
)

/* func PokeReq(client *Client, url *string) {

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		log.Fatal(err)
		// return condition, fmt.Errorf("NewRequest err: %w\n", err)
	}

	var data Keys

	res, err := client.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&data)
	fmt.Println(data.Next)
	fmt.Println(data.Previous)
} */
