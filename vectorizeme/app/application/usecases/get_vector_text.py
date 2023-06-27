from app.domain.ml_models import VectorizerInterface
from app.domain.models import Content

from typing import List


class GetVectorText:
    def __init__(self, vectorizer: VectorizerInterface):
        self.vectorizer = vectorizer

    async def execute(self, text: str) -> Content:
        vector: List[float] = await self.vectorizer.text_to_vector(text=text)

        return Content(
            vector=vector,
            text=text
        )
