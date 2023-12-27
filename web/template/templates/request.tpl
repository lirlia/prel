<!DOCTYPE html>
<html lang="en">
<head>
    <style>
        .container-fluid {
            padding-left: 15px;
            padding-right: 15px;
        }
        .container-fluid h2 {
            padding: 20px 0;
        }
        .table-responsive {
            overflow-x: auto;
            margin-left: -15px;
            margin-right: -15px;
        }
        .action-buttons .btn {
            margin-right: 15px;
        }
        .delete-button {
            margin-left: 30px;
        }
        .expire-time {
            display: none;
        }
    </style>
    {{template "header" .}}
</head>
<body>
    <div class="container-fluid">
        <h2>Pending Requests</h2>
            <div class="table-responsive">
                <table class="table">
                <thead>
                    <tr>
                        <th scope="col">Requester</th>
                        <th scope="col">Project ID</th>
                        <th scope="col">Iam Roles</th>
                        <th scope="col">Reason</th>
                        <th scope="col">Expire Time</th>
                        <th class="text-center" scope="col">Status</th>
                        <th class="text-center" scope="col">Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .RequestPage.Requests}}
                    <tr>
                        <td class="align-middle email">{{.Requester}}</td>
                        <td class="align-middle project-id">{{.ProjectID}}</td>
                        <td class="align-middle iam-roles">
                        {{range .IamRoles}}
                        {{.}}<br>
                        {{end}}
                        </td>
                        <td class="align-middle reason">{{.Reason}}</td>
                        <td class="align-middle expire-time">{{.ExpiredAt}}</td>
                        <td class="align-middle status text-center">{{.Status}}</td>
                        <td class="align-middle text-center">
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
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
    </div>
    <script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
    <script>
        function convertToLocaleTimeString(isoString) {
            var date = new Date(isoString);

            var options = {
                year: 'numeric',
                month: '2-digit',
                day: '2-digit',
                hour: '2-digit',
                minute: '2-digit'
            };

            return date.toLocaleString(undefined, options).replace(/\//g, '-');
        }

        $(document).ready(function() {
            $('.expire-time').each(function() {
                var originalText = $(this).text();
                var convertedText = convertToLocaleTimeString(originalText);
                $(this).text(convertedText).css('display', 'table-cell');
            });
        });

        function updateRequestStatus(requestId, status) {
            $.ajax({
                url: '/api/requests/' + requestId,
                type: 'PATCH',
                contentType: 'application/json',
                data: JSON.stringify({ status: status }),
                success: function(response) {
                    alert('Request updated successfully');
                    location.reload();
                },
                error: function() {
                    alert('Error updating request');
                }
            });
        }

        function deleteRequest(requestId) {
            $.ajax({
                url: '/api/requests/' + requestId,
                type: 'DELETE',
                success: function(response) {
                    alert('Request deleted successfully');
                    var currentPath = window.location.pathname;

                    if (currentPath.includes('/request/')) {
                        window.location.pathname = '/request-form';
                    } else {
                        location.reload();
                    }
                },
                error: function() {
                    alert('Error deleting request');
                }
            });
        }
    </script>
</body>
</html>
