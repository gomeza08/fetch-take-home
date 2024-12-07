package models

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

type Receipt struct {
	ID           uuid.UUID `json:"id"`
	Retailer     string    `json:"retailer" binding:"required"`
	PurchaseDate string    `json:"purchaseDate" binding:"required"`
	PurchaseTime string    `json:"purchaseTime" binding:"required"`
	Items        []Item    `json:"items" binding:"required"`
	Total        string    `json:"total" binding:"required"`
	Points       int       `json:"points"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"`
}

func (r Receipt) CalculatePoints() int {
	return countRetailerNamePoints(r.Retailer) +
		roundDollarAmountPoints(r.Total) +
		multipleOfTwentyFivePoints(r.Total) +
		everyTwoItemsPoints(len(r.Items)) +
		itemDescriptionLengthPoints(r.Items) +
		purchaseDayIsOddPoints(r.PurchaseDate) +
		purchaseTimePoints(r.PurchaseTime)
}

func countRetailerNamePoints(retailerName string) int {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return utf8.RuneCountInString(reg.ReplaceAllString(retailerName, ""))
}

func roundDollarAmountPoints(receiptTotal string) int {
	if strings.HasSuffix(receiptTotal, ".00") {
		return 50
	}
	return 0
}

func multipleOfTwentyFivePoints(receiptTotal string) int {
	if strings.HasSuffix(receiptTotal, ".00") ||
		strings.HasSuffix(receiptTotal, ".25") ||
		strings.HasSuffix(receiptTotal, ".50") ||
		strings.HasSuffix(receiptTotal, ".75") {
		return 25
	}
	return 0
}

func everyTwoItemsPoints(itemCount int) int {
	return (itemCount / 2) * 5
}

func itemDescriptionLengthPoints(receiptItems []Item) int {
	totalPoints := 0
	for _, item := range receiptItems {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints := math.Ceil(itemPrice * .2)
			totalPoints += int(itemPoints)
		}
	}
	return totalPoints
}

func purchaseDayIsOddPoints(purchaseDate string) int {
	purchaseDateInt, _ := strconv.Atoi(purchaseDate[len(purchaseDate)-2:])
	if purchaseDateInt%2 == 1 {
		return 6
	}
	return 0
}

func purchaseTimePoints(purchaseTime string) int {
	if purchaseTime > "14:00" && purchaseTime < "16:00" {
		return 10
	}
	return 0
}
