from fastapi import FastAPI, Depends
from api.routes import router
from config.settings import Settings, get_settings

app = FastAPI(title="Language Service")

@app.get("/health")
def health_check(settings: Settings = Depends(get_settings)):
    return {"status": "OK", "env": settings.env}

app.include_router(router)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8080, reload=True)
