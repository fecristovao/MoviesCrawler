package moviescrawler

import (
	"sync"
)

type Web interface {
	Post(Uri string) string
	Get(Uri string) string
}

type Printer interface {
	PrintFoundMovies()
	PrintFoundMagnetLinks()
}

type Crawler interface {
	SearchMovie(Term string, Page int) FoundMovies
	GetDownloadLinks(Link string) FoundMagnetLinks
	GetNumberOfPages(Term string) int
}

func SearchAll(servico Crawler, Term string) FoundMovies {
	defer recoveryPanic()
	var wg sync.WaitGroup
	var listMovies FoundMovies

	mutex := &sync.Mutex{}
	maxPage := servico.GetNumberOfPages(Term)
	for i := 1; i <= maxPage; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, listMovies *FoundMovies, mutex *sync.Mutex, term string, page int) {
			defer wg.Done()
			service := servico.SearchMovie(term, page)
			if service != nil {
				mutex.Lock()
				(*listMovies) = append((*listMovies), service...)
				mutex.Unlock()
			}
		}(&wg, &listMovies, mutex, Term, i)
	}
	wg.Wait()
	return listMovies
}
