package moviescrawler

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type PirateTorrent struct{}

func (_ PirateTorrent) GetNumberOfPages(Term string) int {
	defer recoveryPanic()
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"

	html := WebRequest{headers, nil}.Get("https://piratetorrent.org/?s=" + url.QueryEscape(Term))
	regex := regexp.MustCompile(`(?ms)<a class="last" href=".*?/([0-9]+)/`)
	match := regex.FindStringSubmatch(html)

	if match == nil {
		return 1
	}

	pageResult, _ := strconv.Atoi(match[1])

	if pageResult == 0 {
		return 1
	} else {
		return pageResult
	}
}

func (_ PirateTorrent) SearchMovie(Term string, Page int) FoundMovies {
	defer recoveryPanic()
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"

	html := WebRequest{headers, nil}.Get("https://piratetorrent.org/page/" + strconv.Itoa(Page) + "/?s=" + url.QueryEscape(Term))
	regex := regexp.MustCompile(`(?s)<h2 class="entry-title" itemprop="headline"><a href="(.+?)">(.+?)<\/a>.*?<img.*?src="(.*?)"`)
	matches := regex.FindAllStringSubmatch(html, -1)

	var movies FoundMovies
	for _, match := range matches {
		movies = append(movies, Movie{Title: match[2], Cover: match[3], Link: match[1]})
	}

	return movies
}

func (_ PirateTorrent) GetDownloadLinks(Link string) FoundMagnetLinks {
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"
	html := WebRequest{headers, nil}.Get(Link)

	regex := regexp.MustCompile(`<h3 style="text-align: center;">(.+?):<\/h3>.*?id=(.+?)&ref=(.+?)&titulo=(.+?)"`)
	matches := regex.FindAllStringSubmatch(string(html), -1)
	var foundLinks FoundMagnetLinks
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	for _, match := range matches {
		wg.Add(1)
		go func(wg *sync.WaitGroup, foundLinks *FoundMagnetLinks, mutex *sync.Mutex) {
			defer wg.Done()
			magnetHeader := make(map[string]string)
			magnetHeader["Cookie"] = "nome=" + match[4] + "; ref=" + match[3] + "; idcriptografada=" + match[2]
			magnetHeader["Referer"] = "https://www.adssuper.com"

			magnetHtml := WebRequest{magnetHeader, nil}.Get("https://www.mastercuriosidadesbr.com/por-que-os-ossos-estalam/")
			magnetRegex := regexp.MustCompile(`(?m)href="(magnet:.+?)"`)
			magnetLink := magnetRegex.FindStringSubmatch(magnetHtml)[1]

			magnetRegex = regexp.MustCompile(`magnet:.*?dn=(.+?)(;|&amp|&)`)

			magnetName := magnetRegex.FindStringSubmatch(magnetLink)
			var magnetTitle string
			if magnetName != nil {
				magnetTitle, _ = url.QueryUnescape(magnetName[1])
				magnetTitle = strings.Trim(magnetTitle, " ")
			} else {
				magnetTitle = match[1]
			}

			mutex.Lock()
			(*foundLinks) = append((*foundLinks), DownloadLink{magnetTitle, magnetLink})
			mutex.Unlock()
		}(&wg, &foundLinks, mutex)
	}

	wg.Wait()

	return foundLinks
}
