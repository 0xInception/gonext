package constants

import "strings"

type Region string

// Technically, Region = EUNE, Platform = eun1

const (
	BR   Region = "br1"
	EUNE Region = "eun1"
	EUW  Region = "euw1"
	JP   Region = "jp1"
	KR   Region = "kr"
	LAN  Region = "la1"
	LAS  Region = "la2"
	NA   Region = "na1"
	OCE  Region = "oc1"
	PBE  Region = "pbe1"
	PH   Region = "ph2"
	RU   Region = "ru"
	SG   Region = "sg2"
	TH   Region = "th2"
	TR   Region = "tr1"
	TW   Region = "tw2"
	VN   Region = "vn2"
)

var regionMap = map[string]Region{
	"BR":   BR,
	"EUNE": EUNE,
	"EUW":  EUW,
	"JP":   JP,
	"KR":   KR,
	"LAN":  LAN,
	"LAS":  LAS,
	"NA":   NA,
	"OCE":  OCE,
	"PBE":  PBE,
	"PH":   PH,
	"RU":   RU,
	"SG":   SG,
	"TH":   TH,
	"TR":   TR,
	"TW":   TW,
	"VN":   VN,
}

// GetRegionByString You can provide a region string such as `EUNE` and it will
// return the equivalent constant value `eun1`
func GetRegionByString(str string) Region {
	return regionMap[strings.ToUpper(str)]
}
