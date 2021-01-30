package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Stock struct {
	Ticker string `json:"symbol"`
	Price  struct {
		RegularMarketPrice struct {
			MarketPrice string `json:"fmt"`
		} `json:"regularMarketPrice"`
		RegularMarketVolume struct {
			MarkeVolume string `json:"fmt"`
		} `json:"regularMarketVolume"`
	} `json:"price"`
}

func (s *Server) Routes() {
	s.r.Get("/", getData)
}

func getData(w http.ResponseWriter, r *http.Request) {

	url := "https://apidojo-yahoo-finance-v1.p.rapidapi.com/stock/v2/get-financials?region=US&symbol=TSLA"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "apidojo-yahoo-finance-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "ce9d761a75mshb1550027513cc2fp124d12jsne0902f94a582")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	stock1 := Stock{}

	if err := json.Unmarshal(body, &stock1); err != nil {
		log.Fatalf("There was an error unmarshaling JSON: %v", err)
	}
	fmt.Println(res)
	fmt.Fprintln(w, "Ticker: "+stock1.Ticker)
	fmt.Fprintln(w, "Current share price is: "+stock1.Price.RegularMarketPrice.MarketPrice)
	fmt.Fprintln(w, "Daily Volume: "+stock1.Price.RegularMarketVolume.MarkeVolume)
}
