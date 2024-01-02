$(document).ready(function () {
    $('#email').select2();
    $('#period').select2();
    $('#project_id').select2().on('change', function () {
        var projectId = $(this).val();
        updateIamRoles(projectId);
    });

    $('#role').select2({
        placeholder: "  Select IAM roles",
        disabled: true
    });

    $('#reason').on('input', function () {
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
            success: function (response) {
                var roles = response.iamRoles;
                roles.forEach(function (role) {
                    $roleSelect.append(new Option(role, role));
                });
                $roleSelect.trigger('change');
                $roleSelect.select2({
                    placeholder: '  Select IAM roles',
                    allowClear: true,
                    disabled: false,
                });
            },
            error: function (xhr, status, error) {
                alert("Error fetching IAM roles:", error);
                location.reload();
            }
        });
    }

    $('#iamRoleRequestForm').on('submit', function (e) {
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
            success: function (response) {
                var requestId = response.requestID;
                window.location.href = '/request/' + requestId;
            },
            error: function (xhr, status, error) {
                console.error("Error: ", error);
                alert("An error occurred: " + error);
            }
        });
    });
});
