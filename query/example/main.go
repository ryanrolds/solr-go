package main

import (
	"context"
	"log"

	"github.com/stevenferrer/solr-go/query"
	"github.com/stevenferrer/solr-go/types"
)

func main() {
	// Initialize JSON query client
	host := "localhost"
	port := 8983
	queryClient := query.NewJSONClient(host, port)

	collection := "techproducts"

	// Simple query string
	resp, err := queryClient.Query(context.Background(),
		collection, types.M{"query": "name:iPod"},
	)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the response
	_ = resp

	// Full-expanded JSON object
	resp, err = queryClient.Query(context.Background(),
		collection, types.M{
			"query": types.M{
				"lucene": types.M{
					"df":    "name",
					"query": "iPod",
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the response
	_ = resp

	// Complex queries
	resp, err = queryClient.Query(context.Background(), collection, types.M{
		"query": types.M{
			"boost": types.M{
				"query": types.M{
					"lucene": types.M{
						"df":    "name",
						"query": "iPod",
					},
				},
				"b": "log(popularity)",
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	// do something with the response
	_ = resp
}
