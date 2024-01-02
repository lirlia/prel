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
    $('.expire-time').each(function () {
        var originalText = $(this).text();
        var convertedText = convertToLocaleTimeString(originalText);
        $(this).text(convertedText).css('display', 'table-cell');
    });
});

function updateRequestStatus(requestId, status) {
    $.ajax({
        url: '/api/requests/' + requestId,
        type: 'PATCH',
        contentType: 'application/json',
        data: JSON.stringify({ status: status }),
        success: function (response) {
            alert('Request updated successfully');
            location.reload();
        },
        error: function () {
            alert('Error updating request');
        }
    });
}

function deleteRequest(requestId) {
    $.ajax({
        url: '/api/requests/' + requestId,
        type: 'DELETE',
        success: function (response) {
            alert('Request deleted successfully');
            var currentPath = window.location.pathname;

            if (currentPath.includes('/request/')) {
                window.location.pathname = '/request-form';
            } else {
                location.reload();
            }
        },
        error: function () {
            alert('Error deleting request');
        }
    });
}
