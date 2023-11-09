package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// Mengambil data dari form body
	fullName := c.PostForm("FullName")
	email := c.PostForm("Email")
	password := c.PostForm("Password")
	accessType := c.PostForm("AccessType")
	// Memeriksa apakah salah satu data kosong
	if fullName == "" || email == "" || password == "" || accessType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Harap isi semua field"})
		return
	}
	// Mengkonversi string "true" menjadi true, yang merupakan tipe bool
	accType, err := strconv.ParseBool(accessType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipe akses tidak valid"})
		return
	}

	// Menginisialisasi koneksi database
	db := connect()
	// Inisialisasi struct newUser
	db.AutoMigrate(&User{})

	// Enkripsi kata sandi menggunakan SHA-256
	hashedPassword := hashPassword(password)

	newUser := User{
		FullName: fullName,
		Email:    email,
		Password: hashedPassword,
		AccType:  accType,
	}

	// Simpan pengguna ke dalam database
	result := db.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Gagal mendaftarkan pengguna"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pengguna berhasil terdaftar"})
}
func hashPassword(password string) string {
	// Membuat instance SHA-256 hasher
	hasher := sha256.New()

	// Mengonversi kata sandi ke byte array
	passwordBytes := []byte(password)

	// Menghitung hash SHA-256 dari kata sandi
	hasher.Write(passwordBytes)
	hashedBytes := hasher.Sum(nil)

	// Mengonversi hasil hash ke string heksadesimal
	hashedPassword := hex.EncodeToString(hashedBytes)

	return hashedPassword
}
func Login(c *gin.Context) {
	// Mengambil data dari form body
	email := c.PostForm("Email")
	password := c.PostForm("Password")

	// Memeriksa apakah data email atau password kosong
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email dan password harus diisi"})
		return
	}

	// Menginisialisasi koneksi database
	db := connect()

	// Mencari pengguna berdasarkan email
	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	// Memeriksa kecocokan password
	// Mengenkripsi kata sandi yang diinputkan oleh pengguna menggunakan SHA-256
	inputPasswordHash := hashPassword(password)

	// Memeriksa apakah kata sandi yang diinputkan sama dengan kata sandi yang ada dalam database
	if inputPasswordHash != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	// Jika login sukses, Anda dapat memanggil generateToken berdasarkan AccType pengguna
	jwtToken, err := generateToken(c, int(user.ID), user.FullName, user.Email, user.AccType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghasilkan token"})
		return
	}

	// Mengirim pesan sukses, nama pengguna, dan token sebagai respons
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil login", "name": user.FullName, "token": jwtToken})
}

func Logout(c *gin.Context) {
	// Menghapus token dengan mengganti nilai token dengan string kosong
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil logout"})
}
