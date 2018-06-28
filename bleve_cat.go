package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/blevesearch/bleve"
)

type Bleve_Data struct {
	AdId                     int64    `json:"ad_id"`
	ShopId                   int64    `json:"shop_id"`
	ItemId                   int64    `json:"item_id"`
	ParentItemId             int64    `json:"parent_item_id"`
	PriceBid                 int      `json:"price_bid"`
	AdTitle                  string   `json:"ad_title"`
	StickerId                int      `json:"sticker_id"`
	Score                    float64  `json:"score"`
	ScoreSearch              float64  `json:"score_search"`
	ScoreBrowse              float64  `json:"score_browse"`
	ScoreSearchV2            float64  `json:"score_search_v2"`
	ScoreBrowseV2            float64  `json:"score_browse_v2"`
	DisplayScore             float64  `json:"display_score"`
	CtrSearch                float64  `json:"ctr_search"`
	CtrBrowse                float64  `json:"ctr_browse"`
	DId                      []int    `json:"d_id"`
	UserId                   int64    `json:"user_id"`
	ClusterId                int64    `json:"cluster_id"`
	ShopName                 string   `json:"shop_name"`
	ShopDomain               string   `json:"shop_domain"`
	ShopCurrencyRate         int64    `json:"shop_currency_rate"`
	ShopLocationId           int64    `json:"shop_location_id"`
	ShopCityId               int64    `json:"shop_city_id"`
	ShopCityName             string   `json:"shop_city_name"`
	ShopIsGmStatus           int      `json:"shop_is_gm_status"`
	ShopIsGmDate             string   `json:"shop_is_gm_date"`
	ShopIsGmScore            int64    `json:"shop_is_gm_score"`
	ShopIsGmBadge            bool     `json:"shop_is_gm_badge"`
	ProductAlias             string   `json:"product_alias"`
	ProductPrice             float64  `json:"product_price"`
	ProductPriceIdr          float64  `json:"product_price_idr"`
	ProductShippingId        []int    `json:"product_shipping_id"`
	ProductCondition         int      `json:"product_condition"`
	ProductFreeReturn        int      `json:"product_free_return"`
	ProductTalk              int      `json:"product_talk"`
	ProductReview            int      `json:"product_review"`
	ProductRating            int      `json:"product_rating"`
	ProductCreateTime        string   `json:"product_create_time"`
	ProductRatingStar        int      `json:"product_rating_star"`
	ProductVariantVuvId      []int64  `json:"product_variant_vuv_id"`
	ProductImageServer       int64    `json:"product_image_server"`
	ProductImageFilepath     string   `json:"product_image_filepath"`
	ProductImageFilename     string   `json:"product_image_filename"`
	BrandId                  int64    `json:"brand_id"`
	BrandName                string   `json:"brand_name"`
	NegativeKeywordExactTag  []string `json:"negative_keyword_exact_tag"`
	NegativeKeywordPhraseTag []string `json:"negative_keyword_phrase_tag"`
}

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	mapping.TypeField = "blog"
	mapping.DefaultType = "blog"
	blogMapping := bleve.NewDocumentMapping()
	nameFieldMapping := bleve.NewTextFieldMapping()
	nameFieldMapping.IncludeTermVectors = true
	blogMapping.AddFieldMappingsAt("ShopName", nameFieldMapping)
	mapping.AddDocumentMapping("blog", blogMapping)

	index, err := bleve.New("example.bleve.5", mapping)
	//index, err := bleve.Open("large.bleve.5") //open existing connection
	if err != nil {
		fmt.Println(err)
		return
	}

	//index some data
	insertData(index)
	//fmt.Println(index.DocCount())  = 129313
	searchUsingCatId(index)
	searchUsingKeyword(index)
	searchUsingKeywordNCategory(index)
}

func insertData(index bleve.Index) {
	// index some data

	inputData := [][]byte{
		[]byte(`{"ad_id":7364051,"shop_id":574630,"item_id":161433142,"parent_item_id":161433142,"price_bid":50,"ad_title":"FRAME TITONT / KACAMATA MINUS / KACAMATA OPTIK / KACAMATA RINGAN","sticker_id":3,"score":1.1915552173211568,"score_search":1.1915552173211568,"score_browse":1.1917552971326137,"score_search_v2":0.022968034825696707,"score_browse_v2":0.022968034825696707,"display_score":1.5915552173211567,"ctr_search":0,"ctr_browse":0,"d_id":[1758,1932,1952],"user_id":4829154,"cluster_id":12933,"label_id":30036,"shop_name":"LTX Homemade Production","shop_domain":"ltxhomemade","shop_currency_rate":1,"shop_location_id":1640,"shop_city_id":146,"shop_city_name":"Kota Tangerang","shop_is_gm_status":0,"shop_is_gm_score":64,"shop_is_gm_badge":false,"product_alias":"frame-titont-kacamata-minus-kacamata-optik-kacamata-ringan","product_price":150000,"product_price_idr":150000,"product_shipping_id":[1,10,13,14],"product_condition":1,"product_free_return":0,"product_talk":0,"product_review":0,"product_rating":0,"product_create_time":"2017-03-18T19:03:38.234315Z","product_rating_star":0,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/3/18/4829154","product_image_filename":"4829154_ded3953e-87c1-416f-aa41-122fa6686cf0_2048_0.jpg","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":[]}`),
		[]byte(`{"ad_id":5787569,"shop_id":1283319,"item_id":110913239,"parent_item_id":110913239,"price_bid":200,"ad_title":"Hippo Sapphire Tempered Glass Xiaomi Note2 / Xiaomi Note 2","sticker_id":3,"score":1.36940983411617,"score_search":1.36940983411617,"score_browse":1.36940983411617,"score_search_v2":0.03110716312854366,"score_browse_v2":0.03110716312854366,"display_score":1.76940983411617,"ctr_search":0,"ctr_browse":0,"d_id":[65,66,76],"user_id":10331914,"cluster_id":3913,"shop_name":"Statisch Distro","shop_domain":"statischdistro","shop_currency_rate":1,"shop_location_id":2254,"shop_city_id":174,"shop_city_name":"Jakarta","shop_is_gm_status":1,"shop_is_gm_date":"2019-03-24T17:00:00Z","shop_is_gm_score":97,"shop_is_gm_badge":true,"product_alias":"hippo-sapphire-tempered-glass-xiaomi-note2-xiaomi-note-2","product_price":27000,"product_price_idr":27000,"product_shipping_id":[1,2,6,10,11,12,13,14],"product_condition":1,"product_free_return":0,"product_talk":2,"product_review":2,"product_rating":90,"product_create_time":"2016-12-14T14:09:13.757885Z","product_rating_star":5,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/10/4/110913239","product_image_filename":"110913239_8fcc5d02-1fa7-4435-b9f6-c1c8e6057056_389_389.jpg","brand_id":2199,"brand_name":"hippo","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":[]}`),
		[]byte(`{"ad_id":7358796,"shop_id":1825916,"item_id":138633624,"parent_item_id":138633624,"price_bid":200,"ad_title":"Casing Blackbery Torch/9810 Original","sticker_id":3,"score":1.36940983411617,"score_search":1.36940983411617,"score_browse":1.36940983411617,"score_search_v2":0.03110716312854366,"score_browse_v2":0.03110716312854366,"display_score":1.76940983411617,"ctr_search":0,"ctr_browse":0,"d_id":[65,66,69],"user_id":16113321,"cluster_id":1364,"label_id":40018,"shop_name":"Awaa Cell","shop_domain":"umarali-","shop_currency_rate":1,"shop_location_id":5573,"shop_city_id":174,"shop_city_name":"Jakarta","shop_is_gm_status":1,"shop_is_gm_date":"2018-08-11T06:21:22Z","shop_is_gm_score":51,"shop_is_gm_badge":false,"product_alias":"casing-blackbery-torch9810-original","product_price":100000,"product_price_idr":100000,"product_shipping_id":[1,11,14],"product_condition":1,"product_free_return":0,"product_talk":2,"product_review":0,"product_rating":0,"product_create_time":"2017-02-05T20:59:31.266577Z","product_rating_star":0,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/2/5/1825916","product_image_filename":"1825916_c93f0de1-96e7-4d25-a442-593dfb03cda0.jpg","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":[]}`),
		[]byte(`{"ad_id":7353134,"shop_id":1267473,"item_id":161253322,"parent_item_id":161253322,"price_bid":300,"ad_title":"Green Angelica Hair Shampo Pencegah Kebotakan Diusia DIni","sticker_id":3,"score":1.4416962179471127,"score_search":1.4416962179471127,"score_browse":1.4416962179471127,"score_search_v2":0.033487705551405504,"score_browse_v2":0.033487705551405504,"display_score":1.8416962179471126,"ctr_search":0,"ctr_browse":0,"d_id":[2133,2147,2174],"user_id":10678699,"cluster_id":2896,"label_id":60010,"shop_name":"green angelica hair","shop_domain":"hairregrowth","shop_currency_rate":1,"shop_location_id":2271,"shop_city_id":175,"shop_city_name":"Jakarta","shop_is_gm_status":1,"shop_is_gm_date":"2018-08-02T03:07:22Z","shop_is_gm_score":99,"shop_is_gm_badge":true,"product_alias":"green-angelica-hair-shampo-pencegah-kebotakan-diusia-dini","product_price":185000,"product_price_idr":185000,"product_shipping_id":[1,2,6,10,12,13,14],"product_condition":1,"product_free_return":0,"product_talk":2,"product_review":5,"product_rating":96,"product_create_time":"2017-03-18T13:29:55.784841Z","product_rating_star":5,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/11/21/161253322","product_image_filename":"161253322_df2c0cff-e605-478e-baa0-2761595d52df_1000_1000.jpeg","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":[]}`),
		[]byte(`{"ad_id":7351496,"shop_id":315078,"item_id":161226825,"parent_item_id":161226825,"price_bid":150,"ad_title":"Otsu Vintage Brown","sticker_id":3,"score":1.3237475295908787,"score_search":1.3237475295908787,"score_browse":1.3237475295908787,"score_search_v2":0.02941814139998203,"score_browse_v2":0.02941814139998203,"display_score":1.7237475295908786,"ctr_search":0,"ctr_browse":0,"d_id":[1759,1800,1850],"user_id":2362432,"cluster_id":12988,"label_id":30063,"shop_name":"Otiv Store","shop_domain":"otiv","shop_currency_rate":1,"shop_location_id":2170,"shop_city_id":165,"shop_city_name":"Kota Bandung","shop_is_gm_status":1,"shop_is_gm_date":"2018-06-30T17:00:00Z","shop_is_gm_score":99,"shop_is_gm_badge":true,"product_alias":"otsu-vintage-brown","product_price":375000,"product_price_idr":375000,"product_wholesale_price":[],"product_shipping_id":[1,4,10],"product_condition":1,"product_free_return":3,"product_talk":8,"product_review":2,"product_rating":100,"product_create_time":"2017-03-18T12:46:05.490807Z","product_rating_star":5,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/3/18/2362432","product_image_filename":"2362432_00c38c27-8528-429b-a8c0-fcfd1398e795_700_700.jpg","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":["sepatu anak","sepatu kulit sintetis","sepatu wanita"]}`),
		[]byte(`{"ad_id":7387720,"shop_id":828853,"item_id":160340299,"parent_item_id":160340299,"price_bid":150,"ad_title":"Thermostat & Sil Jazz Idsi/Vtec, City Idsi/Vtec Jepang","sticker_id":3,"score":1.3232292907585799,"score_search":1.3232292907585799,"score_browse":1.3237475295908787,"score_search_v2":0.02941814139998203,"score_browse_v2":0.02941814139998203,"display_score":1.7232292907585798,"ctr_search":0,"ctr_browse":0,"d_id":[63,38,1342],"user_id":6999031,"cluster_id":24197,"label_id":70030,"shop_name":"Part Honda 123","shop_domain":"parthonda123","shop_currency_rate":1,"shop_location_id":3536,"shop_city_id":252,"shop_city_name":"Kota Surabaya","shop_is_gm_status":1,"shop_is_gm_date":"2018-11-05T17:04:50Z","shop_is_gm_score":99,"shop_is_gm_badge":true,"product_alias":"thermostat-sil-jazz-idsivtec-city-idsivtec-jepang","product_price":115000,"product_price_idr":115000,"product_shipping_id":[1,4,10,13,14],"product_condition":1,"product_free_return":0,"product_talk":5,"product_review":3,"product_rating":93,"product_create_time":"2017-03-17T01:23:14.729785Z","product_rating_star":5,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/7/30/160340299","product_image_filename":"160340299_82ecfaf9-bce0-4d1e-96ac-6cf2e383113c_848_768.jpg","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":[]}`),
		[]byte(`{"ad_id":7386543,"shop_id":170039,"item_id":160990540,"parent_item_id":160990540,"price_bid":50,"ad_title":"Mango MA6705L-95R Original","sticker_id":2,"score":1.1917552971326137,"score_search":1.1917552971326137,"score_browse":1.1917552971326137,"score_search_v2":0.022968034825696707,"score_browse_v2":0.022968034825696707,"display_score":1.5917552971326137,"ctr_search":0,"ctr_browse":0,"d_id":[1758,1946,1949],"user_id":991347,"cluster_id":25956,"label_id":30031,"shop_name":"maredibeli","shop_domain":"maredibeli","shop_currency_rate":1,"shop_location_id":2279,"shop_city_id":176,"shop_city_name":"Jakarta","shop_is_gm_status":1,"shop_is_gm_date":"2018-06-18T04:13:10Z","shop_is_gm_score":74,"shop_is_gm_badge":false,"product_alias":"mango-ma6705l-95r-original","product_price":1355000,"product_price_idr":1355000,"product_shipping_id":[1,2,4,10,11,13,14],"product_condition":1,"product_free_return":0,"product_talk":0,"product_review":0,"product_rating":0,"product_create_time":"2017-03-18T04:33:00.564152Z","product_rating_star":0,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/3/18/991347","product_image_filename":"991347_78891dbd-26cc-4424-bbc9-2a81ad6ef6e5_798_798.jpg","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":[]}`),
		[]byte(`{"ad_id":7386542,"shop_id":170039,"item_id":160990171,"parent_item_id":160990171,"price_bid":50,"ad_title":"Mango MA6701L-81TK Original","sticker_id":2,"score":1.2917552971326137,"score_search":1.1917552971326137,"score_browse":1.1917552971326137,"score_search_v2":0.022968034825696707,"score_browse_v2":0.022968034825696707,"display_score":1.5917552971326137,"ctr_search":0,"ctr_browse":0,"d_id":[1758,1946,1949],"user_id":991347,"cluster_id":25956,"label_id":30031,"shop_name":"maredibeli","shop_domain":"maredibeli","shop_currency_rate":1,"shop_location_id":2279,"shop_city_id":176,"shop_city_name":"Jakarta","shop_is_gm_status":1,"shop_is_gm_date":"2018-06-18T04:13:10Z","shop_is_gm_score":74,"shop_is_gm_badge":false,"product_alias":"mango-ma6701l-81tk-original","product_price":1250000,"product_price_idr":1250000,"product_shipping_id":[1,2,4,10,11,13,14],"product_condition":1,"product_free_return":0,"product_talk":0,"product_review":0,"product_rating":0,"product_create_time":"2017-03-18T04:31:34.100358Z","product_rating_star":0,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/3/18/991347","product_image_filename":"991347_39543c3c-b267-4032-94dd-40c02861a789_798_798.jpg","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":[]}`),
		[]byte(`{"ad_id":7385178,"shop_id":100497,"item_id":162023952,"parent_item_id":162023952,"price_bid":50,"ad_title":"TIDE TO GO STAIN REMOVER PENA PENGHILANG NODA 10 ML","sticker_id":1,"score":1.1917552971326137,"score_search":1.1917552971326137,"score_browse":1.1917552971326137,"score_search_v2":0.022968034825696707,"score_browse_v2":0.022968034825696707,"display_score":1.5917552971326137,"ctr_search":0,"ctr_browse":0,"d_id":[984,987,1014],"user_id":303440,"cluster_id":23287,"label_id":50000,"shop_name":"Lapak Lucu","shop_domain":"milkashop","shop_currency_rate":1,"shop_location_id":1640,"shop_city_id":146,"shop_city_name":"Kota Tangerang","shop_is_gm_status":1,"shop_is_gm_date":"2018-06-01T03:11:55Z","shop_is_gm_score":79,"shop_is_gm_badge":true,"product_alias":"tide-to-go-stain-remover-pena-penghilang-noda-10-ml","product_price":98000,"product_price_idr":98000,"product_shipping_id":[1,6,10,13,14],"product_condition":1,"product_free_return":3,"product_talk":2,"product_review":0,"product_rating":0,"product_create_time":"2017-03-19T20:49:44.801238Z","product_rating_star":0,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/3/19/303440","product_image_filename":"303440_ed2141c3-d585-4413-a389-a56783269b44_600_600.jpg","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":[]}`),
		[]byte(`{"ad_id":7382277,"shop_id":1048989,"item_id":127820989,"parent_item_id":127820989,"price_bid":50,"ad_title":"SANYO ph 175","sticker_id":3,"score":1.1917552971326137,"score_search":1.1917552971326137,"score_browse":1.1917552971326137,"score_search_v2":0.022968034825696707,"score_browse_v2":0.022968034825696707,"display_score":1.5917552971326137,"ctr_search":0,"ctr_browse":0,"d_id":[984,1262,1264],"user_id":7554948,"cluster_id":59,"label_id":50000,"shop_name":"sumberjayatehnik pompa","shop_domain":"sjtpump","shop_currency_rate":1,"shop_location_id":2268,"shop_city_id":175,"shop_city_name":"Jakarta","shop_is_gm_status":1,"shop_is_gm_date":"2018-07-05T13:16:11Z","shop_is_gm_score":99,"shop_is_gm_badge":true,"product_alias":"sanyo-ph-175","product_price":1818000,"product_price_idr":1818000,"product_shipping_id":[1,6,10,13,14],"product_condition":1,"product_free_return":0,"product_talk":6,"product_review":1,"product_rating":100,"product_create_time":"2017-01-13T15:53:44.415102Z","product_rating_star":5,"product_variant_vuv_id":[],"product_image_server":50,"product_image_filepath":"product-1/2017/1/13/1048989","product_image_filename":"1048989_843d14ca-ea7c-4a04-b3eb-75eb65c1cf7a.jpg","negative_keyword_exact_tag":[],"negative_keyword_phrase_tag":[]}`),
	}

	for _, jsonData := range inputData {
		var v Bleve_Data
		json.Unmarshal([]byte(jsonData), &v)
		id := strconv.FormatInt(v.AdId, 10)

		index.Index(id, v)
		index.SetInternal([]byte(id), jsonData)
	}

	fmt.Println("Total Documents")
	fmt.Println(index.DocCount())
}

func searchUsingCatId(index bleve.Index) {
	fmt.Println("SEARCHING USING CATEGORY")
	var val1 float64 = 984
	var val2 float64 = 985
	query := bleve.NewNumericRangeQuery(&val1, &val2)
	query.SetField("d_id")

	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults.Hits.Len())

	for _, res := range searchResults.Hits {
		raw, _ := index.GetInternal([]byte(res.ID))
		var dres Bleve_Data
		json.Unmarshal(raw, &dres)
		fmt.Println(dres.AdId)
	}
}

func searchUsingKeyword(index bleve.Index) {
	fmt.Println("SEARCHING USING KEYWORD")
	start := time.Now()
	query := bleve.NewMatchQuery("iphone")
	query.SetField("ad_title")

	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	search.Size = 9999
	//search.SortBy([]string{"-score"})
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults.Hits.Len())

	for _, res := range searchResults.Hits {
		raw, _ := index.GetInternal([]byte(res.ID))
		var dres Bleve_Data
		json.Unmarshal(raw, &dres)
	}
	fmt.Println("TOTAL TIME", time.Since(start))
}

func searchUsingKeywordNCategory(index bleve.Index) {
	fmt.Println("SEARCHING USING KEYWORD & CATEGORY")
	query1 := bleve.NewMatchQuery("SANYO")
	query1.SetField("ad_title")

	query2 := bleve.NewMatchQuery("175")
	query2.SetField("ad_title")

	var val1 float64 = 984
	var val2 float64 = 985
	query3 := bleve.NewNumericRangeQuery(&val1, &val2)
	query3.SetField("d_id")

	query := bleve.NewConjunctionQuery(query1, query2, query3)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults.Hits.Len())

	for _, res := range searchResults.Hits {
		raw, _ := index.GetInternal([]byte(res.ID))
		var dres Bleve_Data
		json.Unmarshal(raw, &dres)
		fmt.Println(dres.AdId)
	}
}
