from app.domain.ml_models import VectorizerInterface
from app.domain.models import Content
from app.domain.clients import ImageClientInterface

from typing import List


class GetVectorUrl:
    def __init__(
        self,
        vectorizer: VectorizerInterface,
        image_client: ImageClientInterface,
    ):
        self.vectorizer = vectorizer
        self.image_client = image_client

    async def execute(self, url: str) -> Content:
        image: any = await self.image_client.obtain_image(url=url)
        vector: List[float] = await self.vectorizer.image_to_vector(image=image)

        content = Content()
        content.url = url
        content.vector = vector

        return content
