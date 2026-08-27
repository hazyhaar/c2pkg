package c2json

// Constantes des types de nœuds JSON (normatives)
const (
	Err    = 0
	Object = 1
	Array  = 2
	String = 3
	Number = 4
	True   = 5
	False  = 6
	Null   = 7
)

// Validate vérifie la conformité syntaxique d'une chaîne ou flux JSON.
func Validate(json []byte) bool {
	if len(json) == 0 {
		return false
	}
	return C2json_validate(json, len(json)) == 1
}

// FindKey localise une clé dans un objet JSON et retourne la tranche de sa valeur et son type.
func FindKey(json []byte, key string) (val []byte, valType int, ok bool) {
	if len(json) == 0 || len(key) == 0 {
		return nil, Err, false
	}
	var vStart, vLen, vType int
	res := C2json_find_key(json, len(json), []byte(key), len(key), &vStart, &vLen, &vType)
	if res == 0 || vStart+vLen > len(json) {
		return nil, Err, false
	}
	return json[vStart : vStart+vLen], vType, true
}

// GetString extrait une chaîne de caractères associée à une clé et décode ses échappements.
func GetString(json []byte, key string) (string, bool) {
	if len(json) == 0 || len(key) == 0 {
		return "", false
	}
	buf := make([]byte, len(json))
	n := C2json_get_string(json, len(json), []byte(key), len(key), buf, len(buf))
	if n < 0 {
		return "", false
	}
	return string(buf[:n]), true
}

// GetInt extrait un entier 64-bit associé à une clé.
func GetInt(json []byte, key string) (int64, bool) {
	if len(json) == 0 || len(key) == 0 {
		return 0, false
	}
	var out int64
	if C2json_get_int64(json, len(json), []byte(key), len(key), &out) == 0 {
		return 0, false
	}
	return out, true
}

// GetBool extrait un booléen associé à une clé.
func GetBool(json []byte, key string) (bool, bool) {
	if len(json) == 0 || len(key) == 0 {
		return false, false
	}
	var out int
	if C2json_get_bool(json, len(json), []byte(key), len(key), &out) == 0 {
		return false, false
	}
	return out == 1, true
}
