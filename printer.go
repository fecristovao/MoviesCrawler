package moviescrawler

import (
	"fmt"
)

func (movies FoundMovies) PrintFoundMovies() {
	for _, m := range movies {
		fmt.Printf("%s\n%s\n%s\n\n", m.Title, m.Cover, m.Link)
	}
}

func (downloads FoundMagnetLinks) PrintFoundMagnetLinks() {
	for _, download := range downloads {
		fmt.Printf("%s\n%s\n\n", download.Title, download.MagnetLink)
	}
}
