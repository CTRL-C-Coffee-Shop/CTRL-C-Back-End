package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please fill in all fields"})
		return
	}
	// Mengkonversi string "true" menjadi true, yang merupakan tipe bool
	accType, err := strconv.ParseBool(accessType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid access type"})
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
		Url:      "",
	}

	// Simpan pengguna ke dalam database
	result := db.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully registered"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password must be filled"})
		return
	}

	// Menginisialisasi koneksi database
	db := connect()

	// Mencari pengguna berdasarkan email
	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password"})
		return
	}

	// Memeriksa kecocokan password
	// Mengenkripsi kata sandi yang diinputkan oleh pengguna menggunakan SHA-256
	inputPasswordHash := hashPassword(password)

	// Memeriksa apakah kata sandi yang diinputkan sama dengan kata sandi yang ada dalam database
	if inputPasswordHash != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password"})
		return
	}

	// Jika login sukses, Anda dapat memanggil generateToken berdasarkan AccType pengguna
	jwtToken, err := generateToken(int(user.ID), user.FullName, user.Email, user.AccType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.SetCookie(tokenName, jwtToken, 0, "/", "", false, true)

	// Mengirim pesan sukses, nama pengguna, dan token sebagai respons
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user_id": user.ID, "full_name": user.FullName, "email": user.Email, "token": jwtToken, "access_type": user.AccType})
}

func UpdateUser(c *gin.Context) {
	fullName := c.PostForm("FullName")
	email := c.PostForm("Email")
	url := c.PostForm("Url")
	userID, _ := strconv.Atoi(c.PostForm("userID"))

	db := connect()
	var user User
	result := db.Where("email = ?", email).First(&user)

	if result.Error == nil { //periksa lagi, tujuan: ngecek buat jangan ada email yang sama
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email already exist"})
		return
	}

	query := db.Where("id_user = ?", userID).First(&user)
	if query.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	query = db.Table("users").Where("id_user = ?", userID).Select("fullname", "email", "url").Updates(User{FullName: fullName, Email: email, Url: url})
	if query.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to update user's data"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "successfully update data"})
}

func ChangePass(c *gin.Context) { //for change password in edit profile menu
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	password1 := c.PostForm("Password1") //old password
	password2 := c.PostForm("Password2") //new password

	if password1 == password2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter a different password"})
		return
	}

	password2 = hashPassword(password2)
	db := connect()
	var user User
	result := db.Model(user).Where("id_user = ?", userID).Update("password", password2)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to change user's password"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Password updated"})
}
