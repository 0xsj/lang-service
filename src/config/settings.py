from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    app_name: str = "lang-service"
    env: str = "development"
    debug: bool = True

    class Config:
        env_file = ".env"

def get_settings() -> Settings:
    """Dependecy Injection for settings"""
    return Settings()