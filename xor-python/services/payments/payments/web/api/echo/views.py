from fastapi import APIRouter

from payments.web.api.echo.schema import (Message, GetStatusPayload, GetStatusResponse,
                                          CreatePurchaseRequestPayload,
                                          CreatePayoutRequestPayload)

router = APIRouter()


@router.post("/", response_model=Message)
async def send_echo_message(
    incoming_message: Message,
) -> Message:
    """
    Sends echo back to user.

    :param incoming_message: incoming message.
    :returns: message same as the incoming.
    """
    return incoming_message


@router.get("/status", response_model=GetStatusResponse)
async def get_payment_status(
    data: GetStatusPayload,
) -> GetStatusResponse:
    return GetStatusResponse(id=data.payment_id, status="completed")


@router.post("/status", response_model=Message)
async def accept_webhook(
    data: GetStatusPayload,
) -> Message:
    return Message(message="200 ok")


@router.post("/purchase", response_model=Message)
async def create_purchase_request(
    data: CreatePurchaseRequestPayload,
) -> Message:
    return Message(message="201 ok")


@router.post("/payout", response_model=Message)
async def create_payout_request(
    data: CreatePayoutRequestPayload,
) -> Message:
    return Message(message="201 ok")

