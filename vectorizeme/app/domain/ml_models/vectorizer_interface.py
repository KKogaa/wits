from abc import ABC, abstractmethod


class VectorizerInterface(ABC):
    @abstractmethod
    async def image_to_vector(self, image):
        raise NotImplementedError

    @abstractmethod
    async def text_to_vector(self, text: str):
        raise NotImplementedError
