package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "time"
)

func main() {
        // Ouvrir le fichier des logs
        file, err := os.Open("Zentsura.txt")
        if err != nil {
                fmt.Println("Erreur lors de l'ouverture du fichier:", err)
                return
        }
        defer file.Close()

        // Liste pour stocker les horodatages
        var timeList []string

        // Lire les lignes du fichier
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                // Extraire l'horodatage de la ligne (indices 30 à 38)
                if len(line) >= 38 {
                        timeList = append(timeList, line[30:38])
                }
        }

        if err := scanner.Err(); err != nil {
                fmt.Println("Erreur lors de la lecture du fichier:", err)
                return
        }

        // Variables pour construire le mot de passe
        var char string
        var flag string

        // Parcourir la liste des horodatages
        for i := 0; i < len(timeList)-1; i++ {
                // Convertir les horodatages en objets time.Time
                t1, err1 := time.Parse("15:04:05", timeList[i])
                t2, err2 := time.Parse("15:04:05", timeList[i+1])
                if err1 != nil || err2 != nil {
                        fmt.Println("Erreur lors de la conversion des horodatages:", err1, err2)
                        return
                }

                // Calculer la différence de temps entre deux horodatages
                timeDiff := t2.Sub(t1)

                // Déterminer les bits binaires en fonction de la différence de temps
                switch i % 4 {
                case 0, 1, 2:
                        if timeDiff == 0 {
                                char += "00"
                        } else if timeDiff == 2*time.Second {
                                char += "01"
                        } else if timeDiff == 4*time.Second {
                                char += "10"
                        } else if timeDiff == 6*time.Second {
                                char += "11"
                        }
                case 3:
                        if timeDiff == 2*time.Second {
                                char += "0"
                        } else if timeDiff == 4*time.Second {
                                char += "1"
                        }

                        // Convertir la séquence binaire en caractère ASCII
                        charInt, err := strconv.ParseInt(char, 2, 64)
                        if err != nil {
                                fmt.Println("Erreur lors de la conversion binaire en entier:", err)
                                return
                        }

                        flag += string(rune(charInt))
                        char = "" // Réinitialiser pour le caractère suivant
                }
        }

        // Afficher le mot de passe reconstitué
        fmt.Println("Mot de passe:", flag)
