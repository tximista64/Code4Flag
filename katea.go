package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"regexp"
)

func main() {
	// Paramètres de connexion
	host := "XXXXX"
	port := "?????"
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

	// Recherche de la chaîne Base64 dans la réponse
	re := regexp.MustCompile(`my string is '(.*?)'`)
	match := re.FindStringSubmatch(data)

	if len(match) > 1 {
		base64String := match[1]
		fmt.Println("Chaîne Base64 trouvée : ", base64String)

		// Décodage Base64
		decodedMessage, err := base64.StdEncoding.DecodeString(base64String)
		if err != nil {
			log.Fatal("Erreur lors du décodage Base64:", err)
		}

		// Conversion du message décodé en string
		decodedMessageStr := string(decodedMessage)
		fmt.Println("Valeur décodée : ", decodedMessageStr)

		// Envoi de la réponse au serveur
		_, err = conn.Write([]byte(decodedMessageStr + "\n"))
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
		log.Fatal("Chaîne Base64 introuvable.")
	}

	// Fermeture de la connexion
	fmt.Println("Déconnexion.")
}

