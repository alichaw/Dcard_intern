// ad_models.go
package main

import "time"

type GenderType string
type CountryType string
type PlatformType string

type Condition struct {
	AgeStart  *int           `json:"ageStart,omitempty"`
	AgeEnd    *int           `json:"ageEnd,omitempty"`
	Genders   []GenderType   `json:"genders,omitempty"`
	Countries []CountryType  `json:"countries,omitempty"`
	Platforms []PlatformType `json:"platforms,omitempty"`
}

type Ad struct {
	ID         int         `json:"id"`
	Title      string      `json:"title"`
	StartAt    time.Time   `json:"startAt"`
	EndAt      time.Time   `json:"endAt"`
	Conditions []Condition `json:"conditions"`
}
