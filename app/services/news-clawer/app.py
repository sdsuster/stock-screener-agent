import constants    
import httpx
from supabase import Client

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

def run(supabase: Client):
    fetch_tickers()

run(None)