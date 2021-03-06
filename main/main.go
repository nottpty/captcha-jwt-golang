package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jakkrabig/captcha/bankaccount"

	"github.com/robbert229/jwt"

	"github.com/jakkrabig/captcha"
)

type apiHandler struct{}

type tokenHandler struct{}

type validateHandler struct{}

type findAccountHandler struct{}

type createBankAccountHandler struct{}

type depositBankAccountHandler struct{}

type withdrawBankAccountHandler struct{}

type balanceBankAccountHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	str := strings.Split(p, "/")[2]

	pattern, _ := strconv.Atoi(string(str[0]))
	leftOperand, _ := strconv.Atoi(string(str[1]))
	operator, _ := strconv.Atoi(string(str[2]))
	rightOperand, _ := strconv.Atoi(string(str[3]))
	captchaObj := captcha.Captcha(pattern, leftOperand, operator, rightOperand)
	fmt.Fprintf(w, "%s", captchaObj.String())
}

func (tokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	numMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	captcha := r.Header.Get("captcha")
	captchaSplit := strings.Split(captcha, " ")
	leftStr := captchaSplit[0]
	rightStr := captchaSplit[2]
	var leftInt int
	var rightInt int
	if len(leftStr) == 1 {
		leftIntTemp, err := strconv.Atoi(leftStr)
		leftInt = leftIntTemp
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("captcha incorrect!"))
			return
		}
	} else {
		leftInt = numMap[leftStr]
	}
	if len(rightStr) == 1 {
		rightIntTemp, err := strconv.Atoi(rightStr)
		rightInt = rightIntTemp
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("captcha incorrect!"))
			return
		}
	} else {
		rightInt = numMap[rightStr]
	}
	operator := captchaSplit[1]
	var result int
	if operator == "+" {
		result = leftInt + rightInt
	} else if operator == "-" {
		result = leftInt - rightInt
	} else if operator == "*" {
		result = leftInt * rightInt
	}
	userResult, err := strconv.Atoi(captchaSplit[4])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("captcha incorrect!"))
		return
	}
	if result != userResult {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("captcha incorrect!"))
		return
	}

	secret := "ThisIsMyPassword"
	algorithm := jwt.HmacSha256(secret)

	claims := jwt.NewClaim()
	claims.Set("Role", "Admin")
	claims.SetTime("exp", time.Now().Add(time.Minute))

	token, err := algorithm.Encode(claims)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(token))

}

func (validateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	secret := "ThisIsMyPassword"
	algorithm := jwt.HmacSha256(secret)

	loadedClaims, err := algorithm.Decode(token)
	if err != nil {
		panic(err)
	}

	role, err := loadedClaims.Get("Role")
	if err != nil {
		panic(err)
	}

	roleString, ok := role.(string)
	if !ok {
		panic(err)
	}

	if strings.Compare(roleString, "Admin") == 0 {
		//user is an admin
		fmt.Println("User is an admin")
		w.Write([]byte("Login Successful"))
	}
}

func validateToken(r *http.Request) bool {
	token := r.Header.Get("Authorization")
	secret := "ThisIsMyPassword"
	algorithm := jwt.HmacSha256(secret)

	loadedClaims, err := algorithm.Decode(token)
	if err != nil {
		panic(err)
	}

	role, err := loadedClaims.Get("Role")
	if err != nil {
		panic(err)
	}

	roleString, ok := role.(string)
	if !ok {
		panic(err)
	}

	if strings.Compare(roleString, "Admin") == 0 {
		return true
	}
	return false
}

func (createBankAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if validateToken(r) {
		name := r.Header.Get("name")
		newAccount := bankaccount.New(name)
		bankaccount.Save(newAccount)
		w.Write([]byte("Create an account Successful"))
	}

}

func (findAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if validateToken(r) {
		name := r.Header.Get("name")
		account := bankaccount.FindByName(name)

		byteUser, err := json.Marshal(&account)
		if err != nil {
			http.Error(w, fmt.Sprintf("users: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", byteUser)
	}

}

func (depositBankAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if validateToken(r) {
		name := r.Header.Get("name")
		money, err := strconv.Atoi(r.Header.Get("money"))
		if err != nil {
			http.Error(w, fmt.Sprintf("users: %s", err), http.StatusInternalServerError)
			return
		}
		account := bankaccount.FindByName(name)
		account.Deposit(money)
		balanceStr := strconv.Itoa(account.Balance())
		w.Write([]byte(balanceStr))
	}
}

func (withdrawBankAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if validateToken(r) {
		name := r.Header.Get("name")
		money, err := strconv.Atoi(r.Header.Get("money"))
		if err != nil {
			http.Error(w, fmt.Sprintf("users: %s", err), http.StatusInternalServerError)
			return
		}
		account := bankaccount.FindByName(name)
		if money > account.Balance() {
			w.Write([]byte("Not enough money"))
			return
		}
		account.Withdraw(money)
		balanceStr := "Your balance : " + strconv.Itoa(account.Balance())
		w.Write([]byte(balanceStr))
	}
}

func (balanceBankAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if validateToken(r) {
		name := r.Header.Get("name")
		account := bankaccount.FindByName(name)
		balanceStr := strconv.Itoa(account.Balance())
		w.Write([]byte(balanceStr))
	}
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/captcha/", apiHandler{})
	mux.Handle("/tokens/", tokenHandler{})
	mux.Handle("/login/", validateHandler{})
	mux.Handle("/bank/account/create", createBankAccountHandler{})
	mux.Handle("/bank/account/deposit", depositBankAccountHandler{})
	mux.Handle("/bank/account/withdraw", withdrawBankAccountHandler{})
	mux.Handle("/bank/account/balance", balanceBankAccountHandler{})
	mux.Handle("/bank/account/", findAccountHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page")
	})

	http.ListenAndServe(":8080", mux)
}
