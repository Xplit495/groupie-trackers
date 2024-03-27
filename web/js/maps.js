// Declaration of global variables for map, info window, and marker
let map;
let firstInfoWindow;
let firstMarker;

// Function to initialize the Google Map
function initMap() {
    // Creating a new Google Map instance
    map = new google.maps.Map(document.getElementById("map"), {
        center: { lat: 48.8566, lng: 2.3522 }, // This value will be overwritten by the first concert location
        zoom: 3,
    });
}

// Function to geocode the provided address and display it on the map with concert dates
function geocodeAddress(geocoder, address, concertDates) {
    // Using Google Maps Geocoder to convert address into geographic coordinates
    geocoder.geocode({ 'address': address }, function(results, status) {
        if (status === 'OK') {
            // Creating a marker for the geocoded location
            let marker = new google.maps.Marker({
                map: map,
                position: results[0].geometry.location,
                title: address
            });

            // Constructing content for the info window including concert dates
            let contentString = '<div id="content">' +
                '<div id="siteNotice">' +
                '</div>' +
                '<h2 id="firstHeading" class="firstHeading">' + address + '</h2>' +
                '<div id="bodyContent">' +
                '<h2><br>Date de concert : <br><br></h2>';

            concertDates.forEach(function(date) {
                contentString += '<h3>' + date + '<br><br></h3>';
            });

            contentString += '</div></div>';

            // Creating an info window for the marker
            let infoWindow = new google.maps.InfoWindow({
                content: contentString
            });

            // Adding a click event listener to open the info window when marker is clicked
            marker.addListener('click', function() {
                infoWindow.open(map, marker);
            });

            // Setting the center of the map to the first marker's position
            if (!firstMarker) {
                firstMarker = marker;
                firstInfoWindow = infoWindow;
                map.setCenter(results[0].geometry.location);
            }

        } else {
            // Logging error if geocoding fails
            console.error('Geocoding for address "' + address + '" failed with status: ' + status);
        }
    });
}

// Function to iterate through provided locations and geocode them
function codeAddresses(locations) {
    let geocoder = new google.maps.Geocoder();

    Object.keys(locations).forEach((location) => {
        const address = location.replace(/_/g, ' ');
        const concertDates = locations[location];
        geocodeAddress(geocoder, address, concertDates);
    });
}

// Function to fetch concert locations from API and initiate geocoding process
function getConcertLocations() {
    // Retrieving group ID from URL parameters
    const urlParams = new URLSearchParams(window.location.search);
    const groupId = urlParams.get('id');

    // Checking if group ID is available
    if (!groupId) {
        console.error('Group ID is missing in the URL');
        return;
    }

    // Constructing URL for fetching concert locations based on group ID
    let url = new URL('/api/search/locations', window.location.origin);
    url.searchParams.append('id', groupId);

    // Fetching concert locations from API
    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            // Initiating geocoding process with fetched data
            codeAddresses(data);
        })
        .catch(error => console.error(error));
}

// Exposing initMap function to window object for Google Maps API
window.initMap = initMap;

// Calling getConcertLocations function when window is loaded
window.onload = getConcertLocations;
