package moviescrawler

import (
	"net/url"
	"regexp"
	"strconv"
	ht "golang.org/x/net/html"
)

type MegaTorrents struct{}

func (_ MegaTorrents) SearchMovie(Term string, Page int) FoundMovies {
	defer recoveryPanic()
	var headers map[string]string
	headers = make(map[string]string)

	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"
	request := WebRequest{headers, nil}

	html := request.Get("http://www.megatorrentshd.org/page/" + strconv.Itoa(Page) + "/?s=" + url.QueryEscape(Term))

	regex := regexp.MustCompile(`<div class="peli">\n<img src="(.+?)" alt="(.+?)" \/>\n<a href="(.+?)">`)
	matches := regex.FindAllStringSubmatch(string(html), -1)
	var found FoundMovies
	for _, match := range matches {
		found = append(found, Movie{ht.UnescapeString(match[2]), match[1], match[3]})
	}

	return found
}

func (_ MegaTorrents) GetDownloadLinks(Link string) FoundMagnetLinks {
	defer recoveryPanic()
	var headers map[string]string
	headers = make(map[string]string)

	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"
	request := WebRequest{headers, nil}

	html := request.Get(Link)

	regex := regexp.MustCompile(`(?ms)<.*?style=".*?color: #fff;.*?">(.*?)<.*?"(magnet:.+?)"`)
	matches := regex.FindAllStringSubmatch(string(html), -1)

	var found FoundMagnetLinks

	for _, match := range matches {
		found = append(found, DownloadLink{ht.UnescapeString(match[1]), match[2]})
	}

	return found
}

func (_ MegaTorrents) GetNumberOfPages(Term string) int {
	defer recoveryPanic()
	var headers map[string]string
	headers = make(map[string]string)

	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36 OPR/68.0.3618.63"
	request := WebRequest{headers, nil}

	html := request.Get("http://www.megatorrentshd.org/?s=" + url.QueryEscape(Term))
	r := regexp.MustCompile(`(?m)<a class="last" href=".*?([0-9]+).*?">`)
	result := r.FindStringSubmatch(html)
	if result != nil {
		i, _ := strconv.Atoi(result[1])
		return i
	}
	return 1
}
