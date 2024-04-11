from typing import Optional

from aiogram.types import Message, Chat as TelegramChat
from pydantic import BaseModel


class User(BaseModel):
    id: int
    username: str

    @staticmethod
    def from_postgres(row: tuple):
        return User(
            id=row[0],
            username=row[1]
        )

    #
    @staticmethod
    def from_telegram_message(message: Message):
        return User(id=message.from_user.id,
                    username=message.from_user.username)


class Chat(BaseModel):
    id: int
    title: Optional[str]
    invite_link: Optional[str]

    @staticmethod
    def from_postgres(row: tuple):
        return Chat(
            id=row[0],
            title=row[1],
            invite_link=row[2]
        )

    @staticmethod
    def from_telegram_chat(chat: TelegramChat):
        return Chat(id=chat.id,
                    title=chat.title,
                    invite_link=chat.invite_link)
