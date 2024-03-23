let map;
let firstInfoWindow;
let firstMarker;

function initMap() {
    map = new google.maps.Map(document.getElementById("map"), {
        center: { lat: 48.8566, lng: 2.3522 }, // Cette valeur sera écrasée par le premier emplacement de concert
        zoom: 3,
    });
}

function geocodeAddress(geocoder, address, concertDates) {
    geocoder.geocode({ 'address': address }, function(results, status) {
        if (status === 'OK') {
            let marker = new google.maps.Marker({
                map: map,
                position: results[0].geometry.location,
                title: address
            });

            let contentString = '<div id="content">'+
                '<div id="siteNotice">'+
                '</div>'+
                '<h2 id="firstHeading" class="firstHeading">' + address + '</h2>'+
                '<div id="bodyContent">'+
                '<h2><br>Date de concert : <br><br></h2>';

            concertDates.forEach(function(date) {
                contentString += '<h3>' + date + '<br><br></h3>';
            });

            contentString += '</div></div>';

            let infoWindow = new google.maps.InfoWindow({
                content: contentString
            });

            marker.addListener('click', function() {
                infoWindow.open(map, marker);
            });

            if (!firstMarker) {
                firstMarker = marker;
                firstInfoWindow = infoWindow;
                map.setCenter(results[0].geometry.location);
            }

        } else {
            console.error('Le géocodage pour l\'adresse "' + address + '" n\'a pas fonctionné pour la raison suivante: ' + status);
        }
    });
}

function codeAddresses(locations) {
    let geocoder = new google.maps.Geocoder();

    Object.keys(locations).forEach((location) => {
        const address = location.replace(/_/g, ' ');
        const concertDates = locations[location];
        geocodeAddress(geocoder, address, concertDates);
    });
}

function getConcertLocations() {
    const urlParams = new URLSearchParams(window.location.search);
    const groupId = urlParams.get('id');

    if (!groupId) {
        console.error('L\'ID du groupe est manquant dans l\'URL');
        return;
    }

    let url = new URL('/api/search/locations', window.location.origin);
    url.searchParams.append('id', groupId);

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error('La réponse du réseau n\'était pas ok');
            }
            return response.json();
        })
        .then(data => {
            codeAddresses(data);
        })
        .catch(error => console.error(error));
}

window.initMap = initMap;

window.onload = getConcertLocations;
