package moviescrawler

import (
	"net/url"
	"regexp"
	"strconv"
	ht "golang.org/x/net/html"
	
)

type BludTV struct{}

func (_ BludTV) SearchMovie(Term string, Page int) FoundMovies {
	defer recoveryPanic()
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"

	html := WebRequest{headers, nil}.Get("https://www.bludv.tv/page/" + strconv.Itoa(Page) + "/?s=" + url.QueryEscape(Term))
	regex := regexp.MustCompile(`(?s)<div class=post><div class=title> <a href=(.+?) title=".*?">(.+?)<\/a>.*?<img src=(.*?) width=`)
	matches := regex.FindAllStringSubmatch(html, -1)

	var movies FoundMovies
	for _, match := range matches {
		movies = append(movies, Movie{Title: ht.UnescapeString(match[2]), Cover: match[3], Link: match[1]})
	}

	return movies
}

func (_ BludTV) GetNumberOfPages(Term string) int {
	defer recoveryPanic()
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"

	html := WebRequest{headers, nil}.Get("https://www.bludv.tv/?s=" + url.QueryEscape(Term))
	regex := regexp.MustCompile(`(?s)<span class=pages>Página \d* de (\d+)<\/span>`)
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

func (_ BludTV) GetDownloadLinks(Link string) FoundMagnetLinks {
	defer recoveryPanic()
	headers := make(map[string]string)

	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"
	request := WebRequest{headers, nil}

	html := request.Get(Link)

	regex := regexp.MustCompile(`(?s)<span style=".*?font-family: 'arial' , 'helvetica' , sans-serif; font-size: large;">(?:<.*?>)*(.*?)(?:<\/b>|<br>|<\/span>).*?(magnet:.+?)"`)
	matches := regex.FindAllStringSubmatch(string(html), -1)

	var found FoundMagnetLinks

	for _, match := range matches {
		found = append(found, DownloadLink{ht.UnescapeString(match[1]), match[2]})
	}

	return found
}
