from __future__ import annotations

from datetime import date, datetime
from decimal import Decimal

from sqlalchemy import (
    Date,
    DateTime,
    ForeignKey,
    Integer,
    Numeric,
    String,
    BigInteger,
    Text,
    UniqueConstraint,
    Float
)
from sqlalchemy.orm import Mapped, mapped_column, relationship

from database import Base


class Stock(Base):
    __tablename__ = "stocks"

    stock_code: Mapped[str] = mapped_column(
        String(20),
        primary_key=True,
    )

    stock_name: Mapped[str] = mapped_column(
        String(255),
        nullable=False,
    )

    summaries: Mapped[list["StockSummary"]] = relationship(
        back_populates="stock",
        cascade="all, delete-orphan",
    )
class StockSummary(Base):
    __tablename__ = "stock_summaries"

    __table_args__ = (
        UniqueConstraint(
            "stock_code",
            "date",
            name="uq_stock_summary_stock_code_date",
        ),
    )

    id_stock_summary: Mapped[int] = mapped_column(
        Integer,
    )

    no: Mapped[int] = mapped_column(
        Integer,
        nullable=False,
    )

    stock_code: Mapped[str] = mapped_column(
        String(20),
        ForeignKey("stocks.stock_code"),
        nullable=False,
        primary_key=True,
        index=True,
    )

    date: Mapped[datetime] = mapped_column(
        DateTime,
        nullable=False,
    )

    remarks: Mapped[str | None] = mapped_column(
        Text,
        nullable=True,
    )

    previous: Mapped[Decimal] = mapped_column(
        Numeric(18, 4),
        nullable=False,
    )

    open_price: Mapped[Decimal] = mapped_column(
        Numeric(18, 4),
        nullable=False,
    )

    first_trade: Mapped[Decimal] = mapped_column(
        Numeric(18, 4),
        nullable=False,
    )

    high: Mapped[Decimal] = mapped_column(
        Numeric(18, 4),
        nullable=False,
    )

    low: Mapped[Decimal] = mapped_column(
        Numeric(18, 4),
        nullable=False,
    )

    close: Mapped[Decimal] = mapped_column(
        Numeric(18, 4),
        nullable=False,
    )

    change: Mapped[Decimal] = mapped_column(
        Numeric(18, 4),
        nullable=False,
    )

    volume: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    value: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    frequency: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    index_individual: Mapped[Decimal] = mapped_column(
        Numeric(18, 6),
        nullable=False,
    )

    offer: Mapped[Decimal] = mapped_column(
        Numeric(18, 4),
        nullable=False,
    )

    offer_volume: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    bid: Mapped[Decimal] = mapped_column(
        Numeric(18, 4),
        nullable=False,
    )

    bid_volume: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    listed_shares: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    tradeable_shares: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    weight_for_index: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    foreign_sell: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    foreign_buy: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    delisting_date: Mapped[date | None] = mapped_column(
        Date,
        nullable=True,
    )

    non_regular_volume: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    non_regular_value: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    non_regular_frequency: Mapped[int] = mapped_column(
        BigInteger,
        nullable=False,
    )

    persen: Mapped[float | None] = mapped_column(
        Float,
        nullable=True,
    )

    percentage: Mapped[float | None] = mapped_column(
        Float,
        nullable=True,
    )

    stock: Mapped["Stock"] = relationship(
        back_populates="summaries",
    )