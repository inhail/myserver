<html><body>

  {{if .Username}}
          <p><b>{{.Username}}</b>, welcome to your dashboard! <a href="/logout">Logout!</a></p>
  {{else}}
          <p>Either your session has expired or you've logged out! <a href="/login">Login</a></p>
  {{end}}

  </body></html>