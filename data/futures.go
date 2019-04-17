package data

import "github.com/lwl1989/data-cover/types"

//基本档案 113
type GoodsBase struct {
    GoodsId [20]types.X
    RiseLimitPrice1 [9]types.N9
    PeferencePrice [9]types.N9
    FallLimitPrice1 [9]types.N9
    RiseLimitPrice2 [9]types.N9
    FallLimitPrice2 [9]types.N9
    RiseLimitPrice3 [9]types.N9
    FallLimitPrice3 [9]types.N9

    ProdKind        [1]types.X
    DecimalLocator  [1]types.N9
    StrikePriceDL   [1]types.N9

    BeginDate       [8]types.N9
    EndDate         [8]types.N9
    FlowGroup       [2]types.N9
    DeliveryDate    [8]types.N9
    DynamicBanding  [1]types.X
}

//股票参照档案  24
type StockSelectRefer struct {
    IndexKind   [8]types.X
    IndexNumber [6]types.X
    IndexValue  [5]types.N9
    IndexStatus [1]types.X
}

//股票调整档案  154
type StockSelectChange struct {
    BaseDate        [8]types.N9

    BfKindId        [8]types.X
    BfStockId       [6]types.X
    BfStockQnTy     [6]types.N9
    BfStockCash2    [8]types.N9
    BfStockCash3    [6]types.N9
    BfStockId4      [6]types.X
    BfStockQnTy4    [6]types.N9

    AfKindId        [8]types.X
    AfStockId       [6]types.X
    AfStockQnTy     [6]types.N9
    AfStockCash2    [8]types.N9
    AFStockPrice3   [6]types.N9
    AfStockQnTy3    [6]types.N9
    AfStockDate3    [6]types.N9
    AfStockId4      [6]types.X
    AfStockQnTy4    [6]types.N9

    DividendDate    [8]types.N9
}

//契约基本档案  75
type ConTractBase struct {
    KindId  [8]types.X
    Name    [30]types.X
    StockId [6]types.X
    SubType [1]types.X
    ConTractSize    [7]types.N9

    StatusCode      [1]types.X
    CurrencyType    [1]types.X

    DecimalLocator  [1]types.N9
    StrikePriceDL   [1]types.N9

    AcceptQuoteFlag [1]types.X
    BeginDate       [8]types.X
    BlockTradeFlag  [1]types.X
    ExpiryType      [1]types.X
    UnderLyingType  [1]types.X
    MarketCloseGroup    [2]types.N9

    EndSession      [1]types.X
}

//契约风险保证  38
type ConTractCafe struct {
    KindId          [8]types.Char

    AbType          types.Char
    CurrencyType    types.Char

    Margin          [9]types.Char
    MarginMaintain  [9]types.Char
    MarginInitial   [9]types.Char
}

//动态价格  90
type DynamicPrice struct {
    ProdId      [40]types.Char
    Margin      [12]types.Char
    Filler      [38]types.Char
}