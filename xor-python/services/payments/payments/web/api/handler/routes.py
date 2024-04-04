import os

from fastapi import APIRouter
from yookassa import Configuration, Payment, Payout
from yookassa.domain.response import PaymentResponse, PayoutResponse

from payments.web.api.handler.models import (
    CreatePayoutRequestPayload,
    CreatePurchasePayload,
    GetStatusPayload,
    GetStatusResponse,
    Message,
)

router = APIRouter()

Configuration.account_id = os.getenv('YOOKASSA_ACCOUNT_ID')
Configuration.secret_key = os.environ.get('YOOKASSA_KEY')
PAYMENT_RETURN_URL = os.environ.get('PAYMENT_RETURN_URL')


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
    data: CreatePurchasePayload,
) -> PaymentResponse:
    items = []
    for product in data.products:
        items.append(
            {
                "description": product.description,
                "quantity": product.quantity,
                "amount": {
                    "value": product.money,
                    "currency": product.currency
                },
                "payment_mode": product.payment_mode,
                # "supplier": {
                #     "name": "string",
                #     "phone": "string",
                #     "inn": "string"
                # }
            }
        )

    res: PaymentResponse = Payment.create(
        {
            "amount": {
                "value": data.money,
                "currency": data.currency
            },
            "confirmation": {
                "type": "redirect",
                "return_url": PAYMENT_RETURN_URL
            },
            "capture": True,
            "description": data.payment_name,
            "metadata": {
                'orderNumber': data.payment_uuid
            },
            "receipt": {
                "customer": {
                    "full_name": data.full_name,
                    "email": data.email,
                    "phone": data.phone
                },
                "items": items
            }
        }
    )

    return res


@router.post("/payout", response_model=Message)
async def create_payout_request(
    data: CreatePayoutRequestPayload,
) -> PayoutResponse:
    res = Payout.create(
        params={
            "id": data.payment_uuid,
            "amount": {
                "value": data.money,
                "currency": data.currency
            },
            # "status": "succeeded",
            "payout_destination": {
                "type": "bank_card",
                "card": {
                    "first6": data.card_info.first6,
                    "last4": data.card_info.last4,
                    "card_type": data.card_info.card_type,
                    "issuer_country": data.card_info.issuer_country,
                    "issuer_name": data.card_info.issuer_name,
                }
            },
            "description": data.payment_name,
            "created_at": "2021-06-21T16:22:50.512Z",
            # "metadata": {
            #     "order_id": "37"
            # },
            "test": data.is_test
        }
    )
    return res
