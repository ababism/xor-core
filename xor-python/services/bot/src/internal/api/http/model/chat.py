from pydantic import BaseModel


class CreateInviteLinkResponse(BaseModel):
    invite_link: str
