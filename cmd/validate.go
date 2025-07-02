package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/spf13/cobra"
)

type HttpLoader struct {
	client *http.Client
}

func (l *HttpLoader) Load(url string) (any, error) {
	response, err := l.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data any
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if dataMap, ok := data.(map[string]any); ok {
		if schema, exists := dataMap["$schema"]; exists {
			if schema == "https://meta.json-schema.tools/" {
				dataMap["$schema"] = "http://json-schema.org/draft-07/schema#"
			}
		}
	}

	return data, nil
}

func fetchOpenRPCSchema() string {
	schemaURL := "https://meta.open-rpc.org"
	response, err := http.Get(schemaURL)
	if err != nil {
		log.Fatalf("Cannot load OpenRPC Document: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Cannot read OpenRPC Document: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Cannot load OpenRPC Document: %v", response.Status)
	}
	return string(body)
}

var validateCmd = &cobra.Command{
	Use:   "validate [file]",
	Short: "Validate an OpenRPC document",
	Long:  "Validate an OpenRPC document against basic OpenRPC specification requirements. Defaults to 'openrpc.json' if no file is specified.",
	Run: func(cmd *cobra.Command, args []string) {
		filename := "openrpc.json"
		if len(args) > 0 {
			filename = args[0]
		}

		fmt.Printf("Validating OpenRPC document: %s\n", filename)

		schemaJSON := fetchOpenRPCSchema()

		compiler := jsonschema.NewCompiler()

		schemaData, err := jsonschema.UnmarshalJSON(strings.NewReader(schemaJSON))
		if err != nil {
			fmt.Printf("Error parsing schema JSON: %v\n", err)
			return
		}
		compiler.UseLoader(&HttpLoader{client: &http.Client{}})

		err = compiler.AddResource("schema.json", schemaData)
		if err != nil {
			fmt.Printf("Error adding schema: %v\n", err)
			return
		}

		schema, err := compiler.Compile("schema.json")
		if err != nil {
			fmt.Printf("Error compiling schema: %v\n", err)
			return
		}

		openrpc, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", filename, err)
			return
		}

		data, err := jsonschema.UnmarshalJSON(strings.NewReader(string(openrpc)))
		if err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return
		}

		err = schema.Validate(data)
		if err != nil {
			fmt.Printf("❌ Validation failed: %v\n", err)
			return
		}

		fmt.Println("✅ OpenRPC document is valid!")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
