from fastapi import FastAPI

from app.infrastructure.config.settings import Settings

from app.infrastructure.controllers.vector_controller import vectorizer_router

settings = Settings()

app = FastAPI(title=settings.PROJECT_NAME)

app.include_router(router=vectorizer_router)
