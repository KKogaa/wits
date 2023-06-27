from abc import ABC, abstractmethod


class ImageClientInterface(ABC):
    @abstractmethod
    async def obtain_image(self, url: str):
        raise NotImplementedError
