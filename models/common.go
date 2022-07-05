package models

// rules
type Rule struct{
	Val int
	Text string
}

// var RoyalFlush = Rule{ 1, "ロイヤルフラッシュ" }
var StraightFlush = Rule{ 2, "ストレートフラッシュ" }
var FourOfAKIND = Rule{ 3, "フォー・オブ・ア・カインド" }
var FullHouse = Rule{ 4, "フルハウス" }
var Flush = Rule{ 5, "フラッシュ" }
var Straight = Rule{ 6, "ストレート" }
var ThreeOfAKind = Rule{ 7, "スリー・オブ・ア・カインド" }
var TwoPair = Rule{ 8, "ツーペア" }
var OnePair= Rule{ 9, "ワンペア" }
var NoPair = Rule{ 10, "ハイカード" }

// card struct
type Card struct {
	Suit   string
	Number int
}
