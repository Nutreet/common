package types

import (
	"time"
)

type ItemDoc struct {
	Name        string  `firestore:"name"`
	Price       int     `firestore:"price"`
	Company     string  `firestore:"company"`
	PortionSize float32 `firestore:"portionSize"`
	Calories    float32 `firestore:"calories"`
	Protein     float32 `firestore:"protein"`
	Carbs       float32 `firestore:"carbs"`
	Fat         float32 `firestore:"fat"`
}

func NewItemDoc(name string, price int, company string, portionSize float32, calories float32, protein float32, carbs float32, fat float32) *ItemDoc {
	return &ItemDoc{
		Name:        name,
		Price:       price,
		Company:     company,
		PortionSize: portionSize,
		Calories:    calories,
		Protein:     protein,
		Carbs:       carbs,
		Fat:         fat,
	}
}

type MealDoc struct {
	Uid       string    `firestore:"uid"`
	Breakfast []ItemDoc `firestore:"breakfast"`
	Lunch     []ItemDoc `firestore:"lunch"`
	Dinner    []ItemDoc `firestore:"dinner"`
	Snacks    []ItemDoc `firestore:"snacks"`
	Date      time.Time `firestore:"date"`
}

func NewMealDoc(uid string, breakfast []ItemDoc, lunch []ItemDoc, dinner []ItemDoc, snacks []ItemDoc, date time.Time) *MealDoc {
	return &MealDoc{
		Uid:       uid,
		Breakfast: breakfast,
		Lunch:     lunch,
		Dinner:    dinner,
		Snacks:    snacks,
		Date:      date,
	}
}
