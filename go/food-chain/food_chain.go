package foodchain

import "strings"

const testVersion = 2

var Animals = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var ReverseAnimals = []string{"horse", "cow", "goat", "dog", "cat", "bird", "spider", "fly"}
var LastVerse = "I know an old lady who swallowed a horse.\nShe's dead, of course!"
var lastLine = "I don't know why she swallowed the fly. Perhaps she'll die."

func Verse(verse int) string {
	phrases := make([]string, 0)
	animal := Animals[verse-1]

	if verse == 1 {
		phrases = append(phrases, firstLine(animal))
		phrases = append(phrases, lastLine)

		return strings.Join(phrases, "\n")
	}

	if verse == 8 {
		return LastVerse
	}

	phrases = append(phrases, firstLine(animal))
	phrases = append(phrases, catchPhraseFor(animal))

	reverse := foodChainForAnimal(animal)

	for _, a := range reverse {
		l := phraseForFoodChain(a)
		if l == "" {
			break
		}
		phrases = append(phrases, l)
	}

	phrases = append(phrases, lastLine)
	return strings.Join(phrases, "\n")
}

func Verses(start, stop int) string {
	var verses []string
	for i := start; i < stop; i++ {
		verses = append(verses, Verse(i))
		verses = append(verses, "")
	}
	verses = append(verses, Verse(stop))

	return strings.Join(verses, "\n")
}

func Song() string {
	return Verses(1, 8)
}

func firstLine(animal string) string {
	return "I know an old lady who swallowed a " + animal + "."
}

func eatenAnimal(animal string) string {
	for i, a := range Animals {
		if a == animal {
			return Animals[i-1]
		}
	}
	return animal
}

func catchPhraseFor(animal string) string {
	switch {
	case animal == "spider":
		return "It wriggled and jiggled and tickled inside her."
	case animal == "bird":
		return "How absurd to swallow a bird!"
	case animal == "cat":
		return "Imagine that, to swallow a cat!"
	case animal == "dog":
		return "What a hog, to swallow a dog!"
	case animal == "goat":
		return "Just opened her throat and swallowed a goat!"
	case animal == "cow":
		return "I don't know how she swallowed a cow!"
	default:
		return ""
	}
}

func phraseForFoodChain(animal string) string {
	switch {
	case animal == "fly":
		return "She swallowed the spider to catch the fly."
	case animal == "spider":
		return "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her."
	case animal == "bird":
		return "She swallowed the cat to catch the bird."
	case animal == "cat":
		return "She swallowed the dog to catch the cat."
	case animal == "dog":
		return "She swallowed the goat to catch the dog."
	case animal == "goat":
		return "She swallowed the cow to catch the goat."
	default:
		return ""
	}
}

func foodChainForAnimal(animal string) []string {
	for i, a := range ReverseAnimals {
		if a == animal {
			return ReverseAnimals[i+1:]
		}
	}
	return nil
}

func reverse(strings []string) []string {
	for i := 0; i < len(strings)/2; i++ {
		j := len(strings) - i - 1
		strings[i], strings[j] = strings[j], strings[i]
	}
	return strings
}
