<!DOCTYPE html>
<html lang="en">
<head>
    <link rel="stylesheet" href="/static/css/request.css">
    {{template "header" .}}
</head>
<body>
    <div class="container-fluid">
        <h2>Pending Requests</h2>
            <div class="table-responsive">
                <table class="table">
                <thead>
                    <tr>
                        <th class="center-align" scope="col">Actions</th>
                        <th scope="col">Requester</th>
                        <th scope="col">Project ID</th>
                        <th scope="col">Iam Roles</th>
                        <th scope="col" class="center-align">Period</th>
                        <th scope="col">Reason</th>
                        <th scope="col" class="center-align">Request Expiration Time</th>
                        <th class="center-align" scope="col">Status</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .RequestPage.Requests}}
                    <tr>
                        <td class="align-middle center-align">
                            <div class="action-buttons">
                                {{ if eq .Status "pending" }}
                                    {{ if .CanJudge }}
                                        <button id="btn-approve" class="btn btn-success" onclick="updateRequestStatus('{{.ID}}', 'approve')">Approve</button>
                                        <button id="btn-reject" class="btn btn-warning" onclick="updateRequestStatus('{{.ID}}', 'reject')">Reject</button>
                                    {{ end }}
                                    {{ if .CanDelete }}
                                      <button id="btn-delete" class="btn btn-sm btn-outline-danger" onclick="deleteRequest('{{.ID}}')" aria-label="Delete">Delete</button>
                                    {{ end }}
                                {{ end }}
                            </div>
                        </td>
                        <td class="align-middle email">{{.Requester}}</td>
                        <td class="align-middle project-id">{{.ProjectID}}</td>
                        <td class="align-middle iam-roles">
                        {{range .IamRoles}}
                        {{.}}<br>
                        {{end}}
                        </td>
                        <td class="align-middle center-align period">{{.Period}}</td>
                        <td class="align-middle reason">{{.Reason}}</td>
                        <td class="align-middle center-align expire-time">{{.ExpiredAt}}</td>
                        <td class="align-middle center-align status">{{.Status}}</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
    </div>
    <script src="/static/js/request.js"></script>
</body>
</html>
