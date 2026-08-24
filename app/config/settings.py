from singleton.llm import LLMContainer
from .constants import LLM_DEFAULT_NAME
import os
from dotenv import load_dotenv
from langchain.chat_models import init_chat_model

# Load variables from .env file into os.environ
load_dotenv()

def setup_llm():
    LLMContainer().register(LLM_DEFAULT_NAME, init_chat_model(
        "google_genai:gemini-3.5-flash"
    ))

def setup():
    setup_llm()