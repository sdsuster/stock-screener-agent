IDX_TICKERS = "IDX/TICKERS"
HEADERS = {
    "User-Agent": (
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
        "AppleWebKit/537.36 (KHTML, like Gecko) "
        "Chrome/139.0.0.0 Safari/537.36"
    ),
    "Accept": "application/json, text/javascript, */*; q=0.01",
    "Accept-Language": "en-US,en;q=0.9",
    "Referer": "https://www.idx.co.id/",
    "Origin": "https://www.idx.co.id",
    "X-Requested-With": "XMLHttpRequest",
}


ENDPOINTS = {
    IDX_TICKERS: "https://www.idx.co.id/primary/TradingSummary/GetStockSummary"
}
