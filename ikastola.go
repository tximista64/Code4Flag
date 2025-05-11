package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
)

func main() {
	// Paramètres de connexion
	HOST := "XXXXXXXX"
	PORT := "??????"

	// Connexion au serveur
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Erreur de connexion:", err)
		return
	}
	defer conn.Close()

	// Message de connexion réussi
	fmt.Printf("Connexion vers %s:%s réussie.\n", HOST, PORT)

	// Réception des données
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erreur de lecture:", err)
		return
	}

	// Extraction des données
	donnees := string(buffer[:n])
	fmt.Println("Réception:", donnees)

	// Extraction des premiers et deuxièmes nombres
	premierNombre := donnees[171:174]
	deuxiemeNombre := donnees[191:195]

	// Conversion des sous-chaînes en nombres
	a, err := strconv.Atoi(premierNombre)
	if err != nil {
		fmt.Println("Erreur de conversion du premier nombre:", err)
		return
	}

	b, err := strconv.Atoi(deuxiemeNombre)
	if err != nil {
		fmt.Println("Erreur de conversion du deuxième nombre:", err)
		return
	}

	// Calcul des résultats
	sqrtA := math.Sqrt(float64(a))
	result := float64(b) * sqrtA
	roundedResult := math.Round(result*100) / 100 // Arrondir à 2 décimales

	// Conversion du résultat en bytes et envoi au serveur
	resultStr := fmt.Sprintf("%.2f\n", roundedResult)
	_, err = conn.Write([]byte(resultStr))
	if err != nil {
		fmt.Println("Erreur d'envoi:", err)
		return
	}

	// Réception de la réponse du serveur
	n, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Erreur de lecture de la réponse:", err)
		return
	}

	// Affichage de la réponse du serveur
	flag := string(buffer[:n])
	fmt.Println("Réponse du serveur:", flag)

	// Fermeture de la connexion
	fmt.Println("Déconnexion.")
}

