from pydantic import BaseSettings


class Settings(BaseSettings):
    PROJECT_NAME: str = "vectorizeme"

    class Config:
        env_file = ".env"
        case_sensitive = True
