import constants    
import httpx
from schemas import StockSummary as StockSummarySchema
from pydantic import TypeAdapter
from repositories import    StockSummaryRepostiory
from database import get_db


def fetch_tickers():
    params = {
        "length": 9999,
        "start": 0,
    }
    
    with httpx.Client(
        headers=constants.HEADERS,
        timeout=30,
        follow_redirects=True,
    ) as client:
        response = client.get(constants.ENDPOINTS[constants.IDX_TICKERS], params=params)
        response.raise_for_status()

        return response.json()

def run():
    res = fetch_tickers()
    data = TypeAdapter(list[StockSummarySchema]).validate_python(res['data'])
    with get_db() as db:
        repo = StockSummaryRepostiory(db)
        repo.update(data)
        repo.update_ticker(data)

    