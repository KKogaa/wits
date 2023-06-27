from typing import List
from app.domain.ml_models import VectorizerInterface
from PIL import Image
from transformers.models.auto.processing_auto import AutoProcessor
from transformers.models.clip.modeling_clip import CLIPModel
from fastapi import HTTPException


class Vectorizer(VectorizerInterface):
    def __init__(self):
        self.model = CLIPModel.from_pretrained("openai/clip-vit-base-patch32")
        self.processor = AutoProcessor.from_pretrained(
            "openai/clip-vit-base-patch32"
        )

    async def image_to_vector(self, image) -> List[float]:
        try:
            image = Image.open(image)
            inputs = self.processor(
                images=image, return_tensors="pt", padding=True
            )
            image_features= self.model.get_image_features(**inputs)

            return image_features.flatten().tolist()
        except Exception as e:
            raise HTTPException(
                    status_code=500, detail=f"Failed to process image: {e}"
            )

    async def text_to_vector(self, text: str):
        try:
            inputs = self.processor(
                text=[text], return_tensors="pt", padding=True
            )
            text_features = self.model.get_text_features(**inputs)

            return text_features.flatten().tolist()
        except Exception as e:
            raise HTTPException(
                    status_code=500, detail=f"Failed to process text: {e}"
            )
