package es

import (
	"fmt"
	"github.com/olivere/elastic"
)

var ElasticClient *elastic.Client

func init(){
	var err error
	ElasticClient, err = elastic.NewClient(
		elastic.SetURL("http://9200"),
		elastic.SetSniff(false),

		)

	if err != nil {
		fmt.Println(err)
	}
}

func CloseElasticClient(){
	ElasticClient.Stop()
}