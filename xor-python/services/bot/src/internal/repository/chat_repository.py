from abc import ABC, abstractmethod
from typing import Optional

from src.internal.domain.chat import Chat
from src.package.postgres.postgres_client import PostgresClient


class ChatRepository(ABC):
    @abstractmethod
    def get(self, chat_id: int) -> Optional[Chat]:
        raise NotImplementedError()

    @abstractmethod
    def create(self, chat: Chat):
        raise NotImplementedError()


class ChatPostgresRepository(ChatRepository):
    def __init__(self, postgres_client: PostgresClient):
        self._postgres_client = postgres_client

    def get(self, chat_id: int) -> Optional[Chat]:
        query = """
        SELECT id, title, invite_link FROM telegram_chat WHERE id = %s
        """
        chat = self._postgres_client.execute(query, (chat_id,))
        if not chat:
            return None
        return Chat.from_postgres(chat)

    def create(self, chat: Chat):
        query = """
        INSERT INTO telegram_chat (id, title, invite_link) VALUES (%s, %s, %s) RETURNING 1
        """
        self._postgres_client.execute(query, (chat.id, chat.title, chat.invite_link))
