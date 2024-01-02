<!DOCTYPE html>
<html lang="en">
<head>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.13/css/select2.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/request_form.css">
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
    <script src="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.13/js/select2.min.js"></script>
    <script src="/static/js/request_form.js"></script>
</body>
</html>
