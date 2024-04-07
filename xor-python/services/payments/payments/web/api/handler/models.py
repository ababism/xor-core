from typing import List

from pydantic import BaseModel


class Message(BaseModel):
    """Simple message model."""

    message: str


class GetStatusPayload(BaseModel):
    """Get Status Payload model."""

    payment_id: str


class GetStatusResponse(BaseModel):
    """Get Status Response model."""

    id: str
    status: str


class CreatePurchaseProductPayload(BaseModel):
    description: str
    quantity: int
    money: float
    currency: str = "RUB"
    payment_mode: str = "full_payment"


class CreatePurchasePayload(BaseModel):
    """Create Purchase Request Payload model."""
    payment_uuid: str
    payment_name: str
    money: float
    currency: str = "RUB"
    full_name: str
    phone: str
    email: str
    products: List[CreatePurchaseProductPayload]


class CardInfo(BaseModel):
    first6: str
    last4: str
    card_type: str = "MIR"
    issuer_country: str = "RU"
    issuer_name: str


class CreatePayoutRequestPayload(BaseModel):
    """Create Payout Request Payload model."""
    payment_uuid: str
    payment_name: str
    money: str
    currency: str = "RUB"
    full_name: str
    phone: str = ""
    email: str = ""
    card_info: CardInfo
    is_test: bool = False
