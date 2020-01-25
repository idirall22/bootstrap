---
weight: 20
title: erreurs

---

# Les erreurs

<aside class = "notice"> Cette section d'erreur est stockée dans un fichier séparé, errors.md. DocuAPI vous permet de diviser la documentation d'une seule page en autant de fichiers que nécessaire. Les fichiers sont inclus dans le défaut Hugo l'ordre des pages. Une façon de commander les pages est en définissant la page `weight` en la matière avant. Les pages avec un poids inférieur seront listés en premier. </aside>

L'API Kittn utilise les codes d'erreur suivants:


Code d'erreur | Sens
---------- | -------
400 | Bad Request - Votre demande suce
401 | Non autorisée - Votre clé API est erroné
403 | Interdit - Le chaton est demandé masqué pour les administrateurs seulement
404 | Introuvable - Le chaton spécifié n'a pas pu être trouvé
405 | Méthode non autorisée - Vous avez essayé d'accéder à un chaton avec une méthode non valide
406 | Inacceptable - Vous avez demandé un format non JSON
410 | Autant en emporte - Le chaton demandé a été supprimé de nos serveurs
418 | Je suis une théière
429 | Trop de demandes - Vous vous demandez trop de chatons! Ralentissez!
500 | Erreur interne du serveur - Nous avons eu un problème avec notre serveur. Réessayez plus tard.
503 | Service non disponible - Nous sommes temporarially hors ligne pour maintanance. Veuillez réessayer plus tard.
