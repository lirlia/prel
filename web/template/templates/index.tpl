<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .HeaderData.AppName }} - Welcome</title>
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet">
    <style>
        body {
            background-color: #f4f4f4;
            font-family: 'Arial', sans-serif;
            margin: 0;
            padding: 0;
            height: 100vh;
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
    </style>
</head>
<body>
    <div class="container">
        <div class="brand">{{ .HeaderData.AppName }}</div>
        <p>Manage Google Cloud IAM roles with time-limited access.</p>
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
