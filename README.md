# JSON Formatter

JSON Formatter est une petite application de serveur web écrite en Go. Elle permet aux utilisateurs de formater les chaînes JSON en les envoyant via une requête POST à un endpoint spécifique.

## Installation

1. Assurez-vous que Go est installé sur votre machine. Vous pouvez le télécharger à partir de [ici](https://golang.org/dl/).
2. Clonez le dépôt git sur votre machine locale.
3. Naviguez dans le répertoire du projet.
4. Exécutez `go run main.go` pour démarrer le serveur.

## Utilisation

1. Naviguez vers `http://localhost:8080` dans votre navigateur web. Vous verrez un formulaire où vous pouvez entrer du JSON.
2. Entrez votre JSON dans le formulaire et cliquez sur "Submit". Vous serez redirigé vers une page qui affiche le JSON formaté.
3. Vous pouvez également faire une requête POST à `http://localhost:8080/format` avec le JSON que vous souhaitez formater en tant que corps de la requête. L'API retournera le JSON formaté.

## Support

Si vous rencontrez des problèmes, veuillez ouvrir un problème sur GitHub.

## Contribution

Les contributions sont les bienvenues. Veuillez ouvrir une Pull Request.

---
