'use strict'

$(document).ready(function () {
    let hash = window.location.hash;
    const tokenDiv = `<div class="media text-muted pt-3">
                <img data-src="holder.js/32x32?theme=thumb&bg=e83e8c&fg=e83e8c&size=1" alt="" class="mr-2 rounded">
                <p class="media-body pb-3 mb-0 small lh-125 border-bottom border-gray" id="token">
                    <strong class="d-block text-gray-dark">Access Token</strong>
                    ##access_token##
                </p>
            </div>`
    const refreshTokenDiv = `<div class="media text-muted pt-3">
                <img data-src="holder.js/32x32?theme=thumb&bg=e83e8c&fg=e83e8c&size=1" alt="" class="mr-2 rounded">
                <p class="media-body pb-3 mb-0 small lh-125 border-bottom border-gray" id="token">
                    <strong class="d-block text-gray-dark">Refresh Token</strong>
                    ##access_token##
                </p>
            </div>`

    if (hash) {
        hash = hash.slice(1)
        const urlParams = new URLSearchParams(hash);
        const token = urlParams.get('access_token');
        const code = urlParams.get('code');
        if (token) {
            $("#main").append(tokenDiv.replace("##access_token##", token));
            $("#error").remove();
        }

        if (code) {
            $.post("/hybrid", "code=" + code, function (data) {
                $("#main").append(tokenDiv.replace("##access_token##", data.access_token));
                $("#main").append(refreshTokenDiv.replace("##access_token##", data.refresh_token));
                $("#error").remove();
            });
        }
    }
})