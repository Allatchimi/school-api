#!/bin/sh

echo "📁 Contenu du dossier /app :"
ls -al /app

echo "📄 Contenu de /app/app.env (si présent) :"
if [ -f "/app/app.env" ]; then
  cat /app/app.env
else
  echo "⚠️  /app/app.env introuvable!"
fi

echo "📄 Contenu de /app/assets/private/keys/jwt/private.pem (si présent) :"
if [ -f "/app/assets/private/keys/jwt/private.pem" ]; then
  cat /app/assets/private/keys/jwt/private.pem
else
  echo "⚠️  /app/assets/private/keys/jwt/private.pem introuvable!"
fi

echo "📄 Contenu de /app/assets/private/keys/jwt/public.pem (si présent) :"
if [ -f "/app/assets/private/keys/jwt/public.pem" ]; then
  cat /app/assets/private/keys/jwt/public.pem
else
  echo "⚠️  /app/assets/private/keys/jwt/public.pem introuvable!"
fi

echo "✅ En attente..."
tail -f /dev/null