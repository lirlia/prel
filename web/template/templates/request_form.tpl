<!DOCTYPE html>
<html lang="en">
<head>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.13/css/select2.min.css" rel="stylesheet">
    <style>
        .container {
            margin-top: 20px;
        }

        .container h2 {
            padding: 20px 0;
        }

        .select2-container--default .select2-selection--single {
            height: calc(2.25rem + 2px);
        }
        .select2-container--default .select2-selection--single .select2-selection__rendered {
            line-height: 2.25rem;
        }
        .select2-container .select2-selection--single .select2-selection__rendered {
            padding-left: 12px;
        }
        .select2-container--default .select2-results>.select2-results__options {
            max-height: 800px;
            overflow-y: auto;
        }
    </style>
    {{template "header" .}}
</head>
<body>
    <div class="container">
        <h2>IAM Role Request Form</h2>
        <form id="iamRoleRequestForm" action="/request-form" method="post" enctype="multipart/form-data">
            <div class="form-group">
                <label for="email">Email</label>
                <select id="email" name="email" class="form-control" disabled>
                    <option selected disabled>{{.RequestFormPage.Email}}</option>
                </select>
            </div>
            <div class="form-group">
                <label for="project_id">Project ID</label>
                <select id="project_id" name="project_id" class="form-control">
                    <option selected disabled>Select Project</option>
                    {{range .RequestFormPage.Projects }}
                    <option value="{{.ProjectID}}">{{.Name}}</option>
                    {{end}}
                </select>
                <div id="projectId-warning" class="alert alert-warning" style="display: none;">
                    Please select project.
                </div>
            </div>
            <div class="form-group">
                <label for="role">IAM Role</label>
                <select id="role" name="iam_roles" class="form-control" multiple="multiple">
                </select>
                <div id="role-warning" class="alert alert-warning" style="display: none;">
                    Please select at least one role.
                </div>
            </div>
            <div class="form-group">
                <label for="period">Period</label>
                <select id="period" name="period" class="form-control">
                    {{range .RequestFormPage.Periods }}
                        <option value="{{.Key}}">{{.Value}}</option>
                    {{end}}
                </select>
            </div>
            <div class="form-group">
                <label for="reason">Reason</label>
                <textarea id="reason" name="reason" class="form-control" rows="4" maxlength="500"></textarea>
                <div id="charCount">0 / 500</div>
            </div>
            <button type="submit" id="submit-request" class="btn btn-primary">Request</button>
        </form>
    </div>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.13/js/select2.min.js"></script>
    <script>
        $(document).ready(function() {
            $('#email').select2();
            $('#period').select2();
            $('#project_id').select2().on('change', function() {
                var projectId = $(this).val();
                updateIamRoles(projectId);
            });

            $('#role').select2({
                placeholder: "  Select IAM roles",
                disabled: true
            });

            $('#reason').on('input', function() {
                var currentLength = $(this).val().length;
                $('#charCount').text(currentLength + ' / 500');
            });

            function updateIamRoles(projectId) {
                var $roleSelect = $('#role');
                $roleSelect.empty();
                $('#role').select2({
                    placeholder: '  Loading...',
                });
                $.ajax({
                    url: '/api/iam-roles',
                    type: 'GET',
                    data: { projectID: projectId },
                    success: function(response) {
                        var roles = response.iamRoles;
                        roles.forEach(function(role) {
                            $roleSelect.append(new Option(role, role));
                        });
                        $roleSelect.trigger('change');
                        $roleSelect.select2({
                            placeholder: '  Select IAM roles',
                            allowClear: true,
                            disabled: false,
                        });
                    },
                    error: function(xhr, status, error) {
                        alert("Error fetching IAM roles:", error);
                        location.reload();
                    }
                });
            }

            $('#iamRoleRequestForm').on('submit', function(e) {
                e.preventDefault();
                var selectedProjectId = $('#project_id').val();

                if (selectedProjectId === null) {
                    $('#projectId-warning').show();
                    return;
                } else {
                    $('#projectId-warning').hide();
                }

                var selectedRoles = $('#role').val();
                if (selectedRoles.length === 0) {
                    $('#role-warning').show();
                    return;
                } else {
                    $('#role-warning').hide();
                }

                var formData = {
                    projectID: $('#project_id').val(),
                    iamRoles: $('#role').val(),
                    period: parseInt($('#period').val(), 10),
                    reason: $('#reason').val()
                };

                $.ajax({
                    url: '/api/requests',
                    type: 'POST',
                    contentType: 'application/json',
                    data: JSON.stringify(formData),
                    success: function(response) {
                        var requestId = response.requestID;
                        window.location.href = '/request/' + requestId;
                    },
                    error: function(xhr, status, error) {
                        console.error("Error: ", error);
                        alert("An error occurred: " + error);
                    }
                });
            });
        });
    </script>
</body>
</html>
