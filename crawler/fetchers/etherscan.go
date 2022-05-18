package fetchers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	log "github.com/sirupsen/logrus"
)

type EtherscanFetcher struct{}

func NewEtherscanFetcher() *EtherscanFetcher {
	return &EtherscanFetcher{}
}

func (f *EtherscanFetcher) SendRequest(page int) []string {
	url := "https://etherscan.io/tokens?p=" + strconv.Itoa(page)
	log.Println("Fetch Url", url)
	res := make([]string, 0)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Http get err:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	nodes := htmlquery.Find(doc, `//*[@id="tblResult"]/tbody/tr/td[2]/div/div/h3`)
	for _, node := range nodes {
		href := htmlquery.FindOne(node, "./a/@href")
		hash := strings.Split(htmlquery.InnerText(href), "/")[2]
		res = append(res, hash)
	}
	return res
}
