from fastapi import APIRouter

from src.internal.api.http.model.chat import CreateInviteLinkResponse
from src.internal.api.http.model.user import SendMessageRequest, BlockUserRequest
from src.internal.api.http.router.base_router import BaseRouter
from src.internal.api.telegram_bot_service.telegram_bot_service import TelegramBotService


class ChatRouter(BaseRouter):
    def __init__(self, telegram_bot_service: TelegramBotService):
        super().__init__(router=APIRouter(prefix="/chat", tags=["chat"]))
        self._telegram_bot_service = telegram_bot_service

    def mount_handlers(self):
        self.internal_router.post("/{chat_id}/invite-link")(self.create_invite_link)
        self.internal_router.post("/send-message")(self.send_message)
        self.internal_router.post("/block-user")(self.block_user)

    async def create_invite_link(self, chat_id: int) -> CreateInviteLinkResponse:
        invite_link = await self._telegram_bot_service.create_invite_link(chat_id=chat_id)
        return CreateInviteLinkResponse(invite_link=invite_link)

    async def send_message(self, request: SendMessageRequest):
        await self._telegram_bot_service.send_message(request.chat_id, request.message)

    async def block_user(self, request: BlockUserRequest):
        await self._telegram_bot_service.block_user(request.chat_id, request.user_id)
