from models import StockSummary, Stock
from schemas import StockSummary as StockSummarySchema
from sqlalchemy.orm import Session
from sqlalchemy import inspect
from sqlalchemy.dialects.postgresql import insert
from base import update_fields

class StockSummaryRepostiory():
    def __init__(self, session: Session):
        self.session = session

    def update(self, data: list[StockSummarySchema]):
        # orm_data = [
        #     StockSummary(**d.model_dump(exclude={'stock_name'}))
        #     for d in data
        # ]

        for d in data:
            f = self.session.get(StockSummary, d.stock_code)
            if f is not None:
                update_fields(f, d, exclude={'stock_name'})
            else:
                # create
                self.session.add(StockSummary(**d.model_dump(exclude={'stock_name'})))

        self.session.commit()

    def update_ticker(self, data: list[StockSummarySchema]):
        values = [
            {
                "stock_code": item.stock_code,
                "stock_name": item.stock_name,
            }
            for item in data
        ]

        stmt = insert(Stock).values(values)

        stmt = stmt.on_conflict_do_update(
            index_elements=[Stock.stock_code],
            set_={
                "stock_name": stmt.excluded.stock_name,
            },
        )

        self.session.execute(stmt)
        