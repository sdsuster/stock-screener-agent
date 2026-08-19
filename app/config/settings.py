from singleton.llm import LLMContainer
import os
from dotenv import load_dotenv
from langchain.chat_models import init_chat_model

# Load variables from .env file into os.environ
load_dotenv()




def setup():
    # LLMContainer().register("llm", init_chat_model("google_genai:gemini-3.7-flash"))
    print('hei')
    model = init_chat_model(
        "google_genai:gemini-3.5-flash"
    )

    response = model.invoke("Why do parrots talk?")

    print(response.content)    