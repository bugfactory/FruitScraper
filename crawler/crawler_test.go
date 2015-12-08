package crawler

import "testing"

func TestUrlSize(t *testing.T) {
	url := "https://github.com/bugfactory/posts/blob/master/README.md"

	size := UrlSize(url)

	if size != 21 {
		t.Error("Expected url size 21k, got ", size)
	}
}

func TestStringToFloat(t *testing.T) {
	result := StringToFloat("2.5")
	if result != 2.5 {
		t.Error("Expected Float ", result)
	}

	result = StringToFloat("2.")
	if result != 2 {
		t.Error("Expected Float ", result)
	}

	result = StringToFloat("-0.1")
	if result != -0.1 {
		t.Error("Expected Float ", result)
	}
}

func TestNew(t *testing.T) {
	url := "http://patito.github.io"
	c := New(url)

	if c == nil {
		t.Error("Should not be nil")
	}
}

func TestIntToString(t *testing.T) {
	s := IntToString(1)

	if s != "1kb" {
		t.Error("Expected 1kb and got ", s)
	}

	s = IntToString(45555)
	if s != "45555kb" {
		t.Error("Expected 1kb and got ", s)
	}
}

func TestGetPostInfo(t *testing.T) {
	url := "http://patito.github.io/2015/10/17/business-card/"

	s := GetPostInfo(url, ".post-date")
	if s != "17 Oct 2015" {
		t.Error("Expected 17 Oct 2015 and got ", s)
	}

	// Chewbacca Noise
	s = GetPostInfo(url, "growlgrowlgrowl")
	if s != "" {
		t.Error("Expected nothing and got ", s)
	}
}
