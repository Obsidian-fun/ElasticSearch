
package main

import (
	"log"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	es, err := elasticsearch.NewDefaultClient();
	if err != nil {
		log.Fatalf("could not create client: %s", err);
	}
	log.Println(elasticsearch.Version);

	res, err := es.info();
	if err != nil{
		log.Fatalf("error getting response: %s",err);
	}




}
