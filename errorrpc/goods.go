package errorrpc

//商品服务code
const GOODS_ERROR = 12000

var (
	GoodsErr                    = NewCodeError(GOODS_ERROR+1, "商品错误")
	GoodsDBErr                  = NewCodeError(GOODS_ERROR+2, "DB错误")
	GoodsOutOfStock             = NewCodeError(GOODS_ERROR+3, "库存不足")
	GoodsPurchaseFail           = NewCodeError(SYS_ERROR+4, "联联采购失败")
	GoodsWhitelistErr           = NewCodeError(SYS_ERROR+5, "不在白名单内")
	GoodsExistsErr              = NewCodeError(GOODS_ERROR+6, "商品名称已存在")
	GoodsChannelErr             = NewCodeError(GOODS_ERROR+7, "SKU商品类型与供应链商品渠道不一致")
	GoodsStockAlreadyBindErr    = NewCodeError(GOODS_ERROR+8, "批次号已绑定其它SKU")
	GoodsSkuAlreadyBindOtherErr = NewCodeError(GOODS_ERROR+9, "SKU已绑定其它批次号")
	ModuleSpuSortErr            = NewCodeError(GOODS_ERROR+10, "模块商品排序错误")
	ModuleSelectSpuErr          = NewCodeError(GOODS_ERROR+11, "模块选品错误")
	SharingRuleErr              = NewCodeError(GOODS_ERROR+12, "分账规则错误")
	GoodsUpdateStockByThirdSpuIdErr              = NewCodeError(GOODS_ERROR+13, "根据第三方SPU更改库存错误")
)
