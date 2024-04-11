from pydantic import BaseModel


class SendMessageRequest(BaseModel):
    chat_id: int
    message: str


class BlockUserRequest(BaseModel):
    chat_id: int
    user_id: int
