from datetime import datetime

from pydantic import BaseModel, Field, ConfigDict, field_validator

class Stock(BaseModel):
    stock_code: str = Field(alias="StockCode")
    stock_name: str = Field(alias="StockName")



class StockSummary(BaseModel):
    model_config = ConfigDict(
        populate_by_name=True,
        extra="ignore"
    )

    no: int = Field(alias="No")
    id_stock_summary: int = Field(alias="IDStockSummary")
    date: datetime = Field(alias="Date")

    stock_code: str = Field(alias="StockCode")
    stock_name: str = Field(alias="StockName")
    remarks: str = Field(alias="Remarks")

    previous: int = Field(alias="Previous")
    open_price: int = Field(alias="OpenPrice")
    first_trade: int = Field(alias="FirstTrade")

    high: int = Field(alias="High")
    low: int = Field(alias="Low")
    close: int = Field(alias="Close")
    change: int = Field(alias="Change")

    volume: int = Field(alias="Volume")
    value: int = Field(alias="Value")
    frequency: int = Field(alias="Frequency")

    index_individual: float = Field(alias="IndexIndividual")

    offer: int = Field(alias="Offer")
    offer_volume: int = Field(alias="OfferVolume")

    bid: int = Field(alias="Bid")
    bid_volume: int = Field(alias="BidVolume")

    listed_shares: int = Field(alias="ListedShares")
    tradeable_shares: int = Field(alias="TradebleShares")
    weight_for_index: int = Field(alias="WeightForIndex")

    foreign_sell: int = Field(alias="ForeignSell")
    foreign_buy: int = Field(alias="ForeignBuy")

    delisting_date: datetime | None = Field(alias="DelistingDate")

    non_regular_volume: int = Field(alias="NonRegularVolume")
    non_regular_value: int = Field(alias="NonRegularValue")
    non_regular_frequency: int = Field(alias="NonRegularFrequency")

    persen: float | None = None
    percentage: float | None = None


    @field_validator("delisting_date", mode="before")
    @classmethod
    def empty_date_to_none(cls, value):
        if value == "":
            return None
        return value