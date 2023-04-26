package lib

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
)

type EmailVars struct {
	To        string            `csv:"to"`
	OtherData map[string]string `csv:"-"`
}

type EmailVarsConverted struct {
	To           string
	TemplateVars []string
}

func ParsePhishCSV(csvfile string) []EmailVarsConverted {

	dat, err := os.Open(csvfile)
	CheckErr(err)

	csvReader := csv.NewReader(dat)
	dec, err := csvutil.NewDecoder(csvReader)
	CheckErr(err)

	header := dec.Header()

	emailOptions := []EmailVarsConverted{}

	for {
		// u := User{OtherData: make(map[string]string)}
		u := EmailVars{OtherData: make(map[string]string)}
		n := EmailVarsConverted{}
		var tv []string

		if err := dec.Decode(&u); err == io.EOF {
			break
		} else {
			CheckErr(err)
		}

		for _, i := range dec.Unused() {
			templateVar := strings.TrimSpace(header[i]) + "=" + strings.TrimSpace(dec.Record()[i])
			tv = append(tv, templateVar)
			// u.TemplateVars[i] = templateVar
		}
		n.TemplateVars = tv
		n.To = u.To
		emailOptions = append(emailOptions, n)
	}

	// for k := range emailoptions[1].OtherData {
	// 	fmt.Println(k)
	// 	fmt.Println(emailoptions[1].OtherData[k])
	// }
	// fmt.Println(emailoptions[1].OtherData)

	return emailOptions

}
