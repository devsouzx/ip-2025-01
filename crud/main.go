package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Names struct {
	Id int
	Name string
	Email string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "devsouzx"
	dbName := "crud"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("tmpl/*"))

func main() {
	log.Println("Server started on: http://localhost:9000")

	// Gerencia as URLs
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)

	// Ações
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	http.ListenAndServe(":9000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com o banco de dados utilizando a função dbConn()
	db := dbConn()
	// Realiza a consulta com banco de dados e trata erros
	selDB, err := db.Query("SELECT * FROM names ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Names{}

	// Monta um array para guardar os valores da struct
	res := []Names{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variáveis
		var id int
		var name, email string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = id
		n.Name = name
		n.Email = email

		// Junta a Struct com Array
		res = append(res, n)
	}

	// Abre a página Index e exibe todos os registrados na tela
	tmpl.ExecuteTemplate(w, "Index.html", res)

	// Fecha a conexão
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	// Pega o ID do parametro da URL
	nId := r.URL.Query().Get("id")

	// Usa o ID para fazer a consulta e tratar erros
	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	n := Names{}

	for selDB.Next() {
		var id int
		var name, email string

		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		n.Id = id
		n.Name = name
		n.Email = email
	}

	tmpl.ExecuteTemplate(w, "Show.html", n)

	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New.html", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	n := Names{}

	for selDB.Next() {
		var id int
		var name, email string

		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		n.Id = id
		n.Name = name
		n.Email = email
	}

	tmpl.ExecuteTemplate(w, "Edit.html", n)

	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {
		// Pega os campos do formulário
		name := r.FormValue("name")
		email := r.FormValue("email")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("INSERT INTO names(name, email) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulario com a SQL tratada e verifica errors
		insForm.Exec(name, email)

		log.Println("INSERT: Name: " + name + " | E-mail: " + email)
	}

	defer db.Close()

	//Retorna a HOME
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		id := r.FormValue("uid")

		insForm, err := db.Prepare("UPDATE names SET name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, email, id)

		log.Println("UPDATE: Name: " + name + " |E-mail: " + email)
	}

	defer db.Close()

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	nId := r.URL.Query().Get("id")

	delForm, err := db.Prepare("DELETE FROM names WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	delForm.Exec(nId)

	log.Println("DELETE")

	defer db.Close()

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}