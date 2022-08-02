// Copyright © 2019 NAME HERE <andreas.fritzler@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/afritzler/search-conversion/pkg/types"
	"github.com/afritzler/search-conversion/pkg/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

const HelpSearchAPIURL = "https://help.sap.com/http.svc/search"
const HelpBaseURL = "https://help.sap.com"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "search-conversion",
	Short: fmt.Sprintf("search-conversion Version %v", utils.Version),
	Long:  fmt.Sprintf("search-conversion Version %v", utils.Version),
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		log.Printf("starting to serve on http://%v:%v...\n", name, port)
		registerHandlers()
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.help.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".help" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".help")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func registerHandlers() {
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/", returnOK)
}

func returnOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if code, err := w.Write([]byte("OK")); err != nil {
		log.Fatalf("failed to write OK response - got %d: %v", code, err)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		if code, err := w.Write([]byte("OK")); err != nil {
			log.Fatalf("failed to write OK response - got %d: %v", code, err)
		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		var request types.Request
		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		responseType := request.ResponseType
		language := "en-US"
		if request.Language != "" {
			language = request.Language
		}
		var replies []interface{}
		for _, product := range request.Products {
			client := http.Client{
				Timeout: time.Second * 15,
			}

			req, err := http.NewRequest(http.MethodGet, HelpSearchAPIURL, nil)
			if err != nil {
				replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
				break
			}

			req.Header.Set("User-Agent", "search-conversion")
			query := request.Converstation.Memory.Query
			q := req.URL.Query()
			q.Add("state", "PRODUCTION")
			q.Add("product", product.Name)
			q.Add("q", query)
			q.Add("version", product.Version)
			q.Add("language", language)
			req.URL.RawQuery = q.Encode()

			res, getErr := client.Do(req)
			if getErr != nil {
				replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
				break
			}

			body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
				replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
				break
			}
			defer r.Body.Close()

			response := types.Response{}
			err = json.Unmarshal([]byte(body), &response)
			if err != nil {
				replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
				break
			}

			if len(response.Data.Results) == 0 {
				replies = append(replies, generateTextMessage(fmt.Sprintf("Sorry, but I couldn't find anything for '%v' in the documentation! Please try rewording your question!", query), 0))
				break
			}

			var max int
			if product.MaxResults > len(response.Data.Results) {
				max = len(response.Data.Results)
			} else {
				max = product.MaxResults
			}
			switch responseType {
			case types.ButtonsType:
				var buttons []types.Button
				for i := 0; i < max; i++ {
					r := response.Data.Results[i]
					buttons = append(buttons, types.Button{
						Title: r.Title,
						Type:  "web_url",
						Value: HelpBaseURL + r.URL,
					})
				}
				replies = append(replies, types.Buttons{
					Type: types.ButtonsType,
					Content: types.ButtonsContent{
						Title:   "Here is what I found:",
						Buttons: buttons,
					},
				})
			case types.TextType:
				replies = append(replies, generateTextMessage(response.Data.Results[0].URL, 0))
			case types.CardType:
				var buttons []types.Button
				r := response.Data.Results[0]
				buttons = append(buttons, types.Button{
					Title: r.Title,
					Type:  "web_url",
					Value: HelpBaseURL + r.URL,
				})
				replies = append(replies, types.Card{
					Type: types.CardType,
					Content: types.CardContent{
						Title:    response.Data.Results[0].Title,
						SubTitle: response.Data.Results[0].Description,
						ImageURL: "",
						Buttons:  buttons,
					},
				})
			case types.CarouselType:
				var content []types.CardContent
				for i := 0; i < max; i++ {
					var buttons []types.Button
					r := response.Data.Results[i]
					buttons = append(buttons, types.Button{
						Title: r.Title,
						Type:  "web_url",
						Value: HelpBaseURL + r.URL,
					})
					content = append(content, types.CardContent{
						Title:    r.Title,
						SubTitle: r.Description,
						ImageURL: "",
						Buttons:  buttons,
					})
				}
				replies = append(replies, types.Carousel{
					Type:    types.CarouselType,
					Content: content,
				})
			default:
				// didn't find any matching type
				replies = append(replies, generateTextMessage("Sorry, but this response type is not supported!", 0))
			}
		}
		output, err := json.Marshal(types.Replies{Replies: replies})
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		if code, err := w.Write(output); err != nil {
			log.Fatalf("failed to write output - got %d: %v", code, err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		if code, err := w.Write([]byte("StatusMethodNotAllowed")); err != nil {
			log.Fatalf("failed to write default response - got %d: %v", code, err)
		}
	}
}

func generateTextMessage(text string, delay int) types.TextMessage {
	return types.TextMessage{
		Type:    types.ButtonsType,
		Content: text,
		Delay:   delay,
	}
}
