package data

import "strconv"

type GalleryImage struct {
	Path     string
	Alt      string
	Location string
	Year     string
}

type Location struct {
	Name        string
	State       string
	Description string
}

type ImpactStat struct {
	Value string
	Label string
}

const OrgName = "Zoie Foundation"
const YearsActive = 7
const FlagshipEvent = "Feed The Street"

var AboutParagraphs = []string{
	"Zoie Foundation is a community-focused organization committed to meeting real needs, restoring dignity, and giving people a reason to hope. Through practical acts of service, we reach individuals and families in underserved communities with the things that matter in the moments they need them most.",

	"For " + strconv.Itoa(YearsActive) + " years, our work has taken us into communities across Nigeria — providing foodstuffs to families, supporting people facing difficult circumstances, helping with hospital bills, and providing learning materials for primary school children. What began as a simple desire to help has grown into a continuing commitment to show up, serve, and make a tangible difference.",

	"At Zoie Foundation, we believe impact does not always have to begin with something big. Sometimes it begins with a bag of food, a child's school materials, a hospital bill that gets paid, or simply showing someone that their situation has not gone unnoticed.",
}

var ImpactStats = []ImpactStat{
	{Value: strconv.Itoa(YearsActive) + "+", Label: "Years of Service"},
	{Value: "6+", Label: "Communities Reached"},
	{Value: "5+", Label: "States & Locations"},
	{Value: "∞", Label: "Lives Worth Impacting"},
}

var Locations = []Location{
	{
		Name:        "Elewura",
		State:       "Oyo State",
		Description: "Our 2025 outreach in Elewura brought the Feed The Street initiative into the Oyo community, extending practical support to people who needed it.",
	},
	{
		Name:        "Obiagu",
		State:       "Enugu State",
		Description: "In 2024, we reached Obiagu through a school-focused outreach, supporting primary school children with practical learning materials.",
	},
	{
		Name:        "Snake Island",
		State:       "Lagos State",
		Description: "Our 2023 outreach in Snake Island brought volunteers together to provide food support and connect directly with families in the community.",
	},
	{
		Name:        "Makoko",
		State:       "Lagos State",
		Description: "In 2022, Zoie Foundation reached children and families in Makoko with food support and a message of care and dignity.",
	},
	{
		Name:        "Nzam",
		State:       "Anambra State",
		Description: "Our 2021 outreach took us to Nzam, where volunteers came together to provide food support to people within the community.",
	},
	{
		Name:        "Lagos",
		State:       "Lagos State",
		Description: "The journey began in Lagos in 2020, marking the first of a growing series of community outreaches that continue to shape Zoie Foundation's work.",
	},
}

var Gallery = []GalleryImage{
	{
		Path:     "https://59doenkwa7.ufs.sh/f/jhH9iucglD1psBA9ue8QuR0NBCwsYLMl2kVASg7Dt59dciv6",
		Alt:      "Feed The Street outreach in Oyo",
		Location: "Elewura, Oyo",
		Year:     "2025",
	},
	{
		Path:     "https://59doenkwa7.ufs.sh/f/jhH9iucglD1prGhs8UzamMENiBI9WFdK0egtZLxf5Gu4XkQz",
		Alt:      "School outreach with primary school children in Enugu",
		Location: "Obiagu, Enugu",
		Year:     "2024",
	},
	{
		Path:     "https://59doenkwa7.ufs.sh/f/jhH9iucglD1pxodje22TFV4L8BAml7UbSIdOpwtYah9o2ZWK",
		Alt:      "Volunteers preparing food support for the Snake Island outreach",
		Location: "Snake Island, Lagos",
		Year:     "2023",
	},
	{
		Path:     "https://59doenkwa7.ufs.sh/f/jhH9iucglD1ptk7on2LqIZ81WB5MLFrp09xvVdblN3XzgUhs",
		Alt:      "Children receiving food support in Lagos",
		Location: "Makoko, Lagos",
		Year:     "2022",
	},
	{
		Path:     "https://59doenkwa7.ufs.sh/f/jhH9iucglD1p5gZi11pg6eLHpOUzn1suTKhxQ3yvBRaES50D",
		Alt:      "Volunteers providing food support in Nzam",
		Location: "Nzam, Anambra",
		Year:     "2021",
	},
	{
		Path:     "https://59doenkwa7.ufs.sh/f/jhH9iucglD1pff8P7z50eHh1cU8M7CFWGgNwmbrk3soJAqEa",
		Alt:      "Zoie Foundation's first community outreach",
		Location: "Lagos",
		Year:     "2020",
	},
}