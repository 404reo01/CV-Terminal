# Déploiement sur VPS

## Ce que tu envoies à ton pote

Donne-lui l'archive du projet (sans `.env` et `.ssh/`), ou pousse sur GitHub.

## Sur le VPS (commandes à faire une seule fois)

```bash
# 1. Cloner / copier le projet
git clone <ton-repo> cv-terminal
cd cv-terminal

# 2. Générer la clé SSH hôte
mkdir -p .ssh
ssh-keygen -t ed25519 -f .ssh/id_ed25519 -N ""

# 3. Builder et lancer
docker build -t cv-terminal .
docker run -d \
  --name cv-terminal \
  --restart unless-stopped \
  -p 22:22 \
  -v $(pwd)/.ssh:/app/.ssh \
  cv-terminal
```

## Résultat

```bash
ssh ton-domaine.com
# ou
ssh IP_DU_VPS
```

## Si le port 22 est déjà pris par le SSH du VPS

Utilise le port 2222 :
```bash
docker run -d \
  --name cv-terminal \
  --restart unless-stopped \
  -p 2222:22 \
  -v $(pwd)/.ssh:/app/.ssh \
  cv-terminal
```
Connexion : `ssh -p 2222 ton-domaine.com`
