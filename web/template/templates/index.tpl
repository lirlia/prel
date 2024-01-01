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
    <style>
        body {
            background-color: #f4f4f4;
            font-family: 'Arial', sans-serif;
            margin: 0;
            padding: 0;
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
        }
        .container {
            text-align: center;
            background-color: white;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0,0,0,0.05);
            max-width: 600px;
            width: 100%;
            box-sizing: border-box;
            max-height: 90vh;
            overflow-y: auto;
        }
        .container p {
            margin-bottom: 60px;
        }
        .brand {
            font-size: 2.5em;
            color: #333;
            margin-bottom: 20px;
        }
        .signin-button {
            background-color: #4285f4;
            color: white;
            border: none;
            border-radius: 5px;
            padding: 10px 20px;
            font-size: 16px;
            cursor: pointer;
            transition: background-color 0.3s;
        }
        .signin-button:hover {
            background-color: #3069f0;
        }
        .footer {
            text-align: right;
            margin-top: 40px;
            font-size: 0.9em;
            color: #666;
        }

        @media (max-width: 768px) {
            .container {
                margin-top: 20px;
                padding: 20px;
            }
        }
    </style>
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
