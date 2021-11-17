package main

import (
	"bufio"
	"fmt"
	"os"
)

type Side bool

const (
	Left  Side = true
	Right Side = false
)

var (
	keySideMap = map[byte]Side{
		// Left
		'q': Left, 'w': Left, 'e': Left, 'r': Left, 't': Left,
		'a': Left, 's': Left, 'd': Left, 'f': Left, 'g': Left,
		'z': Left, 'x': Left, 'c': Left, 'v': Left, 'b': Left,

		// Right
		'y': Right, 'u': Right, 'i': Right, 'o': Right, 'p': Right,
		'h': Right, 'j': Right, 'k': Right, 'l': Right,
		'n': Right, 'm': Right, ',': Right, '.': Right,
	}
)

func main() {
	var (
		freeHandTime int
		beerHandTime int = 1000
		switchTime   int
		text         string
	)
	fmt.Scanln(&freeHandTime, &switchTime)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}
	leftStart := findBestTime(0, Left, 0, &freeHandTime, &beerHandTime, &switchTime, &text)
	rightStart := findBestTime(0, Right, 0, &freeHandTime, &beerHandTime, &switchTime, &text)
	if leftStart < rightStart {
		fmt.Println(leftStart)
	} else {
		fmt.Println(rightStart)
	}
}

func findBestTime(currentTime int, handSide Side, currentPosition int, freeHandTime *int, beerHandTime *int, switchTime *int, text *string) int {
	if currentPosition >= len(*text) {
		return currentTime
	}
	// Find the letter
	letter := (*text)[currentPosition]
	keySide := keySideMap[letter]
	// == Do possible actions ==
	if keySide == handSide || letter == ' ' { // Possibilities if key is in same side as free hand
		// Don't switch hands
		result := findBestTime(currentTime+(*freeHandTime), handSide, currentPosition+1, freeHandTime, beerHandTime, switchTime, text)
		return result
	} else { // Possibilities if key is in same side as beer hand
		// Don't switch hands
		withoutSwitchTime := findBestTime(currentTime+(*beerHandTime), handSide, currentPosition+1, freeHandTime, beerHandTime, switchTime, text)
		// Switch hands
		withSwitchTime := findBestTime(currentTime+(*switchTime)+(*freeHandTime), !handSide, currentPosition+1, freeHandTime, beerHandTime, switchTime, text)
		// Choose lowest
		if withoutSwitchTime < withSwitchTime {
			return withoutSwitchTime
		} else {
			return withSwitchTime
		}
	}
}
