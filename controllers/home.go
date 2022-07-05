package controllers

import (
    "net/http"
    "strings"
    
    "github.com/gorilla/sessions"
    "github.com/labstack/echo/v4"

    "go-trial/models"
    "go-trial/services"
    "go-trial/validation"
)

const errorsSession = "errors"

// use session key
func getCookieStore() *sessions.CookieStore {
	sessionKey := "test-session-key"
	return sessions.NewCookieStore([]byte(sessionKey))
}

// struct
type ListRquest struct{
    Cards []string `json:"cards"`
}

type ListResponse struct{
	Result []CardResponse `json:"result"`
}

type CardResponse struct{
    Card string `json:"card"`
    Hand string `json:"hand"`
    Best bool `json:"best"`
}

// show home(GET:/)
func Home(c echo.Context) error {
    // get session
    session, _ := getCookieStore().Get(c.Request(), errorsSession)
    texts := session.Flashes("text")
    text := ""
    if (len(texts) > 0) {
        text, _ = texts[0].(string)
    }

    rules := session.Flashes("rule")
    rule := ""
    if (len(rules) > 0) {
        rule, _ = rules[0].(string)
    }

    // get errors
    errorSess := session.Flashes("error")
    var errors interface{}
    if (len(errorSess) > 0) {
        errors = errorSess[0]
    }
    
    // save session
    session.Save(c.Request(), c.Response())

    return c.Render(http.StatusOK, "home", map[string]interface{} {
        "text": text,
        "rule": rule,
        "errors": errors,
    })
}

// API check text (PUT:/check)
func Check(c echo.Context) error {
    // get input value
    text := c.FormValue("text")

    cardList := strings.Split(text, " ")

    // validation
    errorMessages := validation.CheckTextValidation(cardList)

    // check
    checkResult := models.Rule{}
    if (len(errorMessages) == 0) {
        checkResult = services.CheckRole(cardList)
    }

    // add session
    session, _ := getCookieStore().Get(c.Request(), errorsSession)
    session.AddFlash(text, "text")
    session.AddFlash(checkResult.Text, "rule")
    session.AddFlash(errorMessages, "error")

    // save session
    session.Save(c.Request(), c.Response())

    return c.Redirect(http.StatusFound, "/")
}

// API check text list (POST:/check/list)
func ListCheck(c echo.Context) error {
    req := new(ListRquest)
    if err := c.Bind(req); err != nil {
        return err
    }

    response := make([]CardResponse, len(req.Cards))
    bestIndex := 0
    bestRule := 10 // no pair
    for i, v := range req.Cards {
        cardList := strings.Split(v, " ")
        errorMessages := validation.CheckTextValidation(cardList)
        if (len(errorMessages) > 0) {
            return c.JSON(http.StatusBadRequest, "カードリストが不正です")
        }

        // check
        checkResult := services.CheckRole(cardList)
        response[i] = CardResponse{
            Card: v,
            Hand: checkResult.Text,
            Best: false,
        }
        if (bestRule > checkResult.Val) {
            bestIndex = i
            bestRule = checkResult.Val
        }
    }
    response[bestIndex].Best = true

    return c.JSON(http.StatusOK, ListResponse{response})
}
