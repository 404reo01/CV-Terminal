# Le CV de Reo

**Mon CV ne se télécharge pas, il se visite.**

```bash
ssh reo-ssh.duckdns.org
```

Pas de compte, pas de mot de passe : ouvre un terminal, colle la commande, et tu navigues dans mon CV au clavier. Avec de la neige et un renard en ASCII.

<!-- Capture ou GIF à ajouter ici -->

## Navigation

| Touche | Action |
|---|---|
| `↑` `↓` ou `k` `j` | se déplacer / défiler |
| `entrée` | ouvrir une section |
| `esc` ou `b` | revenir au menu |
| `q` | quitter |

## Sous le capot

- **Go**, avec [Bubbletea](https://github.com/charmbracelet/bubbletea) pour l'interface et [Lip Gloss](https://github.com/charmbracelet/lipgloss) pour le style
- **[Wish](https://github.com/charmbracelet/wish)** : un serveur SSH qui sert une application TUI au lieu d'un shell
- **Docker** sur un VPS, avec un DNS DuckDNS

```
cmd/server/main.go        serveur SSH (Wish + Bubbletea)
internal/data/resume.go   tout le contenu du CV
internal/tui/             modèle, vues, animation de neige
internal/ascii/           le renard
```

L'animation se met en pause après 15 s d'inactivité pour ne pas consommer de CPU pour rien.

## Le lancer chez toi

```bash
git clone https://github.com/404reo01/CV-Terminal.git && cd CV-Terminal
mkdir -p .ssh && ssh-keygen -t ed25519 -f .ssh/id_ed25519 -N ""
docker build -t cv-terminal .
docker run -d --name cv-terminal -p 2222:22 -v $(pwd)/.ssh:/app/.ssh cv-terminal
ssh -p 2222 localhost
```

Pour en faire ton propre CV, il suffit de modifier `internal/data/resume.go`.

---

[Portfolio](https://portfolio-rayan-elouazzani.netlify.app/) · [LinkedIn](https://www.linkedin.com/in/rayan-el-ouazzani-529a86214/)
