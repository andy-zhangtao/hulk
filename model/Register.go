package model

type Register struct {
	Name    string `json:"name" bson:"name" bw:"name"`
	Version string `json:"version" bson:"version" bw:"version"`
	IP      string `json:"ip" bson:"ip"`
	Resume  string `json:"resume" bson:"resume"`
	Time    string `json:"time" bson:"time"`
}

