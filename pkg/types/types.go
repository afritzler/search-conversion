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

package types

const (
	ButtonsType      = "buttons"
	TextType         = "text"
	CardType         = "card"
	QuickRepliesType = "quickReplies"
	CarouselType     = "carousel"
	ListType         = "list"
)

// TextMessage defines a response of type text message.
// Example:
// {
// 	"type": "text",
//	"delay": 2,
//	"content": "MY_TEXT",
// }
type TextMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Delay   int    `json:"delay,omitempty"`
}

// QuickReplies defines a response of type text message.
// Example:
// {
// 	"type": "quickReplies",
// 	"content": {
// 		"title": "TITLE",
// 		"buttons": [
// 			{
// 				"title": "BUTTON_TITLE",
// 				"value": "BUTTON_VALUE"
// 			}
// 		]
// 	}
// }
type QuickReplies struct {
	Type    string              `json:"type"`
	Content QuickRepliesContent `json:"content"`
}

// QuickRepliesContent defines a subtype of the QuickReplies type.
type QuickRepliesContent struct {
	Title   string                `json:"title"`
	Buttons []QuickRepliesButtons `json:"buttons"`
}

// QuickRepliesButtons defines a subtype of the QuickRepliesContent type.
type QuickRepliesButtons struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// Card defines a response of type card.
// Example:
// {
//     "type": "card",
//     "content": {
//       "title": "CARD_TITLE",
//       "subtitle": "CARD_SUBTITLE",
//       "imageUrl": "IMAGE_URL",
//       "buttons": [
//         {
//           "title": "BUTTON_TITLE",
//           "type": "BUTTON_TYPE",
//           "value": "BUTTON_VALUE"
//         }
//       ]
//     }
// }
type Card struct {
	Type    string      `json:"type"`
	Content CardContent `json:"content"`
}

// CardContent defines a subtype of the Card type.
type CardContent struct {
	Title    string   `json:"title"`
	SubTitle string   `json:"subtitle,omitempty"`
	ImageURL string   `json:"imageUrl,omitempty"`
	Buttons  []Button `json:"buttons"`
}

// Button defines a subtype for buttons.
type Button struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Buttons defines a response of type buttons.
// Example:
// {
//     "type": "buttons",
//     "content": {
//       "title": "BUTTON_TITLE",
//       "buttons": [
//         {
//           "title": "BUTTON_TITLE",
//           "type": "BUTTON_TYPE",
//           "value": "BUTTON_VALUE"
//         }
//       ]
//     }
// }
type Buttons struct {
	Type    string         `json:"type"`
	Content ButtonsContent `json:"content"`
}

// ButtonsContent defines a subtype of the ButtonsType.
type ButtonsContent struct {
	Title   string   `json:"title"`
	Buttons []Button `json:"buttons"`
}

// Carousel defines a response of type buttons.
// Example:
// {
//     "type": "carousel",
//     "content": [
//       {
//         "title": "CARD_1_TITLE",
//         "subtitle": "CARD_1_SUBTITLE",
//         "imageUrl": "IMAGE_URL",
//         "buttons": [
//           {
//             "title": "BUTTON_1_TITLE",
//             "type": "BUTTON_1_TYPE",
//             "value": "BUTTON_1_VALUE"
//           }
//         ]
//       }
//     ]
// }
type Carousel struct {
	Type    string        `json:"type"`
	Content []CardContent `json:"content"`
}

// List defines a response of type buttons.
// Example:
// {
//     "type": "list",
//     "content": {
//       "elements": [
//         {
//           "title": "ELEM_1_TITLE",
//           "imageUrl": "IMAGE_URL",
//           "subtitle": "ELEM_1_SUBTITLE",
//           "buttons": [
//             {
//               "title": "BUTTON_1_TITLE",
//               "type": "BUTTON_TYPE",
//               "value": "BUTTON_1_VALUE"
//             }
//           ]
//         }
//       ],
//       "buttons": [
//         {
//           "title": "BUTTON_1_TITLE",
//           "type": "BUTTON_TYPE",
//           "value": "BUTTON_1_VALUE"
//         }
//       ]
//     }
// }
type List struct {
	Type    string      `json:"type"`
	Content ListContent `json:"content"`
}

// ListContent defines a subtype of the List type.
type ListContent struct {
	Elements []CardContent `json:"elements"`
	Buttons  []Button      `json:"buttons"`
}

// Picture defines a response of type buttons.
// Example:
// {
//     "type": "picture",
//     "content": "IMAGE_URL",
// }
type Picture struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// Video defines a response of type buttons.
// Example:
// {
//     "type": "video",
//     "content": "VIDEO_URL",
// }
type Video struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// Request is a definition of the request type.
// Example:
// {
//     "response_type": "carousel",
// 		"products":[
// 			{
// 				"name": "PRODUCT_NAME",
// 				"version": "VERSION",
// 				"max_results": 10
// 			}
// 		],
//     "conversation": {
//         "memory": {
// 				"query": "{{memory.query}}"
// 			}
//     }
// }
type Request struct {
	ResponseType  string       `json:"response_type"`
	Products      []Product    `json:"products"`
	Converstation Conversation `json:"conversation"`
}

// Product is a subtype of the Request type.
type Product struct {
	Name       string `json:"name,omitempty"`
	Version    string `json:"version,omitempty"`
	MaxResults int    `json:"max_results,omitempty"`
}

// Conversation is a subtype of the Request type.
type Conversation struct {
	Memory Memory `json:"memory"`
}

// Memory is a subtype of the Conversation type.
type Memory struct {
	Query string `json:"query"`
}

// Replies
type Replies struct {
	Replies []interface{} `json:"replies"`
}

type Reply struct {
}

type Response struct {
	Status string       `json:"status"`
	Data   DataResponse `json:"data"`
}

type DataResponse struct {
	Query          string            `json:"query"`
	MaxResults     int               `json:"maxResults"`
	Results        []ResultResponse  `json:"results"`
	ProductResults string            `json:"productResults"`
	Products       []ProductResponse `json:"products"`
}

type ProductResponse struct {
	Comments     string `json:"comments"`
	Date         string `json:"date"`
	Description  string `json:"description"`
	DocumentType string `json:"documentType"`
	Format       string `json:"format"`
	MimeType     string `json:"mimeType"`
	Product      string `json:"product"`
	Size         string `json:"size"`
	State        string `json:"state"`
	Title        string `json:"title"`
	Transtype    string `json:"transtype"`
	URL          string `json:"url"`
	Version      string `json:"version"`
	Views        string `json:"views"`
}

type ResultResponse struct {
	Date         string `json:"date"`
	Description  string `json:"description"`
	DocumentType string `json:"documentType"`
	Format       string `json:"format"`
	MimeType     string `json:"mimeType"`
	Product      string `json:"product"`
	Size         string `json:"size"`
	State        string `json:"state"`
	Title        string `json:"title"`
	Transtype    string `json:"transtype"`
	URL          string `json:"url"`
	Version      string `json:"version"`
}
