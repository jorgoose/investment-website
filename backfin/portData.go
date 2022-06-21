package backfin

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/piquette/finance-go/equity"
)

type HoldingInfo struct {
	Sector   string
	Name     string
	Symbol   string
	Quantity float64
	AvgPrice string
    CurrPrice float64
}

func RnPCSV() []HoldingInfo {
	csvFile, _ := os.Open("backfin/igdata.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var port []HoldingInfo

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

        qn, _ := strconv.ParseFloat(line[3], 64)

        fmt.Println(line[2])

        lp := priceFetcher(line[2])

		port = append(port, HoldingInfo{
			Sector: line[0],
			Name:  line[1],
			Symbol: line[2],
            Quantity: qn,
            AvgPrice: line[4],
            CurrPrice: lp,
        })
	}
    fmt.Println(port[0].Name, port[0].CurrPrice, port[0].Quantity, port[0].Symbol)
    return port
}

func SectorVals(port []HoldingInfo) map[string]float64 {
    m := make(map[string]float64)
    for _, v := range port {
        if (m[v.Sector] != 0) {
            m[v.Sector] = m[v.Sector] + (float64(v.Quantity) * v.CurrPrice)
        } else {
            m[v.Sector] = (float64(v.Quantity) * v.CurrPrice)
        }
    }
    return m

}

// TODO Add ability to gather live stock price data
func priceFetcher(tick string) (float64) {
    if (tick == "$") {
        return 1.00
    }

    q, err := equity.Get(tick)
    if err != nil { 
        panic(err)
    }

    return q.RegularMarketPrice
}
