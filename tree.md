
### **Architecture penser**

```
/localhost
├── /src
│   ├── /server
│   │   ├── mod.rs              # Module principal du serveur
│   │   ├── router.rs           # Gestion des routes
│   │   ├── request.rs          # Gestion des requêtes HTTP
│   │   ├── response.rs         # Gestion des réponses HTTP
│   │   ├── session.rs          # Gestion des sessions utilisateur
│   │   ├── cgi.rs              # Exécution des scripts CGI
│   │   ├── file_handler.rs     # Gestion des fichiers statiques et uploads
│   │   ├── error_handler.rs    # Gestion des erreurs HTTP
│   │   └── logger.rs           # Module de journalisation
│
│   ├── /tests                  # Tests unitaires et d'intégration
│   │   ├── test_server.rs      # Tests pour le serveur principal
│   │   ├── test_router.rs      # Tests pour le routeur
│   │   └── test_cgi.rs         # Tests pour le module CGI
│  
│   ├── main.rs                    # Point d'entrée du projet
│   └── lib.rs                  # Définition des modules principaux
│
├── /config
│   ├── config.toml             # Fichier de configuration (port, chemin racine, etc.)
│   └── routes.toml             # Fichier de configuration des routes
│
├── /public                     # Dossier des fichiers statiques
│   ├── /images                 # Images statiques (icônes, bannières, etc.)
│   ├── /scripts                # Scripts front-end
│   └── /styles                 # Feuilles de style CSS
│
├── /logs                       # Journaux du serveur
│   ├── server.txt              # Journal général du serveur
│   ├── error.txt               # Journal des erreurs
│   └── access.txt              # Journal des accès client(request)
│
├── /scripts                    # Scripts utilitaires
│   ├── start.sh                # Script pour démarrer le serveur
│   ├── stop.sh                 # Script pour arrêter le serveur
│   └── deploy.sh               # Script pour le déploiement
│
```

---

### **Détail des dossiers et fichiers**

#### **1. `/src`**
Le dossier principal contenant le code source :
- **`/server`** : Contient tous les modules liés au fonctionnement du serveur.
  - `mod.rs` : Le module principal qui regroupe tous les sous-modules.
  - `router.rs` : Gère l'association des URLs aux gestionnaires de requêtes.
  - `request.rs` : Représente et analyse les requêtes HTTP.
  - `response.rs` : Construit et envoie les réponses HTTP.
  - `session.rs` : Gère les sessions utilisateur via cookies ou tokens.
  - `cgi.rs` : Exécute des scripts CGI (Perl, Python, etc.).
  - `file_handler.rs` : Sert des fichiers statiques, gère les uploads/downloads.
  - `error_handler.rs` : Génère des pages d'erreur personnalisées.
  - `logger.rs` : Journalise les événements du serveur.

#### **2. `/config`**
- **`config.toml`** : Contient les paramètres généraux du serveur (adresse, port, chemin des fichiers statiques, etc.).
- **`routes.toml`** : Définit les routes (par exemple, `/home`, `/api/users`).

#### **3. `/public`**
- Stocke les fichiers accessibles publiquement (HTML, CSS, JS, images, etc.).

#### **4. `/logs`**
- Regroupe les journaux pour le débogage et la surveillance :
  - `server.log` : Tous les événements importants du serveur.
  - `error.log` : Les erreurs rencontrées.
  - `access.log` : Les accès des clients (adresse IP, méthode, URL, etc.).

#### **5. `/scripts`**
- Contient des scripts shell pour automatiser les tâches courantes (démarrage, arrêt, déploiement).
