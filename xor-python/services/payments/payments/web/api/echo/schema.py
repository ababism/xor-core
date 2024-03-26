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


class CreatePurchaseRequestPayload(BaseModel):
    """Create Purchase Request Payload model."""
    money: str


class CreatePayoutRequestPayload(BaseModel):
    """Create Payout Request Payload model."""
    money: str
