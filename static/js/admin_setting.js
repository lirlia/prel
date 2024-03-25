$(document).ready(function () {

    // クリックイベントに変更しました
    $('.settingButton').on('click', function () {

        const id = $(this).data('related-field');
        const value = $(`#${id}`).val();
        updateSettingValue(id, value);
    });

    function updateSettingValue(id, value) {

        const data = () => {
            switch (id) {
                case 'notification-message-for-request':
                    return { notificationMessageForRequest: value };
                case 'notification-message-for-judge':
                    return { notificationMessageForJudge: value };
                default:
                    return undefined;
            }
        }

        if (data() === undefined) {
            console.error("Data is undefined for id: ", id);
            return;
        }

        $.ajax({
            url: '/api/settings',
            type: 'PATCH',
            contentType: 'application/json',
            data: JSON.stringify(data()),
            success: function (response) {
                alert("Update successful!");
            },
            error: function (xhr, status, error) {
                console.error("Error: ", error);
                alert("An error occurred: " + error);
            }
        });
    }

});
