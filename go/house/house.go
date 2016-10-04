package house

import "strings"

var phrases = []string{
	"the malt\nthat lay in",
	"the rat\nthat ate",
	"the cat\nthat killed",
	"the dog\nthat worried",
	"the cow with the crumpled horn\nthat tossed",
	"the maiden all forlorn\nthat milked",
	"the man all tattered and torn\nthat kissed",
	"the priest all shaven and shorn\nthat married",
	"the rooster that crowed in the morn\nthat woke",
	"the farmer sowing his corn\nthat kept",
	"the horse and the hound and the horn\nthat belonged to",
}

func Embed(relPhrase, nounPhrase string) string {
	return (relPhrase + " " + nounPhrase)
}

func Verse(subject string, relPhrase []string, nounPhrase string) string {
	if len(relPhrase) <= 0 {
		return Embed(subject, nounPhrase)
	}

	return Verse(Embed(subject, relPhrase[0]), relPhrase[1:], nounPhrase)
}

func Song() string {
	relPhrase := "This is"
	nounPhrase := "the house that Jack built."

	verses := make([]string, 0)

	for i := 0; i <= len(phrases); i++ {
		verse := Verse(relPhrase, reverse(phrases[:i]), nounPhrase)
		verses = append(verses, verse)
	}

	return strings.Join(verses, "\n\n")
}

func reverse(slice []string) []string {
	l := len(slice)
	var reversed = make([]string, l)
	for i, v := range slice {
		reversed[l-i-1] = v
	}
	return reversed
}
