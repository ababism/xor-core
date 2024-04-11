from pydantic import BaseModel, Field


class HttpConfig(BaseModel):
    host: str
    port: int


class TelegramBotConfig(BaseModel):
    token: str


class PostgresConfig(BaseModel):
    host: str
    port: str
    dbname: str
    user: str
    password: str


class Config(BaseModel):
    http_config: HttpConfig = Field(alias="http")
    telegram_bot_config: TelegramBotConfig = Field(alias="telegram-bot")
    postgres_config: PostgresConfig = Field(alias="postgres")
