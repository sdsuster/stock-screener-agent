create table public.ticker (
    stock_code text primary key,
    stock_name text not null,

    date timestamptz not null,
    remarks text,

    previous numeric,
    open_price numeric,
    first_trade numeric,
    high numeric,
    low numeric,
    close numeric,
    change numeric,

    volume bigint,
    value numeric,
    frequency bigint,

    index_individual numeric,

    offer numeric,
    offer_volume bigint,
    bid numeric,
    bid_volume bigint,

    listed_shares bigint,
    tradeable_shares bigint,
    weight_for_index numeric,

    foreign_sell bigint,
    foreign_buy bigint,

    delisting_date timestamptz,

    non_regular_volume bigint,
    non_regular_value numeric,
    non_regular_frequency bigint,

    persen numeric,
    percentage numeric,

    idx_stock_summary_id bigint,

    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);