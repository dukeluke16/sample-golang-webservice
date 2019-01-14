package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/text/language"
)

func TestLocalization_Chinese_Taiwan(t *testing.T) {
	tag := "zh-TW"
	parsedTag := "zh-Hant"
	expectedCode := "US"
	expectedAlert := "完成預訂，即代表您同意 費用規則和限制 以及 危險物質規定 。"
	expectedTitle := "危險物品限制"
	expectedBody := []string{"美國聯邦法律禁止手提行李或您個人物品包含危險物品。如違反規定，可處五年監禁及 $250,000 以上 (49 U.S.C. 5124) 的罰鍰。有害物質包括易爆炸之壓縮氣體、易燃液體和固體、氧化劑、有毒物質、腐蝕物質和放射性物質。例如：油漆、打火機油、煙火、催淚氣體、氧氣瓶及放射線藥物。",
		"醫療用和盥洗用品，以及隨身攜帶特定吸菸用品則屬特殊例外，容許小量 (總重最高 70 盎司) 攜帶於行李。欲了解更多訊息，請聯繫您的航空公司。"}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Chinese_China(t *testing.T) {
	tag := "zh-CN"
	parsedTag := "zh-Hans"
	expectedCode := "US"
	expectedAlert := "完成此预订即表示同意 费用规则和限制 以及 危险物品政策 。"
	expectedTitle := "危险物品限制"
	expectedBody := []string{"联邦法律禁止在行李中携带或随身携带危险物品登机。如有违反，将被处以五年有期徒刑以及至少 250,000 美元的罚款 (49 U.S.C. 5124)。危险物品包括炸药、压缩气体、可燃液体和固体、氧化剂、毒药、腐蚀物和放射性物质。示例：油漆、打火机液、烟花爆竹、催泪瓦斯、氧气瓶和放射性药物。",
		"对于少量药品和化妆品（总计不超过 70 盎司），存在特殊例外，允许在行李中携带，特定烟草制品允许随身携带。有关更多信息，请联系航空公司代表。"}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Turkish(t *testing.T) {
	tag := "tr"
	parsedTag := "tr"
	expectedCode := "US"
	expectedAlert := "Bu rezervasyonu tamamlayarak tarife kurallarını ve kısıtlamaları ve tehlikeli madde ilkesini kabul etmiş olursunuz."
	expectedTitle := "Tehlikeli Madde Kısıtlamaları"
	expectedBody := []string{"Uçak içindeyken bagajınızda veya üzerinizde tehlikeli maddelerin taşınması federal yasalar tarafından yasaklanmıştır. Bu yasağın ihlal edilmesi beş yıllık hapis cezasına ve 250.000 ABD doları veya daha yüksek tutarda para cezalarına yol açabilir (49 U.S.C. 5124). Tehlikeli maddeler arasında patlayıcılar, sıkıştırılmış gazlar, yanıcı sıvı ve katı maddeler, oksitleyiciler, zehirler, aşındırıcılar ve radyoaktif maddeler bulunur. Örnekler: Boyalar, çakmak sıvısı, havai fişekler, göz yaşartıcı gazlar, oksijen şişeleri ve radyofarmasötikler.",
		"Bagajınızda bulunan küçük miktarlardaki tıbbi ve tuvalet malzemeleri (toplam en fazla 70 ons) ve üzerinizde taşıdığınız belirli sigara malzemeleri için özel istisnalar bulunur. Daha fazla bilgi için havayolu temsilcinizle iletişime geçin."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Swedish(t *testing.T) {
	tag := "sv"
	parsedTag := "sv"
	expectedCode := "US"
	expectedAlert := "Genom att slutföra den här bokningen godkänner du reglerna och restriktionerna för biljettpriset samt policyn för farligt gods."
	expectedTitle := "Restriktioner för farligt gods"
	expectedBody := []string{"Enligt federal lagstiftning är transport av farligt gods ombord på flygplanet i bagaget eller på kroppen förbjuden. Brott mot denna lag kan ge fem års fängelse och böter på 250 000 USD eller mer (49 U.S.C. 5124). Farligt gods inkluderar explosiva produkter, komprimerade gaser, brandfarliga vätskor och fasta ämnen, oxidationsmedel, gifter, frätande ämnen och radioaktiva material. Exempel på farligt gods: Färg, tändvätska, fyrverkerier, tårgas, syrgastuber och radioaktiva läkemedel.",
		"Särskilda undantag kan göras för små kvantiteter (upp till totalt 2 kg (70 ounces)) av läkemedel och hygienartiklar i bagaget samt vissa rökverk som du bär på dig. Kontakta en representant för ditt flygbolag om du vill veta mer."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Slovak(t *testing.T) {
	tag := "sk"
	parsedTag := "sk"
	expectedCode := "US"
	expectedAlert := "Dokončením tejto rezervácie súhlasíte s pravidlami a obmedzeniami pre tarify a pravidlami pre nebezpečný tovar."
	expectedTitle := "Obmedzenia nebezpečných látok"
	expectedBody := []string{"Federálne zákony zakazujú prevážanie nebezpečných látok na palube lietadla v batožine alebo pri sebe. Porušenie týchto zákonov sa môže trestať odňatím slobody na 5 rokov a pokutami vo výške 250000 USD alebo vyššie (49 U.S.C. 5124). Medzi nebezpečné látky patria výbušniny, stlačený plyn, horľavé kvapaliny a tuhé látky, okysličovadlá, jedy, žieraviny a rádioaktívne materiály. Napríklad: Farby, náplň do zapaľovačov, ohňostroje, slzné plyny, kyslíkové bomby a rádioaktívne lieky.",
		"Existujú špeciálne výnimky pre malé množstvá (do 70 uncí = 2 litrov) zdravotných alebo hygienických artiklov prevážaných v batožine a pre určitý fajčiarsky tovar prevážaný pri sebe. Ďalšie informácie vám poskytne zástupca leteckej spoločnosti."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Russian(t *testing.T) {
	tag := "ru"
	parsedTag := "ru"
	expectedCode := "US"
	expectedAlert := "Оформляя это бронирование, я подтверждаю свое согласие с правилами и ограничениями по тарифам, а также с политикой провоза опасных веществ."
	expectedTitle := "Ограничения по провозу опасных веществ"
	expectedBody := []string{"Федеральный закон запрещает провоз опасных веществ на борту самолета в багаже или ручной клади. Нарушение этого закона карается тюремным заключением на срок пять лет и штрафами в размере 250 000 долларов США и более (49 Свод законов США 5124). К опасным веществам относят взрывчатые вещества, газы под давлением, воспламеняющиеся жидкости и твердые вещества, окислители, яды, агрессивные и радиоактивные вещества. Например: краски, жидкости для заправки зажигалок, пиротехнические изделия, слезоточивый газ, жидкий кислород и радиофармацевтические препараты.",
		"Предусмотрены исключения для малых количеств (до 70 унций (2070 мл) совокупно) медицинских и косметических препаратов в багаже, а также для определенных курительных принадлежностей в ручной клади. Для получения дополнительной информации обратитесь к представителю авиакомпании."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Romanian(t *testing.T) {
	tag := "ro"
	parsedTag := "ro"
	expectedCode := "US"
	expectedAlert := "Finalizând această rezervare, sunteți de acord cu regulile și restricțiile de călătorie , precum și cu politica privind bunurile periculoase."
	expectedTitle := "Restricții materiale periculoase"
	expectedBody := []string{"Legea federală interzice transportul de materiale periculoase la bordul aeronavelor, în bagaj sau asupra dvs. Încălcarea acestei legi se poate pedepsi cu cinci ani de închisoare și penalizări de cel puțin 250.000 $ (49 U.S.C. 5124). Materialele periculoase includ explozibilii, gazele comprimate, lichidele și solidele inflamabile, oxidanții, substanțele otrăvitoare, agenții corozivi și materialele radioactive. Exemple: vopseluri, combustibilul pentru brichete, artificiile, gazele lacrimogene, recipientele de oxigen și produsele farmaceutice radioactive.",
		"Există excepții speciale pentru cantitățile mici (de până la 70 de uncii în total) de articole medicinale și de igienă transportate în bagajul dvs., precum și anumite materiale destinate fumatului asupra dvs. Pentru mai multe informații, contactați reprezentatul companiei dvs. aeriene."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_EuropeanPortuguese(t *testing.T) {
	tag := "pt-PT"
	parsedTag := "pt-PT"
	expectedCode := "US"
	expectedAlert := "Preenchendo esta reserva você concorda com as regras de tarifa e com as restrições e políticas de materiais perigosos vigentes."
	expectedTitle := "Restrições para Materiais Perigosos"
	expectedBody := []string{"A lei federal proíbe o transporte de materiais perigosos a bordo de aeronaves na bagagem ou com a pessoa. Um violação pode resultar em cinco anos de prisão e multa de US$ 250.000,00 ou mais (49 U.S.C. 5124). Materiais perigosos incluem explosivos, gases comprimidos, líquidos e sólidos inflamáveis, oxidantes, venenos, substâncias corrosivas e materiais radioativos. Exemplos: Tintas, fluido de isqueiro, fogos de artifício, gases lacrimogêneos, garrafas de oxigênio e produtos radiofarmacêuticos.",
		"Há exceções especiais para pequenas quantidades (até 70 onças no total) de artigos medicinais e de higiene pessoal, carregados em sua bagagem, e também alguns materiais de fumo carregados com a própria pessoa. Para obter mais informações, entre em contato com um representante da sua companhia aérea."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}
func TestLocalization_BrazilianPortuguese(t *testing.T) {
	tag := "pt-BR"
	parsedTag := "pt-BR"
	expectedCode := "US"
	expectedAlert := "Preenchendo esta reserva você concorda com as regras de tarifa e com as restrições e políticas de materiais perigosos vigentes."
	expectedTitle := "Restrições para Materiais Perigosos"
	expectedBody := []string{"A lei federal proíbe o transporte de materiais perigosos a bordo de aeronaves na bagagem ou com a pessoa. Um violação pode resultar em cinco anos de prisão e multa de US$ 250.000,00 ou mais (49 U.S.C. 5124). Materiais perigosos incluem explosivos, gases comprimidos, líquidos e sólidos inflamáveis, oxidantes, venenos, substâncias corrosivas e materiais radioativos. Exemplos: Tintas, fluido de isqueiro, fogos de artifício, gases lacrimogêneos, garrafas de oxigênio e produtos radiofarmacêuticos.",
		"Há exceções especiais para pequenas quantidades (até 70 onças no total) de artigos medicinais e de higiene pessoal, carregados em sua bagagem, e também alguns materiais de fumo carregados com a própria pessoa. Para obter mais informações, entre em contato com um representante da sua companhia aérea."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Polish(t *testing.T) {
	tag := "pl"
	parsedTag := "pl"
	expectedCode := "US"
	expectedAlert := "Kończąc niniejszą rezerwację, użytkownik zgadza się z regułami i ograniczeniami opłaty oraz polityką dotyczącą towarów niebezpiecznych."
	expectedTitle := "Ograniczenia dotyczące materiałów bezpośrednich"
	expectedBody := []string{"Przepisy federalne zakazują przewożenia na pokładzie samolotu materiałów niebezpiecznych osobiście lub w bagażu. Naruszenie tych przepisów grozi pięcioma latami więzienia i grzywną w wysokości od 250 000 USD (49 U.S.C. 5124). Materiały niebezpieczne obejmują materiały wybuchowe, sprężone gazy, łatwopalne płyny i substancje stałe, utleniacze, trucizny, środki korozyjne i materiały radioaktywne. Przykłady: farby, paliwo do zapalniczek, ognie sztuczne, gazy łzawiące, butle z tlenem i lekarstwa radioaktywne.",
		"Istnieją specjalne wyjątki dla małych ilości (w sumie do 2 l) artykułów medycznych i toaletowych przewożonych w bagażu i określonych materiałów do palenia przewożonych osobiście. Więcej informacji można uzyskać od przedstawiciela linii lotniczej."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Norwegian(t *testing.T) {
	tag := "no"
	parsedTag := "no"
	expectedCode := "US"
	expectedAlert := "Når du fullfører denne bestillingen, godtar du prisreglene og -restriksjonene og policyen for farlige materialer."
	expectedTitle := "Restriksjoner for farlige materialer"
	expectedBody := []string{"Føderale lover forbyr transport av farlige materialer i fly, enten i sendt bagasje eller håndbagasje. Brudd på denne loven kan føre til fengselsstraff og bøter på 250 000 eller mer (49 U.S.C. 5124). Farlige materialer inkluderer eksplosiver, komprimerte gasser, brannfarlige væsker og legemer, oksidasjonsmidler, giftstoffer, korrosjonsmidler og radioaktivt materiale. Eksempler: malingsstoffer, lightervæske, fyrverkeri, tåregass., oksygenflasker og radiofarmasøytiske stoffer.",
		"Det finnes spesielle unntak for små mengder (opptil 70 unser) med medisiner eller toalettsaker i bagasjen, og bestemte røykerelaterte artikler i håndbagasjen. Du får mer informasjon ved å kontakte flyselskapet."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Dutch(t *testing.T) {
	tag := "nl"
	parsedTag := "nl"
	expectedCode := "US"
	expectedAlert := "Door deze boeking te voltooien, gaat u akkoord met de regels en beperkingen voor vluchttarieven en het beleid inzake gevaarlijke goederen."
	expectedTitle := "Beperkingen ten aanzien van gevaarlijke stoffen"
	expectedBody := []string{"De federale wetgeving verbiedt het vervoer van gevaarlijke stoffen aan boord van een vliegtuig in uw bagage of als handbagage. Een schending van deze wetten kan leiden tot vijf jaar gevangenisstraf en boetes van $250.000 of meer (49 U.S.C. 5124). Gevaarlijke stoffen omvatten explosieven, samengeperste gassen, ontvlambare vloeistoffen en vaste stoffen, oxidatiemiddelen, giftige stoffen, corrosieve stoffen en radioactief materiaal. Voorbeelden: Verf, aanstekervloeistof, vuurwerk, traangas, zuurstofflessen en radiofarmaceutica.",
		"Er zijn speciale uitzonderingen voor kleine hoeveelheden (tot 70 ounce in totaal) medicijnen en toiletartikelen die in uw bagage worden meegenomen en bepaalde rookwaren die u in uw handbagage meeneemt. Voor meer informatie neemt u contact op met de vertegenwoordiger van uw luchtvaartmaatschappij."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Latvian(t *testing.T) {
	tag := "lv"
	parsedTag := "lv"
	expectedCode := "US"
	expectedAlert := "Aizpildot šo rezervāciju, jūs piekrītat braukšanas noteikumiem un ierobežojumiem, un bīstamo kravu politikas."
	expectedTitle := "Bīstamie materiāli ierobežojumi"
	expectedBody := []string{"Federālais likums aizliedz bīstamo materiālu klāja gaisa jūsu bagāžā vai par savu cilvēku pārvadāšanai. Pārkāpums var rasties piecu gadu cietumsodu un sodu 250.000 $ vai vairāk (49 USC 5124). Bīstamie materiāli ietver sprāgstvielu, saspiestas gāzes, uzliesmojoši šķidrumi un cietas vielas, oksidētāji, indes, korodantiem un radioaktīvi materiāli. Piemēri: Krāsām, vieglāks šķidrumu, uguņošana, asaru gāzes, skābekļa pudeles un radio farmaceitisko.",
		"Tur ir īpaši izņēmumi par maziem daudzumiem (kopā līdz 70 unces) par zāļu un tualetes izstrādājumi veic savu bagāžu un dažu smēķēšanas materiāli veic jūsu personu. Lai iegūtu sīkāku informāciju, sazinieties ar aviokompānijas pārstāvis."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Lithuanian(t *testing.T) {
	tag := "lt"
	parsedTag := "lt"
	expectedCode := "US"
	expectedAlert := "Užbaigti šį kartą, jūs sutinkate su kaina taisyklių ir apribojimų ir pavojingų krovinių politikos."
	expectedTitle := "Pavojingų medžiagų apribojimai"
	expectedBody := []string{"Federalinis įstatymas draudžia vežti pavojingų medžiagų laive bagaže ar jūsų asmeniu. Pažeidimas gali sukelti penkerių metų laisvės atėmimo ir bausmės $250,000 ir daugiau (49 U.S.C. 5124). Pavojingų medžiagų, kurios apima sprogmenų, suslėgtoms dujoms, degūs skysčiai ir kietosios medžiagos, oksidatorių, nuodingos, ėsdinančios ir radioaktyviųjų medžiagų. Pavyzdžiai: Dažai, žiebtuvėlis skystis, fejerverkai, ašarinės dujos, deguonies butelius ir radioaktyvieji.",
		"Yra taikomos specialios išimtys nedideliais kiekiais (iki 70 uncijos viso) vaistinių ir tualeto reikmenys atlikti savo bagažą ir tam tikrų rūkymo kilusių vykdoma savo asmenį. Dėl išsamesnės informacijos kreiptis jūsų aviakompanijos atstovas."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_(t *testing.T) {
	tag := "ko"
	parsedTag := "ko"
	expectedCode := "US"
	expectedAlert := "이 예약을 마치면 운임 규칙 및 제한사항 , 그리고 위험 물질 정책 에 동의하는 것입니다."
	expectedTitle := "위험 물질 제한"
	expectedBody := []string{"연방법에서는 위험 물질을 수화물로 맡기거나 직접 소지하고 비행기에 탑승하는 것을 금지합니다. 위반 시에는 5년의 징역형을 받거나 250,000달러 이상의 벌금을 물 수 있습니다(49 U.S.C. 5124). 위험 물질에는 폭발물, 압축가스, 인화성 액체 및 고체, 산화제, 독극물, 부식성 및 방사성 물질이 포함됩니다. 예: 페인트, 라이터용 연료, 불꽃, 최루 가스, 휴대용 산소통, 방사성 의약품.",
		"수화물에 포함하여 운반하는 의료용 물품 및 세면용품, 직접 소지하고 탑승하는 흡연 물질의 적은 용량(최대 합계 70온스)은 특별 예외 사항입니다. 자세한 내용은 항공사 담당자에게 문의하십시오."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Japanese(t *testing.T) {
	tag := "ja"
	parsedTag := "ja"
	expectedCode := "US"
	expectedAlert := "この予約を完了することで、運賃規則、制限事項、および 危険物ポリシー に同意することになります。"
	expectedTitle := "危険物に関する制限"
	expectedBody := []string{"連邦法は、荷物として、あるいは手荷物として、航空機で危険物を運送することを禁止しています。違反した場合、禁固 5 年および $25 万以上 (49 U.S.C. 5124) の罰金となることがあります。危険物には、爆発物、圧縮ガス、可燃性の液体や固体、酸化剤、毒物、腐食性物質、および放射性物質などが含まれます。例:塗料、可燃性の液体、花火、催涙ガス、酸素ボトル、放射性医薬品など。",
		"特別な例外としては、荷物として運送される少量の医薬品とトイレ用品 (合計最大 70 オンス)、そして手荷物として運送される喫煙用具があります。詳細については、各航空会社までお問い合わせください。"}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Italian(t *testing.T) {
	tag := "it"
	parsedTag := "it"
	expectedCode := "US"
	expectedAlert := "Completando questa prenotazione, si accettano le regole e le limitazioni sulle tariffe e sul trasporto di merci pericolose."
	expectedTitle := "Restrizioni su materiali pericolosi"
	expectedBody := []string{"La legge federale vieta il trasporto di materiali pericolosi a bordo del velivolo, nella valigia o con sé. La violazione della legge è un reato punibile con fino a 5 anni di reclusione e multe a partire da 250.000 $ (49 U.S.C. 5124). Sono considerati materiali pericolosi: esplosivi, gas compressi, liquidi e solidi infiammabili, ossidanti, veleni, materiali corrosivi e radioattivi. Esempi: vernici, fluidi per accenditi, fuochi d'artificio, gas lacrimogeni, contenitori di ossigeno e radiofarmaci.",
		"Sono previste delle eccezioni per piccole quantità (fino a circa 2 kg totali) di medicine e articoli per l'igiene personale che possono essere trasportate nel proprio bagaglio e di materiali per fumare con sé. Per ulteriori informazioni contattare il rappresentante della compagnia aerea."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Hungarian(t *testing.T) {
	tag := "hu"
	parsedTag := "hu"
	expectedCode := "US"
	expectedAlert := "A foglalás befejezésével kijelenti, hogy elfogadja a tarifaszabályzatot és a korlátozásokat, továbbá a veszélyes anyagok szállítására vonatkozó szabályzatot."
	expectedTitle := "Veszélyes anyagokra vonatkozó korlátozások"
	expectedBody := []string{"Az USA szövetségi törvényei tiltják a veszélyes anyagok felvitelét repülőgépekre, illetve a veszélyes anyagok légi szállítását, mind poggyászban, mind az emberi testen és testben. A törvény megsértése 5 évig terjedő szabadságvesztéssel, továbbá legalább 250000 USD összegű büntetéssel jár (az USA törvénykönyvének 49. fejezete, 5124. cikkely). Veszélyes anyagnak minősülnek a robbanóanyagok, a sűrített gázok, a gyúlékony folyadékok és más gyúlékony anyagok, az oxidálószerek, a mérgek, a maró hatású és a radioaktív anyagok. Néhány példa: festékek, tűzfokozó folyadékok, tűzijátékok, könnygázok, oxigénnel töltött palackok és radioaktív gyógyszerek.",
		"A poggyászokban szállított kis mennyiségű gyógyszerek és piperecikkek, illetve az utasnál lévő egyes dohánycikkek kivételt képeznek, amennyiben össztömegük nem haladja meg a 70 unciát – kb. 1900 grammot. A légitársaságok munkatársaitól részletesebb felvilágosítást is kérhet."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Croatian(t *testing.T) {
	tag := "hr"
	parsedTag := "hr"
	expectedCode := "US"
	expectedAlert := "Dovršetkom ove rezervacije pristajete na pravila i ograničenja cijene leta te pravilnik o opasnim tvarima."
	expectedTitle := "Ograničenja u pogledu opasnog materijala"
	expectedBody := []string{"Savezno zakonodavstvo zabranjuje unošenje opasnih materijala u zrakoplov sa sobom ili u prtljazi. Kršenje može dovesti do petogodišnje kazne zatvora i novčanih kazni od $250,000 ili više (glava 49. odjeljak 5124. Zakonika SAD-a). Opasni materijali uključuju eksplozive, komprimirane plinove, zapaljive tekućine i suhe tvari, oksidatore, otrove, korozivne tvar i radioaktivne materijale. Primjeri: boje, tekućina za upaljač, pirotehnička sredstva za vatromete, suzavci, boce s kisikom i radiofarmaceutici.",
		"Postoje posebne iznimke za male količine (do ukupno 2,07 litara) lijekova i toaletnog pribora u vašoj prtljazi te za određene materijale za pušenje koje nosite sa sobom. Za dodatne informacije obratite se svojem predstavniku zrakoplovne kompanije."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_French(t *testing.T) {
	tag := "fr"
	parsedTag := "fr"
	expectedCode := "US"
	expectedAlert := "En validant cette réservation, vous acceptez les règles tarifaires et les restrictions, ainsi que la politique sur les marchandises dangereuses."
	expectedTitle := "Restrictions pour substances dangereuses"
	expectedBody := []string{"La loi fédérale interdit le transport de substances dangereuses à bord des avions sur soi ou dans les bagages. Le non-respect de la réglementation en vigueur peut entraîner un emprisonnement de 5 ans et une amende de 250 000 dollars minimum (49 U.S.C. 5124). Les substances dangereuses comprennent les explosifs, les gaz comprimés, les liquides et solides inflammables, les oxydants, les poisons, les produits corrosifs et les matériaux radioactifs. Exemples : peintures, allume-feux liquides, feux d'artifice, gaz lacrymogènes, bouteilles d'oxygène et médicaments radiopharmaceutiques.",
		"Des exceptions spéciales existent pour les petites quantités (moins de 2 kilos) d'articles de toilette et médicinaux transportés dans vos bagages et pour certains produits pour fumeurs transportés sur vous-même. Pour en savoir plus, contactez le représentant de votre compagnie aérienne."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_FrenchCanadian(t *testing.T) {
	tag := "fr-CA"
	parsedTag := "fr-CA"
	expectedCode := "US"
	expectedAlert := "En finalisant cette réservation, vous acceptez les règles et restrictions tarifaires et les politiques relatives aux matières dangereuses."
	expectedTitle := "Restrictions relatives aux matières dangereuses"
	expectedBody := []string{"La législation fédérale interdit quiconque de transporter des matières dangereuses dans ses bagages ou sur lui à bord d'un avion. Toute violation donnera lieu à un emprisonnement de cinq ans et à une amende d'au moins 250 000 $ (49 U.S.C. 5124). Les matières dangereuses comprennent les explosifs, les gaz comprimés, les liquides et solides inflammables, les comburants, les poisons, et les matières radioactives et corrosives. Exemples : peintures, essence pour briquets, feux d'artifice, gaz lacrymogènes, bouteilles d'oxygène et produits radiopharmaceutiques.",
		"Il existe des exceptions particulières pour les faibles quantités (maximum de 70 onces au total) de produits de toilette et de médicaments transportés dans vos bagages et pour certains produits à fumer que vous transportez sur vous. Pour de plus amples renseignements, communiquez avec le représentant de votre compagnie aérienne."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Finnish(t *testing.T) {
	tag := "fi"
	parsedTag := "fi"
	expectedCode := "US"
	expectedAlert := "Tekemällä tämän varauksen valmiiksi hyväksyt hintasäännöt ja rajoitukset sekä vaarallisten aineiden käytännön."
	expectedTitle := "Vaarallisten aineiden rajoitukset"
	expectedBody := []string{"Liittovaltion laki kieltää vaarallisten aineiden kuljettamisen lentokoneessa matkatavaroissa tai matkustajan yllä. Rikkomus voi johtaa viiden vuoden vankeusrangaistukseen ja vähintään 250 000 dollarin sakkoon (49 U.S.C. 5124). Vaarallisia aineita ovat räjähteet, puristetut kaasut, syttyvät nesteet ja kiinteät aineet, hapettimet, myrkyt, syövyttävät aineet ja radioaktiiviset aineet. Esimerkkejä: maalit, sytytysnesteet, ilotulitusvälineet, kyynelkaasut, happipullot ja radiofarmaseuttiset valmisteet.",
		"Poikkeuksena ovat pienet määrät (yhteensä enintään 70 unssia, noin 2 litraa) lääkkeitä ja kosmeettisia valmisteita, jotka kuljetetaan matkatavaroissa, ja tietyt tupakointivälineet, jotka matkustaja kuljettaa yllään. Saat lisätietoja ottamalla yhteyden lentoyhtiön edustajaan."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_LatinAmericanSpanish(t *testing.T) {
	tag := "es-LA"
	parsedTag := "es"
	expectedCode := "US"
	expectedAlert := "Al completar esta reservación, acepta las restricciones y reglas de la tarifa y la política de bienes peligrosos."
	expectedTitle := "Restricciones sobre materiales peligrosos"
	expectedBody := []string{"La ley nacional prohíbe el transporte de materiales peligrosos a bordo de aviones en su equipaje o con usted. El incumplimiento de esta ley puede tener como resultado cinco años de prisión y multas de $250 000 o más (49 U.S.C. 5124). Entre los materiales peligrosos se encuentran explosivos, gases comprimidos, líquidos y sólidos inflamables, oxidantes, sustancias venenosas, corrosivas y materiales radioactivos. Ejemplos: Pinturas, fluido para encendedores, fuegos artificiales, gases lacrimógenos, botellas de oxígeno y productos radiofarmacéuticos.",
		"Existen excepciones especiales para pequeñas cantidades (hasta 70 onzas en total) de artículos medicinales y de aseo personal transportados en su equipaje y algunos elementos para fumar que lleve con usted. Para obtener más información, comuníquese con el representante de su aerolínea."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_EuropeanSpanish(t *testing.T) {
	tag := "es-ES"
	parsedTag := "es-ES"
	expectedCode := "US"
	expectedAlert := "Al efectuar esta reserva, manifiesta que está conforme con las restricciones y reglas de tarifas, así como con la política de productos peligrosos."
	expectedTitle := "Restricciones de materiales peligrosos"
	expectedBody := []string{"Las leyes federales prohíben el transporte de materiales peligrosos a bordo del avión, en el equipaje o en su cuerpo. Cualquier infracción de estas leyes puede tener como consecuencia penas de 5 años de cárcel y multas de 250 000 $ o más (49 U.S.C. 5124). Se consideran materiales peligrosos los siguientes: explosivos, gases comprimidos, líquidos y sólidos inflamables, así como materiales oxidantes, venenosos, corrosivos y radioactivos. Ejemplos: Pinturas, líquidos para encender fuego, materiales pirotécnicos, gases lacrimógenos, botellas de oxígeno y radiofármacos.",
		"Como excepciones especiales, se permite llevar pequeñas cantidades (hasta 1'98 kg en total) de productos médicos y artículos de aseo personal en la maleta y ciertos materiales para fumar que lleve personalmente. Para obtener más información, póngase en contacto con un representante de la línea aérea."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Spanish(t *testing.T) {
	tag := "es"
	parsedTag := "es"
	expectedCode := "US"
	expectedAlert := "Al completar esta reservación, acepta las restricciones y reglas de la tarifa y la política de bienes peligrosos."
	expectedTitle := "Restricciones sobre materiales peligrosos"
	expectedBody := []string{"La ley nacional prohíbe el transporte de materiales peligrosos a bordo de aviones en su equipaje o con usted. El incumplimiento de esta ley puede tener como resultado cinco años de prisión y multas de $250 000 o más (49 U.S.C. 5124). Entre los materiales peligrosos se encuentran explosivos, gases comprimidos, líquidos y sólidos inflamables, oxidantes, sustancias venenosas, corrosivas y materiales radioactivos. Ejemplos: Pinturas, fluido para encendedores, fuegos artificiales, gases lacrimógenos, botellas de oxígeno y productos radiofarmacéuticos.",
		"Existen excepciones especiales para pequeñas cantidades (hasta 70 onzas en total) de artículos medicinales y de aseo personal transportados en su equipaje y algunos elementos para fumar que lleve con usted. Para obtener más información, comuníquese con el representante de su aerolínea."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_GeneralEnglish(t *testing.T) {
	tag := "en"
	parsedTag := "en"
	expectedCode := "US"
	expectedAlert := "By completing this booking, you agree to the fare rules and restrictions and hazardous goods policy."
	expectedTitle := "Hazardous Materials Restrictions"
	expectedBody := []string{"Federal law forbids the carriage of hazardous materials aboard aircraft in your luggage or on your person. A violation can result in five years' imprisonment and penalties of $250,000 or more (49 U.S.C. 5124). Hazardous materials include explosives, compressed gases, flammable liquids and solids, oxidizers, poisons, corrosives and radioactive materials. Examples: Paints, lighter fluid, fireworks, tear gases, oxygen bottles, and radio-pharmaceuticals.",
		"There are special exceptions for small quantities (up to 70 ounces total) of medicinal and toilet articles carried in your luggage and certain smoking materials carried on your person. For further information contact your airline representative."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_AustralianEnglish(t *testing.T) {
	tag := "en-AU"
	parsedTag := "en-GB"
	expectedCode := "US"
	expectedAlert := "By completing this booking, you agree to the fare rules and restrictions and hazardous goods policy."
	expectedTitle := "Hazardous Materials Restrictions"
	expectedBody := []string{"Federal law forbids the carriage of hazardous materials aboard aircraft in your luggage or on your person. A violation can result in five years' imprisonment and penalties of $250,000 or more (49 U.S.C. 5124). Hazardous materials include explosives, compressed gases, flammable liquids and solids, oxidisers, poisons, corrosives and radioactive materials. Examples: Paints, lighter fluid, fireworks, tear gases, oxygen bottles and radio-pharmaceuticals.",
		"There are special exceptions for small quantities (up to 70 ounces total) of medicinal and toilet articles carried in your luggage and certain smoking materials carried on your person. For further information contact your airline representative."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_BritishEnglish(t *testing.T) {
	tag := "en-GB"
	parsedTag := "en-GB"
	expectedCode := "US"
	expectedAlert := "By completing this booking, you agree to the fare rules and restrictions and hazardous goods policy."
	expectedTitle := "Hazardous Materials Restrictions"
	expectedBody := []string{"Federal law forbids the carriage of hazardous materials aboard aircraft in your luggage or on your person. A violation can result in five years' imprisonment and penalties of $250,000 or more (49 U.S.C. 5124). Hazardous materials include explosives, compressed gases, flammable liquids and solids, oxidisers, poisons, corrosives and radioactive materials. Examples: Paints, lighter fluid, fireworks, tear gases, oxygen bottles and radio-pharmaceuticals.",
		"There are special exceptions for small quantities (up to 70 ounces total) of medicinal and toilet articles carried in your luggage and certain smoking materials carried on your person. For further information contact your airline representative."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_AmericanEnglish(t *testing.T) {
	tag := "en-US"
	parsedTag := "en-US"
	expectedCode := "US"
	expectedAlert := "By completing this booking, you agree to the fare rules and restrictions and hazardous goods policy."
	expectedTitle := "Hazardous Materials Restrictions"
	expectedBody := []string{"Federal law forbids the carriage of hazardous materials aboard aircraft in your luggage or on your person. A violation can result in five years' imprisonment and penalties of $250,000 or more (49 U.S.C. 5124). Hazardous materials include explosives, compressed gases, flammable liquids and solids, oxidizers, poisons, corrosives and radioactive materials. Examples: Paints, lighter fluid, fireworks, tear gases, oxygen bottles, and radio-pharmaceuticals.",
		"There are special exceptions for small quantities (up to 70 ounces total) of medicinal and toilet articles carried in your luggage and certain smoking materials carried on your person. For further information contact your airline representative."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Greek(t *testing.T) {
	tag := "el"
	parsedTag := "el"
	expectedCode := "US"
	expectedAlert := "Με την ολοκλήρωση αυτής της κράτησης, συμφωνείτε να το ναύλων περιορισμούς και πολιτική επικίνδυνων εμπορευμάτων."
	expectedTitle := "Επικίνδυνα υλικά περιορισμούς"
	expectedBody := []string{"Ομοσπονδιακός νόμος απαγορεύει τη μεταφορά επικίνδυνων υλικών επί αεροσκάφους στις αποσκευές σας ή στο πρόσωπό σας. Μια παραβίαση μπορεί να οδηγήσει σε φυλάκιση πέντε ετών και των ποινών για $250.000 και άνω (49 U.S.C. 5124). Τα επικίνδυνα υλικά περιλαμβάνουν εκρηκτικών, συμπιεσμένα αέρια, εύφλεκτα υγρά και στερεά, oxidizers, δηλητήρια, διαβρωτικά και ραδιενεργών υλικών. Παραδείγματα: Χρώματα, ελαφρύτερο ρευστό, πυροτεχνήματα, δακρυγόνων αερίων, φιάλες οξυγόνου, και ραδιοφαρμάκων.",
		"Εκεί είναι ειδικές εξαιρέσεις για μικρές ποσότητες (μέχρι 70 συνολικά ουγγιές) φαρμακευτικών και είδη καλλωπισμού που μεταφέρονται στις αποσκευές σας και ορισμένες καπνίζοντας υλικά μεταφέρονται στο πρόσωπό σας. Για περαιτέρω πληροφορίες επικοινωνήστε με σας εκπρόσωπος της αεροπορικής εταιρείας."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_German(t *testing.T) {
	tag := "de"
	parsedTag := "de"
	expectedCode := "US"
	expectedAlert := "Mit Abschluss dieser Buchung stimmen Sie den Tarifbedingungen und -einschränkungen sowie der Gefahrstoffrichtlinie zu."
	expectedTitle := "Gefahrgutbeschränkungen"
	expectedBody := []string{"Die Beförderung von Gefahrgut an Bord von Flugzeugen im Gepäck oder am Passagier ist gesetzlich verboten. Ein Verstoß gegen dieses Verbot kann hohe Strafen zur Folge haben (in den USA gemäß 49 U.S.C. 5124 fünf Jahre Haft und Geldstrafen von 250.000 $ oder mehr). Als Gefahrgut gelten Sprengstoffe, komprimierte Gase, brennbare Flüssigkeiten und Feststoffe, Oxidationsmittel, Gifte, Korrosionsmittel und radioaktive Materialien. Beispiele: Farben, Feuerzeugbenzin, Feuerwerk, Tränengas, Sauerstoffflaschen und Radiopharmaka.",
		"Es gibt spezielle Ausnahmen für kleine Mengen (insgesamt maximal ca. 1,9 kg) Arzneimittel und Körperpflegeartikel im Gepäck sowie für bestimmte Rauchutensilien am Passagier. Weitere Informationen erhalten Sie von Ihrer Fluglinie."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Danish(t *testing.T) {
	tag := "da"
	parsedTag := "da"
	expectedCode := "US"
	expectedAlert := "Ved at færdiggøre denne reservation accepterer du reglerne og begrænsningerne for billetpriser og politikken vedrørende farligt gods."
	expectedTitle := "Begrænsninger vedrørende farlige materialer"
	expectedBody := []string{"Amerikansk lovgivning forbyder, at der medtages farlige materialer i bagagen, eller at du har farlige materialer på dig ombord på et fly. Overtrædelser kan medføre fem års fængsel og bøder på USD 250.000 eller mere (49 U.S.C. 5124). Farlige materialer omfatter sprængstoffer, komprimeret gas, brandfarlige væsker og faste stoffer, iltningsmidler, giftstoffer, ætsende stoffer og radioaktive materialer.",
		"Der er særlige undtagelser for små mængder (op til 70 ounces i alt) af lægemidler og toiletartikler, som du kan have i bagagen, og bestemte rygeartikler, som du kan have på dig. For yderligere oplysninger skal du kontakte din flyselskabsrepræsentant."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Czech(t *testing.T) {
	tag := "cs"
	parsedTag := "cs"
	expectedCode := "US"
	expectedAlert := "Dokončením této rezervace vyjadřujete souhlas s pravidly a omezeními tarifů a omezeními pro nebezpečné látky."
	expectedTitle := "Omezení pro nebezpečné látky"
	expectedBody := []string{"Federální zákony zakazují přítomnost nebezpečných látek na palubě letadla a to jak v zavazadlech, tak u cestujících. Jejich porušení může mít za následek 5 let vězení a pokuty ve výši 250 000 $ a více (49 U.S.C. 5124). Mezi nebezpečné látky patří výbušniny, stlačené plyny, hořlavé kapaliny a pevné látky, oxidační činidla, jedy, žíraviny a radioaktivní materiály. Příklady: barvy, náplně do zapalovačů, pyrotechnika, slzné plyny, láhve s kyslíkem a radioaktivní farmaceutika.",
		"Existují zvláštní výjimky pro malá množství (až do 70 uncí) lékařských a toaletních potřeb převážených v zavazadle a určitých kuřáckých potřeb převážených osobně. Další informace získáte od zástupce letecká společnosti."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func TestLocalization_Bulgarian(t *testing.T) {
	tag := "bg"
	parsedTag := "bg"
	expectedCode := "US"
	expectedAlert := "Със завършването на тази резервация се съгласявате с правилата и ограниченията за тарифата и политиката за опасни вещества."
	expectedTitle := "Ограничения за опасни вещества"
	expectedBody := []string{"Федералният закон забранява пренасянето на опасни вещества на борда на самолета във вашия багаж или лично от вас. Нарушението може да доведе до пет години лишаване от свобода и глоби от $250 000 или повече (49 U.S.C. 5124). Опасните вещества включват експлозиви, компресирани газове, запалими течности и твърди вещества, окислители, токсични вещества, корозивни вещества и радиоактивни материали. Например: бои, течност за запалки, фойерверки, сълзотворен газ, кислородни бутилки и продукти на радиофармацевтичната промишленост.",
		"Има специални изключения за малки количества (до 70 унции/около 2 грама общо) лекарства и козметични артикули в багажа ви и определени продукти за пушене, носени лично от вас. За повече информация се обърнете към представителя на вашата авиокомпания."}

	localizationTest(t, tag, parsedTag, expectedCode, expectedAlert, expectedTitle, expectedBody)
}

func localizationTest(t *testing.T, tag string, parsedTag string, expectedCode string, expectedAlert string, expectedTitle string, expectedBody []string) {
	ts := setupFakeServerUSA()
	w := setupPostRequestAndServe(strings.NewReader(`["sea"]`), &tag)
	defer ts.Close()

	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	if w.HeaderMap["Content-Language"][0] != parsedTag {
		t.Errorf("handler returned wrong Content-Language: got %v want %v", w.HeaderMap, parsedTag)
	}

	parsedData := parseResponse(t, w)
	if len(parsedData) != 1 {
		t.Errorf("Array size error.")
	}

	policy := parsedData[0]
	if policy.Code != expectedCode {
		t.Errorf("handler returned unexpected code: got %v want %v",
			policy.Code, expectedCode)
	}

	if policy.Alert != expectedAlert {
		t.Errorf("handler returned unexpected alert: got %v want %v",
			policy.Alert, expectedAlert)
	}
	if policy.Title != expectedTitle {
		t.Errorf("handler returned unexpected title: got %v want %v",
			policy.Title, expectedTitle)
	}

	if policy.Body[0] != expectedBody[0] || policy.Body[1] != expectedBody[1] {
		t.Errorf("handler returned unexpected body: got \n%v \n\nwant \n%v",
			policy.Body, expectedBody)
	}
}

func TestLocalization_NotAcceptable(t *testing.T) {
	tag := language.Und.String()
	ts := setupFakeServerUSA()
	w := setupPostRequestAndServe(strings.NewReader(`["sea"]`), &tag)
	defer ts.Close()

	if w.Code != http.StatusNotAcceptable {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}
}

type hazardousGoodsPolicyModel struct {
	Code  string
	Alert string
	Title string
	Body  []string
}

func parseResponse(t *testing.T, w *httptest.ResponseRecorder) []hazardousGoodsPolicyModel {
	var parsedData []hazardousGoodsPolicyModel
	if err := json.Unmarshal(w.Body.Bytes(), &parsedData); err != nil {
		t.Errorf("Parsing error.")
	}
	return parsedData
}
