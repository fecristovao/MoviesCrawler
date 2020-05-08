package moviescrawler

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type MeusFilmesTorrent struct{}

func (_ MeusFilmesTorrent) GetDownloadLinks(Link string) FoundMagnetLinks {
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"
	html := WebRequest{headers, nil}.Get(Link)

	regex := regexp.MustCompile(`<a rel="nofollow" target="_blank" href=".*id=(.+?)&ref=(.+?)&titulo=(.+?)"`)
	matches := regex.FindAllStringSubmatch(string(html), -1)

	var foundLinks FoundMagnetLinks
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	for _, match := range matches {
		wg.Add(1)
		go func(wg *sync.WaitGroup, foundLinks *FoundMagnetLinks, mutex *sync.Mutex) {
			defer wg.Done()
			magnetHeader := make(map[string]string)
			magnetHeader["Cookie"] = "nome=" + match[3] + "; ref=" + match[2] + "; idcriptografada=" + match[1]
			magnetHeader["Referer"] = "https://www.adssuper.com"

			magnetHtml := WebRequest{magnetHeader, nil}.Get("https://www.mastercuriosidadesbr.com/resenha-sobre-o-filme-malevola-2-dona-do-mal-parte-final/")
			magnetRegex := regexp.MustCompile(`(?m)href="(magnet:.+?)"`)
			magnetLink := magnetRegex.FindStringSubmatch(magnetHtml)[1]

			magnetRegex = regexp.MustCompile(`magnet:.*?dn=(.+?)(;|&amp|&)`)
			magnetName := magnetRegex.FindStringSubmatch(magnetLink)

			var magnetTitle string
			if magnetName != nil {
				magnetTitle, _ = url.QueryUnescape(magnetName[1])
				magnetTitle = strings.Trim(magnetTitle, " ")
			} else {
				magnetTitle = ""
			}
			mutex.Lock()
			(*foundLinks) = append((*foundLinks), DownloadLink{magnetTitle, magnetLink})
			mutex.Unlock()
		}(&wg, &foundLinks, mutex)
	}

	wg.Wait()

	return foundLinks
}

func (_ MeusFilmesTorrent) SearchMovie(Term string, Page int) FoundMovies {
	defer recoveryPanic()
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"

	html := WebRequest{headers, nil}.Get("https://meusfilmestorrents.com/page/" + strconv.Itoa(Page) + "/?s=" + url.QueryEscape(Term))
	regex := regexp.MustCompile(`(?sm)itemprop="headline"><a href="(.+?)">(.*?)<\/a><\/h2>.*?src="(.+?)"`)
	matches := regex.FindAllStringSubmatch(html, -1)

	var movies FoundMovies
	for _, match := range matches {
		movies = append(movies, Movie{Title: match[2], Cover: match[3], Link: match[1]})
	}

	return movies
}

func (_ MeusFilmesTorrent) GetNumberOfPages(Term string) int {
	defer recoveryPanic()
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"

	html := WebRequest{headers, nil}.Get("https://meusfilmestorrents.com/?s=" + url.QueryEscape(Term))
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
