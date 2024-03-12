package kafkaconsumer

import "time"

type Event struct {
	RequestId       string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	DataContentType string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            []byte    `json:"data"`
}
