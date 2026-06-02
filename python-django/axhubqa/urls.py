from django.http import JsonResponse
from django.urls import path

urlpatterns = [path("", lambda request: JsonResponse({"ok": True}))]
