package moviescrawler

import (
	"net/url"
	"regexp"
	"strconv"
)

type FilmesViaTorrent struct{}

func (_ FilmesViaTorrent) GetNumberOfPages(Term string) int {
	defer recoveryPanic()
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"

	html := WebRequest{headers, nil}.Get("https://filmesviatorrents.biz/?s=" + url.QueryEscape(Term))
	regex := regexp.MustCompile(`(?ms)>(\d+)<\/a>\n<a class="next page-numbers"`)
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

func (_ FilmesViaTorrent) SearchMovie(Term string, Page int) FoundMovies {
	defer recoveryPanic()
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"

	html := WebRequest{headers, nil}.Get("https://filmesviatorrents.biz/page/" + strconv.Itoa(Page) + "/?s=" + url.QueryEscape(Term))
	regex := regexp.MustCompile(`(?s)<div class="TPost B">.*?<a href="(.+?)">.*?<div class="Image">.*?<img src="(.*?)".*?<div class="Title">(.+?)<\/div>`)
	matches := regex.FindAllStringSubmatch(html, -1)

	var movies FoundMovies
	for _, match := range matches {
		movies = append(movies, Movie{Title: match[3], Cover: "http:" + match[2], Link: match[1]})
	}

	return movies
}
