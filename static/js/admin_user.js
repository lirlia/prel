$(document).ready(function () {
    var currentPage = 1;
    var pageSize = $('#pageSize').val();
    var totalPage = 0;

    function updateTable(data) {
        var tbody = $('table tbody');
        tbody.empty();
        $.each(data.users, function (i, user) {
            var roleOptions = roles.map(function (role) {
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
        pagination.append(`<li class="page-item${prevDisabled}"><a class="page-link" href="#" id="prevPage">&lt;</a></li>`);

        var startPage = Math.max(1, currentPage - 2);
        var endPage = Math.min(totalPage, currentPage + 2);

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
        pagination.append(`<li class="page-item${nextDisabled}"><a class="page-link" href="#" id="nextPage">&gt;</a></li>`);

    }

    function fetchData(page, size) {
        $.get('/api/users', { pageID: page, size: size }, function (data) {
            updateTable(data);
            updatePagination(data.totalPage);
        });
    }

    $(document).on('click', '#prevPage', function (e) {
        e.preventDefault();
        if (currentPage > 1) {
            currentPage--;
            fetchData(currentPage, pageSize);
        }
    });

    $(document).on('click', '#nextPage', function (e) {
        e.preventDefault();
        if (currentPage < totalPage) {
            currentPage++;
            fetchData(currentPage, pageSize);
        }
    });

    $(document).on('click', '.page-link[data-page]', function (e) {
        e.preventDefault();
        var page = $(this).data('page');
        if (page) {
            currentPage = page;
            fetchData(currentPage, pageSize);
        }
    });

    $('#pageSize').change(function () {
        pageSize = $(this).val();
        currentPage = 1;
        fetchData(currentPage, pageSize);
    });

    fetchData(currentPage, pageSize);
});

$(document).on('change', '.user-role', function () {
    patchRequest(this);
});

$(document).on('change', '.availability-toggle', function () {
    patchRequest(this);
});

$(document).ready(function () {
    $('#invitationForm').on('submit', function (e) {
        e.preventDefault();
        var email = $('#inviteeEmail').val();
        var role = $('#inviteeRole').val();
        email = email.replace(/</g, "&lt;").replace(/>/g, "&gt;");
        role = role.replace(/</g, "&lt;").replace(/>/g, "&gt;");
        var data = {
            email: email,
            role: role
        };

        $.ajax({
            url: '/api/invitations',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function () {
                $('#invitationForm')[0].reset();
                location.reload();
            },
            error: function (xhr, status, error) {
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
    userRole = userRole.replace(/</g, "&lt;").replace(/>/g, "&gt;");
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
        success: function (response) {
            alert('Request updated successfully');
            console.debug('User updated successfully');
        },
        error: function (xhr, status, error) {
            console.error('Update failed: ' + error);
            alert('Error updating request');
        }
    });
};
