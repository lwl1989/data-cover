package data

type X byte
type N9 byte
type Char byte
//基本档案 113
type GoodsBase struct {
    GoodsId [20]X
    RiseLimitPrice1 [9]N9
    PeferencePrice [9]N9
    FallLimitPrice1 [9]N9
    RiseLimitPrice2 [9]N9
    FallLimitPrice2 [9]N9
    RiseLimitPrice3 [9]N9
    FallLimitPrice3 [9]N9

    ProdKind        [1]X
    DecimalLocator  [1]N9
    StrikePriceDL   [1]N9

    BeginDate       [8]N9
    EndDate         [8]N9
    FlowGroup       [2]N9
    DeliveryDate    [8]N9
    DynamicBanding  [1]X
}

//股票参照档案  24
type StockSelectRefer struct {
    IndexKind   [8]X
    IndexNumber [6]X
    IndexValue  [5]N9
    IndexStatus [1]X
}

//股票调整档案  154
type StockSelectChange struct {
    BaseDate        [8]N9

    BfKindId        [8]X
    BfStockId       [6]X
    BfStockQnTy     [6]N9
    BfStockCash2    [8]N9
    BfStockCash3    [6]N9
    BfStockId4      [6]X
    BfStockQnTy4    [6]N9

    AfKindId        [8]X
    AfStockId       [6]X
    AfStockQnTy     [6]N9
    AfStockCash2    [8]N9
    AFStockPrice3   [6]N9
    AfStockQnTy3    [6]N9
    AfStockDate3    [6]N9
    AfStockId4      [6]X
    AfStockQnTy4    [6]N9

    DividendDate    [8]N9
}

//契约基本档案  75
type ConTractBase struct {
    KindId  [8]X
    Name    [30]X
    StockId [6]X
    SubType [1]X
    ConTractSize    [7]N9

    StatusCode      [1]X
    CurrencyType    [1]X

    DecimalLocator  [1]N9
    StrikePriceDL   [1]N9

    AcceptQuoteFlag [1]X
    BeginDate       [8]X
    BlockTradeFlag  [1]X
    ExpiryType      [1]X
    UnderLyingType  [1]X
    MarketCloseGroup    [2]N9

    EndSession      [1]X
}

//契约风险保证  38
type ConTractCafe struct {
    KindId          [8]Char

    AbType          Char
    CurrencyType    Char

    Margin          [9]Char
    MarginMaintain  [9]Char
    MarginInitial   [9]Char
}

//动态价格  90
type DynamicPrice struct {
    ProdId      [40]Char
    Margin      [12]Char
    Filler      [38]Char
}