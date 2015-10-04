package models

type SomethingModel struct {
  // The struct keys are tagged with json
  // because go uses uppercase to determine if it is public
  // adding the json tag allows you to tag what the response key
  // will be - e.g. remove uppercase
  Some string `json:"some"`
  Thing int `json:"thing"`
}
