package openbd

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBooks(t *testing.T) {
	testCases := []struct {
		jsnStr      string
		str         string
		valid       []bool
		ids         []string
		isbn        []string
		title       []string
		subtitle    []string
		seriestitle []string
		label       []string
		img         []string
		authors     [][]string
		publisher   []string
		pubdate     []string
		desc        []string
	}{
		{
			jsnStr:      `[]`,
			str:         `[]`,
			valid:       []bool{},
			ids:         []string{},
			isbn:        []string{},
			title:       []string{},
			subtitle:    []string{},
			seriestitle: []string{},
			label:       []string{},
			img:         []string{},
			authors:     [][]string{{}},
			publisher:   []string{},
			pubdate:     []string{},
			desc:        []string{},
		},
		{
			jsnStr:      `[{}]`,
			str:         `[{"onix":{"RecordReference":"","NotificationType":"","ProductIdentifier":{"ProductIDType":"","IDValue":""},"DescriptiveDetail":{"ProductComposition":"","ProductForm":"","Collection":{"CollectionType":""},"TitleDetail":{"TitleType":"","TitleElement":{"TitleElementLevel":"","TitleText":{"content":""}}}},"PublishingDetail":{"Imprint":{"ImprintName":""},"Publisher":{"PublishingRole":"","PublisherName":""}},"ProductSupply":{"SupplyDetail":{"ReturnsConditions":{"ReturnsCodeType":"","ReturnsCode":""},"ProductAvailability":""}}},"hanmoto":{"dateshuppan":"","datemodified":"","datecreated":"","datekoukai":""},"summary":{"isbn":"","title":"","volume":"","series":"","publisher":"","pubdate":"","author":"","cover":""}}]`,
			valid:       []bool{false},
			ids:         []string{""},
			isbn:        []string{""},
			title:       []string{""},
			subtitle:    []string{""},
			seriestitle: []string{""},
			label:       []string{""},
			img:         []string{""},
			authors:     [][]string{{}},
			publisher:   []string{""},
			pubdate:     []string{"0001-01-01"},
			desc:        []string{""},
		},
		{
			jsnStr:      `[{"onix": {"RecordReference": "9784535585744", "NotificationType": "03", "ProductIdentifier": {"ProductIDType": "15", "IDValue": "9784535585744"}, "DescriptiveDetail": {"ProductComposition": "00", "ProductForm": "BZ", "Measure": [{"MeasureType": "01", "Measurement": "20", "MeasureUnitCode": "mm"}, {"MeasureType": "02", "Measurement": "0", "MeasureUnitCode": "mm"}], "TitleDetail": {"TitleType": "01", "TitleElement": {"TitleElementLevel": "01", "TitleText": {"collationkey": "ショク ノ リスクガク : ハンランスル アンゼン アンシン オ ヨミトク シテン", "content": "食のリスク学 : 氾濫する「安全・安心」をよみとく視点"}}}, "Contributor": [{"SequenceNumber": "1", "ContributorRole": ["A01"], "PersonName": {"content": "中西 準子"}}], "Language": [{"LanguageRole": "01", "LanguageCode": "jpn", "CountryCode": "JP"}], "Extent": [{"ExtentType": "11", "ExtentValue": "250", "ExtentUnit": "03"}]}, "CollateralDetail": {"TextContent": [{"TextType": "03", "ContentAudience": "00", "Text": "BSE、中国製ギョーザ、有機農業、健康食品など、あらゆる食の問題を俎上に、リスク評価の視点でさばく。環境リスク学を築き上げた著者が、環境問題に取り組む過程で踏み込んだ「食の問題」への明瞭な解答。"}, {"TextType": "04", "ContentAudience": "00", "Text": "第1章 食の安全-その費用と便益(安全を犠牲にすることもある\nさまざまな安全がある ほか)\n第2章 食べもの情報VS.リスク(対談・高橋久仁子)(フードファディズムに出会った\n砂糖は有害か? ほか)\n第3章 食をめぐる論争点-わたしはこう考える(ききて・松永和紀)(中国製の食品は危険ですか\n国産の方が安全で、安心ですか? ほか)\n第4章 さまざまな食の問題(キンメダイはスケープゴートだと考える理由-メチル水銀のリスクについて\n自然と循環は危険-だから、つきあい方には科学が必要 ほか)"}], "SupportingResource": [{"ResourceContentType": "01", "ContentAudience": "01", "ResourceMode": "03", "ResourceVersion": [{"ResourceForm": "02", "ResourceVersionFeature": [{"ResourceVersionFeatureType": "01", "FeatureValue": "D502"}, {"ResourceVersionFeatureType": "04", "FeatureValue": "9784535585744.jpg"}], "ResourceLink": "https://cover.openbd.jp/9784535585744.jpg"}]}]}, "PublishingDetail": {"Imprint": {"ImprintIdentifier": [{"ImprintIDType": "19", "IDValue": "535"}], "ImprintName": "日本評論社"}, "PublishingDate": [{"PublishingDateRole": "01", "Date": ""}]}, "ProductSupply": {"SupplyDetail": {"ReturnsConditions": {"ReturnsCodeType": "04", "ReturnsCode": "02"}, "ProductAvailability": "99"}}}, "hanmoto": {"datecreated": "2015-08-20 03:21:04", "reviews": [{"post_user": "genkina", "reviewer": "", "source_id": 23, "kubun_id": 1, "source": "日本経済新聞", "choyukan": "", "han": "", "link": "", "appearance": "2010-03-07", "gou": ""}], "dateshuppan": "2010-01", "datemodified": "2015-08-20 03:21:04"}, "summary": {"isbn": "9784535585744", "title": "食のリスク学 : 氾濫する「安全・安心」をよみとく視点", "volume": "", "series": "", "publisher": "日本評論社", "pubdate": "2010-01", "cover": "https://cover.openbd.jp/9784535585744.jpg", "author": "中西準子／著"}}]`,
			str:         `[{"onix":{"RecordReference":"9784535585744","NotificationType":"03","ProductIdentifier":{"ProductIDType":"15","IDValue":"9784535585744"},"DescriptiveDetail":{"ProductComposition":"00","ProductForm":"BZ","Measure":[{"MeasureType":"01","Measurement":"20","MeasureUnitCode":"mm"},{"MeasureType":"02","Measurement":"0","MeasureUnitCode":"mm"}],"Collection":{"CollectionType":""},"TitleDetail":{"TitleType":"01","TitleElement":{"TitleElementLevel":"01","TitleText":{"content":"食のリスク学 : 氾濫する「安全・安心」をよみとく視点","collationkey":"ショク ノ リスクガク : ハンランスル アンゼン アンシン オ ヨミトク シテン"}}},"Contributor":[{"SequenceNumber":"1","ContributorRole":["A01"],"PersonName":{"content":"中西 準子"}}],"Language":[{"LanguageRole":"01","LanguageCode":"jpn","CountryCode":"JP"}],"Extent":[{"ExtentType":"11","ExtentValue":"250","ExtentUnit":"03"}]},"CollateralDetail":{"TextContent":[{"TextType":"03","ContentAudience":"00","Text":"BSE、中国製ギョーザ、有機農業、健康食品など、あらゆる食の問題を俎上に、リスク評価の視点でさばく。環境リスク学を築き上げた著者が、環境問題に取り組む過程で踏み込んだ「食の問題」への明瞭な解答。"},{"TextType":"04","ContentAudience":"00","Text":"第1章 食の安全-その費用と便益(安全を犠牲にすることもある\nさまざまな安全がある ほか)\n第2章 食べもの情報VS.リスク(対談・高橋久仁子)(フードファディズムに出会った\n砂糖は有害か? ほか)\n第3章 食をめぐる論争点-わたしはこう考える(ききて・松永和紀)(中国製の食品は危険ですか\n国産の方が安全で、安心ですか? ほか)\n第4章 さまざまな食の問題(キンメダイはスケープゴートだと考える理由-メチル水銀のリスクについて\n自然と循環は危険-だから、つきあい方には科学が必要 ほか)"}],"SupportingResource":[{"ResourceContentType":"01","ContentAudience":"01","ResourceMode":"03","ResourceVersion":[{"ResourceForm":"02","ResourceVersionFeature":[{"ResourceVersionFeatureType":"01","FeatureValue":"D502"},{"ResourceVersionFeatureType":"04","FeatureValue":"9784535585744.jpg"}],"ResourceLink":"https://cover.openbd.jp/9784535585744.jpg"}]}]},"PublishingDetail":{"Imprint":{"ImprintIdentifier":[{"ImprintIDType":"19","IDValue":"535"}],"ImprintName":"日本評論社"},"Publisher":{"PublishingRole":"","PublisherName":""},"PublishingDate":[{"Date":"","PublishingDateRole":"01"}]},"ProductSupply":{"SupplyDetail":{"ReturnsConditions":{"ReturnsCodeType":"04","ReturnsCode":"02"},"ProductAvailability":"99"}}},"hanmoto":{"dateshuppan":"2010-01-01","datemodified":"2015-08-20","datecreated":"2015-08-20","datekoukai":"","reviews":[{"reviewer":"","link":"","appearance":"2010-03-07","kubun_id":1,"source_id":23,"source":"日本経済新聞","choyukan":"","post_user":"genkina","han":"","gou":""}]},"summary":{"isbn":"9784535585744","title":"食のリスク学 : 氾濫する「安全・安心」をよみとく視点","volume":"","series":"","publisher":"日本評論社","pubdate":"2010-01-01","author":"中西準子／著","cover":"https://cover.openbd.jp/9784535585744.jpg"}}]`,
			valid:       []bool{true},
			ids:         []string{"9784535585744"},
			isbn:        []string{"9784535585744"},
			title:       []string{"食のリスク学 : 氾濫する「安全・安心」をよみとく視点"},
			subtitle:    []string{""},
			seriestitle: []string{""},
			label:       []string{""},
			img:         []string{"https://cover.openbd.jp/9784535585744.jpg"},
			authors:     [][]string{{"中西 準子"}},
			publisher:   []string{"日本評論社"},
			pubdate:     []string{"2010-01-01"},
			desc:        []string{"BSE、中国製ギョーザ、有機農業、健康食品など、あらゆる食の問題を俎上に、リスク評価の視点でさばく。環境リスク学を築き上げた著者が、環境問題に取り組む過程で踏み込んだ「食の問題」への明瞭な解答。"},
		},
		{
			jsnStr:      `[{"onix": {"RecordReference": "9784797369915", "NotificationType": "03", "ProductIdentifier": {"ProductIDType": "15", "IDValue": "9784797369915"}, "DescriptiveDetail": {"ProductComposition": "00", "ProductForm": "BZ", "Measure": [{"MeasureType": "01", "Measurement": "150", "MeasureUnitCode": "mm"}, {"MeasureType": "02", "Measurement": "0", "MeasureUnitCode": "mm"}], "Collection": {"CollectionType": "10", "TitleDetail": {"TitleType": "01", "TitleElement": [{"TitleElementLevel": "02", "TitleText": {"content": "GA文庫"}}]}}, "TitleDetail": {"TitleType": "01", "TitleElement": {"TitleElementLevel": "01", "TitleText": {"collationkey": "イノウ バトル ワ ニチジョウケイ ノ ナカ デ", "content": "異能バトルは日常系のなかで"}}}, "Contributor": [{"SequenceNumber": "1", "ContributorRole": ["A01"], "PersonName": {"content": "望 公太"}}], "Language": [{"LanguageRole": "01", "LanguageCode": "jpn", "CountryCode": "JP"}], "Extent": [{"ExtentType": "11", "ExtentValue": "301", "ExtentUnit": "03"}]}, "CollateralDetail": {"TextContent": [{"TextType": "03", "ContentAudience": "00", "Text": "俺を含めた文芸部の五人は半年前、とてつもない能力に目覚めた。そして壮大なる学園異能バトルの世界へ足を踏み入れ-なかった!?「なんも起きねえのかよ!」異能に覚醒してみたものの、日常は完全無欠に平和だ。世界を滅ぼす秘密機関などない!異能戦争もない!勇者も魔王もいやしないっ!だから俺たちはこの超級異能を、「黒炎の龍にヒゲ生やせたーっ!」気軽に無駄遣いすることに決めた。だが異能バトルに憧れ続けた俺には分かる。真なる戦いの刻が…。「はぁ、バッカじゃないの?」神スキルとたわむれる何気ない日常。だが、それだけじゃ終わらない新・異能バトル&ラブコメ、開幕。"}], "SupportingResource": [{"ResourceContentType": "01", "ContentAudience": "01", "ResourceMode": "03", "ResourceVersion": [{"ResourceForm": "02", "ResourceVersionFeature": [{"ResourceVersionFeatureType": "01", "FeatureValue": "D502"}, {"ResourceVersionFeatureType": "04", "FeatureValue": "9784797369915.jpg"}], "ResourceLink": "https://cover.openbd.jp/9784797369915.jpg"}]}]}, "PublishingDetail": {"Imprint": {"ImprintIdentifier": [{"ImprintIDType": "19", "IDValue": "7973"}], "ImprintName": "ソフトバンククリエイティブ"}, "PublishingDate": [{"PublishingDateRole": "01", "Date": ""}]}, "ProductSupply": {"SupplyDetail": {"ReturnsConditions": {"ReturnsCodeType": "04", "ReturnsCode": "02"}, "ProductAvailability": "99"}}}, "hanmoto": {"datecreated": "2016-07-19 14:14:09", "lanove": true, "dateshuppan": "2012-06", "datemodified": "2016-07-19 14:14:09"}, "summary": {"isbn": "9784797369915", "title": "異能バトルは日常系のなかで", "volume": "", "series": "", "publisher": "ソフトバンククリエイティブ", "pubdate": "2012-06", "cover": "https://cover.openbd.jp/9784797369915.jpg", "author": "望公太／著"}}]`,
			str:         `[{"onix":{"RecordReference":"9784797369915","NotificationType":"03","ProductIdentifier":{"ProductIDType":"15","IDValue":"9784797369915"},"DescriptiveDetail":{"ProductComposition":"00","ProductForm":"BZ","Measure":[{"MeasureType":"01","Measurement":"150","MeasureUnitCode":"mm"},{"MeasureType":"02","Measurement":"0","MeasureUnitCode":"mm"}],"Collection":{"CollectionType":"10","TitleDetail":{"TitleType":"01","TitleElement":[{"TitleElementLevel":"02","TitleText":{"content":"GA文庫"}}]}},"TitleDetail":{"TitleType":"01","TitleElement":{"TitleElementLevel":"01","TitleText":{"content":"異能バトルは日常系のなかで","collationkey":"イノウ バトル ワ ニチジョウケイ ノ ナカ デ"}}},"Contributor":[{"SequenceNumber":"1","ContributorRole":["A01"],"PersonName":{"content":"望 公太"}}],"Language":[{"LanguageRole":"01","LanguageCode":"jpn","CountryCode":"JP"}],"Extent":[{"ExtentType":"11","ExtentValue":"301","ExtentUnit":"03"}]},"CollateralDetail":{"TextContent":[{"TextType":"03","ContentAudience":"00","Text":"俺を含めた文芸部の五人は半年前、とてつもない能力に目覚めた。そして壮大なる学園異能バトルの世界へ足を踏み入れ-なかった!?「なんも起きねえのかよ!」異能に覚醒してみたものの、日常は完全無欠に平和だ。世界を滅ぼす秘密機関などない!異能戦争もない!勇者も魔王もいやしないっ!だから俺たちはこの超級異能を、「黒炎の龍にヒゲ生やせたーっ!」気軽に無駄遣いすることに決めた。だが異能バトルに憧れ続けた俺には分かる。真なる戦いの刻が…。「はぁ、バッカじゃないの?」神スキルとたわむれる何気ない日常。だが、それだけじゃ終わらない新・異能バトル\u0026ラブコメ、開幕。"}],"SupportingResource":[{"ResourceContentType":"01","ContentAudience":"01","ResourceMode":"03","ResourceVersion":[{"ResourceForm":"02","ResourceVersionFeature":[{"ResourceVersionFeatureType":"01","FeatureValue":"D502"},{"ResourceVersionFeatureType":"04","FeatureValue":"9784797369915.jpg"}],"ResourceLink":"https://cover.openbd.jp/9784797369915.jpg"}]}]},"PublishingDetail":{"Imprint":{"ImprintIdentifier":[{"ImprintIDType":"19","IDValue":"7973"}],"ImprintName":"ソフトバンククリエイティブ"},"Publisher":{"PublishingRole":"","PublisherName":""},"PublishingDate":[{"Date":"","PublishingDateRole":"01"}]},"ProductSupply":{"SupplyDetail":{"ReturnsConditions":{"ReturnsCodeType":"04","ReturnsCode":"02"},"ProductAvailability":"99"}}},"hanmoto":{"dateshuppan":"2012-06-01","datemodified":"2016-07-19","datecreated":"2016-07-19","datekoukai":"","lanove":true},"summary":{"isbn":"9784797369915","title":"異能バトルは日常系のなかで","volume":"","series":"","publisher":"ソフトバンククリエイティブ","pubdate":"2012-06-01","author":"望公太／著","cover":"https://cover.openbd.jp/9784797369915.jpg"}}]`,
			valid:       []bool{true},
			ids:         []string{"9784797369915"},
			isbn:        []string{"9784797369915"},
			title:       []string{"異能バトルは日常系のなかで"},
			subtitle:    []string{""},
			seriestitle: []string{""},
			label:       []string{"GA文庫"},
			img:         []string{"https://cover.openbd.jp/9784797369915.jpg"},
			authors:     [][]string{{"望 公太"}},
			publisher:   []string{"ソフトバンククリエイティブ"},
			pubdate:     []string{"2012-06-01"},
			desc:        []string{"俺を含めた文芸部の五人は半年前、とてつもない能力に目覚めた。そして壮大なる学園異能バトルの世界へ足を踏み入れ-なかった!?「なんも起きねえのかよ!」異能に覚醒してみたものの、日常は完全無欠に平和だ。世界を滅ぼす秘密機関などない!異能戦争もない!勇者も魔王もいやしないっ!だから俺たちはこの超級異能を、「黒炎の龍にヒゲ生やせたーっ!」気軽に無駄遣いすることに決めた。だが異能バトルに憧れ続けた俺には分かる。真なる戦いの刻が…。「はぁ、バッカじゃないの?」神スキルとたわむれる何気ない日常。だが、それだけじゃ終わらない新・異能バトル&ラブコメ、開幕。"},
		},
		{
			jsnStr:      `[{"onix":{"RecordReference":"9784797391398","NotificationType":"03","ProductIdentifier":{"ProductIDType":"15","IDValue":"9784797391398"},"DescriptiveDetail":{"ProductComposition":"00","ProductForm":"BA","ProductFormDetail":"B119","Collection":{"CollectionType":"10","CollectionSequence":{"CollectionSequenceType":"01","CollectionSequenceTypeName":"完結フラグ","CollectionSequenceNumber":"0"}},"TitleDetail":{"TitleType":"01","TitleElement":{"TitleElementLevel":"01","TitleText":{"collationkey":"スウガクガールノヒミツノートビットトバイナリー","content":"数学ガールの秘密ノート／ビットとバイナリー"}}},"Contributor":[{"SequenceNumber":"1","ContributorRole":["A01"],"PersonName":{"collationkey":"ユウキヒロシ","content":"結城 浩"}}],"Language":[{"LanguageRole":"01","LanguageCode":"jpn"}],"Subject":[{"SubjectSchemeIdentifier":"78","SubjectCode":"0041"},{"SubjectSchemeIdentifier":"79","SubjectCode":"17"}],"Audience":[{"AudienceCodeType":"22","AudienceCodeValue":"00"}]},"CollateralDetail":{"TextContent":[{"TextType":"02","ContentAudience":"00","Text":"コンピュータを支える原理を説き明かす"},{"TextType":"03","ContentAudience":"00","Text":"数学ガールの登場人物が、楽しみつつコンピュータを支える原理について学ぶ数学読み物です。抽象的な解説でなく、具体的な実例を使って解説していますので、中学生から楽しんでいただける内容になっています。プログラミング教育の必要性が叫ばれる現代において、重要な一冊となるでしょう。\n\n●本書の構成\n\n第1章「冒険ビット」では、コンピュータ内部のデータ構造を支える「二進法」について学びます。十進法と二進法の関係、二進法が作り出す0と1のパターン、マイナス二進法などを観察しながら、普段私たちが使っているものとは異なる数の表記法を発見的に楽しみます。\n\n第2章「変幻ピクセル」では、コンピュータグラフィクスの基本である「ピクセル」について学びます。数が生み出すパターンによって画像を構成する原理を、小さなスキャナとプリンタによって具体的に体験します。また、ひとつひとつの手順をコンピュータに実行させるというプログラミングの基礎も実例を通して学びます。\n\n第3章「ラティス・サラダ」では、「束」という代数学について学びます。ビット列が生み出すたくさんのパターンをどのように整理するかを調べていくうちに、代数学の考え方の基本に触れます。\n\n第4章「コンプリメント・コンプレックス」では、コンピュータで用いられる「ビット演算」について学びます。ビット演算と数との関係を考え、計算について自由に発想を広げる経験をします。ルーラー関数を調べながらフラクタル構造を発見する体験も魅力です。\n\n第5章「フリップ・トリップ」では、フリップ・トリップという楽しいゲームで遊びながら、そこに隠れている数学的な構造を理解します。仲間同士で議論しながら、新しい問題に立ち向かうおもしろさを味わいつつ本書を終えます。\n\nどの章も、中学校や高校ではあまり学ばない題材を扱いつつも、決して難しい内容にはならず、登場人物といっしょになって遊びながらコンピュータの原理を学べる内容になっています。\n本書を読むことで、決して古びることのないコンピュータの原理を誰でも楽しみつつ学べるでしょう。"},{"TextType":"04","ContentAudience":"00","Text":"あなたへ\n\nプロローグ\n第1章「冒険ビット」\n第2章「変幻ピクセル」\n第3章「ラティス・サラダ」\n第4章「コンプリメント・コンプレックス」\n第5章「フリップ・トリップ」\nエピローグ\n\nもっと考えたいあなたのために\n"}],"SupportingResource":[{"ResourceContentType":"01","ContentAudience":"01","ResourceMode":"03","ResourceVersion":[{"ResourceForm":"02","ResourceVersionFeature":[{"ResourceVersionFeatureType":"01","FeatureValue":"D502"},{"ResourceVersionFeatureType":"04","FeatureValue":"9784797391398.jpg"}],"ResourceLink":"https:\/\/cover.openbd.jp\/9784797391398.jpg"}]}]},"PublishingDetail":{"Imprint":{"ImprintIdentifier":[{"ImprintIDType":"19","IDValue":"8156"},{"ImprintIDType":"24","IDValue":"4199"}],"ImprintName":"ＳＢクリエイティブ"},"Publisher":{"PublishingRole":"01","PublisherIdentifier":[{"PublisherIDType":"19","IDValue":"8156"},{"PublisherIDType":"24","IDValue":"4199"}],"PublisherName":"ＳＢクリエイティブ"},"PublishingDate":[{"PublishingDateRole":"01","Date":"20190722"}]},"ProductSupply":{"MarketPublishingDetail":{"MarketPublishingStatus":"00"},"SupplyDetail":{"ProductAvailability":"99","Price":[{"PriceType":"03","PriceAmount":"1500","CurrencyCode":"JPY"}]}}},"hanmoto":{"datemodified":"2019-07-18 10:13:03","datecreated":"2019-05-07 16:09:41","datekoukai":"2019-05-07"},"summary":{"isbn":"9784797391398","title":"数学ガールの秘密ノート／ビットとバイナリー","volume":"","series":"","publisher":"ＳＢクリエイティブ","pubdate":"20190722","cover":"https:\/\/cover.openbd.jp\/9784797391398.jpg","author":"結城浩／著"}}]`,
			str:         `[{"onix":{"RecordReference":"9784797391398","NotificationType":"03","ProductIdentifier":{"ProductIDType":"15","IDValue":"9784797391398"},"DescriptiveDetail":{"ProductComposition":"00","ProductForm":"BA","Collection":{"CollectionType":"10","CollectionSequence":{"CollectionSequenceType":"01","CollectionSequenceTypeName":"完結フラグ","CollectionSequenceNumber":"0"}},"TitleDetail":{"TitleType":"01","TitleElement":{"TitleElementLevel":"01","TitleText":{"content":"数学ガールの秘密ノート／ビットとバイナリー","collationkey":"スウガクガールノヒミツノートビットトバイナリー"}}},"Contributor":[{"SequenceNumber":"1","ContributorRole":["A01"],"PersonName":{"content":"結城 浩","collationkey":"ユウキヒロシ"}}],"Language":[{"LanguageRole":"01","LanguageCode":"jpn","CountryCode":""}],"Subject":[{"SubjectSchemeIdentifier":"78","SubjectCode":"0041"},{"SubjectSchemeIdentifier":"79","SubjectCode":"17"}],"Audience":[{"AudienceCodeType":"22","AudienceCodeValue":"00"}]},"CollateralDetail":{"TextContent":[{"TextType":"02","ContentAudience":"00","Text":"コンピュータを支える原理を説き明かす"},{"TextType":"03","ContentAudience":"00","Text":"数学ガールの登場人物が、楽しみつつコンピュータを支える原理について学ぶ数学読み物です。抽象的な解説でなく、具体的な実例を使って解説していますので、中学生から楽しんでいただける内容になっています。プログラミング教育の必要性が叫ばれる現代において、重要な一冊となるでしょう。\n\n●本書の構成\n\n第1章「冒険ビット」では、コンピュータ内部のデータ構造を支える「二進法」について学びます。十進法と二進法の関係、二進法が作り出す0と1のパターン、マイナス二進法などを観察しながら、普段私たちが使っているものとは異なる数の表記法を発見的に楽しみます。\n\n第2章「変幻ピクセル」では、コンピュータグラフィクスの基本である「ピクセル」について学びます。数が生み出すパターンによって画像を構成する原理を、小さなスキャナとプリンタによって具体的に体験します。また、ひとつひとつの手順をコンピュータに実行させるというプログラミングの基礎も実例を通して学びます。\n\n第3章「ラティス・サラダ」では、「束」という代数学について学びます。ビット列が生み出すたくさんのパターンをどのように整理するかを調べていくうちに、代数学の考え方の基本に触れます。\n\n第4章「コンプリメント・コンプレックス」では、コンピュータで用いられる「ビット演算」について学びます。ビット演算と数との関係を考え、計算について自由に発想を広げる経験をします。ルーラー関数を調べながらフラクタル構造を発見する体験も魅力です。\n\n第5章「フリップ・トリップ」では、フリップ・トリップという楽しいゲームで遊びながら、そこに隠れている数学的な構造を理解します。仲間同士で議論しながら、新しい問題に立ち向かうおもしろさを味わいつつ本書を終えます。\n\nどの章も、中学校や高校ではあまり学ばない題材を扱いつつも、決して難しい内容にはならず、登場人物といっしょになって遊びながらコンピュータの原理を学べる内容になっています。\n本書を読むことで、決して古びることのないコンピュータの原理を誰でも楽しみつつ学べるでしょう。"},{"TextType":"04","ContentAudience":"00","Text":"あなたへ\n\nプロローグ\n第1章「冒険ビット」\n第2章「変幻ピクセル」\n第3章「ラティス・サラダ」\n第4章「コンプリメント・コンプレックス」\n第5章「フリップ・トリップ」\nエピローグ\n\nもっと考えたいあなたのために\n"}],"SupportingResource":[{"ResourceContentType":"01","ContentAudience":"01","ResourceMode":"03","ResourceVersion":[{"ResourceForm":"02","ResourceVersionFeature":[{"ResourceVersionFeatureType":"01","FeatureValue":"D502"},{"ResourceVersionFeatureType":"04","FeatureValue":"9784797391398.jpg"}],"ResourceLink":"https://cover.openbd.jp/9784797391398.jpg"}]}]},"PublishingDetail":{"Imprint":{"ImprintIdentifier":[{"ImprintIDType":"19","IDValue":"8156"},{"ImprintIDType":"24","IDValue":"4199"}],"ImprintName":"ＳＢクリエイティブ"},"Publisher":{"PublisherIdentifier":[{"PublisherIDType":"19","IDValue":"8156"},{"PublisherIDType":"24","IDValue":"4199"}],"PublishingRole":"01","PublisherName":"ＳＢクリエイティブ"},"PublishingDate":[{"Date":"2019-07-22","PublishingDateRole":"01"}]},"ProductSupply":{"SupplyDetail":{"ReturnsConditions":{"ReturnsCodeType":"","ReturnsCode":""},"ProductAvailability":"99","Price":[{"PriceType":"03","CurrencyCode":"JPY","PriceAmount":"1500"}]}}},"hanmoto":{"dateshuppan":"","datemodified":"2019-07-18","datecreated":"2019-05-07","datekoukai":"2019-05-07"},"summary":{"isbn":"9784797391398","title":"数学ガールの秘密ノート／ビットとバイナリー","volume":"","series":"","publisher":"ＳＢクリエイティブ","pubdate":"2019-07-22","author":"結城浩／著","cover":"https://cover.openbd.jp/9784797391398.jpg"}}]`,
			valid:       []bool{true},
			ids:         []string{"9784797391398"},
			isbn:        []string{"9784797391398"},
			title:       []string{"数学ガールの秘密ノート／ビットとバイナリー"},
			subtitle:    []string{""},
			seriestitle: []string{""},
			label:       []string{""},
			img:         []string{"https://cover.openbd.jp/9784797391398.jpg"},
			authors:     [][]string{{"結城 浩"}},
			publisher:   []string{"ＳＢクリエイティブ"},
			pubdate:     []string{"2019-07-22"},
			desc:        []string{"数学ガールの登場人物が、楽しみつつコンピュータを支える原理について学ぶ数学読み物です。抽象的な解説でなく、具体的な実例を使って解説していますので、中学生から楽しんでいただける内容になっています。プログラミング教育の必要性が叫ばれる現代において、重要な一冊となるでしょう。\n\n●本書の構成\n\n第1章「冒険ビット」では、コンピュータ内部のデータ構造を支える「二進法」について学びます。十進法と二進法の関係、二進法が作り出す0と1のパターン、マイナス二進法などを観察しながら、普段私たちが使っているものとは異なる数の表記法を発見的に楽しみます。\n\n第2章「変幻ピクセル」では、コンピュータグラフィクスの基本である「ピクセル」について学びます。数が生み出すパターンによって画像を構成する原理を、小さなスキャナとプリンタによって具体的に体験します。また、ひとつひとつの手順をコンピュータに実行させるというプログラミングの基礎も実例を通して学びます。\n\n第3章「ラティス・サラダ」では、「束」という代数学について学びます。ビット列が生み出すたくさんのパターンをどのように整理するかを調べていくうちに、代数学の考え方の基本に触れます。\n\n第4章「コンプリメント・コンプレックス」では、コンピュータで用いられる「ビット演算」について学びます。ビット演算と数との関係を考え、計算について自由に発想を広げる経験をします。ルーラー関数を調べながらフラクタル構造を発見する体験も魅力です。\n\n第5章「フリップ・トリップ」では、フリップ・トリップという楽しいゲームで遊びながら、そこに隠れている数学的な構造を理解します。仲間同士で議論しながら、新しい問題に立ち向かうおもしろさを味わいつつ本書を終えます。\n\nどの章も、中学校や高校ではあまり学ばない題材を扱いつつも、決して難しい内容にはならず、登場人物といっしょになって遊びながらコンピュータの原理を学べる内容になっています。\n本書を読むことで、決して古びることのないコンピュータの原理を誰でも楽しみつつ学べるでしょう。"},
		},
	}

	for _, tc := range testCases {
		bks, err := DecodeBooks([]byte(tc.jsnStr))
		if err != nil {
			t.Errorf("DecodeBooks() is error \"%v\", not want error.", err)
			fmt.Printf("Info(TestError): %+v\n", err)
			continue
		}
		b, err := EncodeBooks(bks)
		if err != nil {
			t.Errorf("EncodeBooks() is error \"%v\", not want error.", err)
			fmt.Printf("Info(TestError): %+v\n", err)
			continue
		}
		str := string(b)
		if str != tc.str {
			t.Errorf("EncodeBooks() is \"%v\"", str)
		}
		if len(bks) != len(tc.valid) {
			t.Errorf("Count of Books is %v, want %v", len(bks), len(tc.valid))
			continue
		}
		for i, bk := range bks {
			if bk.IsValid() != tc.valid[i] {
				t.Errorf("Book[%d] is %v, want %v", i, bk.IsValid(), tc.valid[i])
				continue
			}
			if bk.Id() != tc.ids[i] {
				t.Errorf("Book[%d].Id() is \"%v\", want \"%v\"", i, bk.Id(), tc.ids[i])
			}
			if bk.ISBN() != tc.isbn[i] {
				t.Errorf("Book[%d].ISBN() is \"%v\", want \"%v\"", i, bk.ISBN(), tc.isbn[i])
			}
			if bk.Title() != tc.title[i] {
				t.Errorf("Book[%d].Title() is \"%v\", want \"%v\"", i, bk.Title(), tc.title[i])
			}
			if bk.SubTitle() != tc.subtitle[i] {
				t.Errorf("Book[%d].SubTitle() is \"%v\", want \"%v\"", i, bk.SubTitle(), tc.subtitle[i])
			}
			if bk.SeriesTitle() != tc.seriestitle[i] {
				t.Errorf("Book[%d].SeriesTitle() is \"%v\", want \"%v\"", i, bk.SeriesTitle(), tc.seriestitle[i])
			}
			if bk.Label() != tc.label[i] {
				t.Errorf("Book[%d].Label() is \"%v\", want \"%v\"", i, bk.Label(), tc.label[i])
			}
			if bk.ImageURL() != tc.img[i] {
				t.Errorf("Book[%d].ImageURL() is \"%v\", want \"%v\"", i, bk.ImageURL(), tc.img[i])
			}
			if !reflect.DeepEqual(bk.Authors(), tc.authors[i]) {
				t.Errorf("Book[%d].Authors() is \"%v\", want \"%v\"", i, bk.Authors(), tc.authors[i])
			}
			if bk.Publisher() != tc.publisher[i] {
				t.Errorf("Book[%d].Publisher() is \"%v\", want \"%v\"", i, bk.Publisher(), tc.publisher[i])
			}
			if bk.PublicationDate().String() != tc.pubdate[i] {
				t.Errorf("Book[%d].PublicationDate() is \"%v\", want \"%v\"", i, bk.PublicationDate().String(), tc.pubdate[i])
			}
			if bk.Description() != tc.desc[i] {
				t.Errorf("Book[%d].Description() is \"%v\", want \"%v\"", i, bk.Description(), tc.desc[i])
			}
		}
	}
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
