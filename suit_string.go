// Code generated by "stringer -type=Suit"; DO NOT EDIT

package main

import "fmt"

const _Suit_name = "HeartDiamondSpadeClub"

var _Suit_index = [...]uint8{0, 5, 12, 17, 21}

func (i Suit) String() string {
	i -= 1
	if i < 0 || i >= Suit(len(_Suit_index)-1) {
		return fmt.Sprintf("Suit(%d)", i+1)
	}
	return _Suit_name[_Suit_index[i]:_Suit_index[i+1]]
}
