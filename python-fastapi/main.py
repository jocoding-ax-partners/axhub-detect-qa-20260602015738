import os
from fastapi import FastAPI

app = FastAPI()

@app.get('/')
def root():
    return {'language': 'python', 'message': os.environ.get('APPHUB_QA_PYTHON_MESSAGE', 'python ok')}
