<html><body>
<head>
<title>Welcome Page</title>
</head>
<h1>Hello, World!</h1>
<h2>This is a gps-data testing server.</h2>
  {{if .LoginError}}<p style="color:red">Either username or password is not in our record! Sign Up?</p>{{end}}

  <form method="post" action="/login">
          {{if .Username}}
                   <p><b>{{.Username}}</b>, you're already logged in! <a href="/logout">Logout!</a></p>
          {{else}}
                  <label>Username:</label>
                  <input type="text" name="Username"><br>

                  <label>Password:</label>
                  <input type="password" name="Password">

                  <span style="font-style:italic"> Enter: 'mynakedpassword'</span><br>
                  <input type="submit" name="Login" value="Let me in!">
          {{end}}
  </form>
  </body></html>