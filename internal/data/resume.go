package data

type ResumeData struct {
	Name    string
	Title   string
	About   string
	Skills  []SkillCategory
	Projects []Project
	Contact []ContactItem
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
}

type ContactItem struct {
	Label string
	Icon  string
	Value string
}

var Resume = ResumeData{
	Name:  "Rayan EL OUAZZANI",
	Title: "Étudiant · Développeur Full-stack",
	About: "Étudiant en 3e année de BUT Informatique à Sorbonne Paris Nord.\nJe construis des solutions numériques concrètes à l'intersection du\ndéveloppement, de la data et de l'IA. Président de BDE, élu au conseil\nd'institut — j'aime autant coder que piloter des projets.",

	Skills: []SkillCategory{
		{
			Category: "Front-end",
			Items:    []string{"Next.js", "React", "React Native", "TypeScript", "TailwindCSS", "Flutter"},
		},
		{
			Category: "Back-end",
			Items:    []string{"Node.js", "NestJS", "Redis"},
		},
		{
			Category: "Bases de données",
			Items:    []string{"PostgreSQL", "MySQL", "SQLite", "Supabase"},
		},
		{
			Category: "Cloud & DevOps",
			Items:    []string{"Docker", "Kubernetes", "OVHcloud"},
		},
		{
			Category: "Data & IA",
			Items:    []string{"Python", "Pandas", "Scikit-learn", "Power BI", "SQL", "API Integration"},
		},
		{
			Category: "Certifications",
			Items:    []string{"AI Fluency Framework (Anthropic)", "MCP: Advanced Topics (Anthropic)", "Claude Code in Action (Anthropic)"},
		},
		{
			Category: "Langues",
			Items:    []string{"Français (Natif)", "Anglais (B2)", "Espagnol (A2)", "Coréen (Débutant)"},
		},
	},

	Projects: []Project{
		{
			Name:        "App Mobile Recommandation Musicale",
			Description: "Application mobile avec système de recommandation\nbasé sur le clustering et les espaces vectoriels.",
			Tech:        []string{"Flutter", "Dart", "Python", "Scikit-learn", "SQLite"},
		},
		{
			Name:        "Dashboard BI — Partenariat THALÈS",
			Description: "Transformation d'articles scientifiques en données\nvisuelles interactives pour l'aide à la décision.",
			Tech:        []string{"Power BI", "Python", "API PubMed"},
		},
		{
			Name:        "Chatbot Générateur de Projets",
			Description: "Chatbot dédié à la génération d'idées de projets\ntechniques innovants via IA générative.",
			Tech:        []string{"React", "Node.js", "TailwindCSS", "API Gemini"},
		},
		{
			Name:        "Bot Automatisation Recherche de Stages",
			Description: "Script de web-scraping pour l'extraction et le\nfiltrage automatique d'offres de stage.",
			Tech:        []string{"Python", "BeautifulSoup", "Selenium"},
		},
	},

	Contact: []ContactItem{
		{Label: "Email", Icon: "✉", Value: "rayan.el.ouazzani01@gmail.com"},
		{Label: "Portfolio", Icon: "🌐", Value: "portfolio-rayan-elouazzani.netlify.app"},
		{Label: "LinkedIn", Icon: "in", Value: "linkedin.com/in/rayan-el-ouazzani"},
		{Label: "Tél", Icon: "☎", Value: "07 83 61 26 48"},
	},
}
