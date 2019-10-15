package main

import (
	"context"
	"fmt"
	"net/http"

	paapi5 "github.com/spiegel-im-spiegel/pa-api"
	"github.com/spiegel-im-spiegel/pa-api/entity"
	"github.com/spiegel-im-spiegel/pa-api/query"
)

func main() {
	//Create client
	client := paapi5.New(
		paapi5.WithMarketplace(paapi5.LocaleJapan),
	).CreateClient(
		"mytag-20",             //Amazon associate tag
		"AKIAIOSFODNN7EXAMPLE", //access key for PA-API
		"1234567890",           //secret key for PA-API
		paapi5.WithContext(context.Background()),
		paapi5.WithHttpClient(http.DefaultClient),
	)

	//Make query
	q := query.NewSearchItems(
		client.Marketplace(),
		client.PartnerTag(),
		client.PartnerType(),
	).Search("数学ガール", query.Keywords)

	//Requet and response
	body, err := client.Request(q)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	//io.Copy(os.Stdout, bytes.NewReader(body))

	//Decode JSON
	res, err := entity.DecodeResponse(body)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Println(res.String())
}
