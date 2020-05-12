package moviescrawler

type Movie struct {
	Title, Cover, Link string
}

type DownloadLink struct {
	Title, MagnetLink string
}

type WebRequest struct {
	Header map[string]string
	Params map[string]string
}

type FoundMagnetLinks []DownloadLink
type FoundMovies []Movie

func recoveryPanic() {
	if pan := recover(); pan != nil {
		//fmt.Printf("Panic: %v\n", pan)
	}
}
