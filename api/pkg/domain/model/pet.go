package model

import (
	"errors"
	"time"
)

type Pet struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	Sex      PetSex    `json:"sex"`
	Weight   float64   `json:"weight"`
	Birthday time.Time `json:"birthday"`
}

// (Qiita) Goのinterfaceがわからない人へ | https://qiita.com/rtok/items/46eadbf7b0b7a1b0eb08
// ===========
// PetSex
// ===========
type PetSex interface {
	String() string
	IsFemal() bool
	IsMale() bool
}

func NewPetSex(s string) (PetSex, error) {
	switch s {
	case "1":
		return maleSex, nil
	case "2":
		return femaleSex, nil
	default:
		return nil, errors.New("invalid pet sex")
	}
}

var (
	maleSex   PetSex = male{}
	femaleSex PetSex = female{}
)

type male struct{}

// IsFemal implements PetSex.
func (male) IsFemal() bool {
	return false
}

// IsMale implements PetSex.
func (male) IsMale() bool {
	return true
}

// String implements PetSex.
func (male) String() string {
	return "オス"
}

type female struct{}

// IsFemal implements PetSex.
func (female) IsFemal() bool {
	return true
}

// IsMale implements PetSex.
func (female) IsMale() bool {
	return false
}

// String implements PetSex.
func (female) String() string {
	return "メス"
}

type PetStatus struct {
	Food   PetCondition
	Swaet  PetCondition
	Toilet PetCondition
}

type PetCondition interface {
	Score() int64
	Value() string
}

var (
	ConditionExcellent PetCondition = conditionExcellent{}
	ConditionGood      PetCondition = conditionGood{}
	ConditionPoor      PetCondition = conditionPoor{}
)

type conditionExcellent struct{}

// Score implements PetCondition.
func (conditionExcellent) Score() int64 {
	return 3
}

// Value implements PetCondition.
func (conditionExcellent) Value() string {
	return "◎"
}

type conditionGood struct{}

// Score implements PetCondition.
func (conditionGood) Score() int64 {
	return 2
}

// Value implements PetCondition.
func (conditionGood) Value() string {
	return "○"
}

type conditionPoor struct{}

// Score implements PetCondition.
func (conditionPoor) Score() int64 {
	return 1
}

// Value implements PetCondition.
func (conditionPoor) Value() string {
	return "△"
}

type PetMemo struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
}

type ScheduledEvent struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	Timestamp time.Time `json:"timestamp"`
}

type PetSummary struct {
	Detail    *Pet            `json:"detail"`
	Status    *PetStatus      `json:"status"`
	Memo      *PetMemo        `json:"memo"`
	Schedules *ScheduledEvent `json:"schedules"`
}
