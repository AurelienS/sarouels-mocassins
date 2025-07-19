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
