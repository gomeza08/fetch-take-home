package models

import "testing"

func TestCountRetailerNamePoints(t *testing.T) {
	actual := countRetailerNamePoints("ABC")
	expected := 3
	if actual != expected {
		t.Errorf("TestCountRetailerNamePoints = %d; want %d", actual, expected)
	}

	actual = countRetailerNamePoints("ABC 123")
	expected = 6
	if actual != expected {
		t.Errorf("TestCountRetailerNamePoints = %d; want %d", actual, expected)
	}

	actual = countRetailerNamePoints("ABC, 123")
	expected = 6
	if actual != expected {
		t.Errorf("TestCountRetailerNamePoints = %d; want %d", actual, expected)
	}
}

func TestRoundDollarAmounts(t *testing.T) {
	actual := roundDollarAmountPoints("10.10")
	expected := 0
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}

	actual = roundDollarAmountPoints("10.00")
	expected = 50
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}
}

func TestMultipleOfTwentyFivePoints(t *testing.T) {
	actual := multipleOfTwentyFivePoints("10.10")
	expected := 0
	if actual != expected {
		t.Errorf("TestMultipleOfTwentyFivePoints = %d; want %d", actual, expected)
	}

	actual = multipleOfTwentyFivePoints("10.00")
	expected = 25
	if actual != expected {
		t.Errorf("TestMultipleOfTwentyFivePoints = %d; want %d", actual, expected)
	}

	actual = multipleOfTwentyFivePoints("10.25")
	expected = 25
	if actual != expected {
		t.Errorf("TestMultipleOfTwentyFivePoints = %d; want %d", actual, expected)
	}

	actual = multipleOfTwentyFivePoints("10.75")
	expected = 25
	if actual != expected {
		t.Errorf("TestMultipleOfTwentyFivePoints = %d; want %d", actual, expected)
	}

	actual = multipleOfTwentyFivePoints("10.50")
	expected = 25
	if actual != expected {
		t.Errorf("TestMultipleOfTwentyFivePoints = %d; want %d", actual, expected)
	}
}

func TestEveryTwoItemsPoints(t *testing.T) {
	actual := everyTwoItemsPoints(2)
	expected := 5
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}

	actual = everyTwoItemsPoints(5)
	expected = 10
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}
}

func TestItemDescriptionLengthPoints(t *testing.T) {
	item1 := Item{
		ShortDescription: "Shorty",
		Price:            "10.00",
	}
	actual := itemDescriptionLengthPoints([]Item{item1})
	expected := 1
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}

	item1 = Item{
		ShortDescription: "Shorty1",
		Price:            "10.00",
	}
	actual = itemDescriptionLengthPoints([]Item{item1})
	expected = 0
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}
}

func TestPurchaseDayIsOddPoints(t *testing.T) {
	actual := purchaseDayIsOddPoints("2022-01-01")
	expected := 6
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}

	actual = purchaseDayIsOddPoints("2022-01-02")
	expected = 0
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}
}

func TestPurchaseTimePoints(t *testing.T) {
	actual := purchaseTimePoints("14:01")
	expected := 10
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}

	actual = purchaseTimePoints("15:01")
	expected = 10
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}

	actual = purchaseTimePoints("16:01")
	expected = 0
	if actual != expected {
		t.Errorf("testRoundDollarAmounts = %d; want %d", actual, expected)
	}
}