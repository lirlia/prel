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

$(document).ready(function () {
    var currentPage = 1;
    var pageSize = $('#pageSize').val();
    var totalPage = 0;

    function updateTable(data) {
        var tbody = $('table tbody');
        tbody.empty();
        $.each(data.requests, function (i, request) {
            var roles = request.iamRoles.join('<br>');
            var row = `<tr>
                <td>${request.requester || ''}</td>
                <td>${request.judger || ''}</td>
                <td>${request.projectID}</td>
                <td>${roles}</td>
                <td class="center-align">${request.period}</td>
                <td>${request.reason}</td>
                <td class="center-align">${request.status}</td>
                <td class="center-align">${convertToLocaleTimeString(request.requestTime)}</td>
                <td class="center-align">${request.judgeTime !== '0001-01-01T00:00:00Z' ? convertToLocaleTimeString(request.judgeTime) : ''}</td>
                <td class="center-align">${convertToLocaleTimeString(request.expireTime)}</td>
            </tr>`;
            tbody.append(row);
        });
    }

    function updatePagination(totalPageCount) {
        totalPage = totalPageCount
        var pagination = $('.pagination');
        pagination.empty();

        var prevDisabled = currentPage === 1 ? ' disabled' : '';
        pagination.append(`<li class="page-item${prevDisabled}"><a class="page-link" href="#" id="prevPage">&lt;</a></li>`);

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
        pagination.append(`<li class="page-item${nextDisabled}"><a class="page-link" href="#" id="nextPage">&gt;</a></li>`);

    }

    function fetchData(page, size) {
        $.get('/api/requests', { pageID: page, size: size }, function (data) {
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
