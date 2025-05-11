package main

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"unicode"
)

// Fonction pour appliquer ROT13 à une chaîne
func rot13(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			// Décalage de 13 caractères dans l'alphabet
			offset := rune(13)
			if (unicode.IsLower(r) && r >= 'n') || (unicode.IsUpper(r) && r >= 'N') {
				offset = -13
			}
			result = append(result, r+offset)
		} else {
			// Si ce n'est pas une lettre, on garde le caractère inchangé
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	// Paramètres de connexion
	host := "XXXXXX"
	port := "??????"
	address := host + ":" + port

	// Connexion au serveur
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Erreur de connexion:", err)
	}
	defer conn.Close()

	// Réception des données du serveur
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatal("Erreur de lecture:", err)
	}

	// Conversion des données reçues en string
	data := string(buf)
	fmt.Println("Données reçues : ", data)

	// Recherche de la chaîne encodée en ROT13 dans la réponse
	re := regexp.MustCompile(`my string is '(.*?)'`)
	match := re.FindStringSubmatch(data)

	if len(match) > 1 {
		rot13String := match[1]
		fmt.Println("Chaîne ROT13 trouvée : ", rot13String)

		// Décodage ROT13
		decodedMessage := rot13(rot13String)
		fmt.Println("Valeur décodée avec ROT13 : ", decodedMessage)

		// Envoi de la réponse décryptée au serveur
		_, err = conn.Write([]byte(decodedMessage + "\n"))
		if err != nil {
			log.Fatal("Erreur lors de l'envoi de la réponse:", err)
		}
		fmt.Println("Réponse envoyée.")

		// Attente de la réponse du serveur
		_, err = conn.Read(buf)
		if err != nil {
			log.Fatal("Erreur lors de la réception de la réponse:", err)
		}

		// Affichage de la réponse du serveur
		fmt.Println("Réponse du serveur : ", string(buf))
	} else {
		log.Fatal("Chaîne ROT13 introuvable.")
	}

	// Fermeture de la connexion
	fmt.Println("Déconnexion.")
}

