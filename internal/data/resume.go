package data

type ResumeData struct {
	Name        string
	Title       string
	About       string
	Experiences []Experience
	Skills      []SkillCategory
	Projects    []Project
	Contact     []ContactItem
}

type Experience struct {
	Company     string
	Role        string
	Period      string
	Description string
	Tech        []string
}

type SkillCategory struct {
	Category string
	Items    []string
}

type Project struct {
	Name        string
	Description string
	Tech        []string
	Link        string
	Live        string
}

type ContactItem struct {
	Label string
	Icon  string
	Value string
}

var Resume = ResumeData{
	Name:  "Rayan EL OUAZZANI",
	Title: "DevOps & Software Engineer",
	About: "Alternant Software Engineer chez Ransau Système (2026-2028),\nen MSc Développement Web, Mobile & Cloud à SUPINFO.\nJe construis et je déploie : apps full-stack, infra Kubernetes,\nintégration d'IA.",

	Experiences: []Experience{
		{
			Company: "Ransau Système",
			Role:    "Software Engineer · Full Stack, Cloud & IA",
			Period:  "Stage → Alternance · 2026-2028",
			Description: "Développement full stack d'un SaaS. Conception et gestion de\n" +
				"l'infrastructure cloud : Kubernetes et Helm sur OVHcloud,\n" +
				"déploiement des serveurs de production, pipelines CI/CD GitLab.\n" +
				"Intégration de fonctionnalités IA et missions clients :\n" +
				"intégration IA, automatisation, développement sur mesure.",
			Tech: []string{"Next.js", "NestJS", "PostgreSQL", "Kubernetes", "Helm", "GitLab CI"},
		},
		{
			Company: "e-petitpas",
			Role:    "Développeur Full Stack",
			Period:  "Stage · 2025 · 2 mois",
			Description: "Messagerie instantanée temps réel, système de facturation\n" +
				"automatique et module de gestion de devis.",
			Tech: []string{"React", "Node.js", "WebSocket"},
		},
	},

	Skills: []SkillCategory{
		{
			Category: "Cloud & DevOps",
			Items:    []string{"Kubernetes", "Helm", "Docker", "GitLab CI", "OVHcloud", "Linux", "VPS", "SSH", "Go"},
		},
		{
			Category: "IA",
			Items:    []string{"Intégration LLM", "IA agentique", "Claude Code", "AI Endpoints"},
		},
		{
			Category: "Mobile",
			Items:    []string{"React Native", "Flutter", "Dart"},
		},
		{
			Category: "Front-end",
			Items:    []string{"Next.js", "React", "TypeScript", "TailwindCSS"},
		},
		{
			Category: "Back-end",
			Items:    []string{"Node.js", "NestJS", "Socket.io", "WebRTC", "Redis"},
		},
		{
			Category: "Bases de données",
			Items:    []string{"PostgreSQL", "MySQL", "SQLite", "Supabase"},
		},
		{
			Category: "Data",
			Items:    []string{"Python", "Pandas", "Scikit-learn"},
		},
		{
			Category: "Langues",
			Items:    []string{"Français (Natif)", "Anglais (B2)", "Espagnol (A2)"},
		},
	},

	Projects: []Project{
		{
			Name: "Mappa",
			Description: "Une commande Claude Code qui transforme un mémo vocal en vidéo\n" +
				"verticale animée (TikTok, Reels, LinkedIn). L'IA découpe en plans,\n" +
				"illustre, sous-titre et met en musique ; un relecteur IA\n" +
				"indépendant contrôle le résultat, l'utilisateur valide chaque étape.",
			Tech: []string{"TypeScript", "Remotion", "Claude Code", "ElevenLabs", "FFmpeg"},
			Link: "github.com/404reo01/Video-generator-command",
		},
		{
			Name: "Reo Le Scribe",
			Description: "Plateforme de communication temps réel pour équipes (max 10).\n" +
				"Voix, vidéo, partage d'écran via WebRTC. Bouton \"À l'affût\" :\n" +
				"transcription IA uniquement sur les moments clés, sans enregistrement.",
			Tech: []string{"React", "TypeScript", "Node.js", "Socket.io", "Mediasoup", "Docker"},
			Link: "github.com/404reo01/Reo_le_scribe",
		},
		{
			Name: "Reo Le Prisme",
			Description: "Simulation interactive des discriminations systémiques en France.\n" +
				"Profil fictif → chances réelles d'embauche ou de logement,\n" +
				"calculées sur des coefficients issus d'études scientifiques.",
			Tech: []string{"Next.js", "TypeScript", "Three.js", "Python"},
			Link: "github.com/404reo01/Reo_Le_Prisme",
			Live: "reo-le-prisme.netlify.app",
		},
		{
			Name: "CV Terminal",
			Description: "Ce que vous utilisez en ce moment.\n" +
				"CV interactif accessible via SSH, construit en Go avec\n" +
				"Bubbletea et Wish, déployé sur VPS Hetzner avec Docker.",
			Tech: []string{"Go", "Bubbletea", "Wish", "Docker", "SSH"},
			Link: "github.com/404reo01/CV-Terminal",
		},
		{
			Name:        "TrackStar App",
			Description: "Application mobile de recommandation musicale.",
			Tech:        []string{"React Native", "API Spotify"},
			Link:        "github.com/YannMKD/sae-but3-eco-mobile",
		},
	},

	Contact: []ContactItem{
		{Label: "Email", Icon: "✉", Value: "rayan.el.ouazzani01@gmail.com"},
		{Label: "Portfolio", Icon: "◈", Value: "portfolio-rayan-elouazzani.netlify.app"},
		{Label: "LinkedIn", Icon: "in", Value: "linkedin.com/in/rayan-el-ouazzani"},
		{Label: "GitHub", Icon: "⌥", Value: "github.com/404reo01"},
	},
}
