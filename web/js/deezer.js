// Extracting parameters from the URL query string
const params = new URLSearchParams(window.location.search);

// Getting the value of the 'id' parameter from the URL query string
const artistId = params.get('id');

// Fetching data from a Deezer API endpoint using the artistId
fetch(`/api/deezer?artistId=${artistId}`)
    .then(response => response.text()) // Parsing the response as text
    .then(artistId => {
        // Callback function handling the successful response

        // Finding the container element for the Deezer widget
        let deezerWidgetContainer = document.querySelector('.deezer-widget-container');

        // Creating a container div for the Deezer widget
        const deezerWidget = document.createElement('div');
        deezerWidget.id = 'deezer-widget';

        // Creating an iframe element for the Deezer widget
        const iframe = document.createElement('iframe');
        iframe.setAttribute('title', 'deezer-widget'); // Setting title attribute
        iframe.src = `https://widget.deezer.com/widget/dark/artist/${artistId}/top_tracks`; // Setting source URL
        iframe.width = "500"; // Setting width
        iframe.height = "690"; // Setting height
        iframe.frameBorder = "0"; // Setting frameborder
        iframe.allowTransparency = "true"; // Allowing transparency
        iframe.allow = "encrypted-media; clipboard-write"; // Setting allow attribute

        // Appending iframe to the Deezer widget container
        deezerWidget.appendChild(iframe);

        // Appending the Deezer widget container to the document
        deezerWidgetContainer.appendChild(deezerWidget);

    })
    .catch(error => {
        // Error handling in case the fetch operation fails
        console.error('Error:', error);
    });
