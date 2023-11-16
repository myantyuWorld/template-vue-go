package model

import (
	"errors"
	"fmt"
	"time"
)

type Pet struct {
	ID       uint64
	Name     string
	Sex      PetSex
	Weight   float64
	Birthday time.Time
}

func (pet *Pet) Age(timestamp time.Time) *PetAge {
	// TODO not implemented
	return &PetAge{
		year:  3,
		month: 8,
	}
}

type PetAge struct {
	year  int64
	month int64
}

func (age *PetAge) FormatString() string {
	return fmt.Sprintf("%d歳%dヶ月", age.year, age.month)
}

type PetSex interface {
	String() string
	IsFemale() bool
	IsMale() bool
}

func NewPetSex(s string) (PetSex, error) {
	switch s {
	case maleSex.String():
		return maleSex, nil
	case femaleSex.String():
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

func (sex male) String() string {
	return "オス"
}

func (sex male) IsFemale() bool {
	return false
}

func (sex male) IsMale() bool {
	return true
}

type female struct{}

func (sex female) String() string {
	return "メス"
}

func (sex female) IsFemale() bool {
	return true
}

func (sex female) IsMale() bool {
	return false
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

func (cond conditionExcellent) Score() int64 {
	return 3
}

func (cond conditionExcellent) Value() string {
	return "◎"
}

type conditionGood struct{}

func (cond conditionGood) Score() int64 {
	return 2
}

func (cond conditionGood) Value() string {
	return "◯"
}

type conditionPoor struct{}

func (cond conditionPoor) Score() int64 {
	return 1
}

func (cond conditionPoor) Value() string {
	return "△"
}

type PetMemo struct {
	ID        uint64
	Title     string
	Timestamp time.Time
}

type ScheduledEvent struct {
	ID        uint64
	Title     string
	Location  string
	Timestamp time.Time
}

type PetSummary struct {
	Detail    *Pet
	Status    *PetStatus
	Memo      []*PetMemo
	Schedules []*ScheduledEvent
}
