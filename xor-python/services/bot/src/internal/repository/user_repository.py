from abc import ABC, abstractmethod
from typing import Optional

from src.internal.domain.chat import User
from src.package.postgres.postgres_client import PostgresClient


class UserRepository(ABC):
    @abstractmethod
    def get(self, user_id: int) -> Optional[User]:
        raise NotImplementedError()

    @abstractmethod
    def create(self, user: User):
        raise NotImplementedError()


class UserPostgresRepository(UserRepository):
    def __init__(self, postgres_client: PostgresClient):
        self._postgres_client = postgres_client

    def get(self, user_id: int) -> Optional[User]:
        query = """
        SELECT id, username FROM telegram_user WHERE id = %s
        """
        user = self._postgres_client.execute(query, (user_id,))
        if not user:
            return None
        return User.from_postgres(user)

    def create(self, user: User):
        query = """
        INSERT INTO telegram_user (id, username) VALUES (%s, %s) RETURNING 1;
        """
        self._postgres_client.execute_many(query, (user.id, user.username))
