<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="robots" content="noindex, nofollow">
    <meta name="referrer" content="no-referrer">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{ .HeaderData.AppName }} - Welcome</title>
    <link rel="icon" href="https://raw.githubusercontent.com/lirlia/prel/main/static/favicon.ico">
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/index.css">
</head>
<body>
    <div class="container">
        <div class="brand">
            <img src="https://raw.githubusercontent.com/lirlia/prel/main/images/prel-banner.png" width="60%">
        </div>
        <p class="text-break">Manage Google Cloud IAM roles<br>with time-limited access.</p>
        <form action="/signin" method="post">
            <button type="submit" class="signin-button">Sign in with Google</button>
        </form>
            <div class="footer">
        Created by <a href="https://github.com/lirlia">lirlia</a>
    </div>
    </div>

    <script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
</body>
</html>
