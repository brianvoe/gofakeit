package gofakeit

import "testing"

func TestGetArticle(t *testing.T) {
	// If nothing is passed return empty
	if getArticle("") != "" {
		t.Error("Expected empty string, got: ", getArticle(""))
	}

	// If the word starts with a, e, i, o, use an article
	if getArticle("apple") != "an" {
		t.Error("Expected an, got: ", getArticle("apple"))
	}

	if getArticle("elephant") != "an" {
		t.Error("Expected an, got: ", getArticle("elephant"))
	}

	if getArticle("independent") != "an" {
		t.Error("Expected an, got: ", getArticle("independent"))
	}

	if getArticle("open") != "an" {
		t.Error("Expected an, got: ", getArticle("open"))
	}

	// If the word starts with a u and n or l, use an article
	if getArticle("underwear") != "an" {
		t.Error("Expected an, got: ", getArticle("underwear"))
	}

	if getArticle("ulcer") != "an" {
		t.Error("Expected an, got: ", getArticle("ulcer"))
	}

	// If the word starts with a vowel, use an article
	if getArticle("hippos") != "an" {
		t.Error("Expected an, got: ", getArticle("hippose"))
	}

	// Pass a word that will return "a"
	if getArticle("bus") != "a" {
		t.Error("Expected a, got: ", getArticle("bus"))
	}

	if getArticle("plane") != "a" {
		t.Error("Expected a, got: ", getArticle("plane"))
	}
}
