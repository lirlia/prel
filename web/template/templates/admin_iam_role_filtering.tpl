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
            justify-content: space-between;
            margin-bottom: 15px;
            width: 100%;
        }

        .input-container {
            flex-grow: 1;
            margin-right: 15px;
        }

        .button-container {
            flex-shrink: 0;
        }

        #filterInput, #rule-warning {
            width: 100%;
            max-width: 500px;
        }

        .table th, .table td {
            text-align: left;
        }

        @media (max-width: 768px) {
            .table tr {
                display: flex;
                align-items: center;
                border-top: 1px solid #dee2e6;
            }

            .table td, .table th {
                flex: 1;
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 8px;
                border-top: none !important;
                border-bottom: none !important;
            }

            .table td:first-child, .table th:first-child {
                justify-content: center;
            }

            .table td:last-child, .table th:last-child {
                justify-content: center;
            }

            .removeFilter {
                margin-top: 0;
                margin-left: 10px;
            }

            .controls-container {
                flex-direction: column;
                align-items: stretch;
            }

            .input-container, .button-container {
                width: 100%;
                margin-right: 0;
            }

            .button-container {
                margin-top: 10px;
            }
        }

    </style>
    {{template "header" .}}
</head>
<body>
    <div class="container-fluid">
        <h2>IAM Role Filtering</h2>
        <p>Matching keywords validate IAM roles.</p>
        <p><strong>Note:</strong> No regular expressions.</p>
        <div class="controls-container">
            <div class="input-container">
                <input type="text" id="filterInput" class="form-control" placeholder="Enter a keyword between 3 and 20 characters" minlength="3" maxlength="20">
                <div id="rule-warning" class="alert alert-warning" style="display: none;">
                    Keyword is between 3 ~ 20 characters
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
                        <th scope="col-1">Action</th>
                    </tr>
                </thead>
                <tbody id="rulesList"></tbody>
            </table>
        </div>
    </div>
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
