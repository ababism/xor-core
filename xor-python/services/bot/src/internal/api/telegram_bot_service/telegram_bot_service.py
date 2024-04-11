from aiogram import Bot, Dispatcher
from aiogram.types import Message

from src.internal.config.config import TelegramBotConfig
from src.internal.domain.chat import Chat, User
from src.internal.repository.chat_repository import ChatRepository
from src.internal.repository.user_repository import UserRepository

FOXLERENJOB_ID = 6564949176


class TelegramBotService:
    def __init__(self,
                 config: TelegramBotConfig,
                 chat_repository: ChatRepository,
                 user_repository: UserRepository):
        self._bot = Bot(config.token)
        self._dispatcher = Dispatcher(self._bot)
        self._chat_repository = chat_repository
        self._user_repository = user_repository

    async def run(self):
        self.register_handlers()
        await self._dispatcher.start_polling()

    def register_handlers(self):
        self._dispatcher.register_message_handler(self.start, commands=["start"], run_task=True)

    async def create_invite_link(self, chat_id: int) -> str:
        invite_link = await self._bot.create_chat_invite_link(chat_id)
        return invite_link.invite_link

    async def send_message(self, chat_id: int, message: str):
        await self._bot.send_message(chat_id=chat_id, text=message)

    async def block_user(self, chat_id: int, user_id):
        await self._bot.ban_chat_member(chat_id=chat_id,
                                        user_id=user_id)

    # commands
    async def start(self, message: Message):
        try:
            repo_user = self._user_repository.get(message.from_user.id)
            if not repo_user:
                self._user_repository.create(User.from_telegram_message(message))

            repo_chat = self._chat_repository.get(message.chat.id)
            if not repo_chat:
                self._chat_repository.create(Chat.from_telegram_chat(message.chat))

            await self._bot.unban_chat_member(chat_id=repo_chat.id,
                                              user_id=FOXLERENJOB_ID)
            await self._bot.send_message(message.chat.id, text="TEST")
        except Exception as e:
            print(e)
            await self._bot.send_message(message.chat.id, text="FAILED")
