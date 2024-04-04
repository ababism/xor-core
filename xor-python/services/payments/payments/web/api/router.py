from fastapi.routing import APIRouter

from payments.web.api import docs, handler, monitoring

api_router = APIRouter()
api_router.include_router(monitoring.router)
api_router.include_router(docs.router)
api_router.include_router(handler.router, prefix="/api/v1")
