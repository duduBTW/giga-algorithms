package day5

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type Side = map[int]map[int]bool

type Pages = [][]int
type ManualOrder struct {
	X int
	Y int
}

type Manual struct {
	OrderLeft  Side
	OrderRight Side
	Orders     []ManualOrder
	Pages      Pages
}

type ManualValidPages struct {
	Valid   Pages
	Invalid Pages
}

func ReadInput() (Manual, error) {
	manual := Manual{}
	content, err := os.ReadFile("D:\\Peronal\\giga-algorithms\\advent-2024\\day-5\\data-part-1.txt")
	if err != nil {
		return manual, err
	}

	splitContent := strings.Split(string(content), "\n\n")
	ordersContent := splitContent[0]
	pagesContent := splitContent[1]

	manual.OrderLeft = make(Side)
	manual.OrderRight = make(Side)

	for _, line := range strings.Split(ordersContent, "\n") {
		if line == "" {
			continue
		}

		ordersContent := strings.Split(line, "|")
		leftItem, err := strconv.Atoi(ordersContent[0])
		if err != nil {
			continue
		}

		rightItem, err := strconv.Atoi(ordersContent[1])
		if err != nil {
			continue
		}

		_, okRight := manual.OrderRight[leftItem]
		if !okRight {
			manual.OrderRight[leftItem] = make(map[int]bool)
		}
		manual.OrderRight[leftItem][rightItem] = true

		_, okLeft := manual.OrderLeft[rightItem]
		if !okLeft {
			manual.OrderLeft[rightItem] = make(map[int]bool)
		}
		manual.OrderLeft[rightItem][leftItem] = true

		manual.Orders = append(manual.Orders, ManualOrder{
			X: leftItem,
			Y: rightItem,
		})
	}

	for _, line := range strings.Split(pagesContent, "\n") {
		if line == "" {
			continue
		}

		pagesContent := strings.Split(line, ",")
		pages := make([]int, len(pagesContent))
		for index, pageContent := range pagesContent {
			page, err := strconv.Atoi(pageContent)
			if err != nil {
				continue
			}

			pages[index] = page
		}
		manual.Pages = append(manual.Pages, pages)
	}

	return manual, nil
}

type ValidParams struct {
	manual         Manual
	page           int
	pageCollection []int
	i              int
}

const Invalid = -1

func FindInvalidLeftIndex(params ValidParams) int {
	for j := 0; j < params.i; j++ {
		comparingPage := params.pageCollection[j]
		_, ok := params.manual.OrderRight[params.page][comparingPage]
		if ok {
			return j
		}
	}

	return Invalid
}

func FindInvalidRightIndex(params ValidParams) int {
	for j := params.i; j < len(params.pageCollection); j++ {
		comparingPage := params.pageCollection[j]
		_, ok := params.manual.OrderLeft[params.page][comparingPage]
		if ok {
			return j
		}
	}

	return Invalid
}

func IsValidPageCollection(manual Manual, pageCollection []int) (bool, []int) {
	var invalid = []int{Invalid, Invalid}

	for i := 1; i < len(pageCollection); i++ {
		params := ValidParams{
			manual:         manual,
			pageCollection: pageCollection,
			page:           pageCollection[i],
			i:              i,
		}

		invalidLeftIndex := FindInvalidLeftIndex(params)
		if invalidLeftIndex != -1 {
			return false, []int{i, invalidLeftIndex}
		}

		invalidRightIndex := FindInvalidRightIndex(params)
		if invalidRightIndex != -1 {
			return false, []int{i, invalidRightIndex}
		}

		// Has no middle index
		if len(pageCollection)%2 == 0 {
			return false, invalid
		}
	}

	return true, invalid
}

func FindPageCollectionMiddleIndex(pageCollection []int) int {
	return int(math.Ceil(float64(len(pageCollection))/2)) - 1
}

func ValidateManual(manual Manual) int {
	// manualValidPages := ManualValidPages{}
	total := 0

	for _, pageCollection := range manual.Pages {
		isValidPage, _ := IsValidPageCollection(manual, pageCollection)
		if !isValidPage {
			continue
		}

		middleIndex := FindPageCollectionMiddleIndex(pageCollection)
		total += pageCollection[middleIndex]
	}

	return total
}

func fixPageCollection(pageCollection, invalidPageIndexes []int) []int {
	fixedPagedCollection := make([]int, len(pageCollection))
	copy(fixedPagedCollection, pageCollection)

	firstIndex := invalidPageIndexes[0]
	secondIndex := invalidPageIndexes[1]

	first := fixedPagedCollection[firstIndex]
	second := fixedPagedCollection[secondIndex]

	// Flips index
	fixedPagedCollection[firstIndex] = second
	fixedPagedCollection[secondIndex] = first

	return fixedPagedCollection
}

func FixManual(manual Manual) int {
	total := 0

	for _, originalPageCollection := range manual.Pages {
		fixedPagedCollection := make([]int, len(originalPageCollection))
		copy(fixedPagedCollection, originalPageCollection)

		wasFixed := false
		isValidPageCollection := false

		for !isValidPageCollection {
			isValidFixedPageCollection, invalidPageIndexes := IsValidPageCollection(manual, fixedPagedCollection)
			if isValidFixedPageCollection {
				// Breaks out of the loop if the page is valid
				isValidPageCollection = true
				continue
			}

			// Fix page
			wasFixed = true
			fixedPagedCollection = fixPageCollection(fixedPagedCollection, invalidPageIndexes)
		}

		if wasFixed {
			middleIndex := FindPageCollectionMiddleIndex(fixedPagedCollection)
			total += fixedPagedCollection[middleIndex]
		}
	}

	return total
}

func roundUpInt(x, multiple int) int {
	if multiple == 0 {
		return x
	}
	return int(math.Ceil(float64(x)/float64(multiple)) * float64(multiple))
}
