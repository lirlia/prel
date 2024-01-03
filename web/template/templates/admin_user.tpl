<!DOCTYPE html>
<html lang="en">
<head>
    <link rel="stylesheet" href="/static/css/admin_user.css">
    {{template "header" .}}
</head>
<body>
    <div class="container-fluid">
        <h2>User Management</h2>
        <div class="controls-container">
            <div class="pagination-container">
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
            <div class="invitation-form-container">
                <form id="invitationForm" class="form-inline">
                    <input type="email" class="form-control" id="inviteeEmail" placeholder=" Email" required style="width: 350px;">
                    <select class="form-control" id="inviteeRole">
                        {{range .AdminListPage.UserRoles}}
                        <option value="{{.}}">{{.}}</option>
                        {{end}}
                    </select>
                    <button type="submit" id="invite" class="btn btn-primary">Invite</button>
                </form>
            </div>
        </div>
        <div class="table-responsive">
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col">Email</th>
                        <th scope="col" class="center-align">User Role <i class="fas fa-info-circle" title="Roles include requester, judger, and admin. Requesters can only make requests, judgers can make and approve requests, and admins have judger privileges plus administrative access."></i></th>
                        <th scope="col" class="center-align">Available <i class="fas fa-info-circle" title="Turning this off will prevent the user from sign in."></i></th>
                        <th scope="col" class="center-align">Last SignIn Time</th>
                    </tr>
                </thead>
                <tbody></tbody>
            </table>
        </div>
    </div>
    <script type="text/javascript">
        var roles = [
            {{range .AdminListPage.UserRoles}}
            "{{.}}",
            {{end}}
        ];
    </script>
    <script src="/static/js/admin_user.js"></script>
</body>
</html>
