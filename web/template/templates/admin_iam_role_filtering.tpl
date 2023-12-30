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

        .table-responsive, .rules-table-responsive {
            overflow-x: auto;
            margin-right: -15px;
        }

        .controls-container {
            display: flex;
            align-items: start;
            margin-bottom: 15px;
        }

        .input-container {
            width: 30%;
            margin-right: 15px;
        }

        .button-container {
            display: flex;
            align-items: center;
        }

        #filterInput, #rule-warning {
            width: 100%;
            max-width: 500px;
        }
    </style>
    {{template "header" .}}
</head>
<body>
    <div class="container-fluid">
        <h2>IAM Role Filtering</h2>
        <p>If a keyword matches any of the patterns, it is considered a valid IAM role.</p>
        <p><strong>Note:</strong> Regular expressions are not supported in these patterns.</p>
        <div class="controls-container">
            <div class="input-container">
                <input type="text" id="filterInput" class="form-control" placeholder="Enter a keyword between 3 and 20 characters" minlength="3" maxlength="20">
                <div id="rule-warning" class="alert alert-warning" style="display: none;">
                    Please enter a keyword between 3 and 20 characters
                </div>
            </div>
            <div class="button-container">
                <button id="addFilter" class="btn btn-primary">Add Keyword</button>
            </div>
        </div>
        <div class="rules-table-responsive">
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col-1">Keyword</th>
                        <th scope="col-1"></th>
                    </tr>
                </thead>
                <tbody id="rulesList"></tbody>
            </table>
        </div>
    </div>
    <script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
    <script>
        $(document).ready(function() {

            function fetchRules() {
                $.get('/api/iam-role-filtering-rules', function(data) {
                    var rulesList = $('#rulesList');
                    rulesList.empty();
                    $.each(data.iamRoleFilteringRules, function(i, rule) {
                        var row = `<tr>
                            <td>${rule.pattern}</td>
                            <td><button class="btn btn-danger removeFilter" data-rule-id="${rule.id}">Remove</button></td>
                        </tr>`;
                        rulesList.append(row);
                    });
                });
            }

            $('#addFilter').click(function() {
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
                        success: function() {
                            $('#filterInput').val('');
                            fetchRules();
                        },
                        error: function(xhr, status, error) {
                            alert('Error adding rule: ' + error);
                        }
                    });
                }
            });

            $(document).on('click', '.removeFilter', function() {
                var ruleID = $(this).data('rule-id');
                $.ajax({
                    url: '/api/iam-role-filtering-rules/' + ruleID,
                    type: 'DELETE',
                    success: function() {
                        fetchRules();
                    },
                    error: function(xhr, status, error) {
                        alert('Error removing rule: ' + error);
                    }
                });
            });

            fetchRules();
        });
    </script>
</body>
</html>
