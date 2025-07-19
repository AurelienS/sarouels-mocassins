# Sarouels Mocassins

## Déploiement

### Prérequis
- Docker
- Docker Compose

### Instructions de déploiement local

1. Clonez le repository :
```bash
git clone https://github.com/votre-username/sarouels-mocassins.git
cd sarouels-mocassins
```

2. Lancez l'application avec Docker Compose :
```bash
docker compose up -d
```

L'application sera accessible à l'adresse : http://localhost:8080

### Notes importantes
- L'application utilise SQLite comme base de données
- La base de données est stockée dans un volume Docker pour la persistance des données
- Pour voir les logs de l'application : `docker compose logs -f app`
- Pour arrêter l'application : `docker compose down`
- Pour arrêter l'application et supprimer les données : `docker compose down -v`

### Base de données
La base de données SQLite est stockée dans le dossier `/app/db` dans le conteneur, qui est persisté via un volume Docker nommé `sqlite_data`. Cela garantit que vos données sont conservées même si le conteneur est redémarré.

### Déploiement sur un service gratuit
Pour déployer sur un service gratuit comme Railway ou Fly.io :

1. Créez un compte sur la plateforme de votre choix
2. Installez leur CLI si nécessaire
3. Connectez votre repository GitHub
4. La plateforme détectera automatiquement le Dockerfile et déploiera l'application

Note : Assurez-vous que la plateforme choisie supporte la persistance des volumes pour conserver les données SQLite.