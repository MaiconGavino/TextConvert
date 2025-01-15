package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

// Middleware para configurar CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite todas as origens
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type TextRequest struct {
	Text string `json:"text"`
}

type TextResponse struct {
	Bold       string `json:"bold"`
	Italic     string `json:"italic"`
	BoldItalic string `json:"boldItalic"`
	Underline  string `json:"underline"`
}

func processText(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req TextRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Text == "" {
		http.Error(w, "Request Invalido", http.StatusBadRequest)
		return
	}

	textResp := TextResponse{
		Bold:       ConvertBold(req.Text),
		Italic:     ConvertItalic(req.Text),
		BoldItalic: ConvertBoldItalic(req.Text),
		Underline:  ConvertUnderline(req.Text),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(textResp)
}

func main() {
	// Configurando o FileServer para servir arquivos estáticos
	fileServe := http.FileServer(http.Dir("./template"))
	http.Handle("/", http.StripPrefix("/", fileServe)) // Servir arquivos estáticos sem middleware

	// Configurando a rota "/convert" com middleware de CORS
	http.Handle("/convert", enableCORS(http.HandlerFunc(processText)))

	fmt.Println("Iniciando o servidor na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}

func ConvertBold(input string) string {
	return stringFormat(input, 0x1D400, 0x1D41A)
}

func ConvertItalic(input string) string {
	return stringFormat(input, 0x1D434, 0x1D44E)
}

func ConvertBoldItalic(input string) string {
	return stringFormat(input, 0x1D468, 0x1D482)
}

func ConvertUnderline(input string) string {
	var result strings.Builder
	for _, r := range input {
		if unicode.IsSpace(r) {
			result.WriteRune(r)
		} else {
			result.WriteString(string(r) + "\u0332")
		}
	}
	return result.String()
}

func stringFormat(input string, upperOffset, lowerOffset int) string {
	var result strings.Builder
	for _, r := range input {
		if r >= 'A' && r <= 'Z' {
			result.WriteRune(rune(int(r) - 'A' + upperOffset))
		} else if r >= 'a' && r <= 'z' {
			result.WriteRune(rune(int(r) - 'a' + lowerOffset))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}