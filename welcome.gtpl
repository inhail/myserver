<html lang="en">
  <head>
        <title>Welcome Page</title>
  </head>
  <body>
       <h1>Hello, World!</h1>
       <p>This is a gps-data testing server.</p>
       <form action="http://localhost:9090/gpsdata" method="post">
        Latitude:<input type="digit" name="Latitude"> 
        Longitude:<input type="digit" name="Longitude">
        <input type="submit" value="save">
  </body>
</html>