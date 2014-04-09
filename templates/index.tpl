<html>
<head>
  <title>Test</title>
</head>
<body>
  <table>
  {{ range . }}
  <tr>
    <td> {{ .Name }} </td>
    <td> {{ .Downloads }} </td>
  </tr>
  {{ end }}
  </table>
</body>
</html>
