package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
    // Membuat router
    router := mux.NewRouter()

    // Mendefinisikan route untuk API
    router.HandleFunc("/items", GetItems).Methods("GET")
    router.HandleFunc("/items", CreateItem).Methods("POST")
    router.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
    router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

    // Mengaktifkan server pada port 8080
	fmt.Println("Server berjalan pada port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

// Koneksi ke database MySQL
func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root" // Ganti dengan username MySQL Anda
    dbPass := "" // Ganti dengan password MySQL Anda
    dbName := "akun_API" // Ganti dengan nama database Anda
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        log.Fatal(err)
    }
    return db
}

// Struct untuk representasi objek
type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
	Password string `json:"password"`
}

// Handler untuk mendapatkan semua item
func GetItems(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    defer db.Close()

    var items []Item

    rows, err := db.Query("SELECT * FROM pengguna")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var item Item
        err := rows.Scan(&item.ID, &item.Name, &item.Password)
        if err != nil {
            log.Fatal(err)
        }
        items = append(items, item)
    }
    json.NewEncoder(w).Encode(items)
}

// Handler untuk membuat item baru
func CreateItem(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
	query:= r.URL.Query()
	username:= query.Get("username")
	password:= query.Get("password")
    defer db.Close()

	pengguna:= Item{Name: username, Password: password}
	_, err := db.Exec("INSERT INTO pengguna (username, password) VALUES (?, ?)", pengguna.Name, pengguna.Password)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(pengguna)
}


// Handler untuk mengupdate item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    defer db.Close()

    params := mux.Vars(r)
    id := params["id"]

    query:= r.URL.Query()
	newusername:= query.Get("username")
	newpassword:= query.Get("password")
	

	_, err := db.Exec("UPDATE pengguna SET username=?, password=? WHERE id_pengguna=?", newusername, newpassword, id)
	if err != nil {
		log.Fatal(err)
	}

	updatepengguna := Item{Name: newusername, Password: newpassword}
	json.NewEncoder(w).Encode(updatepengguna)

	fmt.Fprintf(w, "Item updated successfully")
}


// Handler untuk menghapus item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    defer db.Close()

    params := mux.Vars(r)
    id := params["id"]

    _, err := db.Exec("DELETE FROM pengguna WHERE id_pengguna=?", id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(w, "Item deleted successfully")
}
