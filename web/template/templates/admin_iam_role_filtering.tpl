<!DOCTYPE html>
<html lang="en">
<head>
    <link rel="stylesheet" href="/static/css/admin_iam_role_filtering.css">
    {{template "header" .}}
</head>
<body>
    <div class="container-fluid">
        <h2>IAM Role Filtering</h2>
        <p>Matching keywords validate IAM roles.</p>
        <p><strong>Note:</strong> No regular expressions.</p>
        <div class="controls-container">
            <div class="input-container">
                <input type="text" id="filterInput" class="form-control" placeholder="Enter a keyword between 3 and 20 characters" minlength="3" maxlength="20">
                <div id="rule-warning" class="alert alert-warning" style="display: none;">
                    Keyword is between 3 ~ 20 characters
                </div>
            </div>
            <div class="button-container">
                <button id="addFilter" class="btn btn-primary">Add Keyword</button>
            </div>
        </div>
        <div class="rules-table-responsive">
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col-1">Keyword</th>
                        <th scope="col-1">Action</th>
                    </tr>
                </thead>
                <tbody id="rulesList"></tbody>
            </table>
        </div>
    </div>
    <script src="/static/js/admin_iam_role_filtering.js"></script>
</body>
</html>
