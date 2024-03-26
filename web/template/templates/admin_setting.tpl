<!DOCTYPE html>
<html lang="en">
<head>
    <link rel="stylesheet" href="/static/css/admin_user.css">
    {{template "header" .}}
</head>
<body>
    <div class="container-fluid">
        <h2>Setting</h2>
        <div class="table-responsive">
            <table class="table">
				<thead>
					<tr>
						<th scope="col">Key</th>
						<th scope="col">Value</th>
						<!-- button -->
						<th scope="col"></th>
					</tr>
					<tr>
						<td>Request Notification Message</td>
						<td>
							<textarea
								maxlength="1000"
								class="form-control"
								id="notification-message-for-request"
								rows="3">{{.AdminSettingPage.NotificationMessageForRequest}}</textarea>
						</td>
						<td>
							<button
								type="button"
								class="btn btn-primary settingButton"
								data-related-field="notification-message-for-request"
							>Update</button>
						</td>
					</tr>
					<tr>
						<td>Judge Notification Message</td>
						<td>
							<textarea
								maxlength="1000"
								class="form-control"
								id="notification-message-for-judge"
								rows="3">{{.AdminSettingPage.NotificationMessageForJudge}}</textarea>
						</td>
						<td>
							<button type="button" class="btn btn-primary settingButton" data-related-field="notification-message-for-judge">Update</button>
						</td>
					</tr>
				</thead>
            </table>
        </div>
    </div>
    <script src="/static/js/admin_setting.js"></script>
</body>
</html>
