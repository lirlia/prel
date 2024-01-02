<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.Name}}</title>
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
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
            max-width: 600px;
            width: 100%;
            sizing: border-box;
        }

        .container p {
            margin: 30px 0px;
        }

        .brand {
            font-size: 2.5em;
            color: #333;
            margin-bottom: 20px;
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
        <h1>{{.Name}}</h1>
        <p class="text-break">{{.Description}}</p>
        <a href="/" class="signin-button">Go Home</a>
    </div>
    <script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
</body>

</html>
