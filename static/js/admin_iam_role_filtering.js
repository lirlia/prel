$(document).ready(function () {

    function fetchRules() {
        $.get('/api/iam-role-filtering-rules', function (data) {
            var rulesList = $('#rulesList');
            rulesList.empty();
            $.each(data.iamRoleFilteringRules, function (i, rule) {
                var row = `<tr>
                    <td>${rule.pattern}</td>
                    <td><button class="btn btn-danger removeFilter" data-rule-id="${rule.id}">Remove</button></td>
                </tr>`;
                rulesList.append(row);
            });
        });
    }

    $('#addFilter').click(function () {
        var pattern = $('#filterInput').val();
        if (pattern.length < 3 || pattern.length > 20) {
            $('#rule-warning').show();
            return;
        } else {
            $('#rule-warning').hide();
        }

        var pattern = $('#filterInput').val();
        if (pattern) {
            $.ajax({
                url: '/api/iam-role-filtering-rules',
                type: 'POST',
                contentType: 'application/json',
                data: JSON.stringify({ pattern: pattern }),
                success: function () {
                    $('#filterInput').val('');
                    fetchRules();
                },
                error: function (xhr, status, error) {
                    alert('Error adding rule: ' + error);
                }
            });
        }
    });

    $(document).on('click', '.removeFilter', function () {
        var ruleID = $(this).data('rule-id');
        $.ajax({
            url: '/api/iam-role-filtering-rules/' + ruleID,
            type: 'DELETE',
            success: function () {
                fetchRules();
            },
            error: function (xhr, status, error) {
                alert('Error removing rule: ' + error);
            }
        });
    });

    fetchRules();
});
