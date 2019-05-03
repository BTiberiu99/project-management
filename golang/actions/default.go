package actions

var defaultTemplateVUE = `<!doctype html>
<html lang="<%= locale %>">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title><%= t("name") %></title>

    <link href=backend/assets/css/app.css rel=preload as=style>
    <link href=backend/assets/css/app.css rel=stylesheet>

    <link href=backend/assets/css/vendors.css rel=preload as=style>
    <link href=backend/assets/css/vendors.css rel=stylesheet>

    <link href=backend/assets/js/baron.vendors.js rel=preload as=script>
    <link href=backend/assets/js/baron.app.js rel=preload as=script>
</head>

<body>
    <script type="text/javascript">
         <%= if (current_user) { %>
             var user_data = <%=current_user%>
         <% } %>
    </script>

    <noscript><strong>We're sorry but vue doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript>
    <div id="app"></div>

    <script src=backend/assets/js/baron.vendors.js></script>
    <script src=backend/assets/js/baron.app.js></script>
</body>

</html>
`