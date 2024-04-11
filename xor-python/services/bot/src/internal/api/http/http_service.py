from typing import List

import uvicorn
from fastapi import FastAPI

from src.internal.api.http.router.base_router import BaseRouter
from src.internal.api.http.router.chat_router import ChatRouter
from src.internal.api.telegram_bot_service.telegram_bot_service import TelegramBotService
from src.internal.config.config import HttpConfig


class HttpService:
    def __init__(self,
                 config: HttpConfig,
                 telegram_bot_service: TelegramBotService):
        self._config = config
        self._app = FastAPI()
        self._telegram_bot_service = telegram_bot_service

    async def run(self):
        routers = self._get_routers()
        self._mount_routers(routers)

        config = uvicorn.Config(self._app,
                                host=self._config.host,
                                port=self._config.port,
                                log_level="info")
        server = uvicorn.Server(config)
        await server.serve()

    def _get_routers(self) -> List[BaseRouter]:
        return [
            ChatRouter(self._telegram_bot_service)
        ]

    def _mount_routers(self, routers: List[BaseRouter]):
        for router in routers:
            router.mount_handlers()
            self._app.include_router(router.internal_router)
