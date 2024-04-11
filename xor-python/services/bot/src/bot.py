import asyncio
import logging
from typing import Optional

from src.internal.api.http.http_service import HttpService
from src.internal.api.telegram_bot_service.telegram_bot_service import TelegramBotService
from src.internal.config.config import Config
from src.internal.repository.chat_repository import ChatPostgresRepository
from src.internal.repository.user_repository import UserPostgresRepository
from src.package.postgres.postgres_client import PostgresClient

logger = logging.getLogger(__name__)


class Bot:
    def __init__(self, config: Config):
        self._config = config
        self._http_service: Optional[HttpService] = None
        self._telegram_bot_service: Optional[TelegramBotService] = None

    def build(self):
        postgres_client = PostgresClient(self._config.postgres_config)
        chat_repository = ChatPostgresRepository(postgres_client)
        user_repository = UserPostgresRepository(postgres_client)

        self._telegram_bot_service = TelegramBotService(self._config.telegram_bot_config,
                                                        chat_repository,
                                                        user_repository)

        self._http_service = HttpService(self._config.http_config,
                                         self._telegram_bot_service)

        return self

    def run(self):
        loop = asyncio.get_event_loop()

        asyncio.gather(
            self._telegram_bot_service.run(),
            self._http_service.run()
        )

        loop.run_forever()
