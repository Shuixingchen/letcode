package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/PaesslerAG/jsonpath"
)

func main() {
	dateInt, _ := strconv.Atoi("202209")
	aa := int64(dateInt)
	fmt.Println(aa)
}

func PoolHashrate() {
	jsonstr := `{"code":"000000","msg":"","data":{"items":[{"coinType":"BTC","poolHashrate":"49.8 EH/s","networkDiff":35610794164372,"blockReward":6.35872833000000000000,"coinPriceCny":137610.5240,"coinPriceUsd":19033.2675,"miningType":[{"payMethod":"PPLNS","percent":1,"percentStr":"0%"},{"payMethod":"FPPS","percent":0.96,"percentStr":"4%"}],"minimumPayment":0.00500000,"calculateUnit":1000000000000,"coinUnit":"H/s","netHashrate":261.05,"netHashrateUnit":"EH/s","coinCoefficient":4294967296,"supportAnonymous":false,"appPoolTab":["poolDetail"],"nextDiff":"38132064774373.73","nextDiffAdjustTime":"317819567","algorithm":"SHA256d","blockIncentive":6.25000000000000000000,"blockTxFee":0.10872833000000000000,"supportFpps":true,"mergeMiningInfos":[{"coinType":"","mergeRate":0,"explanation":"1 BTC = 2 NMC","coinPriceUsd":null,"coinPriceCny":null}]},{"coinType":"BCH","poolHashrate":"56.9 PH/s","networkDiff":195845605834,"blockReward":6.25000000000000000000,"coinPriceCny":767.0321,"coinPriceUsd":106.0902,"miningType":[{"payMethod":"PPLNS","percent":1,"percentStr":"0%"},{"payMethod":"FPPS","percent":0.96,"percentStr":"4%"}],"minimumPayment":0.01000000,"calculateUnit":1000000000000,"coinUnit":"H/s","netHashrate":1.38,"netHashrateUnit":"EH/s","coinCoefficient":4294967296,"supportAnonymous":false,"appPoolTab":["poolDetail"],"nextDiff":"197721114948.62","nextDiffAdjustTime":"691425522","algorithm":"SHA256d","blockIncentive":6.25000000000000000000,"blockTxFee":0E-20,"supportFpps":true,"mergeMiningInfos":null},{"coinType":"LTC","poolHashrate":"78.2 TH/s","networkDiff":16540224,"blockReward":12.50000000000000000000,"coinPriceCny":367.6455,"coinPriceUsd":50.8500,"miningType":[{"payMethod":"PPS","percent":0.97,"percentStr":"3%"}],"minimumPayment":0.00100000,"calculateUnit":1000000,"coinUnit":"H/s","netHashrate":522.49,"netHashrateUnit":"TH/s","coinCoefficient":4294967296,"supportAnonymous":false,"appPoolTab":["poolDetail"],"nextDiff":"17073961.90","nextDiffAdjustTime":"86133487","algorithm":"Scrypt","blockIncentive":12.50000000000000000000,"blockTxFee":0,"supportFpps":false,"mergeMiningInfos":[{"coinType":"DOGE","mergeRate":1818,"explanation":"DOGE rewards in PPLNS mode","coinPriceUsd":"0.05906000","coinPriceCny":"0.42699789"}]},{"coinType":"ETH","poolHashrate":"0.00 H/s","networkDiff":6476299033491008,"blockReward":0.13135139000000000000,"coinPriceCny":9257.1713,"coinPriceUsd":1280.3833,"miningType":[{"payMethod":"PPS+","percent":0.985,"percentStr":"1.5%"}],"minimumPayment":0.20000000,"calculateUnit":1000000,"coinUnit":"H/s","netHashrate":0.00,"netHashrateUnit":"H/s","coinCoefficient":1,"supportAnonymous":true,"appPoolTab":["poolDetail"],"nextDiff":null,"nextDiffAdjustTime":null,"algorithm":"Ethash","blockIncentive":2.00000000000000000000,"blockTxFee":-1.86864861000000000000,"supportFpps":true,"mergeMiningInfos":null},{"coinType":"ETC","poolHashrate":"7.46 TH/s","networkDiff":1932990348329150,"blockReward":2.56000000000000000000,"coinPriceCny":161.5869,"coinPriceUsd":22.3495,"miningType":[{"payMethod":"PPS","percent":0.97,"percentStr":"3%"}],"minimumPayment":0.20000000,"calculateUnit":1000000,"coinUnit":"H/s","netHashrate":146.57,"netHashrateUnit":"TH/s","coinCoefficient":1,"supportAnonymous":true,"appPoolTab":["poolDetail"],"nextDiff":null,"nextDiffAdjustTime":null,"algorithm":"Etchash","blockIncentive":2.56000000000000000000,"blockTxFee":0,"supportFpps":false,"mergeMiningInfos":null},{"coinType":"ZEC","poolHashrate":"1.21 GSol/s","networkDiff":92740614,"blockReward":2.50000000000000000000,"coinPriceCny":366.5610,"coinPriceUsd":50.7000,"miningType":[{"payMethod":"PPS","percent":0.97,"percentStr":"3%"},{"payMethod":"PPLNS","percent":1,"percentStr":"0%"}],"minimumPayment":0.00100000,"calculateUnit":1000,"coinUnit":"Sol/s","netHashrate":9.01,"netHashrateUnit":"GSol/s","coinCoefficient":8192,"supportAnonymous":true,"appPoolTab":["poolDetail"],"nextDiff":null,"nextDiffAdjustTime":null,"algorithm":"Equihash","blockIncentive":2.50000000000000000000,"blockTxFee":0,"supportFpps":false,"mergeMiningInfos":null},{"coinType":"DASH","poolHashrate":"407 TH/s","networkDiff":62620247,"blockReward":1.08945457000000000000,"coinPriceCny":286.9225,"coinPriceUsd":39.6850,"miningType":[{"payMethod":"PPS","percent":0.98,"percentStr":"2%"}],"minimumPayment":0.00100000,"calculateUnit":1000000000,"coinUnit":"H/s","netHashrate":2.21,"netHashrateUnit":"PH/s","coinCoefficient":4294967296,"supportAnonymous":false,"appPoolTab":["poolDetail"],"nextDiff":"59347318.50","nextDiffAdjustTime":"47238403","algorithm":"X11","blockIncentive":1.08945457000000000000,"blockTxFee":0,"supportFpps":false,"mergeMiningInfos":null},{"coinType":"ETHW","poolHashrate":"985 GH/s","networkDiff":489148802434872,"blockReward":2.00000000000000000000,"coinPriceCny":45.1658,"coinPriceUsd":6.2470,"miningType":[{"payMethod":"PPS+","percent":0.985,"percentStr":"1.5%"}],"minimumPayment":0.20000000,"calculateUnit":1000000,"coinUnit":"H/s","netHashrate":37.40,"netHashrateUnit":"TH/s","coinCoefficient":1,"supportAnonymous":true,"appPoolTab":["poolDetail"],"nextDiff":null,"nextDiffAdjustTime":null,"algorithm":"Ethash","blockIncentive":2.00000000000000000000,"blockTxFee":0E-8,"supportFpps":true,"mergeMiningInfos":null}]}}`
	var object interface{}
	if err := json.Unmarshal([]byte(jsonstr), &object); err != nil {
		fmt.Println("Unmarshal", err)
		return
	}
	h, err := jsonpath.Get(`$.data.items[?(@.coinType=="ETC")].poolHashrate`, object)
	if err != nil {
		fmt.Println("jsonpath", err)
	}
	units, err := jsonpath.Get(`$.data.items[?(@.coinType=="ETC")].poolHashrate`, object)
	if err != nil {
		fmt.Println("aa", err)
	}
	u := ParseUnit(units, "")
	hashreate := ParseHashrate(h, "")
	fmt.Println(hashreate, u)
}

func ParseUnit(object interface{}, poolName string) string {
	var unit string
	unitType := reflect.TypeOf(object)
	switch unitType.String() {
	case "string":
		unit = object.(string)
		break
	case "[]interface {}":
		units := object.([]interface{})
		unit = units[0].(string)
		if strings.Contains(unit, " ") {
			strs := strings.Split(unit, " ")
			unit = strs[1]
		}
		break
	}
	return unit
}

func ParseHashrate(object interface{}, poolName string) float64 {
	var hashrate float64
	hashrateType := reflect.TypeOf(object)
	switch hashrateType.String() {
	case "float64":
		hashrate = object.(float64)
		break
	case "string":
		var err error
		hashrateStr := strings.TrimSpace(object.(string))

		hashrate, err = strconv.ParseFloat(hashrateStr, 64)
		if err != nil {
			output := strings.Map(func(r rune) rune {
				if r >= '0' && r <= '9' {
					return r
				}
				if r == '.' {
					return r
				}
				return -1
			}, hashrateStr)
			hashrate, _ = strconv.ParseFloat(output, 64)
		}
		break
	case "[]interface {}":
		hashrateList := object.([]interface{})
		if len(hashrateList) > 0 {
			var ok bool
			hashrate, ok = hashrateList[0].(float64)
			if !ok {
				hashrateStr := hashrateList[0].(string)
				if strings.Contains(hashrateStr, " ") {
					strs := strings.Split(hashrateStr, " ")
					hashrateStr = strs[0]
				}
				hashrate, _ = strconv.ParseFloat(hashrateStr, 64)
			}
		}
		break
	case "map[string]interface {}":
		hashrateMap := object.(map[string]interface{})
		keys := make([]string, 0)
		for k := range hashrateMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		if r, ok := hashrateMap[keys[len(keys)-1]]; ok {
			hashrate = r.(float64)
		}
	}
	return hashrate
}

func Demo() {

	v := interface{}(nil)

	json.Unmarshal([]byte(`{"data":[
		{"key":"a","value" : "I"},
		{"key":"b","value" : "II"},
		{"key":"c","value" : "III"}
		]}`), &v)

	values, err := jsonpath.Get(`$.data[?(@.key=="b")].value`, v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, value := range values.([]interface{}) {
		fmt.Println(value)
	}

}
