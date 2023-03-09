package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/vision v1.2.0"
)

func main() {
	ctx := context.Background()

	// Initialize the client with Google Application Credentials.
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile("path/to/your/credentials.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Define the request payload as a byte array.
	requestBody := []byte(`{
		"requests": [
			{
				"image": {
					"source": {
						"imageUri": "https://example.com/image.jpg"
					}
				},
				"features": [
					{
						"type": "OBJECT_LOCALIZATION",
						"maxResults": 5
					},
					{
						"type": "FACE_DETECTION",
						"maxResults": 5
					},
					{
						"type": "TEXT_DETECTION",
						"maxResults": 1
					}
				]
			}
		]
	}`)

	// Send the request to the Vision API.
	resp, err := client.AnnotateImage(ctx, &vision.AnnotateImageRequest{
		Image: &vision.Image{
			Source: &vision.ImageSource{
				ImageUri: "https://example.com/image.jpg",
			},
		},
		Features: []*vision.Feature{
			{
				Type:       vision.Feature_OBJECT_LOCALIZATION,
				MaxResults: 5,
			},
			{
				Type:       vision.Feature_FACE_DETECTION,
				MaxResults: 5,
			},
			{
				Type:       vision.Feature_TEXT_DETECTION,
				MaxResults: 1,
			},
		},
	})

	if err != nil {
		log.Fatalf("Failed to annotate image: %v", err)
	}

	// Print the response.
	jsonResponse, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	fmt.Println(string(jsonResponse))
}
