package home

var (
	version   string
	buildDate string
)

const author = "ryan"

type VersionInfo struct {
	Version   string
	Author    string
	BuildDate string
}

type WishInfo struct {
	WishMessage string
}

func ReadVersionInfo() *VersionInfo {

	version = "v0.0.1"
	buildDate = "2022-10-10"

	return &VersionInfo{
		Version:   version,
		Author:    author,
		BuildDate: buildDate,
	}
}

func ReadWishInfo() *WishInfo {

	version = "v0.0.1"
	buildDate = "2022-10-10"

	return &WishInfo{
		WishMessage: "希望家人永远康健，不受老病困扰",
	}
}
