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

        .center-align {
            text-align-last: center;
            -moz-text-align-last: center;
            text-align: center;
        }

        td.center-align {
            text-align: center;
        }

        @media (max-width: 768px) {
            .table-responsive {
                overflow-x: auto;
            }

            .table {
                min-width: 1000px;
            }

            .table th, .table td {
                white-space: nowrap;
            }

            .table .center-align {
                text-align: center;
            }
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
