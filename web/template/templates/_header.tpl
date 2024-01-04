{{define "header"}}
<title>{{ .HeaderData.AppName }}</title>
<meta charset="UTF-8">
<meta name="robots" content="noindex, nofollow">
<meta name="referrer" content="no-referrer">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="icon" href="https://raw.githubusercontent.com/lirlia/prel/main/static/favicon.ico">
<link href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet">
<link rel="stylesheet" href="/static/css/header.css">
<script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
<div class="global-header">
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <a class="navbar-brand" href="/"><img src="https://raw.githubusercontent.com/lirlia/prel/main/images/preln.png" width="50px">{{ .HeaderData.AppName }}</a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav"
            aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item">
                    <a class="nav-link" href="/request-form">📝 Request Form</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/request">📋 Pending Requests</a>
                </li>
                {{ if .HeaderData.IsAdmin }}
                <li class="nav-item">
                    <hr class="d-none d-lg-block vertical-divider">
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/admin/request">📚 Requests List</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/admin/iam-role-filtering">⚙️ IAM Role Filtering</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/admin/user">👥 User Management</a>
                </li>
                {{ end }}
            </ul>
            <ul class="navbar-nav">
                <li class="nav-item">
                    <form action="/signout" method="post" style="display: inline;">
                        <button type="submit" class="btn btn-link nav-link">🚪 Sign out</button>
                    </form>
                </li>
            </ul>
        </div>
    </nav>
</div>
{{end}}
