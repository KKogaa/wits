from app.domain.clients import ImageClientInterface

import requests
from fastapi import HTTPException


class ImageClient(ImageClientInterface):
    async def obtain_image(self, url: str) -> any:
        try:
            image = requests.get(url=url, stream=True).raw
            return image
        except Exception:
            raise HTTPException(status_code=400, detail="Image not found")
