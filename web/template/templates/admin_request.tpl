<!DOCTYPE html>
<html lang="en">
<head>
    <link rel="stylesheet" href="/static/css/admin_request.css">
    {{template "header" .}}
</head>
<body>
    <div class="container-fluid">
        <h2>Requests List</h2>
            <div class="controls-container">
                <div class="select-container">
                    <select id="pageSize" class="form-control">
                        {{range .AdminListPage.Options}}
                            <option value="{{.}}">{{.}}</option>
                        {{end}}
                    </select>
                </div>
                <nav aria-label="Page navigation">
                    <ul class="pagination">
                        <li class="page-item"><a class="page-link" href="#" id="prevPage">&lt;</a></li>
                        <li class="page-item"><a class="page-link" href="#" id="nextPage">&gt;</a></li>
                    </ul>
                </nav>
            </div>
            <div class="table-responsive">
                <table class="table">
                <thead>
                    <tr>
                        <th scope="col">Requester</th>
                        <th scope="col">Judger</th>
                        <th scope="col">Project ID</th>
                        <th scope="col">IAM Roles</th>
                        <th scope="col" class="center-align">Period</th>
                        <th scope="col">Reason</th>
                        <th scope="col" class="center-align">Status</th>
                        <th scope="col" class="center-align">Request Time</th>
                        <th scope="col" class="center-align">Judge Time</th>
                        <th scope="col" class="center-align">Request Expiration Time</th>
                    </tr>
                </thead>
                <tbody></tbody>
            </table>
        </div>
    </div>
    <script src="/static/js/admin_request.js"></script>
</body>
</html>
