from singleton.llm import LLMContainer
from config.constants import LLM_DEFAULT_NAME

class StockScreenerAgent:
    def __init__(self):
        self.graph = self.build_graph()
        self.llm = LLMContainer().get(LLM_DEFAULT_NAME)
        pass

    def build_graph(self):
        pass

    def screen_stocks(self, criteria):
        # Implement stock screening logic based on the provided criteria
        pass