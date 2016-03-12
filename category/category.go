package category

import "log"

type Category int

const (
	Hotentry Category = 1 + iota
	General
	Social
	Economics
	Life
	Knowledge
	It
	Fun
	Entertainment
	Game
)

func (c Category) String() string {
	switch c {
	case Hotentry:
		return "総合"
	case General:
		return "一般"
	case Social:
		return "世の中"
	case Economics:
		return "政治と経済"
	case Life:
		return "暮らし"
	case Knowledge:
		return "学び"
	case It:
		return "テクノロジー"
	case Fun:
		return "おもしろ"
	case Entertainment:
		return "エンタメ"
	case Game:
		return "アニメとゲーム"
	}
	log.Fatal("can't be here")
	return ""
}

func (c Category) Id() string {
	switch c {
	case Hotentry:
		return ""
	case General:
		return "general"
	case Social:
		return "social"
	case Economics:
		return "economics"
	case Life:
		return "life"
	case Knowledge:
		return "knowledge"
	case It:
		return "it"
	case Fun:
		return "fun"
	case Entertainment:
		return "entertainment"
	case Game:
		return "game"
	}
	log.Fatal("can't be here")
	return ""
}
