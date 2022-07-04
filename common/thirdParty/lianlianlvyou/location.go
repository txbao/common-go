package lianlianlvyou

type Location struct {
}

//生成代码
/*
func crateCode()  {

	var sdk = &lianlianlvyou.LianLian{
		ChannelId: 1659,
		Key:       "0uwXooRn4hlI8gERkewq9w==",
		Url:       "https://uapi.lianlianlvyou.com",
	}
	params := make(map[string]interface{})
	params["hasProLocation"] = "0"
	//params["Type"] = 1
	res3, err := sdk.LocationList(params)
	if err != nil {
		fmt.Println("ERR:", err.Error())
	} else {

		//生成代码
		codeStr := `thirdLocation := 0
	Location := 0
	switch thirdLocation {`
		for _, v := range *res3 {

			city := utils.StrReplace("总站", "", v.City, -1)
			city = utils.StrReplace("2站", "", city, -1)
			city = utils.StrReplace("3站", "", city, -1)
			city = utils.StrReplace("4站", "", city, -1)
			city = utils.StrReplace("5站", "", city, -1)
			city = utils.StrReplace("6站", "", city, -1)
			city = utils.StrReplace("7站", "", city, -1)
			city = utils.StrReplace("8站", "", city, -1)
			city = utils.StrReplace("9站", "", city, -1)
			city = utils.StrReplace("10站", "", city, -1)

			city = utils.StrReplace("特惠", "", city, -1)
			city = utils.StrReplace("精选", "", city, -1)
			city = utils.StrReplace("优选", "", city, -1)

			city = utils.StrReplace("新站", "", city, -1)
			city = utils.StrReplace("老站", "", city, -1)
			city = utils.StrReplace("旧站", "", city, -1)

			city = utils.StrReplace("特惠站", "", city, -1)
			city = utils.StrReplace("江北站", "", city, -1)
			city = utils.StrReplace("怒江州站", "怒江", city, -1)
			city = utils.StrReplace("始兴仁化", "始兴", city, -1)
			city = utils.StrReplace("北京通州站", "通州区", city, -1)
			city = utils.StrReplace("兴仁册亨站", "", city, -1)
			city = utils.StrReplace("全国产品", "全国", city, -1)





			city = utils.StrReplace("总站", "", city, -1)
			city = utils.StrReplace("站", "", city, -1)
			r, err := l.svcCtx.BaseLocationModel.FindSeach(city)
			errStr := ""
			if err != nil {
				errStr = err.Error()
			} else {
				errStr = fmt.Sprintf("name=%v,id=%v", r.Name, r.Id)
			}

			codeStr += fmt.Sprintf("\ncase %v:\n		location = %v //原站点名：%s,结果：%v", v.ID, r.Id, v.City, errStr+" | " + city)
		}
		codeStr += `
default:
		location = 0
	}`
		utils.FilePutContents("log/location_code.txt", codeStr, 0777)
	}
	resJson3, _ := utils.StructToJsonStr(res3)
	utils.FilePutContents("log/location.txt", resJson3, 0777)
}
*/

//通过联联站点获取第三方站点
func (o *Location) GetLocation(thirdLocation int64) int64 {
	var location int64 = 0
	switch thirdLocation {
	case 10000:
		location = 3374 //原站点名：全国产品,结果：name=全国,id=3374 | 全国
	case 10026:
		location = 0 //原站点名：小程序,结果：record not found | 小程序
	case 515:
		location = 296 //原站点名：安康站,结果：name=安康市,id=296 | 安康
	case 532:
		location = 1401 //原站点名：安吉站,结果：name=安吉县,id=1401 | 安吉
	case 788:
		location = 1591 //原站点名：安溪站,结果：name=安溪县,id=1591 | 安溪
	case 22:
		location = 259 //原站点名：安顺站,结果：name=安顺市,id=259 | 安顺
	case 29:
		location = 2833 //原站点名：安宁站,结果：name=安宁市,id=2833 | 安宁
	case 290:
		location = 105 //原站点名：安庆站,结果：name=安庆市,id=105 | 安庆
	case 306:
		location = 344 //原站点名：澳门站,结果：name=澳门特别行政区,id=344 | 澳门
	case 576:
		location = 0 //原站点名：安州站,结果：record not found | 安州
	case 358:
		location = 2833 //原站点名：安宁总站,结果：name=安宁市,id=2833 | 安宁
	case 359:
		location = 156 //原站点名：安阳总站,结果：name=安阳市,id=156 | 安阳
	case 371:
		location = 39 //原站点名：鞍山站,结果：name=鞍山市,id=39 | 鞍山
	case 643:
		location = 332 //原站点名：阿克苏站,结果：name=阿克苏地区,id=332 | 阿克苏
	case 391:
		location = 39 //原站点名：鞍山总站,结果：name=鞍山市,id=39 | 鞍山
	case 160:
		location = 156 //原站点名：安阳站,结果：name=安阳市,id=156 | 安阳
	case 197:
		location = 252 //原站点名：资阳2站,结果：name=资阳市,id=252 | 资阳
	case 476:
		location = 2209 //原站点名：安化站,结果：name=安化县,id=2209 | 安化
	case 479:
		location = 2681 //原站点名：安岳站,结果：name=安岳县,id=2681 | 安岳
	case 487:
		location = 4619 //原站点名：安徽站,结果：name=安徽省,id=4619 | 安徽
	case 505:
		location = 1224 //原站点名：安达站,结果：name=安达市,id=1224 | 安达
	case 257:
		location = 2523 //原站点名：巴南站,结果：name=巴南区,id=2523 | 巴南
	case 18:
		location = 26 //原站点名：包头站,结果：name=包头市,id=26 | 包头
	case 24:
		location = 290 //原站点名：宝鸡站,结果：name=宝鸡市,id=290 | 宝鸡
	case 794:
		location = 227 //原站点名：百色市站,结果：name=百色市,id=227 | 百色市
	case 283:
		location = 100 //原站点名：蚌埠站,结果：name=蚌埠市,id=100 | 蚌埠
	case 286:
		location = 262 //原站点名：毕节站,结果：name=毕节地区,id=262 | 毕节
	case 798:
		location = 806 //原站点名：保德站,结果：name=保德县,id=806 | 保德
	case 287:
		location = 251 //原站点名：巴中站,结果：name=巴中市,id=251 | 巴中
	case 35:
		location = 1 //原站点名：北京站,结果：name=北京市,id=1 | 北京
	case 320:
		location = 222 //原站点名：北海站,结果：name=北海市,id=222 | 北海
	case 841:
		location = 2600 //原站点名：北川站,结果：name=北川羌族自治县,id=2600 | 北川
	case 338:
		location = 2531 //原站点名：璧山站,结果：name=璧山县,id=2531 | 璧山
	case 339:
		location = 262 //原站点名：毕节总站,结果：name=毕节地区,id=262 | 毕节
	case 860:
		location = 998 //原站点名：鲅鱼圈站,结果：name=鲅鱼圈区,id=998 | 鲅鱼圈
	case 607:
		location = 41 //原站点名：本溪站,结果：name=本溪市,id=41 | 本溪
	case 96:
		location = 8 //原站点名：保定站,结果：name=保定市,id=8 | 保定
	case 608:
		location = 544 //原站点名：宝坻站,结果：name=宝坻区,id=544 | 宝坻
	case 353:
		location = 290 //原站点名：宝鸡总站,结果：name=宝鸡市,id=290 | 宝鸡
	case 865:
		location = 0 //原站点名：彬州站,结果：record not found | 彬州
	case 407:
		location = 150 //原站点名：滨州站,结果：name=滨州市,id=150 | 滨州
	case 155:
		location = 1155 //原站点名：宝山站,结果：name=宝山区,id=1155 | 宝山
	case 423:
		location = 112 //原站点名：亳州站,结果：name=亳州市,id=112 | 亳州
	case 686:
		location = 32 //原站点名：巴彦淖尔站,结果：name=巴彦淖尔市,id=32 | 巴彦淖尔
	case 180:
		location = 1325 //原站点名：滨海站,结果：name=滨海县,id=1325 | 滨海
	case 194:
		location = 2283 //原站点名：宝安站,结果：name=宝安区,id=2283 | 宝安
	case 452:
		location = 707 //原站点名：霸州站,结果：name=霸州市,id=707 | 霸州
	case 714:
		location = 1915 //原站点名：宝丰站,结果：name=宝丰县,id=1915 | 宝丰
	case 741:
		location = 1215 //原站点名：北安站,结果：name=北安市,id=1215 | 北安
	case 506:
		location = 1334 //原站点名：宝应站,结果：name=宝应县,id=1334 | 宝应
	case 251:
		location = 88 //原站点名：宁波2站,结果：name=宁波市,id=88 | 宁波
	case 507:
		location = 268 //原站点名：保山站,结果：name=保山市,id=268 | 保山
	case 1:
		location = 235 //原站点名：成都站,结果：name=成都市,id=235 | 成都
	case 2:
		location = 234 //原站点名：重庆站,结果：name=重庆市,id=234 | 重庆
	case 516:
		location = 1364 //原站点名：淳安站,结果：name=淳安县,id=1364 | 淳安
	case 517:
		location = 1385 //原站点名：苍南站,结果：name=苍南县,id=1385 | 苍南
	case 775:
		location = 598 //原站点名：成安站,结果：name=成安县,id=598 | 成安
	case 264:
		location = 10 //原站点名：承德站,结果：name=承德市,id=10 | 承德
	case 15:
		location = 183 //原站点名：长沙站,结果：name=长沙市,id=183 | 长沙
	case 528:
		location = 1944 //原站点名：长垣站,结果：name=长垣县,id=1944 | 长垣
	case 531:
		location = 2170 //原站点名：常宁站,结果：name=常宁市,id=2170 | 常宁
	case 791:
		location = 1853 //原站点名：茌平站,结果：name=茌平县,id=1853 | 茌平
	case 281:
		location = 215 //原站点名：潮州站,结果：name=潮州市,id=215 | 潮州
	case 282:
		location = 107 //原站点名：滁州站,结果：name=滁州市,id=107 | 滁州
	case 549:
		location = 1969 //原站点名：长葛站,结果：name=长葛市,id=1969 | 长葛
	case 807:
		location = 49 //原站点名：朝阳站,结果：name=朝阳市,id=49 | 朝阳
	case 809:
		location = 231 //原站点名：崇左站,结果：name=崇左市,id=231 | 崇左
	case 811:
		location = 675 //原站点名：承德县站,结果：name=承德县,id=675 | 承德县
	case 308:
		location = 2569 //原站点名：崇州站,结果：name=崇州市,id=2569 | 崇州
	case 309:
		location = 110 //原站点名：巢湖站,结果：name=巢湖市,id=110 | 巢湖
	case 311:
		location = 197 //原站点名：广州7站,结果：name=广州市,id=197 | 广州
	case 69:
		location = 51 //原站点名：长春站,结果：name=长春市,id=51 | 长春
	case 328:
		location = 183 //原站点名：长沙总站,结果：name=长沙市,id=183 | 长沙
	case 78:
		location = 77 //原站点名：常州站,结果：name=常州市,id=77 | 常州
	case 850:
		location = 2121 //原站点名：赤壁站,结果：name=赤壁市,id=2121 | 赤壁
	case 853:
		location = 2203 //原站点名：慈利站,结果：name=慈利县,id=2203 | 慈利
	case 94:
		location = 189 //原站点名：常德站,结果：name=常德市,id=189 | 常德
	case 365:
		location = 17 //原站点名：长治站,结果：name=长治市,id=17 | 长治
	case 623:
		location = 2119 //原站点名：崇阳站,结果：name=崇阳县,id=2119 | 崇阳
	case 116:
		location = 192 //原站点名：郴州站,结果：name=郴州市,id=192 | 郴州
	case 628:
		location = 589 //原站点名：昌黎站,结果：name=昌黎县,id=589 | 昌黎
	case 373:
		location = 192 //原站点名：郴州总站,结果：name=郴州市,id=192 | 郴州
	case 631:
		location = 2533 //原站点名：城口站,结果：name=城口县,id=2533 | 城口
	case 647:
		location = 601 //原站点名：磁县站,结果：name=磁县,id=601 | 磁县
	case 393:
		location = 28 //原站点名：赤峰站,结果：name=赤峰市,id=28 | 赤峰
	case 138:
		location = 11 //原站点名：沧州站,结果：name=沧州市,id=11 | 沧州
	case 157:
		location = 273 //原站点名：楚雄站,结果：name=楚雄彝族自治州,id=273 | 楚雄
	case 159:
		location = 1294 //原站点名：常熟站,结果：name=常熟市,id=1294 | 常熟
	case 168:
		location = 524 //原站点名：昌平站,结果：name=昌平区,id=524 | 昌平
	case 682:
		location = 2845 //原站点名：澄江站,结果：name=澄江县,id=2845 | 澄江
	case 176:
		location = 1377 //原站点名：慈溪站,结果：name=慈溪市,id=1377 | 慈溪
	case 432:
		location = 234 //原站点名：重庆特惠站,结果：name=重庆市,id=234 | 重庆
	case 435:
		location = 183 //原站点名：长沙精选,结果：name=长沙市,id=183 | 长沙
	case 465:
		location = 329 //原站点名：昌吉站,结果：name=昌吉回族自治州,id=329 | 昌吉
	case 468:
		location = 2525 //原站点名：长寿站,结果：name=长寿区,id=2525 | 长寿
	case 475:
		location = 235 //原站点名：成都特惠站,结果：name=成都市,id=235 | 成都
	case 220:
		location = 2825 //原站点名：呈贡站,结果：name=呈贡县,id=2825 | 呈贡
	case 227:
		location = 183 //原站点名：长沙特惠站,结果：name=长沙市,id=183 | 长沙
	case 739:
		location = 582 //原站点名：曹妃甸站,结果：name=曹妃甸区,id=582 | 曹妃甸
	case 231:
		location = 77 //原站点名：常州2站,结果：name=常州市,id=77 | 常州
	case 761:
		location = 1868 //原站点名：成武站,结果：name=成武县,id=1868 | 成武
	case 260:
		location = 139 //原站点名：东营站,结果：name=东营市,id=139 | 东营
	case 778:
		location = 987 //原站点名：东港站,结果：name=东港市,id=987 | 东港
	case 270:
		location = 65 //原站点名：大庆站,结果：name=大庆市,id=65 | 大庆
	case 536:
		location = 652 //原站点名：定州站,结果：name=定州市,id=652 | 定州
	case 793:
		location = 2563 //原站点名：大邑站,结果：name=大邑县,id=2563 | 大邑
	case 544:
		location = 2472 //原站点名：都安站,结果：name=都安瑶族自治县,id=2472 | 都安
	case 801:
		location = 1466 //原站点名：大通站,结果：name=大通区,id=1466 | 大通
	case 301:
		location = 148 //原站点名：德州站,结果：name=德州市,id=148 | 德州
	case 561:
		location = 1593 //原站点名：德化站,结果：name=德化县,id=1593 | 德化
	case 51:
		location = 249 //原站点名：达州站,结果：name=达州市,id=249 | 达州
	case 826:
		location = 599 //原站点名：大名站,结果：name=大名县,id=599 | 大名
	case 575:
		location = 2713 //原站点名：稻城站,结果：name=稻城县,id=2713 | 稻城
	case 71:
		location = 38 //原站点名：大连站,结果：name=大连市,id=38 | 大连
	case 76:
		location = 239 //原站点名：德阳站,结果：name=德阳市,id=239 | 德阳
	case 851:
		location = 2614 //原站点名：大英站,结果：name=大英县,id=2614 | 大英
	case 87:
		location = 213 //原站点名：东莞站,结果：name=东莞市,id=213 | 东莞
	case 859:
		location = 1854 //原站点名：东阿站,结果：name=东阿县,id=1854 | 东阿
	case 612:
		location = 2665 //原站点名：大竹站,结果：name=大竹县,id=2665 | 大竹
	case 362:
		location = 277 //原站点名：大理站,结果：name=大理白族自治州,id=277 | 大理
	case 874:
		location = 2535 //原站点名：垫江站,结果：name=垫江县,id=2535 | 垫江
	case 635:
		location = 2718 //原站点名：德昌站,结果：name=德昌县,id=2718 | 德昌
	case 381:
		location = 15 //原站点名：大同站,结果：name=大同市,id=15 | 大同
	case 640:
		location = 1001 //原站点名：大石桥站,结果：name=大石桥市,id=1001 | 大石桥
	case 641:
		location = 2494 //原站点名：儋州站,结果：name=儋州市,id=2494 | 儋州
	case 650:
		location = 2529 //原站点名：大足站,结果：name=大足县,id=2529 | 大足
	case 661:
		location = 1018 //原站点名：大洼站,结果：name=大洼区,id=1018 | 大洼
	case 664:
		location = 704 //原站点名：大城站,结果：name=大城县,id=704 | 大城
	case 671:
		location = 3105 //原站点名：定边站,结果：name=定边县,id=3105 | 定边
	case 418:
		location = 72 //原站点名：大兴站,结果：name=大兴安岭地区,id=72 | 大兴
	case 424:
		location = 1415 //原站点名：东阳站,结果：name=东阳市,id=1415 | 东阳
	case 429:
		location = 1329 //原站点名：东台站,结果：name=东台市,id=1329 | 东台
	case 449:
		location = 2808 //原站点名：都匀站,结果：name=都匀市,id=2808 | 都匀
	case 198:
		location = 42 //原站点名：丹东站,结果：name=丹东市,id=42 | 丹东
	case 199:
		location = 2566 //原站点名：都江堰站,结果：name=都江堰市,id=2566 | 都江堰
	case 456:
		location = 1311 //原站点名：东海站,结果：name=东海县,id=1311 | 东海
	case 458:
		location = 1733 //原站点名：德兴站,结果：name=德兴市,id=1733 | 德兴
	case 728:
		location = 1873 //原站点名：东明站,结果：name=东明县,id=1873 | 东明
	case 221:
		location = 1341 //原站点名：丹阳站,结果：name=丹阳市,id=1341 | 丹阳
	case 482:
		location = 1399 //原站点名：德清站,结果：name=德清县,id=1399 | 德清
	case 738:
		location = 2072 //原站点名：当阳站,结果：name=当阳市,id=2072 | 当阳
	case 749:
		location = 2497 //原站点名：东方站,结果：name=东方市,id=2497 | 东方
	case 760:
		location = 1867 //原站点名：单县站,结果：name=单县,id=1867 | 单县
	case 545:
		location = 181 //原站点名：恩施站,结果：name=恩施土家族苗族自治州,id=181 | 恩施
	case 818:
		location = 30 //原站点名：鄂尔多斯站,结果：name=鄂尔多斯市,id=30 | 鄂尔多斯
	case 430:
		location = 152 //原站点名：郑州2站,结果：name=郑州市,id=152 | 郑州
	case 777:
		location = 1507 //原站点名：凤阳站,结果：name=凤阳县,id=1507 | 凤阳
	case 267:
		location = 2299 //原站点名：三水站,结果：name=三水区,id=2299 | 三水
	case 526:
		location = 2339 //原站点名：丰顺站,结果：name=丰顺县,id=2339 | 丰顺
	case 277:
		location = 576 //原站点名：丰润站,结果：name=丰润区,id=576 | 丰润
	case 315:
		location = 115 //原站点名：福州2站,结果：name=福州市,id=115 | 福州
	case 582:
		location = 1451 //原站点名：肥西站,结果：name=肥西县,id=1451 | 肥西
	case 838:
		location = 838 //原站点名：汾阳站,结果：name=汾阳市,id=838 | 汾阳
	case 594:
		location = 595 //原站点名：峰峰站,结果：name=峰峰矿区,id=595 | 峰峰
	case 83:
		location = 115 //原站点名：福州站,结果：name=福州市,id=115 | 福州
	case 86:
		location = 202 //原站点名：佛山站,结果：name=佛山市,id=202 | 佛山
	case 602:
		location = 1275 //原站点名：丰县站,结果：name=丰县,id=1275 | 丰县
	case 91:
		location = 2512 //原站点名：涪陵站,结果：name=涪陵区,id=2512 | 涪陵
	case 350:
		location = 2512 //原站点名：涪陵新站,结果：name=涪陵区,id=2512 | 涪陵
	case 361:
		location = 2512 //原站点名：涪陵总站,结果：name=涪陵区,id=2512 | 涪陵
	case 877:
		location = 521 //原站点名：房山站,结果：name=房山区,id=521 | 房山
	case 652:
		location = 680 //原站点名：丰宁站,结果：name=丰宁满族自治县,id=680 | 丰宁
	case 405:
		location = 1366 //原站点名：富阳站,结果：name=富阳市,id=1366 | 富阳
	case 417:
		location = 40 //原站点名：抚顺站,结果：name=抚顺市,id=40 | 抚顺
	case 674:
		location = 1634 //原站点名：福鼎站,结果：name=福鼎市,id=1634 | 福鼎
	case 675:
		location = 1633 //原站点名：福安站,结果：name=福安市,id=1633 | 福安
	case 167:
		location = 108 //原站点名：阜阳站,结果：name=阜阳市,id=108 | 阜阳
	case 445:
		location = 133 //原站点名：抚州站,结果：name=抚州市,id=133 | 抚州
	case 467:
		location = 2359 //原站点名：佛冈站,结果：name=佛冈县,id=2359 | 佛冈
	case 725:
		location = 2575 //原站点名：富顺站,结果：name=富顺县,id=2575 | 富顺
	case 489:
		location = 4620 //原站点名：福建站,结果：name=福建省,id=4620 | 福建
	case 236:
		location = 1 //原站点名：北京5站,结果：name=北京市,id=1 | 北京
	case 751:
		location = 3102 //原站点名：府谷站,结果：name=府谷县,id=3102 | 府谷
	case 6:
		location = 256 //原站点名：贵阳站,结果：name=贵阳市,id=256 | 贵阳
	case 11:
		location = 248 //原站点名：广安站,结果：name=广安市,id=248 | 广安
	case 529:
		location = 2320 //原站点名：高州站,结果：name=高州市,id=2320 | 高州
	case 533:
		location = 1310 //原站点名：赣榆站,结果：name=赣榆县,id=1310 | 赣榆
	case 556:
		location = 2100 //原站点名：公安站,结果：name=公安县,id=2100 | 公安
	case 10036:
		location = 0 //原站点名：格外享,结果：record not found | 格外享
	case 10037:
		location = 0 //原站点名：格外享站,结果：record not found | 格外享
	case 313:
		location = 2300 //原站点名：高明站,结果：name=高明区,id=2300 | 高明
	case 573:
		location = 1312 //原站点名：灌云站,结果：name=灌云县,id=1312 | 灌云
	case 833:
		location = 3034 //原站点名：高陵站,结果：name=高陵县,id=3034 | 高陵
	case 66:
		location = 197 //原站点名：广州站,结果：name=广州市,id=197 | 广州
	case 329:
		location = 220 //原站点名：桂林总站,结果：name=桂林市,id=220 | 桂林
	case 340:
		location = 248 //原站点名：广安总站,结果：name=广安市,id=248 | 广安
	case 862:
		location = 701 //原站点名：固安站,结果：name=固安县,id=701 | 固安
	case 114:
		location = 220 //原站点名：桂林站,结果：name=桂林市,id=220 | 桂林
	case 387:
		location = 256 //原站点名：贵阳总站,结果：name=贵阳市,id=256 | 贵阳
	case 142:
		location = 130 //原站点名：赣州站,结果：name=赣州市,id=130 | 赣州
	case 404:
		location = 2901 //原站点名：个旧站,结果：name=个旧市,id=2901 | 个
	case 150:
		location = 323 //原站点名：固原站,结果：name=固原市,id=323 | 固原
	case 153:
		location = 2591 //原站点名：广汉站,结果：name=广汉市,id=2591 | 广汉
	case 411:
		location = 225 //原站点名：贵港站,结果：name=贵港市,id=225 | 贵港
	case 677:
		location = 1628 //原站点名：古田站,结果：name=古田县,id=1628 | 古田
	case 680:
		location = 567 //原站点名：藁城站,结果：name=藁城市,id=567 | 藁城
	case 692:
		location = 2325 //原站点名：广宁站,结果：name=广宁县,id=2325 | 广宁
	case 188:
		location = 1336 //原站点名：高邮站,结果：name=高邮市,id=1336 | 高邮
	case 705:
		location = 662 //原站点名：沽源站,结果：name=沽源县,id=662 | 沽源
	case 710:
		location = 606 //原站点名：广平站,结果：name=广平县,id=606 | 广平
	case 228:
		location = 241 //原站点名：广元站,结果：name=广元市,id=241 | 广元
	case 490:
		location = 4626 //原站点名：广东站,结果：name=广东省,id=4626 | 广东
	case 491:
		location = 4627 //原站点名：广西站,结果：name=广西壮族自治区,id=4627 | 广西
	case 256:
		location = 234 //原站点名：重庆8站,结果：name=重庆市,id=234 | 重庆
	case 771:
		location = 313 //原站点名：海东站,结果：name=海东地区,id=313 | 海东
	case 13:
		location = 87 //原站点名：杭州站,结果：name=杭州市,id=87 | 杭州
	case 271:
		location = 1306 //原站点名：海门站,结果：name=海门市,id=1306 | 海门
	case 783:
		location = 2428 //原站点名：合浦站,结果：name=合浦县,id=2428 | 合浦
	case 26:
		location = 294 //原站点名：汉中站,结果：name=汉中市,id=294 | 汉中
	case 28:
		location = 0 //原站点名：鄠邑站,结果：record not found | 鄠邑
	case 284:
		location = 0 //原站点名：虎门站,结果：record not found | 虎门
	case 32:
		location = 98 //原站点名：合肥站,结果：name=合肥市,id=98 | 合肥
	case 289:
		location = 2333 //原站点名：博罗站,结果：name=博罗县,id=2333 | 博罗
	case 553:
		location = 335 //原站点名：和田站,结果：name=和田地区,id=335 | 和田
	case 822:
		location = 2372 //原站点名：惠来站,结果：name=惠来县,id=2372 | 惠来
	case 567:
		location = 2643 //原站点名：洪雅站,结果：name=洪雅县,id=2643 | 洪雅
	case 56:
		location = 232 //原站点名：海口站,结果：name=海口市,id=232 | 海口
	case 825:
		location = 1662 //原站点名：湖口站,结果：name=湖口县,id=1662 | 湖口
	case 828:
		location = 2334 //原站点名：惠东站,结果：name=惠东县,id=2334 | 惠东
	case 578:
		location = 2195 //原站点名：汉寿站,结果：name=汉寿县,id=2195 | 汉寿
	case 70:
		location = 60 //原站点名：哈尔滨站,结果：name=哈尔滨市,id=60 | 哈尔滨
	case 326:
		location = 87 //原站点名：杭州总站,结果：name=杭州市,id=87 | 杭州
	case 842:
		location = 1946 //原站点名：辉县站,结果：name=辉县市,id=1946 | 辉县
	case 75:
		location = 25 //原站点名：呼和浩特站,结果：name=呼和浩特市,id=25 | 呼和浩特
	case 844:
		location = 2345 //原站点名：海丰站,结果：name=海丰县,id=2345 | 海丰
	case 337:
		location = 1301 //原站点名：海安站,结果：name=海安县,id=1301 | 海安
	case 341:
		location = 60 //原站点名：哈尔滨总站,结果：name=哈尔滨市,id=60 | 哈尔滨
	case 597:
		location = 328 //原站点名：哈密站,结果：name=哈密地区,id=328 | 哈密
	case 344:
		location = 98 //原站点名：合肥总站,结果：name=合肥市,id=98 | 合肥
	case 856:
		location = 1761 //原站点名：桓台站,结果：name=桓台县,id=1761 | 桓台
	case 349:
		location = 91 //原站点名：湖州总站,结果：name=湖州市,id=91 | 湖州
	case 863:
		location = 1526 //原站点名：含山站,结果：name=含山县,id=1526 | 含山
	case 352:
		location = 232 //原站点名：海口总站,结果：name=海口市,id=232 | 海口
	case 355:
		location = 13 //原站点名：衡水站,结果：name=衡水市,id=13 | 衡水
	case 867:
		location = 2585 //原站点名：合江站,结果：name=合江县,id=2585 | 合江
	case 103:
		location = 50 //原站点名：葫芦岛站,结果：name=葫芦岛市,id=50 | 葫芦岛
	case 871:
		location = 70 //原站点名：黑河站,结果：name=黑河市,id=70 | 黑河
	case 872:
		location = 1531 //原站点名：霍邱站,结果：name=霍邱县,id=1531 | 霍邱
	case 110:
		location = 207 //原站点名：惠州站,结果：name=惠州市,id=207 | 惠州
	case 622:
		location = 1683 //原站点名：会昌站,结果：name=会昌县,id=1683 | 会昌
	case 878:
		location = 2321 //原站点名：化州站,结果：name=化州市,id=2321 | 化州
	case 112:
		location = 6 //原站点名：邯郸站,结果：name=邯郸市,id=6 | 邯郸
	case 880:
		location = 0 //原站点名：淮口站,结果：record not found | 淮口
	case 630:
		location = 2719 //原站点名：会理站,结果：name=会理县,id=2719 | 会理
	case 377:
		location = 103 //原站点名：淮北站,结果：name=淮北市,id=103 | 淮北
	case 123:
		location = 91 //原站点名：湖州站,结果：name=湖州市,id=91 | 湖州
	case 383:
		location = 1930 //原站点名：鹤山站,结果：name=鹤山区,id=1930 | 鹤山
	case 389:
		location = 101 //原站点名：淮南站,结果：name=淮南市,id=101 | 淮南
	case 390:
		location = 1306 //原站点名：海门总站,结果：name=海门市,id=1306 | 海门
	case 136:
		location = 186 //原站点名：衡阳站,结果：name=衡阳市,id=186 | 衡阳
	case 394:
		location = 1590 //原站点名：惠安站,结果：name=惠安县,id=1590 | 惠安
	case 139:
		location = 81 //原站点名：淮安站,结果：name=淮安市,id=81 | 淮安
	case 653:
		location = 2720 //原站点名：会东站,结果：name=会东县,id=2720 | 会东
	case 408:
		location = 194 //原站点名：怀化站,结果：name=怀化市,id=194 | 怀化
	case 665:
		location = 157 //原站点名：鹤壁站,结果：name=鹤壁市,id=157 | 鹤壁
	case 410:
		location = 170 //原站点名：黄石站,结果：name=黄石市,id=170 | 黄石
	case 666:
		location = 697 //原站点名：黄骅站,结果：name=黄骅市,id=697 | 黄骅
	case 413:
		location = 210 //原站点名：河源站,结果：name=河源市,id=210 | 河源
	case 158:
		location = 274 //原站点名：红河站,结果：name=红河哈尼族彝族自治州,id=274 | 红河
	case 166:
		location = 2265 //原站点名：黄埔站,结果：name=黄埔区,id=2265 | 黄埔
	case 688:
		location = 813 //原站点名：洪洞站,结果：name=洪洞县,id=813 | 洪洞
	case 690:
		location = 2097 //原站点名：汉川站,结果：name=汉川市,id=2097 | 汉川
	case 691:
		location = 824 //原站点名：侯马站,结果：name=侯马市,id=824 | 侯马
	case 436:
		location = 2326 //原站点名：怀集站,结果：name=怀集县,id=2326 | 怀集
	case 696:
		location = 668 //原站点名：怀来站,结果：name=怀来县,id=668 | 怀来
	case 698:
		location = 698 //原站点名：河间站,结果：name=河间市,id=698 | 河间
	case 444:
		location = 178 //原站点名：黄冈站,结果：name=黄冈市,id=178 | 黄冈
	case 192:
		location = 0 //原站点名：菏泽站,结果：record not found | 菏泽
	case 193:
		location = 87 //原站点名：杭州精选,结果：name=杭州市,id=87 | 杭州
	case 204:
		location = 2267 //原站点名：花都站,结果：name=花都区,id=2267 | 花都
	case 463:
		location = 2548 //原站点名：合川站,结果：name=合川市,id=2548 | 合川
	case 209:
		location = 2037 //原站点名：汉阳站,结果：name=汉阳区,id=2037 | 汉阳
	case 721:
		location = 770 //原站点名：怀仁站,结果：name=怀仁县,id=770 | 怀仁
	case 466:
		location = 1786 //原站点名：海阳站,结果：name=海阳市,id=1786 | 海阳
	case 477:
		location = 31 //原站点名：呼伦贝尔站,结果：name=呼伦贝尔市,id=31 | 呼伦贝尔
	case 222:
		location = 2332 //原站点名：惠阳站,结果：name=惠阳区,id=2332 | 惠阳
	case 234:
		location = 1394 //原站点名：海宁站,结果：name=海宁市,id=1394 | 海宁
	case 746:
		location = 825 //原站点名：霍州站,结果：name=霍州市,id=825 | 霍州
	case 492:
		location = 316 //原站点名：海南站,结果：name=海南藏族自治州,id=316 | 海南
	case 237:
		location = 87 //原站点名：杭州优选,结果：name=杭州市,id=87 | 杭州
	case 493:
		location = 534 //原站点名：河北站,结果：name=河北区,id=534 | 河北
	case 494:
		location = 3235 //原站点名：河南总站,结果：name=河南蒙古族自治县,id=3235 | 河南
	case 495:
		location = 1393 //原站点名：海盐站,结果：name=海盐县,id=1393 | 海盐
	case 496:
		location = 4625 //原站点名：湖南站,结果：name=湖南省,id=4625 | 湖南
	case 752:
		location = 794 //原站点名：河津站,结果：name=河津市,id=794 | 河津
	case 497:
		location = 4624 //原站点名：湖北站,结果：name=湖北省,id=4624 | 湖北
	case 498:
		location = 4615 //原站点名：黑龙江站,结果：name=黑龙江省,id=4615 | 黑龙江
	case 265:
		location = 1287 //原站点名：金坛站,结果：name=金坛市,id=1287 | 金坛
	case 266:
		location = 1347 //原站点名：靖江站,结果：name=靖江市,id=1347 | 靖江
	case 273:
		location = 43 //原站点名：锦州站,结果：name=锦州市,id=43 | 锦州
	case 276:
		location = 1596 //原站点名：晋江站,结果：name=晋江市,id=1596 | 晋江
	case 790:
		location = 300 //原站点名：金昌站,结果：name=金昌市,id=300 | 金昌
	case 796:
		location = 0 //原站点名：联联酒景站,结果：record not found | 联联酒景
	case 42:
		location = 135 //原站点名：济南站,结果：name=济南市,id=135 | 济南
	case 812:
		location = 306 //原站点名：酒泉站,结果：name=酒泉市,id=306 | 酒泉
	case 813:
		location = 299 //原站点名：嘉峪关站,结果：name=嘉峪关市,id=299 | 嘉峪关
	case 816:
		location = 781 //原站点名：介休站,结果：name=介休市,id=781 | 介休
	case 562:
		location = 1321 //原站点名：金湖站,结果：name=金湖县,id=1321 | 金湖
	case 55:
		location = 2683 //原站点名：简阳站,结果：name=简阳市,id=2683 | 简阳
	case 312:
		location = 2547 //原站点名：江津站,结果：name=江津市,id=2547 | 江津
	case 569:
		location = 2626 //原站点名：夹江站,结果：name=夹江县,id=2626 | 夹江
	case 332:
		location = 1392 //原站点名：嘉善站,结果：name=嘉善县,id=1392 | 嘉善
	case 333:
		location = 2602 //原站点名：江油站,结果：name=江油市,id=2602 | 江油
	case 592:
		location = 0 //原站点名：蓟州站,结果：record not found | 蓟州
	case 81:
		location = 90 //原站点名：嘉兴站,结果：name=嘉兴市,id=90 | 嘉兴
	case 854:
		location = 1643 //原站点名：进贤站,结果：name=进贤县,id=1643 | 进贤
	case 90:
		location = 127 //原站点名：九江站,结果：name=九江市,id=127 | 九江
	case 93:
		location = 177 //原站点名：荆州站,结果：name=荆州市,id=177 | 荆州
	case 615:
		location = 1422 //原站点名：江山站,结果：name=江山市,id=1422 | 江山
	case 617:
		location = 67 //原站点名：佳木斯站,结果：name=佳木斯市,id=67 | 佳木斯
	case 109:
		location = 203 //原站点名：江门站,结果：name=江门市,id=203 | 江门
	case 117:
		location = 93 //原站点名：金华站,结果：name=金华市,id=93 | 金华
	case 376:
		location = 20 //原站点名：晋中站,结果：name=晋中市,id=20 | 晋中
	case 125:
		location = 1268 //原站点名：江阴站,结果：name=江阴市,id=1268 | 江阴
	case 654:
		location = 2625 //原站点名：井研站,结果：name=井研县,id=2625 | 井研
	case 662:
		location = 546 //原站点名：静海站,结果：name=静海区,id=546 | 静海
	case 667:
		location = 2117 //原站点名：嘉鱼站,结果：name=嘉鱼县,id=2117 | 嘉鱼
	case 414:
		location = 131 //原站点名：吉安站,结果：name=吉安市,id=131 | 吉安
	case 670:
		location = 3104 //原站点名：靖边站,结果：name=靖边县,id=3104 | 靖边
	case 672:
		location = 1751 //原站点名：胶州站,结果：name=胶州市,id=1751 | 胶州
	case 683:
		location = 62 //原站点名：鸡西站,结果：name=鸡西市,id=62 | 鸡西
	case 173:
		location = 1241 //原站点名：嘉定站,结果：name=嘉定区,id=1241 | 嘉定
	case 178:
		location = 142 //原站点名：济宁站,结果：name=济宁市,id=142 | 济宁
	case 181:
		location = 1337 //原站点名：江都站,结果：name=江都市,id=1337 | 江都
	case 184:
		location = 2560 //原站点名：金堂站,结果：name=金堂县,id=2560 | 金堂
	case 442:
		location = 18 //原站点名：晋城站,结果：name=晋城市,id=18 | 晋城
	case 187:
		location = 1258 //原站点名：江宁站,结果：name=江宁区,id=1258 | 江宁
	case 460:
		location = 2371 //原站点名：揭西站,结果：name=揭西县,id=2371 | 揭西
	case 206:
		location = 1343 //原站点名：句容站,结果：name=句容市,id=1343 | 句容
	case 471:
		location = 135 //原站点名：济南特惠站,结果：name=济南市,id=135 | 济南
	case 217:
		location = 52 //原站点名：吉林站,结果：name=吉林市,id=52 | 吉林
	case 730:
		location = 1438 //原站点名：缙云站,结果：name=缙云县,id=1438 | 缙云
	case 734:
		location = 1547 //原站点名：绩溪站,结果：name=绩溪县,id=1547 | 绩溪
	case 486:
		location = 4617 //原站点名：江苏站,结果：name=江苏省,id=4617 | 江苏
	case 748:
		location = 1836 //原站点名：莒南站,结果：name=莒南县,id=1836 | 莒南
	case 501:
		location = 4621 //原站点名：江西站,结果：name=江西省,id=4621 | 江西
	case 246:
		location = 2044 //原站点名：江夏站,结果：name=江夏区,id=2044 | 江夏
	case 250:
		location = 216 //原站点名：揭阳站,结果：name=揭阳市,id=216 | 揭阳
	case 518:
		location = 326 //原站点名：克拉玛依站,结果：name=克拉玛依市,id=326 | 克拉玛依
	case 10:
		location = 265 //原站点名：昆明站,结果：name=昆明市,id=265 | 昆明
	case 800:
		location = 2902 //原站点名：开远站,结果：name=开远市,id=2902 | 开远
	case 810:
		location = 681 //原站点名：宽城站,结果：name=宽城满族自治县,id=681 | 宽城
	case 334:
		location = 2792 //原站点名：凯里站,结果：name=凯里市,id=2792 | 凯里
	case 124:
		location = 1296 //原站点名：昆山站,结果：name=昆山市,id=1296 | 昆山
	case 644:
		location = 334 //原站点名：喀什站,结果：name=喀什地区,id=334 | 喀什
	case 143:
		location = 153 //原站点名：开封站,结果：name=开封市,id=153 | 开封
	case 183:
		location = 3308 //原站点名：库尔勒站,结果：name=库尔勒市,id=3308 | 库尔勒
	case 707:
		location = 3351 //原站点名：奎屯站,结果：name=奎屯市,id=3351 | 奎屯
	case 709:
		location = 574 //原站点名：开平站,结果：name=开平区,id=574 | 开平
	case 742:
		location = 1420 //原站点名：开化站,结果：name=开化县,id=1420 | 开化
	case 513:
		location = 2196 //原站点名：澧县站,结果：name=澧县,id=2196 | 澧县
	case 514:
		location = 24 //原站点名：吕梁站,结果：name=吕梁市,id=24 | 吕梁
	case 7:
		location = 298 //原站点名：兰州站,结果：name=兰州市,id=298 | 兰州
	case 522:
		location = 1544 //原站点名：郎溪站,结果：name=郎溪县,id=1544 | 郎溪
	case 12:
		location = 3185 //原站点名：陇西站,结果：name=陇西县,id=3185 | 陇西
	case 781:
		location = 1791 //原站点名：临朐站,结果：name=临朐县,id=1791 | 临朐
	case 789:
		location = 2177 //原站点名：隆回站,结果：name=隆回县,id=2177 | 隆回
	case 795:
		location = 2802 //原站点名：黎平站,结果：name=黎平县,id=2802 | 黎平
	case 540:
		location = 0 //原站点名：优品站,结果：record not found | 优品
	case 30:
		location = 2556 //原站点名：龙泉驿站,结果：name=龙泉驿区,id=2556 | 龙泉驿
	case 288:
		location = 122 //原站点名：龙岩站,结果：name=龙岩市,id=122 | 龙岩
	case 803:
		location = 54 //原站点名：辽源站,结果：name=辽源市,id=54 | 辽源
	case 804:
		location = 1524 //原站点名：庐江站,结果：name=庐江县,id=1524 | 庐江
	case 43:
		location = 23 //原站点名：临汾站,结果：name=临汾市,id=23 | 临汾
	case 300:
		location = 1389 //原站点名：乐清站,结果：name=乐清市,id=1389 | 乐清
	case 10029:
		location = 0 //原站点名：联联旅游,结果：record not found | 联联旅游
	case 558:
		location = 1035 //原站点名：龙港站,结果：name=龙港区,id=1035 | 龙港
	case 814:
		location = 780 //原站点名：灵石站,结果：name=灵石县,id=780 | 灵石
	case 47:
		location = 2532 //原站点名：梁平站,结果：name=梁平县,id=2532 | 梁平
	case 10031:
		location = 0 //原站点名：联联贝塔站,结果：record not found | 联联贝塔
	case 10032:
		location = 0 //原站点名：联联拼团,结果：record not found | 联联拼团
	case 305:
		location = 281 //原站点名：拉萨站,结果：name=拉萨市,id=281 | 拉萨
	case 10033:
		location = 0 //原站点名：联联拼团小程序,结果：record not found | 联联拼团小程序
	case 10034:
		location = 0 //原站点名：联联周边游APP,结果：record not found | 联联周边游APP
	case 10035:
		location = 0 //原站点名：联联周边游小程序,结果：record not found | 联联周边游小程序
	case 820:
		location = 1444 //原站点名：龙泉站,结果：name=龙泉市,id=1444 | 龙泉
	case 53:
		location = 195 //原站点名：娄底站,结果：name=娄底市,id=195 | 娄底
	case 568:
		location = 2022 //原站点名：鹿邑站,结果：name=鹿邑县,id=2022 | 鹿邑
	case 570:
		location = 146 //原站点名：莱芜站,结果：name=莱芜市,id=146 | 莱芜
	case 572:
		location = 1773 //原站点名：利津站,结果：name=利津县,id=1773 | 利津
	case 317:
		location = 140 //原站点名：烟台2站,结果：name=烟台市,id=140 | 烟台
	case 831:
		location = 1917 //原站点名：鲁山站,结果：name=鲁山县,id=1917 | 鲁山
	case 322:
		location = 74 //原站点名：南京5站,结果：name=南京市,id=74 | 南京
	case 835:
		location = 642 //原站点名：涞源站,结果：name=涞源县,id=642 | 涞源
	case 839:
		location = 827 //原站点名：文水站,结果：name=文水县,id=827 | 文水
	case 588:
		location = 783 //原站点名：临猗站,结果：name=临猗县,id=783 | 临猗
	case 589:
		location = 1318 //原站点名：涟水站,结果：name=涟水县,id=1318 | 涟水
	case 847:
		location = 310 //原站点名：临夏站,结果：name=临夏回族自治州,id=310 | 临夏
	case 596:
		location = 3031 //原站点名：蓝田站,结果：name=蓝田县,id=3031 | 蓝田
	case 598:
		location = 162 //原站点名：漯河站,结果：name=漯河市,id=162 | 漯河
	case 855:
		location = 0 //原站点名：荔浦站,结果：record not found | 荔浦
	case 346:
		location = 23 //原站点名：临汾总站,结果：name=临汾市,id=23 | 临汾
	case 610:
		location = 2659 //原站点名：邻水站,结果：name=邻水县,id=2659 | 邻水
	case 866:
		location = 2314 //原站点名：廉江站,结果：name=廉江市,id=2314 | 廉江
	case 100:
		location = 80 //原站点名：连云港站,结果：name=连云港市,id=80 | 连云港
	case 616:
		location = 2619 //原站点名：隆昌站,结果：name=隆昌县,id=2619 | 隆昌
	case 106:
		location = 244 //原站点名：乐山站,结果：name=乐山市,id=244 | 乐山
	case 618:
		location = 1435 //原站点名：临海站,结果：name=临海市,id=1435 | 临海
	case 876:
		location = 2378 //原站点名：罗定站,结果：name=罗定市,id=2378 | 罗定
	case 629:
		location = 2278 //原站点名：乐昌站,结果：name=乐昌市,id=2278 | 乐昌
	case 119:
		location = 270 //原站点名：丽江站,结果：name=丽江市,id=270 | 丽江
	case 634:
		location = 2900 //原站点名：禄丰站,结果：name=禄丰县,id=2900 | 禄丰
	case 385:
		location = 3185 //原站点名：陇西总站,结果：name=陇西县,id=3185 | 陇西
	case 130:
		location = 238 //原站点名：泸州站,结果：name=泸州市,id=238 | 泸州
	case 131:
		location = 219 //原站点名：柳州站,结果：name=柳州市,id=219 | 柳州
	case 132:
		location = 147 //原站点名：临沂站,结果：name=临沂市,id=147 | 临沂
	case 388:
		location = 46 //原站点名：辽阳站,结果：name=辽阳市,id=46 | 辽阳
	case 133:
		location = 154 //原站点名：洛阳站,结果：name=洛阳市,id=154 | 洛阳
	case 651:
		location = 679 //原站点名：隆化站,结果：name=隆化县,id=679 | 隆化
	case 144:
		location = 12 //原站点名：廊坊站,结果：name=廊坊市,id=12 | 廊坊
	case 402:
		location = 1981 //原站点名：灵宝站,结果：name=灵宝市,id=1981 | 灵宝
	case 154:
		location = 257 //原站点名：六盘水站,结果：name=六盘水市,id=257 | 六盘水
	case 419:
		location = 97 //原站点名：丽水站,结果：name=丽水市,id=97 | 丽水
	case 165:
		location = 2284 //原站点名：龙岗站,结果：name=龙岗区,id=2284 | 龙岗
	case 681:
		location = 1608 //原站点名：龙海站,结果：name=龙海市,id=1608 | 龙海
	case 179:
		location = 111 //原站点名：六安站,结果：name=六安市,id=111 | 六安
	case 438:
		location = 0 //原站点名：滦州站,结果：record not found | 滦州
	case 697:
		location = 2505 //原站点名：陵水站,结果：name=陵水黎族自治县,id=2505 | 陵水
	case 700:
		location = 287 //原站点名：林芝站,结果：name=林芝地区,id=287 | 林芝
	case 446:
		location = 1413 //原站点名：兰溪站,结果：name=兰溪市,id=1413 | 兰溪
	case 706:
		location = 0 //原站点名：联联严选站,结果：record not found | 联联严选
	case 718:
		location = 2639 //原站点名：阆中站,结果：name=阆中市,id=2639 | 阆中
	case 726:
		location = 2682 //原站点名：乐至站,结果：name=乐至县,id=2682 | 乐至
	case 219:
		location = 1286 //原站点名：溧阳站,结果：name=溧阳市,id=1286 | 溧阳
	case 737:
		location = 2335 //原站点名：龙门站,结果：name=龙门县,id=2335 | 龙门
	case 747:
		location = 309 //原站点名：陇南站,结果：name=陇南市,id=309 | 陇南
	case 239:
		location = 149 //原站点名：聊城站,结果：name=聊城市,id=149 | 聊城
	case 499:
		location = 4613 //原站点名：辽宁站,结果：name=辽宁省,id=4613 | 辽宁
	case 758:
		location = 2408 //原站点名：灵川站,结果：name=灵川县,id=2408 | 灵川
	case 759:
		location = 831 //原站点名：柳林站,结果：name=柳林县,id=831 | 柳林
	case 249:
		location = 2489 //原站点名：龙华站,结果：name=龙华区,id=2489 | 龙华
	case 509:
		location = 1367 //原站点名：临安站,结果：name=临安市,id=1367 | 临安
	case 765:
		location = 597 //原站点名：临漳站,结果：name=临漳县,id=597 | 临漳
	case 767:
		location = 2413 //原站点名：龙胜站,结果：name=龙胜各族自治县,id=2413 | 龙胜
	case 773:
		location = 2907 //原站点名：弥勒站,结果：name=弥勒县,id=2907 | 弥勒
	case 9:
		location = 240 //原站点名：绵阳站,结果：name=绵阳市,id=240 | 绵阳
	case 19:
		location = 246 //原站点名：眉山站,结果：name=眉山市,id=246 | 眉山
	case 564:
		location = 2593 //原站点名：绵竹站,结果：name=绵竹市,id=2593 | 绵竹
	case 834:
		location = 1997 //原站点名：民权站,结果：name=民权县,id=1997 | 民权
	case 331:
		location = 240 //原站点名：绵阳总站,结果：name=绵阳市,id=240 | 绵阳
	case 619:
		location = 2190 //原站点名：汨罗站,结果：name=汨罗市,id=2190 | 汨罗
	case 378:
		location = 208 //原站点名：梅州站,结果：name=梅州市,id=208 | 梅州
	case 639:
		location = 2727 //原站点名：冕宁站,结果：name=冕宁县,id=2727 | 冕宁
	case 134:
		location = 205 //原站点名：茂名站,结果：name=茂名市,id=205 | 茂名
	case 655:
		location = 2629 //原站点名：马边站,结果：name=马边彝族自治县,id=2629 | 马边
	case 171:
		location = 1239 //原站点名：闵行站,结果：name=闵行区,id=1239 | 闵行
	case 457:
		location = 69 //原站点名：牡丹江站,结果：name=牡丹江市,id=69 | 牡丹江
	case 208:
		location = 102 //原站点名：马鞍山站,结果：name=马鞍山市,id=102 | 马鞍山
	case 262:
		location = 197 //原站点名：广州6站,结果：name=广州市,id=197 | 广州
	case 16:
		location = 74 //原站点名：南京站,结果：name=南京市,id=74 | 南京
	case 25:
		location = 88 //原站点名：宁波站,结果：name=宁波市,id=88 | 宁波
	case 805:
		location = 628 //原站点名：南宫站,结果：name=南宫市,id=628 | 南宫
	case 38:
		location = 124 //原站点名：南昌站,结果：name=南昌市,id=124 | 南昌
	case 298:
		location = 1597 //原站点名：南安站,结果：name=南安市,id=1597 | 南安
	case 557:
		location = 2550 //原站点名：南川站,结果：name=南川市,id=2550 | 南川
	case 819:
		location = 1999 //原站点名：宁陵站,结果：name=宁陵县,id=1999 | 宁陵
	case 824:
		location = 2634 //原站点名：南部县站,结果：name=南部县,id=2634 | 南部县
	case 65:
		location = 218 //原站点名：南宁站,结果：name=南宁市,id=218 | 南宁
	case 79:
		location = 79 //原站点名：南通站,结果：name=南通市,id=79 | 南通
	case 849:
		location = 1046 //原站点名：农安站,结果：name=农安县,id=1046 | 农安
	case 599:
		location = 1960 //原站点名：南乐站,结果：name=南乐县,id=1960 | 南乐
	case 606:
		location = 2721 //原站点名：宁南站,结果：name=宁南县,id=2721 | 宁南
	case 873:
		location = 1134 //原站点名：讷河站,结果：name=讷河市,id=1134 | 讷河
	case 107:
		location = 245 //原站点名：南充站,结果：name=南充市,id=245 | 南充
	case 875:
		location = 1988 //原站点名：内乡站,结果：name=内乡县,id=1988 | 内乡
	case 879:
		location = 2509 //原站点名：南沙站,结果：name=南沙群岛,id=2509 | 南沙
	case 642:
		location = 2279 //原站点名：南雄站,结果：name=南雄市,id=2279 | 南雄
	case 135:
		location = 164 //原站点名：南阳站,结果：name=南阳市,id=164 | 南阳
	case 649:
		location = 2143 //原站点名：宁乡站,结果：name=宁乡县,id=2143 | 宁乡
	case 149:
		location = 243 //原站点名：内江站,结果：name=内江市,id=243 | 内江
	case 202:
		location = 74 //原站点名：南京江北站,结果：name=南京市,id=74 | 南京
	case 464:
		location = 123 //原站点名：宁德站,结果：name=宁德市,id=123 | 宁德
	case 722:
		location = 2678 //原站点名：南江站,结果：name=南江县,id=2678 | 南江
	case 735:
		location = 279 //原站点名：怒江州站,结果：name=怒江傈僳族自治州,id=279 | 怒江
	case 740:
		location = 1549 //原站点名：宁国站,结果：name=宁国市,id=1549 | 宁国
	case 502:
		location = 4637 //原站点名：宁夏站,结果：name=宁夏回族自治区,id=4637 | 宁夏
	case 510:
		location = 1375 //原站点名：宁海站,结果：name=宁海县,id=1375 | 宁海
	case 258:
		location = 237 //原站点名：攀枝花站,结果：name=攀枝花市,id=237 | 攀枝花
	case 546:
		location = 1560 //原站点名：平潭站,结果：name=平潭县,id=1560 | 平潭
	case 296:
		location = 1395 //原站点名：平湖站,结果：name=平湖市,id=1395 | 平湖
	case 808:
		location = 1058 //原站点名：磐石站,结果：name=磐石市,id=1058 | 磐石
	case 559:
		location = 2451 //原站点名：平果站,结果：name=平果县,id=2451 | 平果
	case 49:
		location = 0 //原站点名：郫都站,结果：record not found | 郫都
	case 57:
		location = 155 //原站点名：平顶山站,结果：name=平顶山市,id=155 | 平顶山
	case 832:
		location = 1916 //原站点名：叶县站,结果：name=叶县,id=1916 | 叶县
	case 836:
		location = 624 //原站点名：平乡站,结果：name=平乡县,id=624 | 平乡
	case 587:
		location = 961 //原站点名：普兰站,结果：name=普兰店市,id=961 | 普兰
	case 595:
		location = 961 //原站点名：普兰店站,结果：name=普兰店市,id=961 | 普兰店
	case 357:
		location = 117 //原站点名：莆田总站,结果：name=莆田市,id=117 | 莆田
	case 360:
		location = 2567 //原站点名：彭州站,结果：name=彭州市,id=2567 | 彭州
	case 621:
		location = 2679 //原站点名：平昌站,结果：name=平昌县,id=2679 | 平昌
	case 626:
		location = 779 //原站点名：平遥站,结果：name=平遥县,id=779 | 平遥
	case 118:
		location = 2874 //原站点名：普洱站,结果：name=普洱哈尼族彝族自治县,id=2874 | 普洱
	case 127:
		location = 160 //原站点名：濮阳站,结果：name=濮阳市,id=160 | 濮阳
	case 384:
		location = 2373 //原站点名：普宁站,结果：name=普宁市,id=2373 | 普宁
	case 386:
		location = 47 //原站点名：盘锦站,结果：name=盘锦市,id=47 | 盘锦
	case 658:
		location = 1019 //原站点名：盘山站,结果：name=盘山县,id=1019 | 盘山
	case 147:
		location = 117 //原站点名：莆田站,结果：name=莆田市,id=117 | 莆田
	case 660:
		location = 527 //原站点名：平谷站,结果：name=平谷区,id=527 | 平谷
	case 200:
		location = 2266 //原站点名：番禺站,结果：name=番禺区,id=2266 | 番禺
	case 715:
		location = 1384 //原站点名：平阳站,结果：name=平阳县,id=1384 | 平阳
	case 462:
		location = 1411 //原站点名：浦江站,结果：name=浦江县,id=1411 | 浦江
	case 724:
		location = 1835 //原站点名：平邑站,结果：name=平邑县,id=1835 | 平邑
	case 750:
		location = 1740 //原站点名：平阴站,结果：name=平阴县,id=1740 | 平阴
	case 756:
		location = 1280 //原站点名：邳州站,结果：name=邳州市,id=1280 | 邳州
	case 769:
		location = 2409 //原站点名：全州站,结果：name=全州县,id=2409 | 全州
	case 774:
		location = 1220 //原站点名：青冈站,结果：name=青冈县,id=1220 | 青冈
	case 780:
		location = 1793 //原站点名：青州站,结果：name=青州市,id=1793 | 青州
	case 525:
		location = 726 //原站点名：清徐站,结果：name=清徐县,id=726 | 清徐
	case 14:
		location = 136 //原站点名：青岛站,结果：name=青岛市,id=136 | 青岛
	case 31:
		location = 119 //原站点名：泉州站,结果：name=泉州市,id=119 | 泉州
	case 297:
		location = 1303 //原站点名：启东站,结果：name=启东市,id=1303 | 启东
	case 44:
		location = 61 //原站点名：齐齐哈尔站,结果：name=齐齐哈尔市,id=61 | 齐齐哈尔
	case 59:
		location = 266 //原站点名：曲靖旧站,结果：name=曲靖市,id=266 | 曲靖
	case 583:
		location = 2526 //原站点名：綦江站,结果：name=綦江县,id=2526 | 綦江
	case 336:
		location = 2666 //原站点名：渠县站,结果：name=渠县,id=2666 | 渠县
	case 852:
		location = 2493 //原站点名：琼海站,结果：name=琼海市,id=2493 | 琼海
	case 869:
		location = 2859 //原站点名：巧家站,结果：name=巧家县,id=2859 | 巧家
	case 614:
		location = 2624 //原站点名：犍为站,结果：name=犍为县,id=2624 | 犍为
	case 395:
		location = 2133 //原站点名：潜江站,结果：name=潜江市,id=2133 | 潜江
	case 673:
		location = 609 //原站点名：曲周站,结果：name=曲周县,id=609 | 曲周
	case 426:
		location = 94 //原站点名：衢州站,结果：name=衢州市,id=94 | 衢州
	case 428:
		location = 1808 //原站点名：曲阜站,结果：name=曲阜市,id=1808 | 曲阜
	case 687:
		location = 686 //原站点名：青县站,结果：name=青县,id=686 | 青县
	case 441:
		location = 2568 //原站点名：邛崃站,结果：name=邛崃市,id=2568 | 邛崃
	case 191:
		location = 2557 //原站点名：青白江站,结果：name=青白江区,id=2557 | 青白江
	case 447:
		location = 266 //原站点名：曲靖站,结果：name=曲靖市,id=266 | 曲靖
	case 196:
		location = 5 //原站点名：秦皇岛站,结果：name=秦皇岛市,id=5 | 秦皇岛
	case 454:
		location = 224 //原站点名：钦州站,结果：name=钦州市,id=224 | 钦州
	case 211:
		location = 1256 //原站点名：栖霞站,结果：name=栖霞区,id=1256 | 栖霞
	case 470:
		location = 307 //原站点名：庆阳站,结果：name=庆阳市,id=307 | 庆阳
	case 215:
		location = 584 //原站点名：迁安站,结果：name=迁安市,id=584 | 迁安
	case 225:
		location = 212 //原站点名：清远站,结果：name=清远市,id=212 | 清远
	case 232:
		location = 1747 //原站点名：黄岛站,结果：name=黄岛区,id=1747 | 黄岛
	case 500:
		location = 4636 //原站点名：青海站,结果：name=青海省,id=4636 | 青海
	case 248:
		location = 1750 //原站点名：城阳站,结果：name=城阳区,id=1750 | 城阳
	case 764:
		location = 2741 //原站点名：清镇站,结果：name=清镇市,id=2741 | 清镇
	case 511:
		location = 1437 //原站点名：青田站,结果：name=青田县,id=1437 | 青田
	case 770:
		location = 0 //原站点名：雄安站,结果：record not found | 雄安
	case 524:
		location = 1819 //原站点名：荣成站,结果：name=荣成市,id=1819 | 荣成
	case 535:
		location = 696 //原站点名：任丘站,结果：name=任丘市,id=696 | 任丘
	case 285:
		location = 1304 //原站点名：如皋站,结果：name=如皋市,id=1304 | 如皋
	case 291:
		location = 145 //原站点名：日照站,结果：name=日照市,id=145 | 日照
	case 817:
		location = 2759 //原站点名：仁怀站,结果：name=仁怀市,id=2759 | 仁怀
	case 324:
		location = 2641 //原站点名：仁寿新站,结果：name=仁寿县,id=2641 | 仁寿
	case 581:
		location = 1664 //原站点名：瑞昌站,结果：name=瑞昌市,id=1664 | 瑞昌
	case 85:
		location = 2641 //原站点名：仁寿站,结果：name=仁寿县,id=2641 | 仁寿
	case 345:
		location = 0 //原站点名：仁寿总站 ,结果：record not found | 仁寿
	case 870:
		location = 2443 //原站点名：容县站,结果：name=容县,id=2443 | 容县
	case 372:
		location = 2937 //原站点名：瑞丽站,结果：name=瑞丽市,id=2937 | 瑞丽
	case 657:
		location = 1388 //原站点名：瑞安站,结果：name=瑞安市,id=1388 | 瑞安
	case 699:
		location = 284 //原站点名：日喀则站,结果：name=日喀则地区,id=284 | 日喀则
	case 703:
		location = 2276 //原站点名：乳源站,结果：name=乳源瑶族自治县,id=2276 | 乳源
	case 453:
		location = 1302 //原站点名：如东站,结果：name=如东县,id=1302 | 如东
	case 455:
		location = 2530 //原站点名：荣昌站,结果：name=荣昌县,id=2530 | 荣昌
	case 768:
		location = 2105 //原站点名：松滋站,结果：name=松滋市,id=2105 | 松滋
	case 521:
		location = 2596 //原站点名：三台站,结果：name=三台县,id=2596 | 三台
	case 527:
		location = 209 //原站点名：汕尾站,结果：name=汕尾市,id=209 | 汕尾
	case 784:
		location = 2829 //原站点名：石林站,结果：name=石林彝族自治县,id=2829 | 石林
	case 17:
		location = 165 //原站点名：商丘站,结果：name=商丘市,id=165 | 商丘
	case 274:
		location = 2516 //原站点名：沙坪坝站,结果：name=沙坪坝区,id=2516 | 沙坪坝
	case 21:
		location = 118 //原站点名：三明站,结果：name=三明市,id=118 | 三明
	case 279:
		location = 73 //原站点名：上海10站,结果：name=上海市,id=73 | 上海
	case 792:
		location = 2312 //原站点名：遂溪站,结果：name=遂溪县,id=2312 | 遂溪
	case 802:
		location = 2750 //原站点名：绥阳站,结果：name=绥阳县,id=2750 | 绥阳
	case 292:
		location = 199 //原站点名：深圳5站,结果：name=深圳市,id=199 | 深圳
	case 40:
		location = 3 //原站点名：石家庄站,结果：name=石家庄市,id=3 | 石家庄
	case 41:
		location = 78 //原站点名：苏州站,结果：name=苏州市,id=78 | 苏州
	case 299:
		location = 1595 //原站点名：石狮站,结果：name=石狮市,id=1595 | 石狮
	case 303:
		location = 198 //原站点名：韶关站,结果：name=韶关市,id=198 | 韶关
	case 304:
		location = 109 //原站点名：宿州站,结果：name=宿州市,id=109 | 宿州
	case 560:
		location = 1327 //原站点名：射阳站,结果：name=射阳县,id=1327 | 射阳
	case 563:
		location = 1353 //原站点名：泗阳站,结果：name=泗阳县,id=1353 | 泗阳
	case 52:
		location = 2561 //原站点名：双流站,结果：name=双流县,id=2561 | 双流
	case 58:
		location = 233 //原站点名：三亚站,结果：name=三亚市,id=233 | 三亚
	case 60:
		location = 73 //原站点名：上海站,结果：name=上海市,id=73 | 上海
	case 316:
		location = 37 //原站点名：沈阳2站,结果：name=沈阳市,id=37 | 沈阳
	case 61:
		location = 242 //原站点名：遂宁站,结果：name=遂宁市,id=242 | 遂宁
	case 318:
		location = 0 //原站点名：松山湖站,结果：record not found | 松山湖
	case 574:
		location = 339 //原站点名：石河子站,结果：name=石河子市,id=339 | 石河子
	case 321:
		location = 1352 //原站点名：沭阳站,结果：name=沭阳县,id=1352 | 沭阳
	case 577:
		location = 2613 //原站点名：射洪站,结果：name=射洪县,id=2613 | 射洪
	case 67:
		location = 199 //原站点名：深圳站,结果：name=深圳市,id=199 | 深圳
	case 68:
		location = 37 //原站点名：沈阳站,结果：name=沈阳市,id=37 | 沈阳
	case 837:
		location = 1998 //原站点名：睢县站,结果：name=睢县,id=1998 | 睢县
	case 327:
		location = 78 //原站点名：苏州总站,结果：name=苏州市,id=78 | 苏州
	case 330:
		location = 76 //原站点名：徐州总站,结果：name=徐州市,id=76 | 徐州
	case 846:
		location = 3114 //原站点名：石泉站,结果：name=石泉县,id=3114 | 石泉
	case 82:
		location = 92 //原站点名：绍兴站,结果：name=绍兴市,id=92 | 绍兴
	case 604:
		location = 3101 //原站点名：神木站,结果：name=神木县,id=3101 | 神木
	case 864:
		location = 2204 //原站点名：桑植站,结果：name=桑植县,id=2204 | 桑植
	case 97:
		location = 163 //原站点名：三门峡站,结果：name=三门峡市,id=163 | 三门峡
	case 356:
		location = 134 //原站点名：上饶总站,结果：name=上饶市,id=134 | 上饶
	case 102:
		location = 523 //原站点名：顺义站,结果：name=顺义区,id=523 | 顺义
	case 366:
		location = 2330 //原站点名：四会站,结果：name=四会市,id=2330 | 四会
	case 374:
		location = 165 //原站点名：商丘总站,结果：name=商丘市,id=165 | 商丘
	case 636:
		location = 629 //原站点名：沙河站,结果：name=沙河市,id=629 | 沙河
	case 638:
		location = 600 //原站点名：涉县站,结果：name=涉县,id=600 | 涉县
	case 129:
		location = 201 //原站点名：汕头站,结果：name=汕头市,id=201 | 汕头
	case 396:
		location = 171 //原站点名：十堰站,结果：name=十堰市,id=171 | 十堰
	case 145:
		location = 134 //原站点名：上饶站,结果：name=上饶市,id=134 | 上饶
	case 669:
		location = 2830 //原站点名：嵩明站,结果：name=嵩明县,id=2830 | 嵩明
	case 161:
		location = 1245 //原站点名：青浦站,结果：name=青浦区,id=1245 | 青浦
	case 169:
		location = 73 //原站点名：上海4站,结果：name=上海市,id=73 | 上海
	case 425:
		location = 57 //原站点名：松原站,结果：name=松原市,id=57 | 松原
	case 170:
		location = 73 //原站点名：上海精选,结果：name=上海市,id=73 | 上海
	case 172:
		location = 2298 //原站点名：顺德站,结果：name=顺德区,id=2298 | 顺德
	case 434:
		location = 1407 //原站点名：嵊州站,结果：name=嵊州市,id=1407 | 嵊州
	case 448:
		location = 187 //原站点名：邵阳站,结果：name=邵阳市,id=187 | 邵阳
	case 704:
		location = 283 //原站点名：山南站,结果：name=山南地区,id=283 | 山南
	case 450:
		location = 71 //原站点名：绥化站,结果：name=绥化市,id=71 | 绥化
	case 711:
		location = 2273 //原站点名：始兴仁化站,结果：name=始兴县,id=2273 | 始兴
	case 472:
		location = 708 //原站点名：三河站,结果：name=三河市,id=708 | 三河
	case 218:
		location = 86 //原站点名：宿迁站,结果：name=宿迁市,id=86 | 宿迁
	case 483:
		location = 1479 //原站点名：濉溪站,结果：name=濉溪县,id=1479 | 濉溪
	case 484:
		location = 4622 //原站点名：山东站,结果：name=山东省,id=4622 | 山东
	case 485:
		location = 4611 //原站点名：山西站,结果：name=山西省,id=4611 | 山西
	case 503:
		location = 4630 //原站点名：四川站,结果：name=四川省,id=4630 | 四川
	case 762:
		location = 1354 //原站点名：泗洪站,结果：name=泗洪县,id=1354 | 泗洪
	case 261:
		location = 0 //原站点名：塘厦站,结果：record not found | 塘厦
	case 782:
		location = 854 //原站点名：土默特右旗站,结果：name=土默特右旗,id=854 | 土默特右旗
	case 275:
		location = 85 //原站点名：泰州2站,结果：name=泰州市,id=85 | 泰州
	case 538:
		location = 2198 //原站点名：桃源站,结果：name=桃源县,id=2198 | 桃源
	case 33:
		location = 2 //原站点名：天津站,结果：name=天津市,id=2 | 天津
	case 37:
		location = 289 //原站点名：铜川站,结果：name=铜川市,id=289 | 铜川
	case 815:
		location = 777 //原站点名：太谷站,结果：name=太谷县,id=777 | 太谷
	case 310:
		location = 2304 //原站点名：台山站,结果：name=台山市,id=2304 | 台山
	case 63:
		location = 14 //原站点名：太原站,结果：name=太原市,id=14 | 太原
	case 72:
		location = 4 //原站点名：唐山站,结果：name=唐山市,id=4 | 唐山
	case 843:
		location = 55 //原站点名：通化站,结果：name=通化市,id=55 | 通化
	case 590:
		location = 3269 //原站点名：同心站,结果：name=同心县,id=3269 | 同心
	case 343:
		location = 1298 //原站点名：太仓总站,结果：name=太仓市,id=1298 | 太仓
	case 857:
		location = 48 //原站点名：铁岭站,结果：name=铁岭市,id=48 | 铁岭
	case 858:
		location = 2798 //原站点名：天柱站,结果：name=天柱县,id=2798 | 天柱
	case 111:
		location = 522 //原站点名：北京通州站,结果：name=通州区,id=522 | 通州区
	case 120:
		location = 85 //原站点名：泰州站,结果：name=泰州市,id=85 | 泰州
	case 121:
		location = 96 //原站点名：台州站,结果：name=台州市,id=96 | 台州
	case 392:
		location = 1363 //原站点名：桐庐站,结果：name=桐庐县,id=1363 | 桐庐
	case 400:
		location = 0 //原站点名：坦洲站,结果：record not found | 坦洲
	case 415:
		location = 29 //原站点名：通辽站,结果：name=通辽市,id=29 | 通辽
	case 420:
		location = 522 //原站点名：通州站,结果：name=通州区,id=522 | 通州
	case 182:
		location = 1396 //原站点名：桐乡站,结果：name=桐乡市,id=1396 | 桐乡
	case 185:
		location = 104 //原站点名：铜陵站,结果：name=铜陵市,id=104 | 铜陵
	case 702:
		location = 2677 //原站点名：通江站,结果：name=通江县,id=2677 | 通江
	case 713:
		location = 327 //原站点名：吐鲁番站,结果：name=吐鲁番地区,id=327 | 吐鲁番
	case 716:
		location = 2120 //原站点名：通山站,结果：name=通山县,id=2120 | 通山
	case 461:
		location = 1769 //原站点名：滕州站,结果：name=滕州市,id=1769 | 滕州
	case 723:
		location = 1831 //原站点名：郯城站,结果：name=郯城县,id=1831 | 郯城
	case 212:
		location = 143 //原站点名：泰安站,结果：name=泰安市,id=143 | 泰安
	case 473:
		location = 1892 //原站点名：通许站,结果：name=通许县,id=1892 | 通许
	case 732:
		location = 2528 //原站点名：铜梁站,结果：name=铜梁县,id=2528 | 铜梁
	case 224:
		location = 1298 //原站点名：太仓站,结果：name=太仓市,id=1298 | 太仓
	case 753:
		location = 260 //原站点名：铜仁站,结果：name=铜仁地区,id=260 | 铜仁
	case 763:
		location = 2527 //原站点名：潼南站,结果：name=潼南县,id=2527 | 潼南
	case 508:
		location = 302 //原站点名：天水站,结果：name=天水市,id=302 | 天水
	case 3:
		location = 169 //原站点名：武汉站,结果：name=武汉市,id=169 | 武汉
	case 776:
		location = 33 //原站点名：乌兰察布站,结果：name=乌兰察布市,id=33 | 乌兰察布
	case 534:
		location = 1410 //原站点名：武义站,结果：name=武义县,id=1410 | 武义
	case 280:
		location = 221 //原站点名：梧州站,结果：name=梧州市,id=221 | 梧州
	case 537:
		location = 2316 //原站点名：吴川站,结果：name=吴川市,id=2316 | 吴川
	case 34:
		location = 99 //原站点名：芜湖站,结果：name=芜湖市,id=99 | 芜湖
	case 36:
		location = 292 //原站点名：渭南老站,结果：name=渭南市,id=292 | 渭南
	case 806:
		location = 303 //原站点名：武威站,结果：name=武威市,id=303 | 武威
	case 39:
		location = 75 //原站点名：无锡站,结果：name=无锡市,id=75 | 无锡
	case 551:
		location = 275 //原站点名：文山站,结果：name=文山壮族苗族自治州,id=275 | 文山
	case 10028:
		location = 0 //原站点名：联联微课,结果：record not found | 联联微课
	case 46:
		location = 2511 //原站点名：万州站,结果：name=万州区,id=2511 | 万州
	case 565:
		location = 1732 //原站点名：婺源站,结果：name=婺源县,id=1732 | 婺源
	case 54:
		location = 2559 //原站点名：温江站,结果：name=温江区,id=2559 | 温江
	case 566:
		location = 2781 //原站点名：望谟站,结果：name=望谟县,id=2781 | 望谟
	case 827:
		location = 608 //原站点名：魏县站,结果：name=魏县,id=608 | 魏县
	case 64:
		location = 325 //原站点名：乌鲁木齐站,结果：name=乌鲁木齐市,id=325 | 乌鲁木齐
	case 585:
		location = 0 //原站点名： 武安站,结果：record not found |  武安
	case 586:
		location = 625 //原站点名：威县站,结果：name=威县,id=625 | 威县
	case 845:
		location = 2495 //原站点名：文昌站,结果：name=文昌市,id=2495 | 文昌
	case 591:
		location = 322 //原站点名：吴忠站,结果：name=吴忠市,id=322 | 吴忠
	case 600:
		location = 784 //原站点名：万荣站,结果：name=万荣县,id=784 | 万荣
	case 601:
		location = 1386 //原站点名：文成站,结果：name=文成县,id=1386 | 文成
	case 347:
		location = 89 //原站点名：温州总站,结果：name=温州市,id=89 | 温州
	case 603:
		location = 2617 //原站点名：威远站,结果：name=威远县,id=2617 | 威远
	case 348:
		location = 2511 //原站点名：万州总站,结果：name=万州区,id=2511 | 万州
	case 105:
		location = 89 //原站点名：温州站,结果：name=温州市,id=89 | 温州
	case 108:
		location = 144 //原站点名：威海站,结果：name=威海市,id=144 | 威海
	case 627:
		location = 682 //原站点名：围场站,结果：name=围场满族蒙古族自治县,id=682 | 围场
	case 633:
		location = 2275 //原站点名：翁源站,结果：name=翁源县,id=2275 | 翁源
	case 140:
		location = 141 //原站点名：潍坊站,结果：name=潍坊市,id=141 | 潍坊
	case 403:
		location = 1434 //原站点名：温岭站,结果：name=温岭市,id=1434 | 温岭
	case 659:
		location = 711 //原站点名：武邑站,结果：name=武邑县,id=711 | 武邑
	case 663:
		location = 705 //原站点名：文安站,结果：name=文安县,id=705 | 文安
	case 668:
		location = 2658 //原站点名：武胜站,结果：name=武胜县,id=2658 | 武胜
	case 427:
		location = 292 //原站点名：渭南站,结果：name=渭南市,id=292 | 渭南
	case 684:
		location = 1805 //原站点名：汶上站,结果：name=汶上县,id=1805 | 汶上
	case 174:
		location = 1297 //原站点名：吴江站,结果：name=吴江市,id=1297 | 吴江
	case 431:
		location = 288 //原站点名：西安2站,结果：name=西安市,id=288 | 西安
	case 437:
		location = 169 //原站点名：武汉特惠站,结果：name=武汉市,id=169 | 武汉
	case 693:
		location = 960 //原站点名：瓦房店站,结果：name=瓦房店市,id=960 | 瓦房店
	case 207:
		location = 169 //原站点名：武汉精选,结果：name=武汉市,id=169 | 武汉
	case 216:
		location = 27 //原站点名：乌海站,结果：name=乌海市,id=27 | 乌海
	case 733:
		location = 797 //原站点名：五台站,结果：name=五台县,id=797 | 五台
	case 223:
		location = 75 //原站点名：无锡精选,结果：name=无锡市,id=75 | 无锡
	case 744:
		location = 0 //原站点名：五兴站,结果：record not found | 五兴
	case 247:
		location = 183 //原站点名：长沙3站,结果：name=长沙市,id=183 | 长沙
	case 253:
		location = 543 //原站点名：武清站,结果：name=武清区,id=543 | 武清
	case 4:
		location = 288 //原站点名：西安站,结果：name=西安市,id=288 | 西安
	case 786:
		location = 1682 //原站点名：兴国站,结果：name=兴国县,id=1682 | 兴国
	case 23:
		location = 291 //原站点名：咸阳站,结果：name=咸阳市,id=291 | 咸阳
	case 27:
		location = 34 //原站点名：兴安盟站,结果：name=兴安盟,id=34 | 兴安盟
	case 797:
		location = 676 //原站点名：兴隆站,结果：name=兴隆县,id=676 | 兴隆
	case 547:
		location = 0 //原站点名：襄县站,结果：record not found | 襄县
	case 294:
		location = 1346 //原站点名：兴化站,结果：name=兴化市,id=1346 | 兴化
	case 555:
		location = 22 //原站点名：忻州站,结果：name=忻州市,id=22 | 忻州
	case 45:
		location = 2715 //原站点名：西昌站,结果：name=西昌市,id=2715 | 西昌
	case 302:
		location = 161 //原站点名：许昌站,结果：name=许昌市,id=161 | 许昌
	case 50:
		location = 2558 //原站点名：新都站,结果：name=新都区,id=2558 | 新都
	case 307:
		location = 343 //原站点名：香港站,结果：name=香港特别行政区,id=343 | 香港
	case 10038:
		location = 0 //原站点名：小蜜蜂,结果：record not found | 小蜜蜂
	case 571:
		location = 566 //原站点名：辛集站,结果：name=辛集市,id=566 | 辛集
	case 829:
		location = 1573 //原站点名：仙游站,结果：name=仙游县,id=1573 | 仙游
	case 74:
		location = 76 //原站点名：徐州站,结果：name=徐州市,id=76 | 徐州
	case 77:
		location = 312 //原站点名：西宁站,结果：name=西宁市,id=312 | 西宁
	case 335:
		location = 2776 //原站点名：兴义站,结果：name=兴义市,id=2776 | 兴义
	case 84:
		location = 116 //原站点名：厦门站,结果：name=厦门市,id=116 | 厦门
	case 605:
		location = 926 //原站点名：锡林浩特站,结果：name=锡林浩特市,id=926 | 锡林浩特
	case 861:
		location = 1404 //原站点名：新昌站,结果：name=新昌县,id=1404 | 新昌
	case 354:
		location = 185 //原站点名：湘潭总站,结果：name=湘潭市,id=185 | 湘潭
	case 99:
		location = 2076 //原站点名：襄阳站,结果：name=襄阳区,id=2076 | 襄阳
	case 613:
		location = 2016 //原站点名：西华站,结果：name=西华县,id=2016 | 西华
	case 364:
		location = 176 //原站点名：孝感站,结果：name=孝感市,id=176 | 孝感
	case 881:
		location = 0 //原站点名：西部新城站,结果：record not found | 西部新城
	case 637:
		location = 2277 //原站点名：新丰站,结果：name=新丰县,id=2277 | 新丰
	case 141:
		location = 158 //原站点名：新乡站,结果：name=新乡市,id=158 | 新乡
	case 399:
		location = 0 //原站点名：新加坡站,结果：record not found | 新加坡
	case 401:
		location = 1320 //原站点名：盱眙站,结果：name=盱眙县,id=1320 | 盱眙
	case 146:
		location = 166 //原站点名：信阳站,结果：name=信阳市,id=166 | 信阳
	case 676:
		location = 1627 //原站点名：霞浦站,结果：name=霞浦县,id=1627 | 霞浦
	case 422:
		location = 114 //原站点名：宣城站,结果：name=宣城市,id=114 | 宣城
	case 678:
		location = 3064 //原站点名：兴平站,结果：name=兴平市,id=3064 | 兴平
	case 679:
		location = 812 //原站点名：襄汾站,结果：name=襄汾县,id=812 | 襄汾
	case 175:
		location = 276 //原站点名：西双版纳站,结果：name=西双版纳傣族自治州,id=276 | 西双版纳
	case 433:
		location = 1279 //原站点名：新沂站,结果：name=新沂市,id=1279 | 新沂
	case 701:
		location = 2946 //原站点名：香格里拉站,结果：name=香格里拉县,id=2946 | 香格里拉
	case 190:
		location = 185 //原站点名：湘潭站,结果：name=湘潭市,id=185 | 湘潭
	case 451:
		location = 2002 //原站点名：夏邑站,结果：name=夏邑县,id=2002 | 夏邑
	case 720:
		location = 703 //原站点名：香河站,结果：name=香河县,id=703 | 香河
	case 214:
		location = 7 //原站点名：邢台站,结果：name=邢台市,id=7 | 邢台
	case 478:
		location = 1374 //原站点名：象山站,结果：name=象山县,id=1374 | 象山
	case 481:
		location = 128 //原站点名：新余站,结果：name=新余市,id=128 | 新余
	case 229:
		location = 2565 //原站点名：新津站,结果：name=新津县,id=2565 | 新津
	case 743:
		location = 952 //原站点名：新民站,结果：name=新民市,id=952 | 新民
	case 242:
		location = 0 //原站点名：小榄站,结果：record not found | 小榄
	case 244:
		location = 179 //原站点名：咸宁站,结果：name=咸宁市,id=179 | 咸宁
	case 245:
		location = 124 //原站点名：南昌2站,结果：name=南昌市,id=124 | 南昌
	case 252:
		location = 288 //原站点名：西安精选,结果：name=西安市,id=288 | 西安
	case 259:
		location = 2549 //原站点名：永川站,结果：name=永川市,id=2549 | 永川
	case 772:
		location = 3028 //原站点名：阎良站,结果：name=阎良区,id=3028 | 阎良
	case 5:
		location = 320 //原站点名：银川站,结果：name=银川市,id=320 | 银川
	case 520:
		location = 1416 //原站点名：永康站,结果：name=永康市,id=1416 | 永康
	case 523:
		location = 66 //原站点名：伊春站,结果：name=伊春市,id=66 | 伊春
	case 779:
		location = 1383 //原站点名：永嘉站,结果：name=永嘉县,id=1383 | 永嘉
	case 785:
		location = 2895 //原站点名：姚安站,结果：name=姚安县,id=2895 | 姚安
	case 787:
		location = 1592 //原站点名：永春站,结果：name=永春县,id=1592 | 永春
	case 539:
		location = 1092 //原站点名：延吉站,结果：name=延吉市,id=1092 | 延吉
	case 543:
		location = 2364 //原站点名：英德站,结果：name=英德市,id=2364 | 英德
	case 799:
		location = 761 //原站点名：阳城站,结果：name=阳城县,id=761 | 阳城
	case 550:
		location = 1966 //原站点名：鄢陵站,结果：name=鄢陵县,id=1966 | 鄢陵
	case 552:
		location = 129 //原站点名：鹰潭站,结果：name=鹰潭市,id=129 | 鹰潭
	case 554:
		location = 1870 //原站点名：郓城站,结果：name=郓城县,id=1870 | 郓城
	case 10030:
		location = 0 //原站点名：联联永兴站,结果：record not found | 联联永兴
	case 823:
		location = 529 //原站点名：延庆站,结果：name=延庆区,id=529 | 延庆
	case 10039:
		location = 0 //原站点名：亦鲜生站,结果：record not found | 亦鲜生
	case 62:
		location = 250 //原站点名：雅安站,结果：name=雅安市,id=250 | 雅安
	case 830:
		location = 2870 //原站点名：永胜站,结果：name=永胜县,id=2870 | 永胜
	case 319:
		location = 234 //原站点名：重庆9站,结果：name=重庆市,id=234 | 重庆
	case 579:
		location = 3350 //原站点名：伊宁站,结果：name=伊宁市,id=3350 | 伊宁
	case 73:
		location = 140 //原站点名：烟台站,结果：name=烟台市,id=140 | 烟台
	case 593:
		location = 603 //原站点名：永年站,结果：name=永年区,id=603 | 永年
	case 342:
		location = 0 //原站点名：联联永昌站,结果：record not found | 联联永昌
	case 88:
		location = 247 //原站点名：宜宾站,结果：name=宜宾市,id=247 | 宜宾
	case 95:
		location = 267 //原站点名：玉溪站,结果：name=玉溪市,id=267 | 玉溪
	case 98:
		location = 172 //原站点名：宜昌站,结果：name=宜昌市,id=172 | 宜昌
	case 868:
		location = 1681 //原站点名：于都站,结果：name=于都县,id=1681 | 于都
	case 104:
		location = 83 //原站点名：扬州站,结果：name=扬州市,id=83 | 扬州
	case 363:
		location = 211 //原站点名：阳江总站,结果：name=阳江市,id=211 | 阳江
	case 620:
		location = 2082 //原站点名：宜城站,结果：name=宜城市,id=2082 | 宜城
	case 368:
		location = 2539 //原站点名：云阳站,结果：name=云阳县,id=2539 | 云阳
	case 624:
		location = 2828 //原站点名：宜良站,结果：name=宜良县,id=2828 | 宜良
	case 369:
		location = 2003 //原站点名：永城站,结果：name=永城市,id=2003 | 永城
	case 625:
		location = 3052 //原站点名：杨凌站,结果：name=杨凌区,id=3052 | 杨凌
	case 370:
		location = 211 //原站点名：阳江新站,结果：name=阳江市,id=211 | 阳江
	case 632:
		location = 2406 //原站点名：阳朔站,结果：name=阳朔县,id=2406 | 阳朔
	case 122:
		location = 188 //原站点名：岳阳站,结果：name=岳阳市,id=188 | 岳阳
	case 379:
		location = 1269 //原站点名：宜兴新站,结果：name=宜兴市,id=1269 | 宜兴
	case 126:
		location = 1269 //原站点名：宜兴站,结果：name=宜兴市,id=1269 | 宜兴
	case 128:
		location = 1414 //原站点名：义乌站,结果：name=义乌市,id=1414 | 义乌
	case 646:
		location = 1376 //原站点名：余姚站,结果：name=余姚市,id=1376 | 余姚
	case 648:
		location = 2210 //原站点名：沅江站,结果：name=沅江市,id=2210 | 沅江
	case 398:
		location = 191 //原站点名：益阳站,结果：name=益阳市,id=191 | 益阳
	case 148:
		location = 82 //原站点名：盐城站,结果：name=盐城市,id=82 | 盐城
	case 152:
		location = 1335 //原站点名：仪征站,结果：name=仪征市,id=1335 | 仪征
	case 409:
		location = 132 //原站点名：宜春站,结果：name=宜春市,id=132 | 宜春
	case 412:
		location = 217 //原站点名：云浮站,结果：name=云浮市,id=217 | 云浮
	case 162:
		location = 1238 //原站点名：杨浦站,结果：name=杨浦区,id=1238 | 杨浦
	case 421:
		location = 1968 //原站点名：禹州站,结果：name=禹州市,id=1968 | 禹州
	case 689:
		location = 1600 //原站点名：云霄站,结果：name=云霄县,id=1600 | 云霄
	case 695:
		location = 664 //原站点名：蔚县站,结果：name=蔚县,id=664 | 蔚县
	case 440:
		location = 2355 //原站点名：阳西站,结果：name=阳西县,id=2355 | 阳西
	case 186:
		location = 0 //原站点名：燕郊站,结果：record not found | 燕郊
	case 443:
		location = 16 //原站点名：阳泉站,结果：name=阳泉市,id=16 | 阳泉
	case 195:
		location = 295 //原站点名：榆林站,结果：name=榆林市,id=295 | 榆林
	case 708:
		location = 564 //原站点名：元氏站,结果：name=元氏县,id=564 | 元氏
	case 712:
		location = 2360 //原站点名：阳山站,结果：name=阳山县,id=2360 | 阳山
	case 459:
		location = 0 //原站点名：联联永欣教育站,结果：record not found | 联联永欣教育
	case 717:
		location = 2186 //原站点名：岳阳县站,结果：name=岳阳县,id=2186 | 岳阳县
	case 719:
		location = 2071 //原站点名：宜都站,结果：name=宜都市,id=2071 | 宜都
	case 210:
		location = 21 //原站点名：运城站,结果：name=运城市,id=21 | 运城
	case 213:
		location = 211 //原站点名：阳江站,结果：name=阳江市,id=211 | 阳江
	case 469:
		location = 2657 //原站点名：岳池站,结果：name=岳池县,id=2657 | 岳池
	case 474:
		location = 1430 //原站点名：玉环站,结果：name=玉环县,id=1430 | 玉环
	case 731:
		location = 1832 //原站点名：沂水站,结果：name=沂水县,id=1832 | 沂水
	case 480:
		location = 3092 //原站点名：洋县站,结果：name=洋县,id=3092 | 洋县
	case 230:
		location = 2138 //原站点名：岳麓站,结果：name=岳麓区,id=2138 | 岳麓
	case 235:
		location = 293 //原站点名：延安站,结果：name=延安市,id=293 | 延安
	case 241:
		location = 44 //原站点名：营口站,结果：name=营口市,id=44 | 营口
	case 755:
		location = 1802 //原站点名：鱼台站,结果：name=鱼台县,id=1802 | 鱼台
	case 757:
		location = 2635 //原站点名：营山站,结果：name=营山县,id=2635 | 营山
	case 504:
		location = 4632 //原站点名：云南站,结果：name=云南省,id=4632 | 云南
	case 254:
		location = 226 //原站点名：玉林站,结果：name=玉林市,id=226 | 玉林
	case 255:
		location = 84 //原站点名：镇江2站,结果：name=镇江市,id=84 | 镇江
	case 512:
		location = 1166 //原站点名：肇源站,结果：name=肇源县,id=1166 | 肇源
	case 519:
		location = 1225 //原站点名：肇东站,结果：name=肇东市,id=1225 | 肇东
	case 8:
		location = 152 //原站点名：郑州站,结果：name=郑州市,id=152 | 郑州
	case 530:
		location = 1880 //原站点名：中牟站,结果：name=中牟县,id=1880 | 中牟
	case 20:
		location = 252 //原站点名：资阳站,结果：name=资阳市,id=252 | 资阳
	case 541:
		location = 0 //原站点名：周边游信息站,结果：record not found | 周边游信息
	case 542:
		location = 190 //原站点名：张家界站,结果：name=张家界市,id=190 | 张家界
	case 548:
		location = 168 //原站点名：驻马店站,结果：name=驻马店市,id=168 | 驻马店
	case 48:
		location = 137 //原站点名：淄博站,结果：name=淄博市,id=137 | 淄博
	case 821:
		location = 1794 //原站点名：诸城站,结果：name=诸城市,id=1794 | 诸城
	case 314:
		location = 152 //原站点名：郑州精选,结果：name=郑州市,id=152 | 郑州
	case 580:
		location = 2349 //原站点名：紫金站,结果：name=紫金县,id=2349 | 紫金
	case 584:
		location = 1165 //原站点名：肇州站,结果：name=肇州县,id=1165 | 肇州
	case 840:
		location = 2599 //原站点名：梓潼站,结果：name=梓潼县,id=2599 | 梓潼
	case 80:
		location = 84 //原站点名：镇江站,结果：name=镇江市,id=84 | 镇江
	case 848:
		location = 651 //原站点名：涿州站,结果：name=涿州市,id=651 | 涿州
	case 89:
		location = 236 //原站点名：自贡站,结果：name=自贡市,id=236 | 自贡
	case 92:
		location = 258 //原站点名：遵义站,结果：name=遵义市,id=258 | 遵义
	case 609:
		location = 2618 //原站点名：资中站,结果：name=资中县,id=2618 | 资中
	case 611:
		location = 2589 //原站点名：中江站,结果：name=中江县,id=2589 | 中江
	case 101:
		location = 95 //原站点名：舟山站,结果：name=舟山市,id=95 | 舟山
	case 113:
		location = 200 //原站点名：珠海站,结果：name=珠海市,id=200 | 珠海
	case 115:
		location = 214 //原站点名：中山站,结果：name=中山市,id=214 | 中山
	case 645:
		location = 1743 //原站点名：章丘站,结果：name=章丘市,id=1743 | 章丘
	case 137:
		location = 184 //原站点名：株洲站,结果：name=株洲市,id=184 | 株洲
	case 397:
		location = 138 //原站点名：枣庄站,结果：name=枣庄市,id=138 | 枣庄
	case 656:
		location = 2090 //原站点名：钟祥站,结果：name=钟祥市,id=2090 | 钟祥
	case 151:
		location = 269 //原站点名：昭通站,结果：name=昭通市,id=269 | 昭通
	case 156:
		location = 1405 //原站点名：诸暨站,结果：name=诸暨市,id=1405 | 诸暨
	case 416:
		location = 167 //原站点名：周口站,结果：name=周口市,id=167 | 周口
	case 163:
		location = 204 //原站点名：湛江站,结果：name=湛江市,id=204 | 湛江
	case 164:
		location = 9 //原站点名：张家口站,结果：name=张家口市,id=9 | 张家口
	case 685:
		location = 304 //原站点名：张掖站,结果：name=张掖市,id=304 | 张掖
	case 177:
		location = 2287 //原站点名：斗门站,结果：name=斗门区,id=2287 | 斗门
	case 694:
		location = 324 //原站点名：中卫站,结果：name=中卫市,id=324 | 中卫
	case 439:
		location = 2537 //原站点名：忠县站,结果：name=忠县,id=2537 | 忠县
	case 189:
		location = 120 //原站点名：漳州站,结果：name=漳州市,id=120 | 漳州
	case 201:
		location = 1295 //原站点名：张家港站,结果：name=张家港市,id=1295 | 张家港
	case 203:
		location = 2268 //原站点名：增城站,结果：name=增城市,id=2268 | 增城
	case 727:
		location = 583 //原站点名：遵化站,结果：name=遵化市,id=583 | 遵化
	case 736:
		location = 2073 //原站点名：枝江站,结果：name=枝江市,id=2073 | 枝江
	case 226:
		location = 206 //原站点名：肇庆站,结果：name=肇庆市,id=206 | 肇庆
	case 488:
		location = 4618 //原站点名：浙江站,结果：name=浙江省,id=4618 | 浙江
	case 745:
		location = 2068 //原站点名：秭归站,结果：name=秭归县,id=2068 | 秭归
	case 754:
		location = 3079 //原站点名：子长站,结果：name=子长县,id=3079 | 子长
	case 766:
		location = 2780 //原站点名：贞丰兴仁册亨站,结果：name=贞丰县,id=2780 | 贞丰
	default:
		location = 0
	}
	if location == 0 {
		//拉果没找到 - 默认全国站
		//location = 3374
	}
	return location
}
