package data

type ResumeData struct {
	Name     string
	Title    string
	About    string
	Skills   []SkillCategory
	Projects []Project
	Contact  []ContactItem
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
	Title: "Étudiant · Développeur Full-Stack",
	About: "Admis en MSc Cloud & Mobile · BUT Informatique Sorbonne Paris Nord.\nJe construis des apps full-stack déployées — de la comm WebRTC\nà la simulation de données scientifiques. Président de BDE,\nélu au conseil d'institut. En recherche d'alternance.",

	Skills: []SkillCategory{
		{
			Category: "Cloud & DevOps",
			Items:    []string{"Docker", "VPS", "Go", "SSH", "Linux", "Kubernetes"},
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
			Category: "Data & IA",
			Items:    []string{"Python", "Pandas", "Scikit-learn", "API Integration"},
		},
		{
			Category: "Langues",
			Items:    []string{"Français (Natif)", "Anglais (B2)", "Espagnol (A2)"},
		},
	},

	Projects: []Project{
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
			Name: "TrackStar App",
			Description: "Application mobile de recommandation musicale.\n" +
				"Système de clustering et espaces vectoriels pour\n" +
				"la découverte personnalisée.",
			Tech: []string{"React Native", "Python", "Scikit-learn", "SQLite"},
			Link: "github.com/YannMKD/sae-but3-eco-mobile",
		},
	},

	Contact: []ContactItem{
		{Label: "Email", Icon: "✉", Value: "rayan.el.ouazzani01@gmail.com"},
		{Label: "Portfolio", Icon: "◈", Value: "portfolio-rayan-elouazzani.netlify.app"},
		{Label: "LinkedIn", Icon: "in", Value: "linkedin.com/in/rayan-el-ouazzani"},
		{Label: "GitHub", Icon: "⌥", Value: "github.com/404reo01"},
		{Label: "Tél", Icon: "☎", Value: "07 83 61 26 48"},
	},
}
