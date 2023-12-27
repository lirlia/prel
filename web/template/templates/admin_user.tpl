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

        .controls-container {
            display: flex;
            align-items: center;
            justify-content: space-between;
            margin-bottom: 15px;
        }

        .pagination {
            margin-bottom: 0;
        }

        .pagination-container {
            display: flex;
            align-items: center;
        }

        .select-container {
            margin-right: 15px;
        }

        #pageSize {
            width: auto;
        }

        .switch {
            position: relative;
            display: inline-block;
            width: 60px;
            height: 34px;
        }

        .switch input {
            opacity: 0;
            width: 0;
            height: 0;
        }

        .slider {
            position: absolute;
            cursor: pointer;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background-color: #ccc;
            -webkit-transition: .4s;
            transition: .4s;
        }

        .slider:before {
            position: absolute;
            content: "";
            height: 26px;
            width: 26px;
            left: 4px;
            bottom: 4px;
            background-color: white;
            -webkit-transition: .4s;
            transition: .4s;
        }

        input:checked + .slider {
            background-color: #2196F3;
        }

        .slider-disable {
            background-color: #E9ECEF !important;
        }

        input:focus + .slider {
            box-shadow: 0 0 1px #2196F3;
        }

        input:checked + .slider:before {
            -webkit-transform: translateX(26px);
            -ms-transform: translateX(26px);
            transform: translateX(26px);
        }

        .slider.round {
            border-radius: 34px;
        }

        .slider.round:before {
            border-radius: 50%;
        }

        .fas.fa-info-circle {
            color: #17a2b8;
            font-size: 0.8em;
            margin-left: 5px;
        }

        th, td {
            vertical-align: middle !important;
        }

        .center-align {
            text-align-last: center;
            -moz-text-align-last: center;
            text-align: center;
        }

        td.center-align {
            text-align: center;
        }
    </style>
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
                        <li class="page-item"><a class="page-link" href="#" id="prevPage">Previous</a></li>
                        <li class="page-item"><a class="page-link" href="#" id="nextPage">Next</a></li>
                    </ul>
                </nav>
            </div>
            <div class="invitation-form-container">
                <form id="invitationForm" class="form-inline">
                    <input type="email" class="form-control" id="inviteeEmail" placeholder="Email" required style="width: 350px;">
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
                        <th scope="col">User Role <i class="fas fa-info-circle" title="Roles include requester, judger, and admin. Requesters can only make requests, judgers can make and approve requests, and admins have judger privileges plus administrative access."></i></th>
                        <th scope="col" class="center-align">Available <i class="fas fa-info-circle" title="Turning this off will prevent the user from sign in."></i></th>
                        <th scope="col" class="center-align">Last SignIn Time</th>
                    </tr>
                </thead>
                <tbody></tbody>
            </table>
        </div>
    </div>
    <script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
    <script type="text/javascript">
        var roles = [
            {{range .AdminListPage.UserRoles}}
            "{{.}}",
            {{end}}
        ];
    </script>
    <script>
        $(document).ready(function() {
            var currentPage = 1;
            var pageSize = $('#pageSize').val();
            var totalPage = 0;

            function updateTable(data) {
                var tbody = $('table tbody');
                tbody.empty();
                $.each(data.users, function(i, user) {
                    var roleOptions = roles.map(function(role) {
                        var selected = role === user.role ? ' selected' : '';
                        return `<option value="${role}"${selected}>${role}</option>`;
                    }).join('');

                    var checkedAttribute = user.isAvailable ? 'checked' : '';
                    var row = `<tr data-user-id="${user.id}" data-role="${user.role}" data-is-invited="${user.isInvited}">
                        <td class="email">${user.email || ''} ${user.isInvited ? '<span class="badge badge-info">Invited</span>' : ''}</td>
                        <td class="role">
                            <select ${user.isInvited ? 'disabled' : ''} class="user-role form-control">${roleOptions}</select>
                        </td>
                        <td class="center-align available">
                            <label class="switch">
                                <input ${user.isInvited ? 'disabled' : ''} type="checkbox" class="availability-toggle" ${checkedAttribute}>
                                <span class="slider round ${user.isInvited ? 'slider-disable' : ''}"></span>
                            </label>
                        </td>
                        // Display the last sign-in time in local format if available
                        <td class="center-align last-signin-time">${user.lastSigninTime ? convertToLocaleTimeString(user.lastSigninTime) : ''}</td>
                    </tr>`;
                    tbody.append(row);
                });
            }

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

            function updatePagination(totalPageCount) {
                totalPage = totalPageCount
                var pagination = $('.pagination');
                pagination.empty();

                var prevDisabled = currentPage === 1 ? ' disabled' : '';
                pagination.append(`<li class="page-item${prevDisabled}"><a class="page-link" href="#" id="prevPage">Previous</a></li>`);

                var startPage = Math.max(1, currentPage - 4);
                var endPage = Math.min(totalPage, currentPage + 4);

                if (startPage > 1) {
                    pagination.append(`<li class="page-item"><a class="page-link" href="#" data-page="1">1</a></li>`);
                    if (startPage > 2) {
                        pagination.append(`<li class="page-item disabled"><span class="page-link">...</span></li>`);
                    }
                }

                for (var i = startPage; i <= endPage; i++) {
                    var active = i === currentPage ? ' active' : '';
                    pagination.append(`<li class="page-item${active}"><a class="page-link" href="#" data-page="${i}">${i}</a></li>`);
                }

                if (endPage < totalPage) {
                    if (endPage < totalPage - 1) {
                        pagination.append(`<li class="page-item disabled"><span class="page-link">...</span></li>`);
                    }
                    pagination.append(`<li class="page-item"><a class="page-link" href="#" data-page="${totalPage}">${totalPage}</a></li>`);
                }

                var nextDisabled = currentPage === totalPage ? ' disabled' : '';
                pagination.append(`<li class="page-item${nextDisabled}"><a class="page-link" href="#" id="nextPage">Next</a></li>`);

            }

            function fetchData(page, size) {
                $.get('/api/users', { pageID: page, size: size }, function(data) {
                    updateTable(data);
                    updatePagination(data.totalPage);
                });
            }

            $(document).on('click', '#prevPage', function(e) {
                e.preventDefault();
                if (currentPage > 1) {
                    currentPage--;
                    fetchData(currentPage, pageSize);
                }
            });

            $(document).on('click', '#nextPage', function(e) {
                e.preventDefault();
                 if (currentPage < totalPage) {
                    currentPage++;
                    fetchData(currentPage, pageSize);
                 }
            });

            $(document).on('click', '.page-link[data-page]', function(e) {
                e.preventDefault();
                var page = $(this).data('page');
                if (page) {
                    currentPage = page;
                    fetchData(currentPage, pageSize);
                }
            });

            $('#pageSize').change(function() {
                pageSize = $(this).val();
                currentPage = 1;
                fetchData(currentPage, pageSize);
            });

            fetchData(currentPage, pageSize);
        });

        $(document).on('change', '.user-role', function() {
            patchRequest(this);
        });

        $(document).on('change', '.availability-toggle', function() {
            patchRequest(this);
        });

        $(document).ready(function() {
            $('#invitationForm').on('submit', function(e) {
                e.preventDefault();
                var email = $('#inviteeEmail').val();
                var role = $('#inviteeRole').val();
                var data = {
                    email: email,
                    role: role
                };

                $.ajax({
                    url: '/api/invitations',
                    type: 'POST',
                    contentType: 'application/json',
                    data: JSON.stringify(data),
                    success: function() {
                        $('#invitationForm')[0].reset();
                        location.reload();
                    },
                    error: function(xhr, status, error) {
                        alert('Error sending invitation: ' + error);
                    }
                });
            });
        });

        function patchRequest(element) {
            var row = $(element).closest('tr');
            var userId = row.data('user-id');
            var userRole = row.find('.user-role').val();
            var userIsInvited = row.data('is-invited');
            var isAvailable = row.find('.availability-toggle').is(':checked');
            var data = {
                isAvailable: isAvailable,
                role: userRole
            };

            if (userIsInvited) {
                alert('Can\'t change invited user status');
                location.reload();
            }

            $.ajax({
                url: '/api/users/' + userId,
                type: 'PATCH',
                contentType: 'application/json',
                data: JSON.stringify(data),
                success: function(response) {
                    alert('Request updated successfully');
                    console.debug('User updated successfully');
                },
                error: function(xhr, status, error) {
                    console.error('Update failed: ' + error);
                    alert('Error updating request');
                }
            });
        };
    </script>
</body>
</html>
