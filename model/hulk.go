package model

type Hulk struct {
	Name      string                 `json:"name" bson:"name" bw:"name"`
	Version   string                 `json:"version" bson:"version" bw:"version"`
	Time      string                 `json:"time" bson:"time"`
	Configure map[string]interface{} `json:"configure" bson:"configure"`
}
