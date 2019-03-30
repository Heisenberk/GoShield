package crypto

import "crypto/sha256"
import "math/rand"
import "time"
import "github.com/Heisenberk/goshield/structure"
//import "encoding/hex"

// CreateSalt va créer un sel pseudo-aléatoire de 15 octets. 
func CreateSalt() []byte {
	rand.Seed(time.Now().UnixNano())
	salt := make([]byte, 15)
	rand.Read(salt)
	return salt
}

// CreateHash va calculer SHA256(salt|password).
func CreateHash(doc *structure.Documents) {

	password := doc.Password

	// création du sel et initialisation du hash. 
	salt := CreateSalt() 
	doc.Salt=salt
    hash := sha256.New()

    // concaténation du sel avec le mot de passe. 
    concat := append(salt, password...)

    // génération du hash. 
    hash.Write(concat)
    doc.Hash = hash.Sum(nil)

}