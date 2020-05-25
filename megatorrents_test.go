package moviescrawler

import (
	"testing"
	ht "golang.org/x/net/html"
)

func TestMegaTorrentsSearch(t *testing.T) {
	var expectedMovies = FoundMovies{
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2020/05/capa-ricos-de-amor-hd-filmes-200x300.jpg",Title: "Ricos de Amor (2020) HD WEB-DL 1080p Nacional",Link: "http://www.megatorrentshd.org/ricos-de-amor-2020-hd-web-dl-1080p-nacional/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2019/11/capa-cadeado-do-amor-hd-filmes-200x300.jpg",Title: "Cadeado do Amor (2019) HD WEB-DL 1080p FULL Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/cadeado-do-amor-2019-hd-web-dl-1080p-full-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2019/09/capa-Historia-de-Amor-No-Inverno-hd-filmes-200x300.jpg",Title: "História de Amor No Inverno (2019) HD WEB-DL 720p e 1080p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/historia-de-amor-no-inverno-2019-hd-web-dl-720p-e-1080p-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2019/09/capa-Minha-Namorada-e-uma-Vampira-hd-filmes.jpg",Title: "Minha Namorada é Uma Vampira (2016) &#8211; HD BluRay 720p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/minha-namorada-e-uma-vampira-2016-hd-bluray-720p-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2019/08/capa-Piranhas–Os-Meninos-da-Camorra-hd-filmes-200x300.jpg",Title: "Piranhas – Os Meninos da Camorra (2019) BluRay 720p e 1080p Legendado",Link: "http://www.megatorrentshd.org/piranhas-os-meninos-da-camorra-2019-bluray-720p-e-1080p-legendado/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2019/07/capa-sem-amor-hd-filmes-200x300.jpg",Title: "Sem Amor (2018) HD BluRay 720p e 1080p 5.1 Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/sem-amor-2018-hd-bluray-720p-e-1080p-5-1-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2019/04/capa-sexo-musica-e-amor-hd-filmes-200x300.jpg",Title: "Sexo, Música e Amor (2019) HD WEB-DL 720p e 1080p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/sexo-musica-e-amor-2019-hd-web-dl-720p-e-1080p-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2018/08/capa-o-melhor-presente-e-o-amor-hd-200x300.jpg",Title: "O Melhor Presente é o Amor (2018) – HD BluRay 720p e 1080p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/o-melhor-presente-e-o-amor-2018-hd-bluray-720p-e-1080p-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2018/06/capa-Sense8-Amor-Vincit-Omnia-hd-200x300.jpg",Title: "Sense8 – Amor Vincit Omnia (2018) – HD WEB-DL 720p e 1080p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/sense8-amor-vincit-omnia-2018-hd-web-dl-720p-e-1080p-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2018/05/capa-com-amor-simon-hd-200x300.jpg",Title: "Com Amor Simon (2018) &#8211; HD Bluray 720p e 1080p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/com-amor-simon-2018-hd-web-dl-720p-e-1080p-legendado/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2018/05/capa-Vamps-A-Morte-Nao-Existe-para-o-Amor-200x300.jpg",Title: "Vamps – A Morte Não Existe para o Amor (2018) – HD BluRay 720p e 1080p Dublado / Legendado",Link: "http://www.megatorrentshd.org/vamps-a-morte-nao-existe-para-o-amor-2018-hd-bluray-720p-e-1080p-dublado-legendado/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2018/05/capa-Kites-As-Barreiras-do-Amor-hd.jpg",Title: "Kites: As Barreiras do Amor (2010) – HD BluRay 720p e 1080p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/kites-as-barreiras-do-amor-2010-hd-bluray-720p-e-1080p-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2018/01/capa-Bruno-e-Boots–No-Amor-e-na-Guerra.jpg",Title: "Bruno e Boots – No Amor e na Guerra (2018) – HD WEB-DL 720p e 1080p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/bruno-e-boots-no-amor-e-na-guerra-2018-hd-web-dl-720p-e-1080p-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2017/12/capa-amor-e-tulipas-hd.jpg",Title: "Amor e Tulipas (2017) – HD BluRay 720p e 1080p Dual Áudio",Link: "http://www.megatorrentshd.org/amor-e-tulipas-2017-hd-bluray-720p-e-1080p-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2017/09/capa-Amores-Canibais-2017-hd.jpg",Title: "Amores Canibais (2017) – HD BluRay 720p e 1080p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/amores-canibais-2017-hd-bluray-720p-e-1080p-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2017/08/capa-amor-em-tempos-de-guerra-hd-1-200x300.jpg",Title: "Amor em Tempos de Guerra (2017) – HD BluRay 720p e 1080p Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/amor-em-tempos-de-guerra-2017-hd-bluray-720p-e-1080p-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2017/09/oscar-e-wilder-capa-um-amor-improvavel-hd-bluray-720p-e-1080p-dublado-200x300.jpg",Title: "Oscar e Wilder: Um Amor Improvável &#8211; HD BluRay 720p e 1080p Dublado",Link: "http://www.megatorrentshd.org/oscar-e-wilder-um-amor-improvavel-hd-bluray-720p-e-1080p-dublado/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2017/08/capa-uma-laço-de-amor-200x300.jpg",Title: "Um Laço de Amor (2017) – HD BluRay 1080p e 720p 5.1 Dublado / Dual Áudio",Link: "http://www.megatorrentshd.org/um-laco-de-amor-2017-hd-bluray-1080p-e-720p-5-1-dublado-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2017/06/capa-Em-Guerra-por-Amor-hd-200x300.jpg",Title: "Em Guerra por Amor (2017) &#8211; HD BluRay 720p e 1080p",Link: "http://www.megatorrentshd.org/em-guerra-por-amor-2017-hd-bluray-720p-e-1080p/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2017/06/capa-Amor.com-HD-200x300.jpg",Title: "Amor.com (2017) &#8211; HD 720p e 1080p Nacional",Link: "http://www.megatorrentshd.org/amor-com-2017-hd-720p-e-1080p-nacional/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2015/05/Patch-Adams-O-Amor-e-Contagioso.png",Title: "Patch Adams: O Amor é Contagioso – HD 720p e 1080p",Link: "http://www.megatorrentshd.org/patch-adams-o-amor-e-contagioso-hd-720p/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2017/01/capa-Amor-Amizade-HD-200x300.jpg",Title: "Amor &#038; Amizade 2017 &#8211; HD BluRay 720p e 1080p Dual Áudio",Link: "http://www.megatorrentshd.org/amor-amizade-2017-hd-bluray-720p-e-1080p-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2016/11/capa-Deusa-do-amor-hd-200x300.jpg",Title: "Deusa do Amor (2016) &#8211; HD 720p",Link: "http://www.megatorrentshd.org/deusa-do-amor-2016-hd-720p/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2016/11/capa-Um-Namorado-para-Minha-Mulher.jpg",Title: "Um Namorado Para Minha Mulher (2016) &#8211; HD 720p e 1080p Dublado",Link: "http://www.megatorrentshd.org/um-namorado-para-minha-mulher-2016-hd-720p-e-1080p-dublado/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2016/11/1aVlWQJtIRKaf6lYzNDyjxF5oad-200x300.jpg",Title: "Amor de Irmão (2016) – HD 720p e 1080p Dual Áudio",Link: "http://www.megatorrentshd.org/amor-de-irmao-2016-hd-720p-e-1080p-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2016/10/capa-minha-namorada-fora-de-controle-182x300.jpg",Title: "Minha Namorada Fora de Controle &#8211; HD 720p e 1080p Dual Áudio",Link: "http://www.megatorrentshd.org/minha-namorada-fora-de-controle-hd-720p-e-1080p-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2016/08/capa-Honey-3-200x300.jpg",Title: "Honey 3: No Ritmo do Amor (2016) &#8211; HD 720p e 1080p Dual Áudio",Link: "http://www.megatorrentshd.org/honey-3-no-ritmo-do-amor-2016-hd-720p-e-1080p-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2016/08/capa-Amor-Por-Direito-200x300.jpg",Title: "Amor Por Direito (2016) &#8211; HD 720p e 1080p Dual Áudio",Link: "http://www.megatorrentshd.org/amor-por-direito-2016-hd-720p-e-1080p-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2016/07/capa-O-Maior-Amor-do-Mundo-HD-200x300.jpg",Title: "O Maior Amor do Mundo (2016) &#8211; HD 720p e 1080p Dual Áudio",Link: "http://www.megatorrentshd.org/o-maior-amor-do-mundo-2016-hd-720p-e-1080p-dual-audio/"},
		Movie{Cover: "http://www.megatorrentshd.org/wp-content/uploads/2016/07/capa-Uma-Canção-de-Amor-para-Bobby-Long-2004-HD-min-211x300.png",Title: "Uma Canção de Amor para Bobby Long (2004) &#8211; HD 720p e 1080p Dual Áudio",Link: "http://www.megatorrentshd.org/uma-cancao-de-amor-para-bobby-long-2004-hd-720p-e-1080p-dual-audio/"},
	}
	foundMovies := MegaTorrents{}.SearchMovie("amor", 1)
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

func TestMegaTorrentsGetDownloadLinksActual(t *testing.T) {
	var expectedLinks = FoundMagnetLinks{
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:F564D1711F5817589A5C7E8CED467F65B9A1D93F&dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%20%202019%20%281080p%20-%20BluRay%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce", Title: "BluRay Rip HD 1080p Dual Áudio"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:7296D8B865D711B3E00FC2CDBE97F54890875F14&dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%202019%20%28720p%20-%20BluRay%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce", Title: "BluRay Rip HD 720p Dual Áudio"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:9BFEC0D85CF6B9A4B8BC66CBCAF4B409D6377023&dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%202019%20%281080p%20-%20BluRay%29%20%5bDUBLADO%5d%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce", Title: "BluRay Rip HD 1080p Dublado"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:8826B08649746CE6C7554F985F081A6F9F7E8A97&dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%202019%20%28720p%20-%20BluRay%29%20%5bDUBLADO%5d%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce", Title: "BluRay Rip HD 720p Dublado"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:E1658C93B1660E35DCFFA122DC654A41337385B4&dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%202019%20%281080p%20-%20BluRay%20-%20Legendado%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2740%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce", Title: "HD 1080p Legenda FIXA"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:11A2670079B5D448AD2FCE21AD24E0147B5A58B0&dn=John.Wick.3.2019.1080p.Bluray.X264-EVO%5BTGx%5D+%E2%AD%90&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fbt.xxx-tracker.com%3A2710%2Fannounce&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=udp%3A%2F%2Fopen.demonii.si%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce&tr=udp%3A%2F%2Ftracker.pirateparty.gr%3A6969%2Fannounce&tr=udp%3A%2F%2Fipv4.tracker.harry.lu%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.cyberia.is%3A6969%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2710%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce", Title: "BluRay ULTRA HD 1080p Legendado"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:C67627D20ED5BA7D7EC002A3610C33BAE83298F4&dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%20%202019%20%281080p%20-%20BluRay%20-%20FULL%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce", Title: "BluRay ULTRA HD 1080p Dual Áudio"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:419b2c9566bc2969a57d669ebdcdd35f14f9f2af&dn=John.Wick.Chapter.3.Parabellum.2019.INTERNAL.HDR.2160p.WEB.X265-DEFLATE&tr=http%3A%2F%2Ftracker.trackerfix.com%3A80%2Fannounce&tr=udp%3A%2F%2F9.rarbg.me%3A2720&tr=udp%3A%2F%2F9.rarbg.to%3A2730", Title: "4K Legendado"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:844E212E0E48CCA59470FCCFD9F07E98A1C96EFD&dn=%5bWWW.BLUDV.TV%5d%20John%20Wick%203%20-%20Parabellum%20%202019%20%284K%20-%20BluRay%29%20Acesse%20o%20ORIGINAL%20WWW.BLUDV.TV&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce", Title: "4K Dual Áudio"},
	}

	foundLinks := MegaTorrents{}.GetDownloadLinks("http://www.megatorrentshd.org/john-wick-3-parabellum-2019-hd-bluray-1080p-e-720p-2160p-dublado-e-legendado-5-1/")
	if len(expectedLinks) != len(foundLinks) {
		t.Errorf("The number of found links is not equal to the expected\nExpected: %d\nFound: %d\n", len(expectedLinks), len(foundLinks))
		return
	}

	for _, foundLink := range foundLinks {
		flag := false
		for _, expectedLink := range expectedLinks {
			if foundLink.Title == ht.UnescapeString(expectedLink.Title) || foundLink.MagnetLink == expectedLink.MagnetLink {
				if 	foundLink.Title != ht.UnescapeString(expectedLink.Title) {
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

func TestMegaTorrentsGetDownloadLinksElder(t *testing.T) {
	var expectedLinks = FoundMagnetLinks{
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:FD3E77DB2575FE09DF4BEC017292F0A27B0E1E61&dn=Donnie%20Darko%202001%20%281080p%29%20WWW.BLUDV.COM&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce", Title: "BluRay Rip HD 1080p Dual Áudio"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:F6B70D6425DC824C69EB787226C095C71C72719E&dn=Donnie%20Darko%202001%20%28720p%29%20WWW.BLUDV.COM&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce", Title: "BluRay Rip HD 720p Dual Áudio"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:29F6FB05116AA2DEE0711EA3411CE59F7FDFA736&dn=Donnie%20Darko%202001%20%28720p%29%20DUBLADO%20WWW.BLUDV.COM&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2780%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=http%3a%2f%2fglotorrents.pw%3a80%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2ftorrent.gresille.org%3a80%2fannounce&tr=udp%3a%2f%2ftracker.aletorrenty.pl%3a2710%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.piratepublic.com%3a1337%2fannounce", Title: "BluRay Rip HD 720p Dublado"},
	}

	foundLinks := MegaTorrents{}.GetDownloadLinks("http://www.megatorrentshd.org/donnie-darko-hd-bluray-720p-e-1080p-dublado-dual-audio/")
	if len(expectedLinks) != len(foundLinks) {
		t.Errorf("The number of found links is not equal to the expected\nExpected: %d\nFound: %d\n", len(expectedLinks), len(foundLinks))
		return
	}

	for _, foundLink := range foundLinks {
		flag := false
		for _, expectedLink := range expectedLinks {
			if foundLink.Title == ht.UnescapeString(expectedLink.Title) || foundLink.MagnetLink == expectedLink.MagnetLink {
				if 	foundLink.Title != ht.UnescapeString(expectedLink.Title) {
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

func TestMegaTorrentsGetDownloadLinksSerie(t *testing.T) {
	var expectedLinks = FoundMagnetLinks{
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:0388b8a8cb2238120211fa7300d454dbf43429f6&dn=LAPUMiA.Org%20-%20La%20Casa%20de%20Papel%202020%20-%20S04%20(1080p)&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.trackerfix.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.leechers-paradise.org%3a6969%2fannounce&tr=udp%3a%2f%2feddie4.nl%3a6969%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=udp%3a%2f%2fzer0day.ch%3a1337%2fannounce", Title: "4° Temporada Completa Dual Audio | WEB-DL 1080p Dual Áudio"},
		DownloadLink{MagnetLink: "magnet:?xt=urn:btih:74e224e745e45217a95b250413649bf4cb5059e2&dn=LAPUMiA.Org%20-%20La%20Casa%20de%20Papel%202020%20-%20S04%20(720p)&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2ftracker.openbittorrent.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.trackerfix.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2ftracker.leechers-paradise.org%3a6969%2fannounce&tr=udp%3a%2f%2feddie4.nl%3a6969%2fannounce&tr=udp%3a%2f%2fp4p.arenabg.com%3a1337%2fannounce&tr=udp%3a%2f%2fexplodie.org%3a6969%2fannounce&tr=udp%3a%2f%2fzer0day.ch%3a1337%2fannounce", Title: "4° Temporada Completa Dual Audio | WEB-DL 720p Dual Áudio"},
	}

	foundLinks := MegaTorrents{}.GetDownloadLinks("http://www.megatorrentshd.org/la-casa-de-papel-4a-temporada-completa-2020-web-dl-720p-e-1080p-5-1-hd-dublado-legendado/")
	if len(expectedLinks) != len(foundLinks) {
		t.Errorf("The number of found links is not equal to the expected\nExpected: %d\nFound: %d\n", len(expectedLinks), len(foundLinks))
		return
	}

	for _, foundLink := range foundLinks {
		flag := false
		for _, expectedLink := range expectedLinks {
			if foundLink.Title == ht.UnescapeString(expectedLink.Title) || foundLink.MagnetLink == expectedLink.MagnetLink {
				if 	foundLink.Title != ht.UnescapeString(expectedLink.Title) {
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