# Conversor de Texto com Golang e Vue.js 🚀

Este projeto é uma aplicação simples que permite aos usuários converter texto em diferentes estilos, como negrito, itálico, sublinhado e combinações desses. Podendo ser testado [Através desse Link](https://magenta-selkie-0df725.netlify.app), esse projeto utiliza **Golang** no backend e **Vue.js** no frontend, ele demonstra como integrar essas tecnologias de maneira eficiente e funcional.

---

## 🛠️ **Tecnologias Utilizadas**

### **Backend**
- **Linguagem:** Golang
- **Servidor HTTP:** Embutido no Go
- **Manipulação de Texto:** Unicode

### **Frontend**
- **Framework:** Vue.js (via CDN)
- **HTML:** Interface simples e responsiva
- **TailwindCSS:** Estilização da interface

---

## 🚀 **Como Rodar o Projeto**

### **Pré-requisitos**
- [Golang instalado](https://golang.org/dl/)
- Um navegador web para acessar a interface

### **Passo a Passo**

1. Clone o repositório:
   ```bash
   git clone https://github.com/MaiconGavino/TextConvert.git
   cd TextConvert
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Inicie o servidor backend:
   ```bash
   go run main.go
   ```

4. Acesse a interface no navegador:
   ```plaintext
   http://localhost:8080
   ```

---

## 📂 **Estrutura do Projeto**

```plaintext
TextConvert/

├── main.go          # Código do servidor em Golang
├── go.mod           # Gerenciador de dependências do Go
└── template/
    ├── index.html   # Interface do usuário
```

---

## ✨ **Melhorias Futuras**

- [ ] Adicionar suporte a novos estilos de texto.
- [ ] Implementar opções de exportação (PDF, TXT).
- [ ] Criar testes unitários para as funções de backend.

---

## 📥 **Contribua**

Sinta-se à vontade para melhorar este projeto! Envie um pull request ou abra uma issue no repositório.

---

## 💡 **Autor**

Desenvolvido por **Maicon Gavino**  
[LinkedIn](https://www.linkedin.com/in/maicongavino) | [GitHub](https://github.com/MaiconGavino)

---

## 🔗 **Licença**

Este projeto é licenciado sob a [MIT License](LICENSE).

