const params = new URLSearchParams(window.location.search);
const artistId = params.get('id');

fetch(`/api/deezer?artistId=${artistId}`)
    .then(response => response.text())
    .then(artistId => {

            let deezerWidgetContainer = document.querySelector('.deezer-widget-container');

            const deezerWidget = document.createElement('div');
            deezerWidget.id = 'deezer-widget';

            const iframe = document.createElement('iframe');
            iframe.setAttribute('title', 'deezer-widget');
            iframe.src = `https://widget.deezer.com/widget/dark/artist/${artistId}/top_tracks`;
            iframe.width = "500";
            iframe.height = "690";
            iframe.frameBorder = "0";
            iframe.allowTransparency = "true";
            iframe.allow = "encrypted-media; clipboard-write";

            deezerWidget.appendChild(iframe);

            deezerWidgetContainer.appendChild(deezerWidget);

    })

    .catch(error => {
            console.error('Error:', error);
    });

