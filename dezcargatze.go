package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
)

func main() {
	// Paramètres de connexion
	host := "X.X.X.X"
	port := "?????"
	address := host + ":" + port

	// Connexion au serveur
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Erreur de connexion:", err)
	}
	defer conn.Close()

	// Lecture des données initiales du serveur
	buf := readResponse(conn)
	fmt.Printf("Données reçues : %s\n", buf)

	// Boucle d'itération : traitement et réponse jusqu'à ce qu'il n'y ait plus de chaîne Base64
	for buf != "" {
		buf = processMessage(conn, buf)
	}

	// Fermeture de la connexion
	fmt.Println("Déconnexion.")
}

func sendCmd(conn net.Conn, cmd string) {
	// Envoie une commande au serveur via la connexion
	fmt.Printf(" => %s\n", cmd)
	conn.Write([]byte(cmd + "\r\n"))
}

func processMessage(conn net.Conn, buf string) string {
	// Recherche de la chaîne Base64 dans la réponse
	re := regexp.MustCompile(`my string is '(.*?)'`)
	match := re.FindStringSubmatch(buf)

	if len(match) > 1 {
		encodedString := match[1]
		fmt.Printf("Chaîne Base64 trouvée : %s\n", encodedString)

		// Décodage Base64
		decodedData, err := base64.StdEncoding.DecodeString(encodedString)
		if err != nil {
			log.Fatal("Erreur lors du décodage Base64:", err)
		}

		// Décompression zlib des données décodées
		reader, err := zlib.NewReader(bytes.NewReader(decodedData))
		if err != nil {
			log.Fatal("Erreur lors de la décompression zlib:", err)
		}
		defer reader.Close()

		// Lecture des données décompressées
		var decompressedData bytes.Buffer
		_, err = io.Copy(&decompressedData, reader)
		if err != nil {
			log.Fatal("Erreur lors de la copie des données décompressées:", err)
		}

		// Conversion du message décompressé en string
		decompressedMessage := decompressedData.String()
		fmt.Printf("Message décompressé : %s\n", decompressedMessage)

		// Envoi de la réponse décompressée au serveur
		conn.Write([]byte(decompressedMessage + "\n"))
		fmt.Printf("Réponse envoyée : %s\n", decompressedMessage)

		// Lecture de la réponse du serveur après l'envoi
		buf = readResponse(conn)
		fmt.Printf("Réponse du serveur après l'envoi : %s\n", buf)

		return buf
	} else {
		fmt.Println("Chaîne Base64 introuvable.")
		return ""
	}
}

func readResponse(conn net.Conn) string {
	// Lit les données du serveur
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Erreur de lecture:", err)
	}

	return string(buf[:n])
}

