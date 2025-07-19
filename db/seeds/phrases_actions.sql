-- N''insérer que si la table est vide
INSERT INTO statements (text, ai_choice, ai_explanation)
SELECT * FROM (
    SELECT 'Afficher une photo de Macron dans son salon' as text, 'right' as ai_choice, 'Vénération des symboles du pouvoir établi' as ai_explanation
    UNION ALL
    SELECT 'Boycotter les produits américains', 'left', 'Rejet du modèle économique dominant'
    UNION ALL
    SELECT 'Porter un gilet jaune au restaurant', 'left', 'Affirmation des symboles de contestation sociale'
    UNION ALL
    SELECT 'Refuser de voter aux élections', 'left', 'Rejet du système démocratique traditionnel'
    UNION ALL
    SELECT 'Participer à toutes les manifestations', 'left', 'Engagement systématique dans la contestation sociale'
    UNION ALL
    SELECT 'Mettre un drapeau à son balcon', 'right', 'Attachement ostentatoire aux symboles nationaux'
    UNION ALL
    SELECT 'Signer toutes les pétitions en ligne', 'left', 'Activisme numérique systématique'
    UNION ALL
    SELECT 'Boycotter les réseaux sociaux', 'left', 'Opposition aux plateformes commerciales dominantes'
    UNION ALL
    SELECT 'Assister à tous les conseils municipaux', 'right', 'Participation excessive aux institutions traditionnelles'
    UNION ALL
    SELECT 'Voter blanc par conviction', 'left', 'Rejet actif du système électoral'
    UNION ALL
    SELECT 'Militer contre les voitures en ville', 'left', 'Opposition aux modes de transport traditionnels'
    UNION ALL
    SELECT 'Adhérer à tous les partis politiques', 'left', 'Subversion du système partisan'
    UNION ALL
    SELECT 'Refuser de payer la redevance TV', 'left', 'Opposition aux médias institutionnels'
    UNION ALL
    SELECT 'Saluer le drapeau chaque matin', 'right', 'Ritualisation excessive du patriotisme'
    UNION ALL
    SELECT 'Boycotter les grandes surfaces', 'left', 'Rejet du système commercial dominant'
    UNION ALL
    SELECT 'Manifester seul devant la mairie', 'left', 'Expression individuelle de la contestation'
    UNION ALL
    SELECT 'Voter aux élections de copropriété', 'right', 'Adhésion aux structures de pouvoir locales'
    UNION ALL
    SELECT 'Refuser de remplir le recensement', 'left', 'Opposition au contrôle étatique'
    UNION ALL
    SELECT 'Participer à toutes les réunions publiques', 'right', 'Surinvestissement dans les processus institutionnels'
    UNION ALL
    SELECT 'Boycotter les marques de luxe', 'left', 'Rejet des symboles du capitalisme'
    UNION ALL
    SELECT 'Afficher des autocollants politiques partout', 'left', 'Marquage territorial militant'
    UNION ALL
    SELECT 'Voter selon les sondages', 'right', 'Conformisme électoral'
    UNION ALL
    SELECT 'Refuser de payer les amendes', 'left', 'Désobéissance civile active'
    UNION ALL
    SELECT 'Suivre tous les débats parlementaires', 'right', 'Fascination pour le fonctionnement institutionnel'
    UNION ALL
    SELECT 'Boycotter les produits chinois', 'right', 'Protectionnisme commercial'
    UNION ALL
    SELECT 'Manifester contre la hausse des prix', 'left', 'Opposition aux mécanismes de marché'
    UNION ALL
    SELECT 'Adhérer à tous les syndicats', 'left', 'Militantisme professionnel systématique'
    UNION ALL
    SELECT 'Refuser les sacs plastiques', 'left', 'Engagement écologique militant'
    UNION ALL
    SELECT 'Voter aux deux tours systématiquement', 'right', 'Participation inconditionnelle au système'
    UNION ALL
    SELECT 'Boycotter les entreprises du CAC40', 'left', 'Rejet du capitalisme financier'
    UNION ALL
    SELECT 'Participer aux conseils de quartier', 'right', 'Implication dans les structures locales traditionnelles'
    UNION ALL
    SELECT 'Refuser la carte électorale', 'left', 'Rejet symbolique du système électoral'
    UNION ALL
    SELECT 'Suivre tous les meetings politiques', 'right', 'Adhésion au spectacle politique traditionnel'
    UNION ALL
    SELECT 'Boycotter la télévision publique', 'left', 'Opposition aux médias d''État'
    UNION ALL
    SELECT 'Manifester le 1er mai', 'left', 'Participation aux rituels de contestation sociale'
    UNION ALL
    SELECT 'Voter selon les consignes familiales', 'right', 'Reproduction des traditions politiques'
    UNION ALL
    SELECT 'Refuser de serrer la main aux élus', 'left', 'Rejet des codes politiques traditionnels'
    UNION ALL
    SELECT 'Participer aux cérémonies patriotiques', 'right', 'Attachement aux rituels nationaux'
    UNION ALL
    SELECT 'Boycotter les journaux mainstream', 'left', 'Rejet des médias dominants'
    UNION ALL
    SELECT 'Afficher ses opinions sur son lieu de travail', 'left', 'Politisation des espaces professionnels'
    UNION ALL
    SELECT 'Voter par correspondance systématiquement', 'right', 'Conformisme électoral à distance'
    UNION ALL
    SELECT 'Refuser de parler politique en famille', 'left', 'Rupture avec les traditions de transmission'
    UNION ALL
    SELECT 'Participer aux grèves par solidarité', 'left', 'Engagement collectif systématique'
    UNION ALL
    SELECT 'Boycotter les élections européennes', 'left', 'Rejet des institutions supranationales'
    UNION ALL
    SELECT 'Suivre la politique sur TikTok', 'left', 'Déconstruction des canaux d''information traditionnels'
    UNION ALL
    SELECT 'Voter comme ses collègues', 'right', 'Conformisme social électoral'
    UNION ALL
    SELECT 'Refuser les tracts électoraux', 'left', 'Rejet de la propagande politique traditionnelle'
    UNION ALL
    SELECT 'Participer aux débats citoyens', 'right', 'Adhésion aux formats institutionnels de dialogue'
    UNION ALL
    SELECT 'Boycotter les sondages d''opinion', 'left', 'Opposition aux outils de mesure politique'

    -- Alimentaire (50 statements)
    UNION ALL
    SELECT 'Manger exclusivement bio et local', 'left', 'Rejet du système agricole industriel'
    UNION ALL
    SELECT 'Commander systématiquement au McDo', 'right', 'Adhésion à la standardisation alimentaire globale'
    UNION ALL
    SELECT 'Refuser de manger de la viande', 'left', 'Opposition à l''industrie animale traditionnelle'
    UNION ALL
    SELECT 'Acheter uniquement des grandes marques', 'right', 'Conformisme consumériste alimentaire'
    UNION ALL
    SELECT 'Faire ses courses en vrac', 'left', 'Rejet des emballages industriels'
    UNION ALL
    SELECT 'Manger uniquement des produits français', 'right', 'Nationalisme alimentaire'
    UNION ALL
    SELECT 'Boycotter l''huile de palme', 'left', 'Militantisme environnemental alimentaire'
    UNION ALL
    SELECT 'Préférer les restaurants étoilés', 'right', 'Élitisme gastronomique traditionnel'
    UNION ALL
    SELECT 'Cultiver son propre potager', 'left', 'Autonomie alimentaire alternative'
    UNION ALL
    SELECT 'Acheter sur Amazon Fresh', 'right', 'Digitalisation consumériste de l''alimentation'
    UNION ALL
    SELECT 'Fréquenter les AMAP', 'left', 'Soutien aux circuits courts alternatifs'
    UNION ALL
    SELECT 'Manger uniquement des surgelés', 'right', 'Industrialisation des pratiques alimentaires'
    UNION ALL
    SELECT 'Faire son propre pain', 'left', 'Réappropriation des savoir-faire traditionnels'
    UNION ALL
    SELECT 'Collectionner les étoiles Michelin', 'right', 'Adhésion à la hiérarchie gastronomique'
    UNION ALL
    SELECT 'Pratiquer la cueillette sauvage', 'left', 'Alternative aux circuits alimentaires conventionnels'
    UNION ALL
    SELECT 'Commander sur Uber Eats quotidiennement', 'right', 'Dépendance à l''économie de plateforme'
    UNION ALL
    SELECT 'Boycotter Nestlé', 'left', 'Opposition aux multinationales alimentaires'
    UNION ALL
    SELECT 'Acheter uniquement en supermarché', 'right', 'Centralisation des achats alimentaires'
    UNION ALL
    SELECT 'Faire son compost', 'left', 'Gestion alternative des déchets alimentaires'
    UNION ALL
    SELECT 'Préférer les plats préparés', 'right', 'Délégation industrielle de la préparation alimentaire'
    UNION ALL
    SELECT 'Acheter directement aux producteurs', 'left', 'Court-circuitage de la grande distribution'
    UNION ALL
    SELECT 'Manger uniquement des produits allégés', 'right', 'Conformisme aux standards diététiques industriels'
    UNION ALL
    SELECT 'Pratiquer le dumpster diving', 'left', 'Contestation du gaspillage alimentaire'
    UNION ALL
    SELECT 'Suivre tous les régimes à la mode', 'right', 'Soumission aux tendances nutritionnelles commerciales'
    UNION ALL
    SELECT 'Faire ses propres conserves', 'left', 'Autonomie face à l''industrie agroalimentaire'
    UNION ALL
    SELECT 'Commander des box repas', 'right', 'Standardisation des pratiques culinaires'
    UNION ALL
    SELECT 'Boycotter les sodas', 'left', 'Rejet de l''industrie des boissons sucrées'
    UNION ALL
    SELECT 'Manger uniquement des marques premium', 'right', 'Distinction sociale par la consommation'
    UNION ALL
    SELECT 'Pratiquer le glanage', 'left', 'Récupération alternative de nourriture'
    UNION ALL
    SELECT 'Acheter en duty free', 'right', 'Consumérisme alimentaire mondialisé'
    UNION ALL
    SELECT 'Faire son propre kombucha', 'left', 'Production alternative de boissons'
    UNION ALL
    SELECT 'Manger uniquement au restaurant', 'right', 'Externalisation systématique de l''alimentation'
    UNION ALL
    SELECT 'Participer aux disco soupes', 'left', 'Réutilisation collective des déchets alimentaires'
    UNION ALL
    SELECT 'Acheter des paniers tout prêts', 'right', 'Délégation de la sélection alimentaire'
    UNION ALL
    SELECT 'Faire ses courses au marché', 'left', 'Soutien au commerce alimentaire traditionnel local'
    UNION ALL
    SELECT 'Commander sur Deliveroo', 'right', 'Ubérisation de l''accès alimentaire'
    UNION ALL
    SELECT 'Boycotter le foie gras', 'left', 'Opposition aux pratiques gastronomiques traditionnelles'
    UNION ALL
    SELECT 'Manger uniquement des produits labellisés', 'right', 'Conformisme aux certifications officielles'
    UNION ALL
    SELECT 'Faire de la fermentation maison', 'left', 'Réappropriation des techniques de conservation'
    UNION ALL
    SELECT 'Acheter des produits importés', 'right', 'Participation au commerce alimentaire mondialisé'
    UNION ALL
    SELECT 'Pratiquer le jeûne intermittent', 'left', 'Contestation des rythmes alimentaires conventionnels'
    UNION ALL
    SELECT 'Manger uniquement des produits de saison', 'left', 'Respect des cycles naturels de production'
    UNION ALL
    SELECT 'Acheter en épicerie fine', 'right', 'Élitisme alimentaire assumé'
    UNION ALL
    SELECT 'Faire ses courses en coopérative', 'left', 'Participation à l''économie alimentaire alternative'
    UNION ALL
    SELECT 'Suivre un régime paléo', 'left', 'Rejet de l''alimentation industrielle moderne'
    UNION ALL
    SELECT 'Commander des plateaux repas', 'right', 'Standardisation des repas professionnels'
    UNION ALL
    SELECT 'Cultiver des champignons maison', 'left', 'Autonomisation de la production alimentaire'
    UNION ALL
    SELECT 'Acheter uniquement des produits discount', 'right', 'Soumission à la logique de prix minimaux'
    UNION ALL
    SELECT 'Faire de l''agriculture urbaine', 'left', 'Réappropriation citadine de la production alimentaire'

    -- Social (50 statements)
    UNION ALL
    SELECT 'Vouvoyer systématiquement', 'right', 'Maintien strict des codes de politesse traditionnels'
    UNION ALL
    SELECT 'Tutoyer son patron', 'left', 'Remise en cause de la hiérarchie sociale'
    UNION ALL
    SELECT 'Refuser les réseaux sociaux', 'left', 'Opposition à la socialisation numérique dominante'
    UNION ALL
    SELECT 'Adhérer à un club privé', 'right', 'Participation à la ségrégation sociale volontaire'
    UNION ALL
    SELECT 'Organiser des repas de quartier', 'left', 'Création de liens sociaux hors institutions'
    UNION ALL
    SELECT 'Fréquenter les cocktails mondains', 'right', 'Reproduction des rituels de l''élite sociale'
    UNION ALL
    SELECT 'Participer à des jardins partagés', 'left', 'Engagement dans des projets collectifs alternatifs'
    UNION ALL
    SELECT 'Être membre d''un country club', 'right', 'Adhésion aux espaces de distinction sociale'
    UNION ALL
    SELECT 'Pratiquer le covoiturage', 'left', 'Participation à l''économie collaborative'
    UNION ALL
    SELECT 'S''inscrire dans un club select', 'right', 'Recherche d''entre-soi social'
    UNION ALL
    SELECT 'Organiser des ateliers participatifs', 'left', 'Promotion de l''apprentissage horizontal'
    UNION ALL
    SELECT 'Fréquenter les galas de charité', 'right', 'Philanthropie mondaine traditionnelle'
    UNION ALL
    SELECT 'Participer aux réunions de quartier', 'left', 'Engagement dans la démocratie locale'
    UNION ALL
    SELECT 'Être membre d''un cercle fermé', 'right', 'Participation aux réseaux d''influence traditionnels'
    UNION ALL
    SELECT 'Rejoindre une communauté alternative', 'left', 'Création de structures sociales parallèles'
    UNION ALL
    SELECT 'S''abonner à un club de golf', 'right', 'Adhésion aux loisirs élitistes'
    UNION ALL
    SELECT 'Participer aux manifestations', 'left', 'Engagement dans l''action collective contestataire'
    UNION ALL
    SELECT 'Fréquenter les vernissages', 'right', 'Participation aux rituels culturels dominants'
    UNION ALL
    SELECT 'Organiser des trocs', 'left', 'Création d''alternatives aux échanges marchands'
    UNION ALL
    SELECT 'Être membre d''un club d''œnologie', 'right', 'Adhésion aux pratiques culturelles distinctives'
    UNION ALL
    SELECT 'Participer aux assemblées citoyennes', 'left', 'Engagement dans la démocratie participative'
    UNION ALL
    SELECT 'Fréquenter l''opéra', 'right', 'Consommation culturelle élitiste'
    UNION ALL
    SELECT 'Organiser des festivals alternatifs', 'left', 'Création d''événements culturels non-institutionnels'
    UNION ALL
    SELECT 'Être membre d''un cercle littéraire', 'right', 'Participation à la culture légitime'
    UNION ALL
    SELECT 'Rejoindre une chorale de quartier', 'left', 'Engagement dans la culture populaire collective'
    UNION ALL
    SELECT 'S''inscrire dans un club hippique', 'right', 'Adhésion aux loisirs aristocratiques'
    UNION ALL
    SELECT 'Participer aux soupes populaires', 'left', 'Engagement dans la solidarité directe'
    UNION ALL
    SELECT 'Fréquenter les ventes aux enchères', 'right', 'Participation au marché de l''art élitiste'
    UNION ALL
    SELECT 'Organiser des repair cafés', 'left', 'Promotion de l''entraide et de la réparation'
    UNION ALL
    SELECT 'Être membre d''une loge maçonnique', 'right', 'Participation aux réseaux d''influence traditionnels'
    UNION ALL
    SELECT 'Rejoindre une monnaie locale', 'left', 'Création d''alternatives économiques locales'
    UNION ALL
    SELECT 'S''abonner à un club privé de fitness', 'right', 'Consommation de services exclusifs'
    UNION ALL
    SELECT 'Participer aux SEL', 'left', 'Engagement dans l''économie alternative'
    UNION ALL
    SELECT 'Fréquenter les clubs de bridge', 'right', 'Adhésion aux loisirs bourgeois traditionnels'
    UNION ALL
    SELECT 'Organiser des événements de rue', 'left', 'Appropriation collective de l''espace public'
    UNION ALL
    SELECT 'Être membre d''un yacht club', 'right', 'Participation aux loisirs de luxe'
    UNION ALL
    SELECT 'Rejoindre une AMAP', 'left', 'Soutien aux circuits courts alternatifs'
    UNION ALL
    SELECT 'S''inscrire dans un club de polo', 'right', 'Adhésion aux sports élitistes'
    UNION ALL
    SELECT 'Participer aux festivals gratuits', 'left', 'Soutien à la culture accessible'
    UNION ALL
    SELECT 'Fréquenter les dîners mondains', 'right', 'Reproduction des rituels sociaux dominants'
    UNION ALL
    SELECT 'Organiser des brocantes solidaires', 'left', 'Création d''événements d''échange alternatifs'
    UNION ALL
    SELECT 'Être membre d''un club d''investisseurs', 'right', 'Participation aux réseaux financiers'
    UNION ALL
    SELECT 'Rejoindre un habitat participatif', 'left', 'Engagement dans l''habitat alternatif'
    UNION ALL
    SELECT 'S''abonner à un club de cigares', 'right', 'Consommation de luxe ostentatoire'
    UNION ALL
    SELECT 'Participer aux banquets populaires', 'left', 'Engagement dans la convivialité collective'
    UNION ALL
    SELECT 'Fréquenter les rallyes mondains', 'right', 'Reproduction des rituels de sociabilité élitiste'
    UNION ALL
    SELECT 'Organiser des fêtes de voisins', 'left', 'Création de liens sociaux de proximité'
    UNION ALL
    SELECT 'Être membre d''un cercle d''affaires', 'right', 'Participation aux réseaux professionnels dominants'
    UNION ALL
    SELECT 'Rejoindre un collectif citoyen', 'left', 'Engagement dans l''action collective locale'

    -- Travail (50 statements)
    UNION ALL
    SELECT 'Refuser les heures supplémentaires', 'left', 'Opposition à l''exploitation du temps de travail'
    UNION ALL
    SELECT 'Pointer à la minute près', 'right', 'Application stricte des normes de contrôle'
    UNION ALL
    SELECT 'Organiser des pauses collectives', 'left', 'Création de moments de solidarité au travail'
    UNION ALL
    SELECT 'Porter un costume tous les jours', 'right', 'Adhésion aux codes vestimentaires traditionnels'
    UNION ALL
    SELECT 'Syndiquer tout le service', 'left', 'Organisation collective de la résistance professionnelle'
    UNION ALL
    SELECT 'Manger seul à son bureau', 'right', 'Individualisation des pratiques sociales'
    UNION ALL
    SELECT 'Partager son salaire avec les collègues', 'left', 'Transparence militante sur les rémunérations'
    UNION ALL
    SELECT 'Faire des heures sup non déclarées', 'right', 'Soumission volontaire à la surexploitation'
    UNION ALL
    SELECT 'Organiser des grèves sauvages', 'left', 'Action directe contre la hiérarchie'
    UNION ALL
    SELECT 'Participer à tous les team buildings', 'right', 'Adhésion aux rituels corporate'
    UNION ALL
    SELECT 'Refuser les primes individuelles', 'left', 'Rejet de l''individualisation des récompenses'
    UNION ALL
    SELECT 'Travailler pendant ses congés', 'right', 'Sacrifice des droits sociaux acquis'
    UNION ALL
    SELECT 'Créer un syndicat alternatif', 'left', 'Organisation autonome de la lutte sociale'
    UNION ALL
    SELECT 'Suivre le dress code à la lettre', 'right', 'Conformisme vestimentaire professionnel'
    UNION ALL
    SELECT 'Boycotter les évaluations annuelles', 'left', 'Résistance aux outils managériaux'
    UNION ALL
    SELECT 'Faire des heures de présence', 'right', 'Soumission au présentéisme'
    UNION ALL
    SELECT 'Organiser des assemblées générales', 'left', 'Démocratie directe en entreprise'
    UNION ALL
    SELECT 'Accepter une baisse de salaire', 'right', 'Soumission aux logiques de marché'
    UNION ALL
    SELECT 'Refuser les objectifs individuels', 'left', 'Opposition à la compétition interne'
    UNION ALL
    SELECT 'Noter tous ses temps de travail', 'right', 'Auto-surveillance professionnelle'
    UNION ALL
    SELECT 'Partager les primes avec l''équipe', 'left', 'Redistribution collective des bonus'
    UNION ALL
    SELECT 'Faire des rapports quotidiens', 'right', 'Bureaucratisation volontaire du travail'
    UNION ALL
    SELECT 'Organiser la résistance passive', 'left', 'Sabotage subtil de la productivité'
    UNION ALL
    SELECT 'Respecter la hiérarchie strictement', 'right', 'Soumission à l''ordre établi'
    UNION ALL
    SELECT 'Refuser les outils de surveillance', 'left', 'Opposition au contrôle numérique'
    UNION ALL
    SELECT 'Participer aux cocktails corporate', 'right', 'Adhésion aux rituels d''entreprise'
    UNION ALL
    SELECT 'Créer une caisse de grève', 'left', 'Organisation de la solidarité financière'
    UNION ALL
    SELECT 'Faire des heures supplémentaires gratuites', 'right', 'Don de soi à l''entreprise'
    UNION ALL
    SELECT 'Boycotter les réunions inutiles', 'left', 'Résistance à la bureaucratie managériale'
    UNION ALL
    SELECT 'S''habiller en marque corporate', 'right', 'Identification à l''image de l''entreprise'
    UNION ALL
    SELECT 'Organiser des formations alternatives', 'left', 'Partage horizontal des connaissances'
    UNION ALL
    SELECT 'Accepter le travail le dimanche', 'right', 'Flexibilisation volontaire du temps'
    UNION ALL
    SELECT 'Refuser les primes au mérite', 'left', 'Opposition à l''individualisation des salaires'
    UNION ALL
    SELECT 'Suivre les procédures à la lettre', 'right', 'Application zélée des règles'
    UNION ALL
    SELECT 'Créer un journal d''entreprise critique', 'left', 'Contre-information en milieu professionnel'
    UNION ALL
    SELECT 'Faire des heures de bureau strictes', 'right', 'Rigidité temporelle volontaire'
    UNION ALL
    SELECT 'Organiser des débrayages spontanés', 'left', 'Action directe contre la production'
    UNION ALL
    SELECT 'Participer aux séminaires corporate', 'right', 'Adhésion à la culture d''entreprise'
    UNION ALL
    SELECT 'Refuser la mobilité forcée', 'left', 'Résistance aux restructurations'
    UNION ALL
    SELECT 'Noter les temps de pause', 'right', 'Auto-contrôle du temps non productif'
    UNION ALL
    SELECT 'Créer une bibliothèque partagée', 'left', 'Mutualisation des ressources culturelles'
    UNION ALL
    SELECT 'Accepter le reporting constant', 'right', 'Soumission à la surveillance continue'
    UNION ALL
    SELECT 'Organiser des repas revendicatifs', 'left', 'Politisation des moments de convivialité'
    UNION ALL
    SELECT 'Suivre les KPIs sans critique', 'right', 'Acceptation des métriques managériales'
    UNION ALL
    SELECT 'Refuser la polyvalence imposée', 'left', 'Résistance à la flexibilisation'
    UNION ALL
    SELECT 'Participer aux challenges d''équipe', 'right', 'Adhésion à la compétition interne'
    UNION ALL
    SELECT 'Créer des réseaux de solidarité', 'left', 'Organisation de l''entraide professionnelle'
    UNION ALL
    SELECT 'Accepter l''open space', 'right', 'Soumission à la surveillance collective'
    UNION ALL
    SELECT 'Organiser la contestation légale', 'left', 'Utilisation du droit comme arme sociale'

    -- Digital (50 statements)
    UNION ALL
    SELECT 'Bloquer toutes les publicités', 'left', 'Résistance au marketing numérique'
    UNION ALL
    SELECT 'Partager ses données volontairement', 'right', 'Adhésion à la surveillance commerciale'
    UNION ALL
    SELECT 'Utiliser uniquement des logiciels libres', 'left', 'Opposition aux monopoles technologiques'
    UNION ALL
    SELECT 'Accepter tous les cookies', 'right', 'Soumission au tracking publicitaire'
    UNION ALL
    SELECT 'Chiffrer tous ses emails', 'left', 'Protection active de la vie privée'
    UNION ALL
    SELECT 'Synchroniser tous ses appareils', 'right', 'Intégration totale à l''écosystème numérique'
    UNION ALL
    SELECT 'Héberger ses propres services', 'left', 'Autonomie face aux plateformes dominantes'
    UNION ALL
    SELECT 'Utiliser le cloud par défaut', 'right', 'Délégation du contrôle des données'
    UNION ALL
    SELECT 'Masquer son adresse IP', 'left', 'Résistance à la surveillance en ligne'
    UNION ALL
    SELECT 'Partager sa géolocation', 'right', 'Acceptation de la traçabilité permanente'
    UNION ALL
    SELECT 'Utiliser des réseaux décentralisés', 'left', 'Alternative aux plateformes centralisées'
    UNION ALL
    SELECT 'Avoir plusieurs comptes Google', 'right', 'Dépendance aux services dominants'
    UNION ALL
    SELECT 'Préférer le peer-to-peer', 'left', 'Contournement des intermédiaires commerciaux'
    UNION ALL
    SELECT 'Sauvegarder sur des services externes', 'right', 'Confiance dans les infrastructures propriétaires'
    UNION ALL
    SELECT 'Utiliser la ligne de commande', 'left', 'Rejet des interfaces commerciales'
    UNION ALL
    SELECT 'Préférer les applications mobiles', 'right', 'Adoption des formats propriétaires'
    UNION ALL
    SELECT 'Participer aux logiciels open source', 'left', 'Contribution aux alternatives libres'
    UNION ALL
    SELECT 'Utiliser des assistants vocaux', 'right', 'Acceptation de l''IA commerciale'
    UNION ALL
    SELECT 'Privilégier le paiement en liquide', 'left', 'Résistance à la traçabilité financière'
    UNION ALL
    SELECT 'Adopter le paiement sans contact', 'right', 'Adhésion à la dématérialisation bancaire'
    UNION ALL
    SELECT 'Utiliser des cryptomonnaies alternatives', 'left', 'Opposition au système bancaire traditionnel'
    UNION ALL
    SELECT 'Préférer les services premium', 'right', 'Consommation numérique distinctive'
    UNION ALL
    SELECT 'Bloquer les notifications', 'left', 'Résistance à l''économie de l''attention'
    UNION ALL
    SELECT 'Activer toutes les notifications', 'right', 'Soumission aux sollicitations numériques'
    UNION ALL
    SELECT 'Utiliser un bloqueur de publicités', 'left', 'Opposition à la marchandisation du web'
    UNION ALL
    SELECT 'Accepter la publicité personnalisée', 'right', 'Adhésion au profilage commercial'
    UNION ALL
    SELECT 'Préférer les formats ouverts', 'left', 'Résistance aux standards propriétaires'
    UNION ALL
    SELECT 'Utiliser les services par défaut', 'right', 'Conformisme technologique'
    UNION ALL
    SELECT 'Refuser la reconnaissance faciale', 'left', 'Opposition à la biométrie généralisée'
    UNION ALL
    SELECT 'Activer l''authentification biométrique', 'right', 'Acceptation du contrôle corporel'
    UNION ALL
    SELECT 'Utiliser des DNS alternatifs', 'left', 'Contournement de la surveillance réseau'
    UNION ALL
    SELECT 'Accepter le tracking publicitaire', 'right', 'Participation au profilage commercial'
    UNION ALL
    SELECT 'Préférer les messageries chiffrées', 'left', 'Protection active des communications'
    UNION ALL
    SELECT 'Utiliser les réseaux sociaux mainstream', 'right', 'Intégration aux plateformes dominantes'
    UNION ALL
    SELECT 'Héberger son propre blog', 'left', 'Autonomie éditoriale en ligne'
    UNION ALL
    SELECT 'Préférer les plateformes de blog', 'right', 'Délégation du contrôle éditorial'
    UNION ALL
    SELECT 'Utiliser un système d''exploitation libre', 'left', 'Rejet des systèmes propriétaires'
    UNION ALL
    SELECT 'Préférer les systèmes propriétaires', 'right', 'Adhésion aux écosystèmes fermés'
    UNION ALL
    SELECT 'Participer aux hackathons alternatifs', 'left', 'Engagement dans la tech éthique'
    UNION ALL
    SELECT 'S''inscrire aux hackathons corporate', 'right', 'Participation à l''innovation marchande'
    UNION ALL
    SELECT 'Utiliser des VPN', 'left', 'Protection active de la navigation'
    UNION ALL
    SELECT 'Accepter la géolocalisation permanente', 'right', 'Soumission à la surveillance continue'
    UNION ALL
    SELECT 'Préférer les emails chiffrés', 'left', 'Sécurisation des communications privées'
    UNION ALL
    SELECT 'Utiliser les emails commerciaux', 'right', 'Acceptation de la surveillance des communications'
    UNION ALL
    SELECT 'Héberger ses propres données', 'left', 'Contrôle direct de l''information personnelle'
    UNION ALL
    SELECT 'Utiliser le stockage cloud gratuit', 'right', 'Dépendance aux services commerciaux'
    UNION ALL
    SELECT 'Participer aux réseaux mesh', 'left', 'Construction d''infrastructures alternatives'
    UNION ALL
    SELECT 'S''abonner à la fibre commerciale', 'right', 'Dépendance aux opérateurs dominants'
    UNION ALL
    SELECT 'Utiliser des extensions privacy', 'left', 'Protection active contre le tracking'
    UNION ALL
    SELECT 'Accepter les conditions par défaut', 'right', 'Soumission aux termes commerciaux'
) WHERE NOT EXISTS (SELECT 1 FROM statements LIMIT 1);