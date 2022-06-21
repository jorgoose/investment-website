package backfin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RawCompanyProfile []struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Beta              float64 `json:"beta"`
	VolAvg            int     `json:"volAvg"`
	MktCap            int64   `json:"mktCap"`
	LastDiv           float64 `json:"lastDiv"`
	Range             string  `json:"range"`
	Changes           float64 `json:"changes"`
	CompanyName       string  `json:"companyName"`
	Currency          string  `json:"currency"`
	Cik               string  `json:"cik"`
	Isin              string  `json:"isin"`
	Cusip             string  `json:"cusip"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Industry          string  `json:"industry"`
	Website           string  `json:"website"`
	Description       string  `json:"description"`
	Ceo               string  `json:"ceo"`
	Sector            string  `json:"sector"`
	Country           string  `json:"country"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Zip               string  `json:"zip"`
	DcfDiff           float64 `json:"dcfDiff"`
	Dcf               float64 `json:"dcf"`
	Image             string  `json:"image"`
	IpoDate           string  `json:"ipoDate"`
	DefaultImage      bool    `json:"defaultImage"`
	IsEtf             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
	IsAdr             bool    `json:"isAdr"`
	IsFund            bool    `json:"isFund"`
}

func CompanyProfileFetch(ticker string) RawCompanyProfile{
    url := "https://financialmodelingprep.com/api/v3/profile/" + ticker + "?apikey=b405b3fe5bb816a2deed52e84d1eb84d"
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var result RawCompanyProfile

    if err := json.Unmarshal(responseData, &result); err != nil {   // Parse []byte to go struct pointer
        fmt.Println("Can not unmarshal JSON")
    }

    return result
}

type CoreCompanyProfile struct {
	Symbol string
	Price float64
	Mktcap int64
	CompanyName string
	Industry string
	Website string
	Description string
	Ceo string
	Sector string
	Country string
	Image string
	DefaultImage string
}

func InitCompanyProfile(s string) CoreCompanyProfile{
	rawProf := CompanyProfileFetch(s)
	coreProf := CoreCompanyProfile{}

	for _, val := range rawProf {
		coreProf.Symbol = val.Symbol
		coreProf.Price = val.Price
		coreProf.Mktcap = val.MktCap
		coreProf.CompanyName = val.CompanyName
		coreProf.Industry = val.Industry
		coreProf.Website = val.Website
		coreProf.Description = val.Description
        coreProf.Ceo = val.Ceo
		coreProf.Sector = val.Sector
		coreProf.Country = val.Country
		coreProf.Image = val.Image
    }

	return coreProf
}