package constants

import "time"

type Request struct {
	Data string `json:"data"`
	Elid string `json:"elid"`
	Slid string `json:"slid"`
	Bid  string `json:"bid"`
	Id   string `json:"id"`
}

type Commands struct {
	Id      string    `json:"ID"`
	BSid    string    `json:"BID"`
	Data    string    `json:"DATA"`
	Sent    string    `json:"SENT"`
	Ack     string    `json:"ACK"`
	NeedAck string    `json:"NEEDACK"`
	Sendon  time.Time `json:"SENDON"`
	Ts      time.Time `json:"TS"`
}

type SatData struct {
	ID   string `db:"id"`
	Key  string `db:"k"`
	Val  string `db:"v"`
	MID  string `db:"mid"`
	TS   string `db:"ts"`
	T    string `db:"t"`
	Name string `db:"t"`
}

type Stations struct {
	ID string `db:"id"`
}

type GPSData struct {
	Lat string `db:"lat"`
	Lon string `db:"lon"`
	TS  string `db:"ts"`
}

type SatImages struct {
	Id       string `json:"ID"`
	Filename string `json:"FILENAME"`
	Block    string `json:"BLOCK"`
	Part     string `json:"PART"`

	Data []byte    `json:"DATA"`
	Ts   time.Time `json:"TS"`
}
