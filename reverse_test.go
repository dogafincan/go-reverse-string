package reverse

import (
	"testing"
)

func TestStringKorean(t *testing.T) {
	expected := "요세하녕안"
	actual := String("안녕하세요")

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}

func TestStringJapanese(t *testing.T) {
	expected := "はちにんこ"
	actual := String("こんにちは")

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}

func TestStringMandarin(t *testing.T) {
	expected := "好你"
	actual := String("你好")

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}

func TestStringThai(t *testing.T) {
	expected := "บรัคดีสวัส"
	actual := String("สวัสดีครับ")

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}

func TestStringVietnamese(t *testing.T) {
	expected := "oàhc nix"
	actual := String("xin chào")

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}

func TestStringEnglish(t *testing.T) {
	expected := "olleh"
	actual := String("hello")

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}
