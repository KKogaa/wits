from typing import List
from app.domain.models import Content


class TextVectorizerPresenter:
    vector: List[float]
    text: str

    def __init__(self, content: Content):
        self.vector = content.vector
        self.text = content.text


class UrlVectorizerPresenter:
    vector: List[float]
    url: str

    def __init__(self, content: Content):
        self.vector = content.vector
        self.url = content.url
