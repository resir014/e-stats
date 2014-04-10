<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>PPA Stats</title>
  <link rel="stylesheet" type="text/css" href="css/normalize.css">
  <link rel="stylesheet" type="text/css" href="css/style.css">
  <link rel="stylesheet" href="http://fonts.googleapis.com/css?family=Open+Sans:300,300italic,400,400italic,700,700italic">
</head>
<body>

  <header>
    <div class="container">
      <h1>elementary OS PPA</h1>
      <p>Package Download Stats</p>
    </div>
  </header>

  <section>
    <div class="container">
      <table>
        <thead>
          <tr>
            <th>Package Name</th>
            <th>Downloads</th>
          </tr>
        </thead>
        
        <tbody>
          {{ range . }}
          <tr>
            <td>{{ .Name }} </td>
            <td>{{ .Downloads }} </td>
          </tr>
          {{ end }}
        </tbody>
      </table>
    </div>
  </section>
  
  <footer>
    <div class="container">
      <p>here be footer - feel free to add something here.</p>
    </div>
  </footer>
</body>
</html>
