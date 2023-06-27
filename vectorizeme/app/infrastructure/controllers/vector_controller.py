from fastapi import APIRouter, Depends
from app.domain.models import Content
from app.application.usecases import GetVectorUrl, GetVectorText
from app.infrastructure.ml_models import Vectorizer
from app.infrastructure.clients import ImageClient
from app.infrastructure.controllers.presenters import (
    UrlVectorizerPresenter,
    TextVectorizerPresenter,
)
import urllib.parse

vectorizer_router = APIRouter()


def setup_vector_url_usecase():
    return GetVectorUrl(vectorizer=Vectorizer(), image_client=ImageClient())


@vectorizer_router.get("/vectorize/image-url/", tags=["vectorize"])
async def get_vector_from_image_url(
    url: str,
    get_vector_url_usecase: GetVectorUrl = Depends(setup_vector_url_usecase),
):
    url = urllib.parse.unquote(url)
    content: Content = await get_vector_url_usecase.execute(url=url)
    return UrlVectorizerPresenter(content=content)


def setup_vector_text_usecase():
    return GetVectorText(vectorizer=Vectorizer())


@vectorizer_router.get("/vectorize/text/", tags=["vectorize"])
async def get_vector_from_text(
    text: str,
    get_vector_text_usecase: GetVectorText = Depends(
        setup_vector_text_usecase),
):
    content: Content = await get_vector_text_usecase.execute(text=text)
    return TextVectorizerPresenter(content=content)
