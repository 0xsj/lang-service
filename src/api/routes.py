from fastapi import APIRouter, Depends
from services.translation import translate_text
from config.settings import Settings, get_settings

router = APIRouter()

@router.get("/")
def home(settings: Settings = Depends(get_settings)):
    return {"message": f"{settings.app_name} is running in {settings.env} mode"}

@router.post("/translate")
def translate(text: str, source_lang: str, target_lang: str):
    return {"translated_text": translate_text(text, source_lang, target_lang)}

