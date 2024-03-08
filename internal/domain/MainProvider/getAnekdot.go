package MainProvider

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

var anekdotik string

func (d *Domain) GetAnekdot(c echo.Context) error {
	response, err := http.Get(d.anekdotUrl)
	if err != nil {
		fmt.Println("Ошибка запроса.")
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка чтения тела.")
	}
	rresponse := "<span style=\"font-size: 21px;\">" + string(body) + "</span>"
	anekdotik = string(body)
	html := `
	<!DOCTYPE html>
	  <html>
	    <head>
	      <style>
	        body {
	          background-color: white;
	          font-size: 25px;Z
	          text-align: center;
	        }
	      </style>
	    </head>
	    <body>
	<script async src="https://telegram.org/js/telegram-widget.js?22" data-telegram-login="Rjaka_prikol_bot" data-size="large" data-auth-url="https://bba3icajjuulniimgp15.containers.yandexcloud.net/getcalmar" data-request-access="write"></script>
<script type="text/javascript">
  function handleTelegramAuth(user) {
    alert('Logged in as ' + user.first_name + ' ' + user.last_name + ' (' + user.id + (user.username ? ', @' + user.username : '') + ')');
    // Здесь можно добавить код для отправки данных пользователя на сервер или выполнения других действий при успешной авторизации
  }
</script>

	  
      <div id="joke">
<div style="display: flex; justify-content: center;">
  <h1>Анекдоты</h1>
</div>
<div style="display: flex; justify-content: center;">` + rresponse + `</div>
<div style="display: flex; justify-content: center;">
  <button onclick="generateJoke()">Сгенерировать новый анекдот</button>
</div>
<a href="https://t.me/Rjaka_prikol_bot" target="_blank">
  <img src="https://www.info-expert.ru/upload/medialibrary/74e/74e6e22523f9eb3f13472ca2039257fd.png" alt="Telegram Bot Icon" style="width: 50px;">
  <button>Перейти к боту</button>
</a>
</div>

	      
	      
	      <script>
	        function generateJoke() {
	          fetch('/getcalmar')
	            .then(response => response.text())
	            .then(joke => {
	              document.getElementById('joke').innerHTML = joke;
	            });
	        }
	
	        generateJoke();
	      </script>
	    </body>
	  </html>
	`
	SendAnecdoteToTelegram(anekdotik)
	return c.HTML(http.StatusOK, html)
}
