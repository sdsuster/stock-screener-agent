from __future__ import annotations
from langchain.chat_models import BaseChatModel
from typing import  Optional

class LLMContainer:
    _instance: Optional[LLMContainer] = None
    _llms: dict[str, BaseChatModel]

    def __new__(cls, *args, **kwargs):
        if cls._instance is None:
            cls._instance = super(LLMContainer, cls).__new__(cls)
            cls._instance._llms = {}
        return cls._instance

    def __init__(self):
        if not hasattr(self, '_llms'):
            self._llms = {}

    def register(self, name: str, llm: BaseChatModel):
        self._llms[name] = llm

    def get(self, name: str) -> BaseChatModel:
        return self._llms[name]