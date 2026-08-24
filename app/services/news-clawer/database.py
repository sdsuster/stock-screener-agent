from pathlib import Path
import os

from dotenv import load_dotenv
from sqlalchemy import create_engine
from sqlalchemy.orm import DeclarativeBase, sessionmaker
from contextlib import contextmanager


BASE_DIR = Path(__file__).resolve().parent.parent
load_dotenv()

DATABASE_URL = os.environ["DATABASE_URL"]


class Base(DeclarativeBase):
    pass


engine = create_engine(
    DATABASE_URL,
    pool_pre_ping=True,
)


SessionLocal = sessionmaker(
    bind=engine,
    autoflush=False,
    autocommit=False,
)


@contextmanager
def get_db():
    db = SessionLocal()

    try:
        yield db
    finally:
        db.close()