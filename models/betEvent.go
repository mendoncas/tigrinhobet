package models

type BetEvent struct {
	finished    bool
	happened    bool
	title       string
	description string
	bet         *[]Bet
}
