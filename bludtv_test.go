package moviescrawler

import (
	"testing"
	"strings"
	"fmt"
	ht "golang.org/x/net/html"
)

func strcmp(s1, s2 string) int {
    lens := len(s1)
    if lens > len(s2) {
        lens = len(s2)
    }
    for i := 0; i < lens; i++ {
        if s1[i] != s2[i] {
            return int(s1[i]) - int(s2[i])
        }
    }
    return len(s1) - len(s2)
}

func TestBludTVSearch(t *testing.T) {
	var expectedMovies = FoundMovies{
		Movie{Title: "Divino Amor – 2019 Nacional (BluRay) 720p e 1080p – Download", Cover: "https://1.bp.blogspot.com/-pI573KBTyuU/XWQL6DaRxVI/AAAAAAAAOug/C7KyYPxA5sgFFA8iLuAtakF0_fp7LvzVgCLcBGAs/s1600/cover.jpg", Link: "https://www.bludv.tv/divino-amor-2019-nacional-bluray-720p-e-1080p-download/"},
		Movie{Title: "Primeira Vez Amor – 1ª Temporada Completa Torrent – 2019 Dual Áudio (WEB-DL) 720p – Download", Cover: "https://4.bp.blogspot.com/-eElXj3eQ6Jc/XLh5_qEDfvI/AAAAAAAANvs/6uz_oWH3JBYBUMLQzVshqgom28v960b3ACLcBGAs/s1600/cover.jpg", Link: "https://www.bludv.tv/primeira-vez-amor-1a-temporada-completa-torrent-2019-dublado-dual-audio-legendado-web-dl-720p-e-1080p-download/"},
		Movie{Title: "A Quimica do Amor Torrent – 2019 Dublado / Dual Áudio (BluRay) 720p e 1080p – Download", Cover: "https://2.bp.blogspot.com/-B-Peva34UUA/Wz0rtQL4boI/AAAAAAAAIhY/FRKnzToJndg-GVwRq3CvOckF1MMRJkOVwCLcBGAs/s1600/The-Female-Brain.jpg", Link: "https://www.bludv.tv/a-quimica-do-amor-torrent-2019-dublado-dual-audio-legendado-bluray-720p-e-1080p-download-2/"},
		Movie{Title: "Amor de Irmão [Brotherly Love] – 2015 Dublado / Dual Áudio / Legendado (WEB-DL) 720p e 1080p – Download", Cover: "https://4.bp.blogspot.com/-C5n_eY3zBl0/XIKEmHn3RuI/AAAAAAAADvE/-pInXSj22nklSYWO7o5PIucMcBsm6f_EgCLcBGAs/s1600/cover.jpg", Link: "https://www.bludv.tv/amor-de-irmao-brotherly-love-2015-dublado-dual-audio-legendado-web-dl-720p-e-1080p-download/"},
		Movie{Title: "Amor em Little Italy Torrent – 2018 Dublado / Dual Áudio (BluRay) 720p e 1080p – Download", Cover: "https://i.imgur.com/i7sdN5V.jpg", Link: "https://www.bludv.tv/amor-em-little-italy-torrent-2018-dublado-dual-audio-bluray-download/"},
		Movie{Title: "Amor Sonia Torrent – 2019 – Legendado (WEB-DL) 1080p – Download", Cover: "https://3.bp.blogspot.com/-N2asMwEuKRk/XEErJMkMsJI/AAAAAAAAJX0/4bR4nih2Fj8P7GUtXzghqFNWGd-O3pHvQCLcBGAs/s1600/Amor-Sonia.png", Link: "https://www.bludv.tv/amor-sonia-torrent-2019-dual-audio-dublado-legendado-bluray-720p-e-1080p-download/"},
		Movie{Title: "Todo Clichê do Amor Torrent – 2018 Nacional (WEB-DL) 720p e 1080p – Download", Cover: "https://4.bp.blogspot.com/-ZTDJw35RHjs/XDTTz2tphyI/AAAAAAAADgA/nTQZnHTUjG03ckHCO79Ejy9w6446kIznQCLcBGAs/s1600/cover.jpg", Link: "https://www.bludv.tv/todo-cliche-do-amor-torrent-2018-nacional-web-dl-720p-e-1080p-download/"},
		Movie{Title: "De Carona para o Amor Torrent – 2018 Dublado / Dual Áudio (BluRay) 720p e 1080p – Download", Cover: "https://2.bp.blogspot.com/-34tU7sphg8c/W-53Ik82ncI/AAAAAAAAM94/sGzavbsvZ18eEh3FuEy7utDSeGCf5WT5QCLcBGAs/s1600/cover.jpg", Link: "https://www.bludv.tv/de-carona-para-o-amor-torrent-2018-dublado-dual-audio-legendado-bluray-720p-e-1080p-download/"},
		Movie{Title: "O Melhor Presente é o Amor Torrent – 2018 (BluRay) 720p e 1080p Dublado / Dual Áudio", Cover: "https://2.bp.blogspot.com/-5kdkyWLFZXA/W4c2q2Ct6fI/AAAAAAAAI4k/O5ltBfXIz8wBqSLjs__sxW743vFRf-VeACLcBGAs/s1600/O-Melhor-Presente-%25C3%25A9-o-Amor.jpg", Link: "https://www.bludv.tv/o-melhor-presente-e-o-amor-torrent-2018-bluray-720p-e-1080p-dublado-dual-audio/"},
		Movie{Title: "O Diabo e o Padre Amorth Torrent – 2018 (WEB-DL) 1080p Legendado", Cover: "https://2.bp.blogspot.com/-6AUJ6-coKLw/W1m1XWq802I/AAAAAAAAMJM/xbrxh-RjLhwvK3afonkr54U040zQbT-EQCLcBGAs/s1600/cover.jpg", Link: "https://www.bludv.tv/o-diabo-e-o-padre-amorth-torrent-2018-web-dl-1080p-legendado/"},
	}
	foundMovies := BludTV{}.SearchMovie("amor", 1)
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

func TestBludTVGetDownloadLinksActual(t *testing.T) {
	var expectedLinks = FoundMagnetLinks{
		DownloadLink{Title: string([]byte{86,101,114,115,195,163,111,32,68,117,97,108,32,195,129,117,100,105,111,194,160,88,50,54,53,32,40,77,75,86,41}), MagnetLink: "magnet:?xt=urn:btih:CE29F2873DD258067D4BE35B8E05C691B690936C&amp;dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%20%202019%20%281080p%20-%20BluRay%20-%20x265%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&amp;tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&amp;tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&amp;tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&amp;tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&amp;tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&amp;tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&amp;tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&amp;tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce"},
		DownloadLink{Title: string([]byte{86,101,114,115,195,163,111,32,68,117,97,108,32,195,129,117,100,105,111,194,160,40,77,75,86,41}), MagnetLink: "magnet:?xt=urn:btih:7296D8B865D711B3E00FC2CDBE97F54890875F14&amp;dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%202019%20%28720p%20-%20BluRay%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&amp;tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&amp;tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&amp;tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&amp;tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&amp;tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&amp;tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&amp;tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&amp;tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce"},
		DownloadLink{Title: "Versão Dublado (MP4)", MagnetLink: "magnet:?xt=urn:btih:8826B08649746CE6C7554F985F081A6F9F7E8A97&amp;dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%202019%20%28720p%20-%20BluRay%29%20%5bDUBLADO%5d%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&amp;tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&amp;tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&amp;tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&amp;tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&amp;tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&amp;tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&amp;tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&amp;tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce"},
		DownloadLink{Title: "Versão Dual Áudio | Sem Compactação | 10 GB (MKV)", MagnetLink: "magnet:?xt=urn:btih:C67627D20ED5BA7D7EC002A3610C33BAE83298F4&amp;dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%20%202019%20%281080p%20-%20BluRay%20-%20FULL%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&amp;tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&amp;tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce"},
		DownloadLink{Title: "Versão Legenda FIXA (MKV)", MagnetLink: "magnet:?xt=urn:btih:E1658C93B1660E35DCFFA122DC654A41337385B4&amp;dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%202019%20%281080p%20-%20BluRay%20-%20Legendado%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&amp;tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&amp;tr=udp%3a%2f%2f9.rarbg.to%3a2740%2fannounce&amp;tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&amp;tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&amp;tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&amp;tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&amp;tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&amp;tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce"},
		DownloadLink{Title: "Versão Legendado (MKV)", MagnetLink: "magnet:?xt=urn:btih:02bd2fbb03b5911f3d4936dc44bc9bcdff8f47bb&amp;dn=John.Wick.3.2019.720p.Bluray.X264-EVO%5BTGx%5D&amp;xl=6151005201&amp;tr=udp%3A%2F%2Ftracker.coppersurfer.tk:6969/announce&amp;tr=udp%3A%2F%2Ftracker.leechers-paradise.org:6969/announce&amp;tr=udp%3A%2F%2Fbt.xxx-tracker.com:2710/announce&amp;tr=udp%3A%2F%2Ftracker.internetwarriors.net:1337/announce&amp;tr=udp%3A%2F%2Ftracker.openbittorrent.com:80/announce&amp;tr=udp%3A%2F%2Fexplodie.org:6969/announce&amp;tr=udp%3A%2F%2Ftracker.opentrackr.org:1337/announce&amp;tr=udp%3A%2F%2Ftracker.tiny-vps.com:6969/announce&amp;tr=udp%3A%2F%2Fopen.demonii.si:1337/announce&amp;tr=udp%3A%2F%2Ftracker.torrent.eu.org:451/announce&amp;tr=udp%3A%2F%2Ftracker.pirateparty.gr:6969/announce&amp;tr=udp%3A%2F%2Fipv4.tracker.harry.lu:80/announce&amp;tr=udp%3A%2F%2Ftracker.cyberia.is:6969/announce&amp;tr=udp%3A%2F%2F9.rarbg.to:2710/announce&amp;tr=udp%3A%2F%2Fdenis.stalker.upeer.me:6969/announce"},
		DownloadLink{Title: "Versão Dual Áudio | 4K UHD (MKV)", MagnetLink: "magnet:?xt=urn:btih:844E212E0E48CCA59470FCCFD9F07E98A1C96EFD&amp;dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%20%202019%20%284K%20-%20BluRay%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&amp;tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&amp;tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce"},
	}

	foundLinks := BludTV{}.GetDownloadLinks("https://www.bludv.tv/john-wick-3-parabellum-2019-dublado-dual-audio-legendado-bluray-720p-1080p-2160p-4k-baixar/")
	if len(expectedLinks) != len(foundLinks) {
		t.Errorf("The number of found links is not equal to the expected\nExpected: %d\nFound: %d\n", len(expectedLinks), len(foundLinks))
		return
	}

	for _, foundLink := range foundLinks {
		flag := false
		for _, expectedLink := range expectedLinks {
			if foundLink.Title == expectedLink.Title || foundLink.MagnetLink == expectedLink.MagnetLink {
				if strings.Compare(foundLink.Title, expectedLink.Title) != 0 {
					fmt.Println([]byte(foundLink.Title))
					fmt.Println([]byte(expectedLink.Title))
					t.Errorf("The title don't match\nFound: %s\nExpected: %s\n", foundLink.Title, expectedLink.Title)
				} else if foundLink.MagnetLink != expectedLink.MagnetLink {
					t.Errorf("The link don't match\nFound: %s\nExpected: %s\n", foundLink.MagnetLink, expectedLink.MagnetLink)
				}
				flag = true
			}
		}
		if !flag {
			t.Errorf("Found Links is not equal to the expected\n")
			return
		}
	}
}

func TestBludTVGetDownloadLinksSerie1(t *testing.T) {
	var expectedLinks = FoundMagnetLinks{
		DownloadLink{Title: "Versão Legendado (MKV)", MagnetLink: "magnet:?xt=urn:btih:921FA60E90916AF8EB61BA66F30CCE05D919BC9D&amp;dn=Mr.Pickles.S04.720p.LEG-RICKSZ&amp;tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce"},
	}

	foundLinks := BludTV{}.GetDownloadLinks("https://www.bludv.tv/mr-pickles-2019-4a-temporada-torrent-web-dl-720p-legendado-download/")
	if len(expectedLinks) != len(foundLinks) {
		t.Errorf("The number of found links is not equal to the expected\nExpected: %d\nFound: %d\n", len(expectedLinks), len(foundLinks))
		return
	}

	for _, foundLink := range foundLinks {
		flag := false
		for _, expectedLink := range expectedLinks {
			if foundLink.Title == expectedLink.Title || foundLink.MagnetLink == expectedLink.MagnetLink {
				if strcmp(foundLink.Title, expectedLink.Title) != 0 {
					fmt.Println([]byte(foundLink.Title))
					fmt.Println([]byte(expectedLink.Title))
					t.Errorf("The title don't match\nFound: %s\nExpected: %s\n", foundLink.Title, expectedLink.Title)
				} else if foundLink.MagnetLink != expectedLink.MagnetLink {
					t.Errorf("The link don't match\nFound: %s\nExpected: %s\n", foundLink.MagnetLink, expectedLink.MagnetLink)
				}
				flag = true
			}
		}
		if !flag {
			t.Errorf("Found Links is not equal to the expected\n")
			return
		}
	}
}

func TestGetDownloadLinksSerie2(t *testing.T) {
	var expectedLinks = FoundMagnetLinks{
		DownloadLink{Title: "Versão Dual Áudio | Sem Compactação (MKV)", MagnetLink: "magnet:?xt=urn:btih:B183921174CDB394FB3B290DB65776A0FC290DBC&amp;dn=Mr.Pickles.T01.2014.1080p.WEB-DL.H264.AC3.2.0-5.1.DUAL-MLD-BLUDV&amp;tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&amp;tr=udp%3a%2f%2f9.rarbg.to%3a2790%2fannounce&amp;tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&amp;tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&amp;tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&amp;tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&amp;tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&amp;tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce"},
	}

	foundLinks := BludTV{}.GetDownloadLinks("https://www.bludv.tv/mr-pickles-2014-1a-temporada-completa-torrent-bluray-full-1080p-dual-audio-download/")
	if len(expectedLinks) != len(foundLinks) {
		t.Errorf("The number of found links is not equal to the expected\nExpected: %d\nFound: %d\n", len(expectedLinks), len(foundLinks))
		return
	}

	for _, foundLink := range foundLinks {
		flag := false
		for _, expectedLink := range expectedLinks {
			if foundLink.Title == expectedLink.Title || foundLink.MagnetLink == expectedLink.MagnetLink {
				if strcmp(foundLink.Title, expectedLink.Title) != 0 {
					fmt.Println([]byte(foundLink.Title))
					fmt.Println([]byte(expectedLink.Title))
					t.Errorf("The title don't match\nFound: %s\nExpected: %s\n", foundLink.Title, expectedLink.Title)
				} else if foundLink.MagnetLink != expectedLink.MagnetLink {
					t.Errorf("The link don't match\nFound: %s\nExpected: %s\n", foundLink.MagnetLink, expectedLink.MagnetLink)
				}
				flag = true
			}
		}
		if !flag {
			t.Errorf("Found Links is not equal to the expected\n")
			return
		}
	}
}
