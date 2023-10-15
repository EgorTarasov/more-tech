from pydantic_settings import BaseSettings, SettingsConfigDict
from pydantic import Field


class Config(BaseSettings):
    model_config = SettingsConfigDict(
        env_file=".env",
    )

    mongo_uri: str = Field(..., alias="MONGO_URI")
    mongo_db: str = Field(..., alias="MONGO_DB")


CONFIG = Config(_env_file=".env")
