package moviescrawler

import (
	"testing"
	ht "golang.org/x/net/html"
	"fmt"
)

func TestComandoTorrentsSearch(t *testing.T) {
	var expectedMovies = FoundMovies{
		Movie{Title: "John Wick 3 – Parabellum Torrent (2019) Dual Áudio / Dublado BluRay 720p | 1080p | 2160p 4K – Download", Cover: "https://www.comandotorrents.org/wp-content/uploads/2019/05/John.Wick_.3.Parabellum.poster.torrent.2019-min.png", Link: "https://comandotorrents.org/john-wick-3-parabellum-torrent-2019-dublado-hd-720p-download/"},
		Movie{Title: "John Wick – Um Novo Dia para Matar (2017) Dublado e Dual Áudio WEB-DL 720p | 1080p – Torrent Download", Cover: "https://www.comandotorrents.org/wp-content/uploads/2017/05/John.Wick_UmNovoDiaParaMatar_Legendado.jpg", Link: "https://comandotorrents.org/john-wick-um-novo-dia-para-matar-2017-dublado-e-dual-audio-web-dl-720p-1080p-torrent-download/"},
		Movie{Title: "John Wick – Um Novo Dia para Matar (2017) Legendado WEB-DL 720p | 1080p – Torrent Download", Cover: "https://www.comandotorrents.org/wp-content/uploads/2017/05/John.Wick_UmNovoDiaParaMatar_Legendado.jpg", Link: "https://comandotorrents.org/john-wick-um-novo-dia-para-matar-2017-legendado-web-dl-720p-1080p-torrent-download/"},
		Movie{Title: "John Wick – Um Novo Dia para Matar (2017) Legendado – Torrent Download", Cover: "https://www.comandotorrents.org/wp-content/uploads/2017/03/John-Wick-Um-Novo-Dia-para-Matar-250x365.jpg", Link: "https://comandotorrents.org/john-wick-um-novo-dia-para-matar-2017-legendado-torrent-download/"},
		Movie{Title: "Nada Santo Torrent (2019) Dual Áudio 5.1 / Dublado WEB-DL 720p | 1080p – Download", Cover: "https://www.comandotorrents.org/wp-content/uploads/2019/04/nada.santo_.poster.torrent.2019-min.png", Link: "https://comandotorrents.org/nada-santo-torrent-2019-dual-audio-5-1-dublado/"},
		Movie{Title: "Wayne 1ª Temporada Completa Torrent (2019) Legendado WEB-DL 720p – Download", Cover: "https://www.comandotorrents.org/wp-content/uploads/2019/01/wayne.s01.poster.torrent.2019.png", Link: "https://comandotorrents.org/wayne-1a-temporada-completa-torrent-2019-dual-audio-dublado-web-dl-720p-download/"},
		Movie{Title: "De Volta ao Jogo BluRay 1080p Dual Áudio Torrent (2015)", Cover: "https://www.comandotorrents.org/wp-content/uploads/2015/02/De-Volta-ao-Jogo-e1424898128363.jpg", Link: "https://comandotorrents.org/de-volta-ao-jogo-bluray-1080p-dual-audio-torrent-2015/"},
		Movie{Title: "De Volta ao Jogo (2014) WEB-DL 720p – 1080p Legendado Download Torrent", Cover: "https://www.comandotorrents.org/wp-content/uploads/2015/01/De-2BVolta-2Bao-2BJogo-2B-2014-2BWEB-DL-2B720p-2B-2B1080p-2BLegendado-2BDownload-2BTorrent.jpg", Link: "https://comandotorrents.org/de-volta-ao-jogo-2014-web-dl-720p-1080p-legendado-download-torrent/"},
	}
	foundMovies := ComandoTorrents{}.SearchMovie("john wick", 1)
	if len(expectedMovies) != len(foundMovies) {
		t.Errorf("The number of found movies is not equal to the expected\nExpected: %d\nFound: %d\n", len(expectedMovies), len(foundMovies))
	}

	for _, foundMovie := range foundMovies {
		flag := false
		var tempExpct Movie
		for _, expectedMovie := range expectedMovies {
			if foundMovie.Title == ht.UnescapeString(expectedMovie.Title) && foundMovie.Cover == expectedMovie.Cover && foundMovie.Link == expectedMovie.Link {
				flag = true
			}
			tempExpct = expectedMovie
		}
		if !flag {
			if foundMovie.Title != ht.UnescapeString(tempExpct.Title) {
				t.Errorf("The title don't match\nFound: %s\nExpected: %s\n", foundMovie.Title, tempExpct.Title)
			} else if foundMovie.Cover != tempExpct.Cover {
				t.Errorf("The cover don't match\nFound: %s\nExpected: %s\n", foundMovie.Cover, tempExpct.Cover)
			} else if foundMovie.Link != tempExpct.Link {
				t.Errorf("The link don't match\nFound: %s\nExpected: %s\n", foundMovie.Link, tempExpct.Link)
			}
			t.Errorf("Found Movies is not equal to the expected\n")
		}
	}
}

func TestComandoTorrentsGetDownloadLinksActual(t *testing.T) {
	var expectedLinks = FoundMagnetLinks{
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:4bdf3018ea471e15d6941966a09da0c880b73d3a&dn=%5bACESSE%20comandotorrents.org%5d%20John%20Wick%203%20-%20Parabellum%202019%20REPACK%20%5b720p%5d%20%5bBluRay%5d%20%5bDUAL%5d&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2fglotorrents.pw%3a6969%2fannounce&tr=udp%3a%2f%2ftracker4.piratux.com%3a6969%2fannounce&tr=udp%3a%2f%2fcoppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2ft2.pow7.com%2fannounce&tr=udp%3a%2f%2ftracker.yify-torrents.com%3a80%2fannounce&tr=http%3a%2f%2fwww.h33t.com%3a3310%2fannounce&tr=http%3a%2f%2fexodus.desync.com%2fannounce&tr=http%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2fbt.careland.com.cn%3a6969%2fannounce&tr=http%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.istole.it%3a80%2fannounce&tr=http%3a%2f%2ftracker.blazing.de%2fannounce&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=http%3a%2f%2ftracker.yify-torrents.com%2fannounce&tr=udp%3a%2f%2ftracker.prq.to%2fannounce&tr=udp%3a%2f%2ftracker.1337x.org%3a80%2fannounce&tr=udp%3a%2f%2f9.rarbg.com%3a2740%2fannounce&tr=udp%3a%2f%2ftracker.ex.ua%3a80%2fannounce&tr=udp%3a%2f%2fipv4.tracker.harry.lu%3a80%2fannounce&tr=udp%3a%2f%2f12.rarbg.me%3a80%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f11.rarbg.com%2fannounce&tr=udp%3a%2f%2ftracker.ccc.de%3a80%2fannounce&tr=udp%3a%2f%2ffr33dom.h33t.com%3a3310%2fannounce&tr=udp%3a%2f%2fpublic.popcorn-tracker.org%3a6969%2fannounce", Title: "BluRay 720p REPACK Dual Áudio 5.1 (MKV)"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:9d206173973452ec80de945c61a87d7430d44fa7&dn=%5bACESSE%20comandotorrents.org%5d%20John%20Wick%203%20-%20Parabellum%202019%20%5b1080p%5d%20%5bBluRay%5d%20%5bDUAL%5d&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2fglotorrents.pw%3a6969%2fannounce&tr=udp%3a%2f%2ftracker4.piratux.com%3a6969%2fannounce&tr=udp%3a%2f%2fcoppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2ft2.pow7.com%2fannounce&tr=udp%3a%2f%2ftracker.yify-torrents.com%3a80%2fannounce&tr=http%3a%2f%2fwww.h33t.com%3a3310%2fannounce&tr=http%3a%2f%2fexodus.desync.com%2fannounce&tr=http%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2fbt.careland.com.cn%3a6969%2fannounce&tr=http%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.istole.it%3a80%2fannounce&tr=http%3a%2f%2ftracker.blazing.de%2fannounce&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=http%3a%2f%2ftracker.yify-torrents.com%2fannounce&tr=udp%3a%2f%2ftracker.prq.to%2fannounce&tr=udp%3a%2f%2ftracker.1337x.org%3a80%2fannounce&tr=udp%3a%2f%2f9.rarbg.com%3a2740%2fannounce&tr=udp%3a%2f%2ftracker.ex.ua%3a80%2fannounce&tr=udp%3a%2f%2fipv4.tracker.harry.lu%3a80%2fannounce&tr=udp%3a%2f%2f12.rarbg.me%3a80%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f11.rarbg.com%2fannounce&tr=udp%3a%2f%2ftracker.ccc.de%3a80%2fannounce&tr=udp%3a%2f%2ffr33dom.h33t.com%3a3310%2fannounce&tr=udp%3a%2f%2fpublic.popcorn-tracker.org%3a6969%2fannounce", Title: "BluRay 1080p Dual Áudio 5.1 (MKV)"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:7a0b384ab8e372a4f7437b223b75ef231258fcc4&dn=%5bACESSE%20comandotorrents.org%5d%20John%20Wick%203%20-%20Parabellum%202019%20%5bx265%5d%20%5b1080p%5d%20%5bBluRay%5d%20%5bDUAL%5d&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2fglotorrents.pw%3a6969%2fannounce&tr=udp%3a%2f%2ftracker4.piratux.com%3a6969%2fannounce&tr=udp%3a%2f%2fcoppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2ft2.pow7.com%2fannounce&tr=udp%3a%2f%2ftracker.yify-torrents.com%3a80%2fannounce&tr=http%3a%2f%2fwww.h33t.com%3a3310%2fannounce&tr=http%3a%2f%2fexodus.desync.com%2fannounce&tr=http%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2fbt.careland.com.cn%3a6969%2fannounce&tr=http%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.istole.it%3a80%2fannounce&tr=http%3a%2f%2ftracker.blazing.de%2fannounce&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=http%3a%2f%2ftracker.yify-torrents.com%2fannounce&tr=udp%3a%2f%2ftracker.prq.to%2fannounce&tr=udp%3a%2f%2ftracker.1337x.org%3a80%2fannounce&tr=udp%3a%2f%2f9.rarbg.com%3a2740%2fannounce&tr=udp%3a%2f%2ftracker.ex.ua%3a80%2fannounce&tr=udp%3a%2f%2fipv4.tracker.harry.lu%3a80%2fannounce&tr=udp%3a%2f%2f12.rarbg.me%3a80%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f11.rarbg.com%2fannounce&tr=udp%3a%2f%2ftracker.ccc.de%3a80%2fannounce&tr=udp%3a%2f%2ffr33dom.h33t.com%3a3310%2fannounce&tr=udp%3a%2f%2fpublic.popcorn-tracker.org%3a6969%2fannounce", Title: "BluRay 1080p Dual Áudio 5.1 x265 (MKV)"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:8dd6a8687c0aa3c860db675cbf9b672c0742c1cf&dn=%5bACESSE%20comandotorrents.org%5d%20John%20Wick%203%20-%20Parabellum%202019%20%5b1080p%5d%20%5bBluRay%5d%20%5bFULL%5d%20%5bDUAL%5d&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2fglotorrents.pw%3a6969%2fannounce&tr=udp%3a%2f%2ftracker4.piratux.com%3a6969%2fannounce&tr=udp%3a%2f%2fcoppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2ft2.pow7.com%2fannounce&tr=udp%3a%2f%2ftracker.yify-torrents.com%3a80%2fannounce&tr=http%3a%2f%2fwww.h33t.com%3a3310%2fannounce&tr=http%3a%2f%2fexodus.desync.com%2fannounce&tr=http%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2fbt.careland.com.cn%3a6969%2fannounce&tr=http%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.istole.it%3a80%2fannounce&tr=http%3a%2f%2ftracker.blazing.de%2fannounce&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=http%3a%2f%2ftracker.yify-torrents.com%2fannounce&tr=udp%3a%2f%2ftracker.prq.to%2fannounce&tr=udp%3a%2f%2ftracker.1337x.org%3a80%2fannounce&tr=udp%3a%2f%2f9.rarbg.com%3a2740%2fannounce&tr=udp%3a%2f%2ftracker.ex.ua%3a80%2fannounce&tr=udp%3a%2f%2fipv4.tracker.harry.lu%3a80%2fannounce&tr=udp%3a%2f%2f12.rarbg.me%3a80%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f11.rarbg.com%2fannounce&tr=udp%3a%2f%2ftracker.ccc.de%3a80%2fannounce&tr=udp%3a%2f%2ffr33dom.h33t.com%3a3310%2fannounce&tr=udp%3a%2f%2fpublic.popcorn-tracker.org%3a6969%2fannounce", Title: "BluRay 1080p Dual Áudio 5.1 [ULTRA FULL HD] (MKV)"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:28B66D9A55FB0084EA48716BE129637E0A300E34&dn=%5bACESSE%20comandotorrents.org%5d%20John%20Wick%203%20-%20Parabellum%202019%20%5b2160p%5d%20%5b4K%5d%20%5bWEB-DL%5d%20%5bDUAL%5d&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&tr=udp%3a%2f%2ftracker4.piratux.com%3a6969%2fannounce&tr=udp%3a%2f%2fcoppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2ft2.pow7.com%2fannounce&tr=http%3a%2f%2ftracker.yify-torrents.com%2fannounce&tr=http%3a%2f%2fwww.h33t.com%3a3310%2fannounce&tr=udp%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=http%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2fbt.careland.com.cn%3a6969%2fannounce&tr=udp%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.istole.it%3a80%2fannounce&tr=http%3a%2f%2ftracker.blazing.de%2fannounce&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.yify-torrents.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.prq.to%2fannounce&tr=udp%3a%2f%2ftracker.1337x.org%3a80%2fannounce&tr=udp%3a%2f%2f9.rarbg.com%3a2780%2fannounce&tr=udp%3a%2f%2ftracker.ex.ua%3a80%2fannounce&tr=udp%3a%2f%2fipv4.tracker.harry.lu%3a80%2fannounce&tr=udp%3a%2f%2f12.rarbg.me%3a80%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f11.rarbg.com%2fannounce&tr=udp%3a%2f%2ftracker.ccc.de%3a80%2fannounce&tr=udp%3a%2f%2ffr33dom.h33t.com%3a3310%2fannounce&tr=udp%3a%2f%2fpublic.popcorn-tracker.org%3a6969%2fannounce", Title: "WEB-DL 2160p Dual Áudio 5.1 [4K] (MKV)"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:6f266a0f1fa67ba090151c111768c7c594816c77&dn=%5bACESSE%20comandotorrents.org%5d%20John%20Wick%203%20-%20Parabellum%202019%20REPACK%20%5b720p%5d%20%5bBluRay%5d%20%5bDUBLADO%5d&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2fglotorrents.pw%3a6969%2fannounce&tr=udp%3a%2f%2ftracker4.piratux.com%3a6969%2fannounce&tr=udp%3a%2f%2fcoppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2ft2.pow7.com%2fannounce&tr=udp%3a%2f%2ftracker.yify-torrents.com%3a80%2fannounce&tr=http%3a%2f%2fwww.h33t.com%3a3310%2fannounce&tr=http%3a%2f%2fexodus.desync.com%2fannounce&tr=http%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2fbt.careland.com.cn%3a6969%2fannounce&tr=http%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.istole.it%3a80%2fannounce&tr=http%3a%2f%2ftracker.blazing.de%2fannounce&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=http%3a%2f%2ftracker.yify-torrents.com%2fannounce&tr=udp%3a%2f%2ftracker.prq.to%2fannounce&tr=udp%3a%2f%2ftracker.1337x.org%3a80%2fannounce&tr=udp%3a%2f%2f9.rarbg.com%3a2740%2fannounce&tr=udp%3a%2f%2ftracker.ex.ua%3a80%2fannounce&tr=udp%3a%2f%2fipv4.tracker.harry.lu%3a80%2fannounce&tr=udp%3a%2f%2f12.rarbg.me%3a80%2fannounce&tr=udp%3a%2f%2ftracker.publicbt.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f11.rarbg.com%2fannounce&tr=udp%3a%2f%2ftracker.ccc.de%3a80%2fannounce&tr=udp%3a%2f%2ffr33dom.h33t.com%3a3310%2fannounce&tr=udp%3a%2f%2fpublic.popcorn-tracker.org%3a6969%2fannounce", Title: string([]byte{66,108,117,82,97,121,32,55,50,48,112,194,160,82,69,80,65,67,75,32,68,117,98,108,97,100,111,32,40,77,80,52,41})},
	}

	foundLinks := ComandoTorrents{}.GetDownloadLinks("https://comandotorrents.org/john-wick-3-parabellum-torrent-2019-dublado-hd-720p-download/")
	if len(expectedLinks) != len(foundLinks) {
		t.Errorf("The number of found links is not equal to the expected\nExpected: %d\nFound: %d\n", len(expectedLinks), len(foundLinks))
	}

	for _, foundLink := range foundLinks {
		flag := false
		var expctLinks DownloadLink
		for _, expectedLink := range expectedLinks {
			if foundLink.Title == expectedLink.Title && foundLink.MagnetLink == expectedLink.MagnetLink {
				flag = true
			}
			expctLinks = expectedLink
		}
		if !flag {
			if 	foundLink.Title != expctLinks.Title {
				t.Errorf("The title don't match\nFound: %s\nExpected: %s\n", foundLink.Title, expctLinks.Title)
				
			} 
			
			if foundLink.MagnetLink != expctLinks.MagnetLink {
				t.Errorf("The link don't match\nFound: %s\nExpected: %s\n", foundLink.MagnetLink, expctLinks.MagnetLink)
			}
			t.Errorf("Found Links is not equal to the expected\n")
			return
		}
	}
}

func TestComandoTorrentsGetDownloadLinksElder(t *testing.T) {
	var expectedLinks = FoundMagnetLinks{
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:2F70ABDD53A7587CEBEAB6FF0A8A1840EB392599&dn=Os%20Embalos%20de%20S%c3%a1bado%20%c3%a0%20Noite%20720p%20%281977%29%20Dual%20%c3%81udio%20BluRay%205.1%20--%20By%20-%20Lucas%20Firmo&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.ccc.de%3a80%2fannounce&tr=http%3a%2f%2ftracker1.wasabii.com.tw%3a6969%2fannounce&tr=http%3a%2f%2ftracker.tfile.me%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2fmgtracker.org%3a2710%2fannounce&tr=udp%3a%2f%2f9.rarbg.com%3a2710%2fannounce&tr=udp%3a%2f%2f9.rarbg.me%3a2800%2fannounce&tr=http%3a%2f%2ft2.pow7.com%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce", Title: "720p – AC3 5.1 Surround"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:76ADBA3A9283B02400A8DC89B8D683B3860A4247&dn=Os%20Embalos%20de%20S%c3%a1bado%20%c3%a0%20Noite%201080p%20%281977%29%20Dual%20%c3%81udio%20BluRay%205.1%20--%20By%20-%20Lucas%20Firmo&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.ccc.de%3a80%2fannounce&tr=http%3a%2f%2ftracker1.wasabii.com.tw%3a6969%2fannounce&tr=http%3a%2f%2ftracker.tfile.me%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=http%3a%2f%2fmgtracker.org%3a2710%2fannounce&tr=udp%3a%2f%2f9.rarbg.com%3a2710%2fannounce&tr=udp%3a%2f%2f9.rarbg.me%3a2800%2fannounce&tr=http%3a%2f%2ft2.pow7.com%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce", Title: string([]byte{49,48,56,48,112,32,226,128,147,32,65,67,51,32,53,46,49,194,160,83,117,114,114,111,117,110,100})},
	}

	foundLinks := ComandoTorrents{}.GetDownloadLinks("https://comandotorrents.org/os-embalos-de-sabado-a-noite-torrent-bluray-720p-e-1080p-dual-audio-5-1-download-1977/")
	if len(expectedLinks) != len(foundLinks) {
		t.Errorf("The number of found links is not equal to the expected\nExpected: %d\nFound: %d\n", len(expectedLinks), len(foundLinks))
	}

	for _, foundLink := range foundLinks {
		flag := false
		var expctLinks DownloadLink
		for _, expectedLink := range expectedLinks {
			if foundLink.Title == expectedLink.Title && foundLink.MagnetLink == expectedLink.MagnetLink {
				flag = true
			}
			expctLinks = expectedLink
		}
		if !flag {
			if 	foundLink.Title != expctLinks.Title {
				fmt.Println([]byte(foundLink.Title))
				fmt.Println([]byte(expctLinks.Title))
				t.Errorf("The title don't match\nFound: %s\nExpected: %s\n", foundLink.Title, expctLinks.Title)
				
			} 
			
			if foundLink.MagnetLink != expctLinks.MagnetLink {
				t.Errorf("The link don't match\nFound: %s\nExpected: %s\n", foundLink.MagnetLink, expctLinks.MagnetLink)
			}
			t.Errorf("Found Links is not equal to the expected\n")
			return
		}
	}
}