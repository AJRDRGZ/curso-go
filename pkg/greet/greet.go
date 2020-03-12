// Package greet es de prueba y retorna
// los saludos en ingles, italiano y español
package greet

var emoji = "🙋‍♀️"

// English retorna el saludo en ingles
func English() string {
	return "Hi " + emoji
}

// Italian retorna el saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}
